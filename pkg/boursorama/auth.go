package boursorama

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kpotier/banking/internal/hparser"
	"github.com/kpotier/banking/internal/hquery"
	"github.com/kpotier/banking/pkg/bank"

	"golang.org/x/net/html"
)

func (b *Boursorama) login(username, pwd []byte, ctx context.Context, q chan<- string, a <-chan string) error {
	// To connect we first need to get the form in the login page. Then, we need
	// to find the matrix random challenge and add it to the form. Finally, we
	// must convert the password into a code.

	// We are no longer logged in and we generate a new instance of Client with
	// an empty cookie jar.
	b.loggedIn = false
	b.http = hparser.Default()

	_, n, err := hparser.Get(b.http, url_base+url_login)
	if err != nil {
		return err
	}
	cookieNode, err := hquery.FindFirst("script", n)
	if err != nil {
		return fmt.Errorf("cookie: %w", err)
	}
	if cookieNode.FirstChild == nil {
		return fmt.Errorf("cookie: could not find __brs_mit cookie")
	}
	idx := strings.Index(cookieNode.FirstChild.Data, "__brs_mit=")
	idx2 := strings.IndexByte(cookieNode.FirstChild.Data, ';')
	cookie := cookieNode.FirstChild.Data[idx+len("__brs_mit=") : idx2]
	url, err := url.Parse(url_base)
	if err != nil {
		return fmt.Errorf("url.Parse: %w", err)
	}
	b.http.Jar.SetCookies(url, []*http.Cookie{{Name: "__brs_mit", Value: cookie}})

	_, n, err = hparser.Get(b.http, url_base+url_login)
	if err != nil {
		return err
	}
	form, err := hquery.NewForm("form", "", n)
	if err != nil {
		return fmt.Errorf("form: %w", err)
	}
	form.Set("form[clientNumber]", string(username))

	_, n, err = hparser.Get(b.http, url_base+url_login_keyboard)
	if err != nil {
		return err
	}
	ch, err := matrixRandomChallenge(n)
	if err != nil {
		return fmt.Errorf("matrixRandomChallenge: %w", err)
	}
	form.Set("form[matrixRandomChallenge]", ch)
	c, err := vKeyboardCode(pwd, n)
	if err != nil {
		return fmt.Errorf("vKeyboardCode: %w", err)
	}
	form.Set("form[password]", c)

	// Now we post the form and we check if the user is logged in. For now we do
	// not support 2FA authentification and some checks may be missing.
	r, n, err := hparser.PostForm(b.http, url_base+url_login_POST, form)
	if err != nil {
		return err
	}
	switch r.Request.URL.Path {
	case url_account_locked:
		e, _ := hquery.FindFirst("h2[class*=\"error\"]", n)
		if e == nil || e.FirstChild == nil {
			return fmt.Errorf("post %s: could not find error message", url_base+url_account_locked)
		}
		return bank.UnrecognizedError(e.FirstChild.Data)
	case url_login_POST:
		e, _ := hquery.FindFirst("h2[class*=\"error\"]+div[class*=\"msg\"]", n)
		if e == nil || e.FirstChild == nil {
			return fmt.Errorf("post %s: could not find error message", url_base+url_login_POST)
		}
		msg := strings.TrimSpace(e.FirstChild.Data)
		switch msg {
		case "Il semble que votre identifiant ou votre mot de passe n'est pas valide.":
			return bank.ErrBadLoginPwd
		default:
			return bank.UnrecognizedError(msg)
		}
	}

	// By default, all other urls are considered as successful
	b.loggedIn = true
	return nil
}

func (b *Boursorama) logout() error {
	if !b.loggedIn {
		return bank.ErrNotLoggedIn
	}
	_, _, err := b.getLoggedIn(url_base + url_accounts)
	if err != nil {
		return err
	}
	// The returned page is the same even if we are not logged in. That's why we
	// first perform a check in a random url.
	_, _, err = b.getLoggedIn(url_base + url_logout)
	if err != nil && !errors.Is(err, bank.ErrNotLoggedIn) {
		return err
	}
	return nil
}

// getLoggedIn issues a GET to the specified URL while making sure the user is
// connected.
func (b *Boursorama) getLoggedIn(url string) (r *http.Response, n *html.Node, err error) {
	r, n, err = hparser.Get(b.http, url)
	if err == nil && r.Request.URL.Path == url_login {
		b.loggedIn = false
		err = fmt.Errorf("get %s: %w", url, bank.ErrNotLoggedIn)
	}
	return
}

// matrixRandomChallenge returns the challenge code that must be given when
// connecting to Boursorama Banque.
func matrixRandomChallenge(n *html.Node) (challenge string, err error) {
	scripts, _ := hquery.FindAll("script", n)
	var found bool
	for _, s := range scripts {
		if s.FirstChild == nil {
			continue
		}
		data := s.FirstChild.Data
		idx1 := strings.Index(data, ".val(\"")
		if idx1 != -1 {
			idx2 := strings.Index(data[idx1:], "\")")
			if idx2 != -1 {
				challenge = data[idx1+6 : idx1+idx2]
				found = true
				break
			}
		}
	}
	if !found {
		err = fmt.Errorf("could not find challenge")
	}
	return
}

// vKeyboardCode returns the code corresponding to the specified password.
func vKeyboardCode(pwd []byte, n *html.Node) (code string, err error) {
	m, err := matchvKeyboardHashes(n)
	if err != nil {
		return
	}
	list := []string{}
	for _, r := range string(pwd) {
		c, ok := m[r]
		if !ok {
			err = bank.ErrBadPwd
			return
		}
		list = append(list, c)
	}
	code = strings.Join(list, "|")
	return
}

// matchPwdHashesCode returns the code corresponding to each rune by comparing the
// md5 hashes of each image. This method is called by getPwdCode.
func matchvKeyboardHashes(n *html.Node) (map[rune]string, error) {
	buttons, _ := hquery.FindAll("ul.password-input button", n)
	if len(buttons) == 0 {
		return nil, fmt.Errorf("could not find buttons")
	}
	// We extract the key and the img of each button.
	m := make(map[rune]string, len(vKeyboardHashes))
	for _, b := range buttons {
		key, ok := hquery.FindAttr("data-matrix-key", b)
		if !ok {
			return nil, fmt.Errorf("could not find attr[data-matrix-key] of button")
		}
		// For each button we search the src of the image
		img, _ := hquery.FindFirst("img", b)
		if img == nil {
			return nil, fmt.Errorf("could not find img")
		}
		src, _ := hquery.FindAttr("src", img)
		sum := md5.Sum([]byte(src))
		var found bool
		for i, h := range vKeyboardHashes {
			if h == sum {
				found = true
				m[i] = key
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find corresponding md5 sum, got %v for %s", sum, src)
		}
	}
	return m, nil
}
