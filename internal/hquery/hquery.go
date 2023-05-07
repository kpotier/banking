// Package hquery provides tools to query a parsed HTML document with CSS
// selectors. It is a light-weight alternative to goquery.
package hquery

import (
	"fmt"
	"net/url"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// FindAll returns a slice of the nodes that match the selector, from n and its
// children.
func FindAll(sel string, n *html.Node) ([]*html.Node, error) {
	s, err := cascadia.Compile(sel)
	if err != nil {
		return nil, err
	}
	return s.MatchAll(n), nil
}

// FindFirst returns the first node that matches the selector, from n and its
// children.
func FindFirst(sel string, n *html.Node) (*html.Node, error) {
	s, err := cascadia.Compile(sel)
	if err != nil {
		return nil, err
	}
	return s.MatchFirst(n), nil
}

// FindAttr returns the attribute value of n.
func FindAttr(attr string, n *html.Node) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}

// NewForm returns a pre-filled map from the selector. A specific submit can be
// specified can be specified in cases where there is more than one. If it is
// not specified, the first submit will be chosen.
func NewForm(sel, selSubmit string, n *html.Node) (url.Values, error) {
	form := url.Values{}
	f, err := FindFirst(sel, n)
	if err != nil {
		return nil, fmt.Errorf("find `%s`: %w", sel, err)
	}
	if f == nil {
		return nil, fmt.Errorf("could not find form `%s`", sel)
	}

	elems, _ := FindAll("input, textarea, select", f)

	var submit *html.Node
	if selSubmit != "" {
		submit, err = FindFirst(selSubmit, f)
		if err != nil {
			return nil, fmt.Errorf("find submit `%s` %w", selSubmit, err)
		}
		if submit == nil {
			return nil, fmt.Errorf("could not find submit `%s`", selSubmit)
		}
	}

	for _, e := range elems {
		name, ok := FindAttr("name", e)
		if !ok {
			continue
		}

		typ, _ := FindAttr("type", e)
		if typ == "checkbox" || typ == "radio" {
			if _, checked := FindAttr("checked", e); !checked {
				continue
			}
		} else if typ == "submit" {
			if submit == nil {
				submit = e
			}
			if submit != e {
				continue
			}
		}

		if e.Data == "select" {
			opts, _ := FindAll("option[selected]", e)
			for _, o := range opts {
				form.Add(name, getValueOrText(o))
			}
		} else {
			form.Set(name, getValueOrText(e))
		}
	}
	return form, nil
}

func getValueOrText(n *html.Node) string {
	val, ok := FindAttr("value", n)
	if !ok && n.FirstChild != nil {
		val = n.FirstChild.Data
	}
	return val
}
