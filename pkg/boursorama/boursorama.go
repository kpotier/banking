package boursorama

import (
	"context"
	"net/http"
	"time"

	"github.com/kpotier/banking/pkg/bank"
)

var url_base = "https://clients.boursorama.com"

const (
	url_dashboard = "/"

	url_login          = "/connexion/"
	url_login_keyboard = "/connexion/clavier-virtuel?_hinclude=1"
	url_login_POST     = "/connexion/saisie-mot-de-passe"

	url_logout = "/se-deconnecter"

	url_account_locked = "/connexion/compte-verrouille"

	url_accounts = "/dashboard/liste-comptes?rumroute=dashboard.new_accounts&_hinclude=1"

	url_transactions_checking = "/compte/cav/%s/mouvements"
	url_transactions_savings  = "/budget/compte/%s/mouvements"

	// url_transactions_next is meant to be appended to
	// url_transactions_checking or url_transactions_savings
	url_transactions_next = "?rumroute=accounts.bank.movements&continuationToken=%s"
)

// months in French used to decode a date.
var months = [...]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}

// vKeyboardHashes contains the md5 sum of each image that composes the password
// virtual keyboard. Each image represents a rune.
var vKeyboardHashes = map[rune][16]byte{
	'0': {56, 202, 243, 49, 55, 44, 7, 7, 105, 227, 112, 48, 215, 235, 152, 233},
	'1': {96, 30, 24, 60, 144, 205, 238, 84, 12, 105, 34, 215, 22, 145, 102, 232},
	'2': {242, 126, 243, 158, 207, 236, 87, 184, 223, 230, 143, 73, 95, 242, 208, 16},
	'3': {174, 104, 168, 176, 130, 224, 3, 148, 208, 6, 143, 92, 14, 53, 53, 240},
	'4': {38, 34, 113, 66, 192, 50, 25, 193, 23, 115, 217, 76, 209, 3, 211, 77},
	'5': {119, 69, 89, 92, 164, 123, 87, 79, 10, 40, 235, 65, 198, 255, 221, 43},
	'6': {246, 218, 173, 31, 123, 120, 26, 121, 201, 0, 56, 109, 129, 215, 125, 189},
	'7': {154, 208, 96, 223, 212, 51, 129, 66, 126, 124, 147, 226, 187, 67, 8, 27},
	'8': {184, 208, 254, 219, 218, 54, 178, 65, 195, 200, 148, 242, 51, 26, 135, 26},
	'9': {147, 225, 127, 66, 162, 186, 169, 66, 202, 116, 165, 28, 58, 137, 195, 107},
}

// Boursorama enables to connect to Boursorama Bank and to retrieve the list of
// accounts as well as the list of operations for each account. This structure
// can be instanced by "hand" (Username and Password must be given) or through
// the New method.
type Boursorama struct {
	http     *http.Client
	loggedIn bool
}

// New returns an instance of Boursorama with the specified credentials. It is
// meant to be used by the banks module as it returns the Bank interface
func New() bank.Bank {
	return &Boursorama{nil, false}
}

// Login connects to Boursorama Banque. Credentials have been given when
// instancing Boursorama. Questions like 2FA code may be asked through the
// prompt channel (this method will block until it gets an answer).
func (b *Boursorama) Login(username, pwd []byte,
	ctx context.Context, q chan<- string, a <-chan string) error {
	return b.login(username, pwd, ctx, q, a)
}

// Logout disconnects the user from Boursorama Banque. It returns an error if
// the user was already disconnected.
func (b *Boursorama) Logout() error {
	return b.logout()
}

// Accounts returns the list of accounts detained by the user.
func (b *Boursorama) Accounts() ([]*bank.Account, error) {
	return b.accounts()
}

// Transactions returns a list of transactions that have been performed for a
// specific account after a specified date.
func (b *Boursorama) Transactions(account *bank.Account, after time.Time) ([]*bank.Transaction, error) {
	return b.transactions(account, after)
}

func init() {
	// Register the bank instancer called "boursorama"
	bank.SetBank("boursorama", New)
}
