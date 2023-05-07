package hparser_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"syscall"
	"testing"

	"github.com/kpotier/banking/internal/hparser"

	"golang.org/x/net/html"
)

func TestClient_Get(t *testing.T) {
	t.Run("refuse connection", func(t *testing.T) {
		// Should refuse connection and returns syscall.ECONNREFUSED
		ts := httptest.NewServer(nil)
		ts.Close()
		_, _, err := hparser.Get(hparser.Default(), ts.URL)
		if !errors.Is(err, syscall.ECONNREFUSED) {
			t.Errorf("Get() error = %v, wantErr %v", nil, "connection should be refused")
		}
	})

	t.Run("working conditions", func(t *testing.T) {
		text := "foo"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(text)) }))
		defer ts.Close()
		_, n, err := hparser.Get(hparser.Default(), ts.URL)
		if err != nil {
			t.Errorf("Get() error = %v, wantErr %v", err, false)
		} else if err := checkBody(text, n); err != nil {
			t.Errorf("Get(), error = %v, wantErr %v", err, false)
		}
	})

	t.Run("check cookie", func(t *testing.T) {
		c := hparser.Default()
		err := cookieServer(func(url string) (statusCode int, err error) {
			r, _, err := hparser.Get(c, url)
			return r.StatusCode, err
		})
		if err != nil {
			t.Errorf("cookieServer() error = %s, wantErr %v", err, false)
		}
	})
}

func TestClient_PostForm(t *testing.T) {
	t.Run("refuse connection", func(t *testing.T) {
		ts := httptest.NewServer(nil)
		ts.Close()
		_, _, err := hparser.PostForm(hparser.Default(), ts.URL, nil)
		if !errors.Is(err, syscall.ECONNREFUSED) {
			t.Errorf("PostForm() error = %v, wantErr %v", nil, "connection should be refused")
		}
	})

	t.Run("working conditions", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			w.Write([]byte(r.PostForm.Encode()))
		}))
		defer ts.Close()
		values := url.Values{"foo": {"bar", "baz"}, "foo2": {"bar2"}}
		_, n, err := hparser.PostForm(hparser.Default(), ts.URL, values)
		if err != nil {
			t.Errorf("PostForm() error = %v, wantErr %v", err, false)
		} else if err := checkBody(values.Encode(), n); err != nil {
			t.Errorf("PostForm(), error = %v, wantErr %v", err, false)
		}
	})

	t.Run("check cookie", func(t *testing.T) {
		c := hparser.Default()
		err := cookieServer(func(url string) (statusCode int, err error) {
			r, _, err := hparser.PostForm(c, url, nil)
			return r.StatusCode, err
		})
		if err != nil {
			t.Fatalf("cookieServer() error = %s, wantErr %v", err, false)
		}
	})
}

func checkBody(text string, n *html.Node) error {
	if n.FirstChild == nil || n.FirstChild.Data != "html" {
		return fmt.Errorf("cannot find html")
	} else if n.FirstChild.LastChild == nil || n.FirstChild.LastChild.Data != "body" {
		return fmt.Errorf("cannot find body")
	} else if n.FirstChild.LastChild.FirstChild == nil || n.FirstChild.LastChild.FirstChild.Data != text {
		return fmt.Errorf("cannot find %s", text)
	}
	return nil
}

func cookieServer(fn func(url string) (statusCode int, err error)) error {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("foo"); err != nil {
			http.SetCookie(w, &http.Cookie{Name: "foo", Value: "bar"})
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.Write([]byte("cookie found"))
		}
	}))
	defer ts.Close()
	if _, err := fn(ts.URL); err != nil {
		return fmt.Errorf("set cookie: %w", err)
	}
	if statusCode, err := fn(ts.URL); err != nil {
		return fmt.Errorf("get cookie: %w", err)
	} else if statusCode != http.StatusOK {
		return fmt.Errorf("could not find cookie")
	}
	return nil
}
