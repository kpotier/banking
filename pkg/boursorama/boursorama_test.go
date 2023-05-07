package boursorama

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/kpotier/banking/internal/hparser"
	"github.com/kpotier/banking/pkg/bank"
	"github.com/kpotier/banking/pkg/money"
)

func TestNew(t *testing.T) {
	bank := New()
	_, ok := bank.(*Boursorama)
	if !ok {
		t.Fatal("returned bank is not a pointer of Boursorama")
	}
}

func newAuthServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}
		switch url {
		case url_login:
			w.Write([]byte(`
				<form>
					<input type="hidden" name="form[_token]" value="somerandomtoken" />
					<input type="hidden" name="form[ajx]" value="1" />
					<input type="text"   name="form[clientNumber]" value="" />
					<input type="hidden" name="form[fakePassword]" value="" />
					<input type="hidden" name="form[matrixRandomChallenge]" value="" />
					<input type="hidden" name="form[passwordAck]" value="{&quot;js&quot;:false}" />
					<input type="hidden" name="form[password]" value="" />
					<input type="hidden" name="form[platformAuthenticatorAvailable]" value="" />
				</form>
			`))
		case url_login_keyboard:
			w.Write([]byte(authHTML + `
				<script>
            		$(function () {
                		$("[data-matrix-random-challenge]").val("somerandomchallenge")
            		})
        		</script>
			`))
		case url_login_POST:
			fn(w, r)
		case url_account_locked:
			w.Write([]byte(`<h2 class="narrow-modal__title u-text-center narrow-modal__title--error">Compte vérouillé</h2>`))
		case url_dashboard:
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func TestBoursorama_Login(t *testing.T) {
	const username = "123"
	const password = "456"
	const passwordAuth = "FOUR|FIVE|SIX"

	t.Run("ok login", func(t *testing.T) {
		ts := newAuthServer(func(w http.ResponseWriter, r *http.Request) {
			checks := map[string]string{
				"form[_token]":                         "somerandomtoken",
				"form[ajx]":                            "1",
				"form[clientNumber]":                   username,
				"form[fakePassword]":                   "",
				"form[matrixRandomChallenge]":          "somerandomchallenge",
				"form[passwordAck]":                    "{\"js\":false}",
				"form[password]":                       passwordAuth,
				"form[platformAuthenticatorAvailable]": "",
			}
			for k, v := range checks {
				if r.FormValue(k) != v {
					t.Fatalf("%s = %s, want %s", k, r.FormValue(k), v)
				}
			}
			http.Redirect(w, r, url_dashboard, http.StatusFound)
		})
		defer ts.Close()
		url_base = ts.URL
		b := &Boursorama{}
		err := b.Login([]byte(username), []byte(password), context.Background(), nil, nil)
		if err != nil {
			t.Errorf("b.Login() err = %v, wantErr %v", err, false)
		}
	})

	t.Run("bad pwd from server", func(t *testing.T) {
		ts := newAuthServer(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
				<h2 class="narrow-modal__title u-text-center narrow-modal__title--error">
					Erreur
				</h2>
				<div class="narrow-modal-window__msg">
					Il semble que votre identifiant ou votre mot de passe n&#039;est pas valide.
				</div>
			`))
		})
		defer ts.Close()
		url_base = ts.URL
		b := &Boursorama{}
		err := b.Login([]byte{}, []byte{}, context.Background(), nil, nil)
		if err == nil {
			t.Errorf("b.Login() err = %v, wantErr %v", nil, true)
		} else if !errors.Is(err, bank.ErrBadLoginPwd) {
			t.Errorf("b.Login() err = %v, want %v", err, bank.ErrBadLoginPwd)
		}
	})

	t.Run("unrecognized error", func(t *testing.T) {
		ts := newAuthServer(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
				<h2 class="narrow-modal__title u-text-center narrow-modal__title--error">
					Erreur
				</h2>
				<div class="narrow-modal-window__msg">
					Une autre erreur...
				</div>
			`))
		})
		defer ts.Close()
		url_base = ts.URL
		b := &Boursorama{}
		err := b.Login([]byte{}, []byte{}, context.Background(), nil, nil)
		var unrecognizedErr bank.UnrecognizedError
		if err == nil {
			t.Errorf("b.Login() err = %v, wantErr %v", nil, true)
		} else if !errors.As(err, &unrecognizedErr) {
			t.Errorf("b.Login() err = %v, want banks.UnrecognizedError", err)
		} else if unrecognizedErr.Error() != "Une autre erreur..." {
			t.Errorf("b.Login() err = %v, want %v", err, "Une autre erreur...")
		}
	})

	t.Run("locked account", func(t *testing.T) {
		ts := newAuthServer(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url_account_locked, http.StatusFound)
		})
		defer ts.Close()
		url_base = ts.URL
		b := &Boursorama{}
		err := b.Login([]byte{}, []byte{}, context.Background(), nil, nil)
		var unrecognizedErr bank.UnrecognizedError
		if err == nil {
			t.Errorf("b.Login() err = %v, wantErr %v", nil, true)
		} else if !errors.As(err, &unrecognizedErr) {
			t.Errorf("b.Login() err = %v, want banks.UnrecognizedError", err)
		} else if unrecognizedErr.Error() != "Compte vérouillé" {
			t.Errorf("b.Login() err = %v, want %v", err, "Compte vérouillé")
		}
	})
}

func TestBoursorama_Logout(t *testing.T) {
	t.Run("not logged in", func(t *testing.T) {
		b := &Boursorama{}
		err := b.Logout()
		if err == nil {
			t.Errorf("b.Logout() err = %v, wantErr %v", nil, true)
		} else if !errors.Is(err, bank.ErrNotLoggedIn) {
			t.Errorf("b.Logout() err = %v, want %v", err, bank.ErrNotLoggedIn)
		}
	})

	t.Run("logged in but expired session", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case url_login:
			default:
				http.Redirect(w, r, url_login, http.StatusFound)
			}
		}))
		defer ts.Close()
		b := &Boursorama{loggedIn: true, http: hparser.Default()}
		url_base = ts.URL
		err := b.Logout()
		if err == nil {
			t.Errorf("b.Logout() err = %v, wantErr %v", nil, true)
		} else if !errors.Is(err, bank.ErrNotLoggedIn) {
			t.Errorf("b.Logout() err = %v, want %v", err, bank.ErrNotLoggedIn)
		}
	})

	t.Run("logged in", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case url_logout:
				http.Redirect(w, r, url_login, http.StatusFound)
			default:
			}
		}))
		defer ts.Close()
		b := &Boursorama{loggedIn: true, http: hparser.Default()}
		url_base = ts.URL
		err := b.Logout()
		if err != nil {
			t.Errorf("b.Logout() err = %v, wantErr %v", err, false)
		}
	})
}

func TestBoursorama_Accounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
				<a class="c-info-box__link-wrapper" href="/compte/epargne/livret-a/idLivretA/" data-tag-commander-click='{"label": "application::customer.dashboard::click_accounts_saving", "s2": 1, "type": "N"}' aria-label="Détails du compte LIVRET A - Solde : 9 019,22 €" title="LIVRET A"></a>
				<a class="c-info-box__link-wrapper" href="/compte/cav/idBQ/" data-tag-commander-click='{"label": "application::customer.dashboard::click_accounts_cav", "s2": 1, "type": "N"}' aria-label="Détails du compte BOURSORAMA BANQUE - Solde : 100,11 €" title="BOURSORAMA BANQUE"></a>
				<a class="c-info-box__link-wrapper" href="/budget/compte/idlivret/" data-tag-commander-click='{"label": "application::customer.dashboard::click_accounts_pfm_saving", "s2": 1, "type": "N"}' aria-label="Détails du compte Livret - Solde : 9 222,01 €" title="Livret"></a>
				<a class="c-info-box__link-wrapper" href="/compte/pea/idpea/" data-tag-commander-click='{"label": "application::customer.dashboard::click_accounts_investement", "s2": 1, "type": "N"}' aria-label="Détails du compte PEA - Solde : 1 090,00 €" title="PEA"></a>
			`))
	}))
	defer ts.Close()
	b := &Boursorama{loggedIn: true, http: hparser.Default()}
	url_base = ts.URL
	a, err := b.Accounts()
	if err != nil {
		t.Fatalf("b.Accounts() err = %v, wantErr %v", err, false)
	}
	expected := []*bank.Account{
		{ID: "idLivretA", Name: "LIVRET A", Type: bank.AccountSavings, Balance: money.Money{Code: money.EUR, Amount: 901922}},
		{ID: "idBQ", Name: "BOURSORAMA BANQUE", Type: bank.AccountChecking, Balance: money.Money{Code: money.EUR, Amount: 10011}},
		{ID: "idlivret", Name: "Livret", Type: bank.AccountSavings, Balance: money.Money{Code: money.EUR, Amount: 922201}},
		{ID: "idpea", Name: "PEA", Type: bank.AccountStocks, Balance: money.Money{Code: money.EUR, Amount: 109000}},
	}
	if !reflect.DeepEqual(expected, a) {
		t.Errorf("b.Accounts() = %v, want %v", a, expected)
	}
}

func TestBoursorama_Transactions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}
		switch url {
		case fmt.Sprintf(url_transactions_checking+url_transactions_next, "account1", "next-page-1"):
			w.Write([]byte(transactionsSecondHTML))
		case fmt.Sprintf(url_transactions_checking, "account1"):
			w.Write([]byte(transactionsFirstHTML))
		}
	}))
	defer ts.Close()
	b := &Boursorama{loggedIn: true, http: hparser.Default()}
	url_base = ts.URL
	tr, err := b.Transactions(&bank.Account{Type: bank.AccountChecking, ID: "account1"}, time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("b.Transactions() err = %v, wantErr %v", err, false)
	}
	expected := []*bank.Transaction{
		{ID: "id1", Pending: false, DateDebit: time.Date(2022, time.July, 28, 0, 0, 0, 0, time.UTC), DateDone: time.Time{}, RawName: "VIR SEPA MR OU MME CHOU", Name: "MR OU MME CHOU", Type: bank.TransactionTrsfSEPA, Category: "Non catégorisé", Value: money.Money{Code: money.EUR, Amount: 123029}},
		{ID: "id2-1", Pending: false, DateDebit: time.Date(2022, time.July, 28, 0, 0, 0, 0, time.UTC), DateDone: time.Date(2021, time.November, 21, 0, 0, 0, 0, time.UTC), Name: "MR POU", RawName: "CARTE 21/11/21 MR POU CB*2100", Card: "CB*2100", Type: bank.TransactionCard, Category: "Carburant", Value: money.Money{Code: money.EUR, Amount: 4999}},
		{ID: "id2-2", Pending: false, DateDebit: time.Date(2022, time.July, 28, 0, 0, 0, 0, time.UTC), RawName: "Splitted", Name: "Splitted", Category: "Auto", Value: money.Money{Code: money.EUR, Amount: 1000}},
		{ID: "id3", Pending: false, DateDebit: time.Date(2022, time.January, 22, 0, 0, 0, 0, time.UTC), DateDone: time.Time{}, RawName: "PRLV SEPA MR OU MME CHOUPETTE", Name: "MR OU MME CHOUPETTE", Type: bank.TransactionDDebitSEPA, Category: "Catégorie 2", Value: money.Money{Code: money.EUR, Amount: -2049}},
		{ID: "id4", Pending: true, DateDebit: time.Date(2021, time.August, 21, 0, 0, 0, 0, time.UTC), DateDone: time.Date(2021, time.August, 21, 0, 0, 0, 0, time.UTC), RawName: "POULET", Name: "POULET", Card: "CB*2100", Type: bank.TransactionNone, Value: money.Money{Code: money.EUR, Amount: -1916}},
	}
	if !reflect.DeepEqual(expected, tr) {
		t.Errorf("b.Transactions() = %v, want %v", tr, expected)
	}

	tr, err = b.Transactions(&bank.Account{Type: bank.AccountChecking, ID: "account1"}, time.Time{})
	if err != nil {
		t.Fatalf("b.Transactions() err = %v, wantErr %v", err, false)
	}
	expected = append(expected, &bank.Transaction{ID: "id5", Pending: false, DateDebit: time.Date(2021, time.January, 22, 0, 0, 0, 0, time.UTC), RawName: "VIR INST MR OU MME CHOUPETTEL", Name: "MR OU MME CHOUPETTEL", Type: bank.TransactionTrsfINST, Category: "Catégorie 5", Value: money.Money{Code: money.EUR, Amount: -1011}})
	expected = append(expected, &bank.Transaction{ID: "id9", Pending: false, DateDebit: time.Date(2021, time.January, 22, 0, 0, 0, 0, time.UTC), RawName: "VIR Corner - Amazoune", Name: "Corner - Amazoune", Type: bank.TransactionTrsf, Category: "Catégorie 1", Value: money.Money{Code: money.EUR, Amount: -10141}})
	expected = append(expected, &bank.Transaction{ID: "id2", Pending: false, DateDebit: time.Date(2021, time.January, 22, 0, 0, 0, 0, time.UTC), DateDone: time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC), RawName: "AVOIR 01/02/19 BOUBI BOUBI CB*29", Name: "BOUBI BOUBI", Card: "CB*29", Type: bank.TransactionCredit, Category: "Catégorie 2", Value: money.Money{Code: money.EUR, Amount: -921}})
	if !reflect.DeepEqual(expected, tr) {
		t.Errorf("b.Transactions() = %v, want %v", tr, expected)
	}
}
