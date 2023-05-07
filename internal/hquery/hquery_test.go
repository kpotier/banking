package hquery_test

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/kpotier/banking/internal/hquery"

	"golang.org/x/net/html"
)

func TestFindAll(t *testing.T) {
	test := `
		<html>
			<strong>1</strong>
			<div>
				<strong>
					2
					<strong>3</strong>
				</strong>
			</div>
			<strong/>
			<strong />
		</html>
		`
	node, err := html.Parse(strings.NewReader(test))
	if err != nil {
		panic(err)
	}

	t.Run("find strong", func(t *testing.T) {
		n, err := hquery.FindAll("strong", node)
		if err != nil {
			t.Fatalf("FindAll() error = %v, wantErr %v", err, false)
		}
		if len(n) != 5 {
			t.Errorf("FindAll() len = %d, want %d", len(n), 5)
		}
		for i := 0; i < 3; i++ {
			c := n[i]
			number, err := strconv.Atoi(strings.TrimSpace(c.FirstChild.Data))
			if err != nil {
				t.Errorf("FindAll(), error = node %d does not have a number, got %s", i, c.FirstChild.Data)
			} else if number-1 != i {
				t.Errorf("FindAll(), error = expected order %d, got %d", i, number-1)
			}
		}
	})

	t.Run("wrong selector", func(t *testing.T) {
		sel := "?"
		_, err := hquery.FindAll(sel, node)
		if err == nil {
			t.Fatalf("FindAll() error = %v, wantErr %v", err, "expected identifier, found ? instead")
		}
	})
}

func TestFindFirst(t *testing.T) {
	test := `
		<html>
			<strong>1</strong>
			<div>
				<strong>
					2
					<strong>3</strong>
				</strong>
			</div>
		</html>
		`
	node, err := html.Parse(strings.NewReader(test))
	if err != nil {
		panic(err)
	}

	t.Run("find first strong", func(t *testing.T) {
		n, err := hquery.FindFirst("strong", node)
		if err != nil {
			t.Fatalf("FindFirst() error = %v, wantErr %v", err, false)
		}

		number, err := strconv.Atoi(strings.TrimSpace(n.FirstChild.Data))
		if err != nil {
			t.Errorf("FindFirst(), error = expected a number, got %s", n.FirstChild.Data)
		} else if number != 1 {
			t.Errorf("FindFirst(), error = order must be %d, got %d", 1, number)
		}

	})

	t.Run("wrong selector", func(t *testing.T) {
		sel := "?"
		_, err := hquery.FindFirst(sel, node)
		if err == nil {
			t.Fatalf("FindFirst() error = %v, wantErr %v", err, "expected identifier, found ? instead")
		}
	})
}

func TestFindAttr(t *testing.T) {
	test := "<input name=\"first\" disabled />"
	node, err := html.Parse(strings.NewReader(test))
	if err != nil {
		panic(err)
	}
	n, err := hquery.FindFirst("input", node)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		attr   string
		want   string
		wantOk bool
	}{
		{
			"name",
			"first",
			true,
		},
		{
			"disabled",
			"",
			true,
		},
		{
			"foo",
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run("find attr "+tt.attr, func(t *testing.T) {
			attr, ok := hquery.FindAttr(tt.attr, n)
			if tt.wantOk != ok {
				t.Errorf("FindAttr() ok = %v, want %v", ok, tt.wantOk)
			} else if attr != tt.want {
				t.Errorf("FindAttr() = %v, want %v", attr, tt.want)
			}
		})
	}
}

func TestNewForm(t *testing.T) {
	type args struct {
		sel       string
		selSubmit string
		html      string
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr bool
	}{
		{
			"empty form",
			args{
				"form#bar",
				"",
				`<form id="foo">
					<input type="text" name="i1" value="v1" />
				</form>
				<form id="bar"></form>
				<form id="baz">
					<input type="text" name="i2" value="v2" />
				</form>`,
			},
			url.Values{},
			false,
		},
		{
			"inputs",
			args{
				"form#foo",
				"",
				`<form id="foo">
					<input type="text" name="i1" value="v1" />
					<input type="text" name="i2"></input>

					<input type="password" value="v3" />
					<input type="password" name="i4" value="v4" />

					<textarea name="i5">Some <b>content</b></textarea>

					<input type="radio" name="i6" value="v6i">
					<input type="radio" name="i6" value="v6ii" checked>
					<input type="radio" name="i6" value="v6iii">

					<input type="checkbox" name="i7" value="v7">
					<input type="checkbox" name="i8" value="v8" checked>

					<input type="radio" name="i9" value="v9">

					<select name="i10">
						<option>v10i</option>
						<option>v10ii</option>
					</select>

					<select name="i11">
						<option selected>v11i</option>
						<option selected>v11ii</option>
						<option>v11iii</option>
					</select>

					<select name="i12"></select>
				</form>`,
			},
			url.Values{"i1": {"v1"}, "i2": {""}, "i4": {"v4"}, "i5": {"Some <b>content</b>"}, "i6": {"v6ii"}, "i8": {"v8"}, "i11": {"v11i", "v11ii"}},
			false,
		},
		{
			"submit",
			args{
				"form#foo",
				"",
				`<form id="foo">
					<input type="text" name="i1" value="v1" />
					<input type="submit" name="i2" value="v2"></input>
					<input type="submit" name="i3" value="v3"></input>
				</form>`,
			},
			url.Values{"i1": {"v1"}, "i2": {"v2"}},
			false,
		},
		{
			"specified submit",
			args{
				"form#foo",
				"input[type=\"submit\"]#bar",
				`<form id="foo">
					<input type="text" name="i1" value="v1" />
					<input type="submit" name="i2" value="v2"></input>
					<input type="submit" id="bar" name="i3" value="v3"></input>
				</form>`,
			},
			url.Values{"i1": {"v1"}, "i3": {"v3"}},
			false,
		},
		{
			"wrong selector",
			args{
				"?",
				"",
				"",
			},
			nil,
			true,
		},
		{
			"cannot find form",
			args{
				"form",
				"",
				"",
			},
			nil,
			true,
		},
		{
			"wrong submit selector",
			args{
				"form",
				"?",
				"<form></form>",
			},
			nil,
			true,
		},
		{
			"cannot find submit",
			args{
				"form",
				"input",
				"<form></form>",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := html.Parse(strings.NewReader(tt.args.html))
			if err != nil {
				panic(err)
			}
			got, err := hquery.NewForm(tt.args.sel, tt.args.selSubmit, parsed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
