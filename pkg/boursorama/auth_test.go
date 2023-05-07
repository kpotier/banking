package boursorama

import (
	"errors"
	"strings"
	"testing"

	"github.com/kpotier/banking/pkg/bank"

	"golang.org/x/net/html"
)

func Test_vKeyboardCode(t *testing.T) {
	type args struct {
		pwd string
		n   string
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
		wantErr  bool
		err      error
	}{
		{
			"check good code",
			args{"91230456877", authHTML},
			"NINE|ONE|TWO|THREE|ZERO|FOUR|FIVE|SIX|EIGHT|SEVEN|SEVEN",
			false,
			nil,
		},
		{
			"bad pwd",
			args{"01?", authHTML},
			"",
			true,
			bank.ErrBadPwd,
		},
		{
			"cannot find buttons",
			args{"", ""},
			"",
			true,
			nil,
		},
		{
			"cannot find attr",
			args{"", `<ul class="password-input"><li data-matrix-list-item data-matrix-list-item-index="0"><button type="button" class="sasmap__key"><img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMS42IDM2LjFjLjMuNC43LjcgMS40LjcuOSAwIDEuNC0uNiAxLjQtMS41di01aC45djVjMCAxLjYtMSAyLjMtMi4zIDIuMy0uOCAwLTEuNC0uMi0xLjktLjh6Ii8+PHBhdGggZD0ibTIwLjcgMzQuMy0uNy44djIuNGgtLjl2LTcuMmguOXYzLjdsMy4yLTMuN2gxLjFsLTMgMy40IDMuMiAzLjhoLTEuMXoiLz48cGF0aCBkPSJtMjcuNyAzMC4zaC45djYuNGgzLjR2LjhoLTQuMnYtNy4yeiIvPjwvZz48cGF0aCBkPSJtMTcuNCAyMC4xYzEuMSAxLjYgMi42IDIuNSA0LjggMi41IDIuNSAwIDQuMy0xLjggNC4zLTQuMiAwLTIuNi0xLjgtNC4yLTQuMy00LjItMS42IDAtMi45LjUtNC4yIDEuN2wtMS0uNnYtOWgxMHYxLjNoLTguNXY2LjhjLjktLjggMi4zLTEuNiA0LjEtMS42IDIuOSAwIDUuNSAxLjkgNS41IDUuNSAwIDMuNC0yLjYgNS42LTUuOCA1LjYtMi45IDAtNC42LTEuMS01LjgtMi44eiIvPjwvZz48L3N2Zz4="></button></li></ul>`},
			"",
			true,
			nil,
		},
		{
			"cannot find img",
			args{"", `<ul class="password-input"><li data-matrix-list-item data-matrix-list-item-index="0"><button type="button" data-matrix-key="FIVE" class="sasmap__key"></button></li></ul>`},
			"",
			true,
			nil,
		},
		{
			"cannot find corresponding md5 sum",
			args{"", `<ul class="password-input"><li data-matrix-list-item data-matrix-list-item-index="0"><button type="button" data-matrix-key="FIVE" class="sasmap__key"><img alt="" class="sasmap__img" src="foo"></button></li></ul>`},
			"",
			true,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := html.Parse(strings.NewReader(tt.args.n))
			if err != nil {
				panic(err)
			}
			gotCode, err := vKeyboardCode([]byte(tt.args.pwd), n)
			if (err != nil) != tt.wantErr {
				t.Errorf("vKeyboardCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.err != nil && !errors.Is(err, tt.err) {
				t.Errorf("vKeyboardCode() error = %v, want %v", err, tt.err)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("vKeyboardCode() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_matrixRandomChallenge(t *testing.T) {
	t.Run("check good challenge", func(t *testing.T) {
		p, err := html.Parse(strings.NewReader(`
			<script></script>
			<script>
				$(function () {
					$("[data-matrix-random-challenge]").val("somerandomchallenge")
				})
			</script>
		`))
		if err != nil {
			panic(err)
		}
		gotChallenge, err := matrixRandomChallenge(p)
		if err != nil {
			t.Errorf("matrixRandomChallenge() error = %v, wantErr %v", err, false)
		} else if gotChallenge != "somerandomchallenge" {
			t.Errorf("matrixRandomChallenge() = %v, want %v", gotChallenge, "somerandomchallenge")
		}
	})

	t.Run("check non existing challenge", func(t *testing.T) {
		p, err := html.Parse(strings.NewReader("<script></script><script></script>"))
		if err != nil {
			panic(err)
		}
		gotChallenge, err := matrixRandomChallenge(p)
		if err == nil {
			t.Errorf("matrixRandomChallenge() = %s, wantErr %v", gotChallenge, true)
		}
	})
}
