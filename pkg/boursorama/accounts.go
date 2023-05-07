package boursorama

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/kpotier/banking/internal/hquery"

	"github.com/kpotier/banking/pkg/bank"
	"github.com/kpotier/banking/pkg/money"
)

func (b *Boursorama) accounts() ([]*bank.Account, error) {
	if !b.loggedIn {
		return nil, bank.ErrNotLoggedIn
	}
	_, n, err := b.getLoggedIn(url_base + url_accounts)
	if err != nil {
		return nil, err
	}
	var accounts []*bank.Account
	elems, _ := hquery.FindAll("a[class*=\"link-wrapper\"]", n)
	for i, e := range elems {
		var typ bank.AccountType
		tcc, _ := hquery.FindAttr("data-tag-commander-click", e)
		switch {
		case strings.Contains(tcc, "_cav"):
			typ = bank.AccountChecking
		case strings.Contains(tcc, "_saving"):
			typ = bank.AccountSavings
		case strings.Contains(tcc, "_investement"):
			typ = bank.AccountStocks
		default:
			return nil, fmt.Errorf("account #%d: no matching type for %s", i, tcc)
		}

		var id string
		href, ok := hquery.FindAttr("href", e)
		if !ok {
			return nil, fmt.Errorf("account #%d: could not find id", i)
		}
		if len(href) > 0 && href[len(href)-1] == '/' {
			href = href[:len(href)-1]
		}
		hrefSplit := strings.Split(href, "/")
		id = hrefSplit[len(hrefSplit)-1]

		var name string
		name, ok = hquery.FindAttr("title", e)
		if !ok {
			return nil, fmt.Errorf("account #%d: could not find name", i)
		}

		arialabel, ok := hquery.FindAttr("aria-label", e)
		if !ok {
			return nil, fmt.Errorf("account #%d: could not find balance", i)
		}
		substr := "Solde : "
		idx := strings.Index(arialabel, substr)
		if idx < 0 || len(arialabel) <= idx+len(substr) {
			return nil, fmt.Errorf("account #%d: invalid balance index", i)
		}
		balance, err := parseValue(arialabel[idx+len(substr):])
		if err != nil {
			return nil, fmt.Errorf("account #%d: error = %w, balance %s", i, err, arialabel)
		}

		accounts = append(accounts, &bank.Account{ID: id, Name: name, Type: typ, Balance: balance})
	}
	return accounts, nil
}

func (b *Boursorama) transactions(account *bank.Account, after time.Time) ([]*bank.Transaction, error) {
	if !b.loggedIn {
		return nil, bank.ErrNotLoggedIn
	}
	var url string
	switch account.Type {
	case bank.AccountChecking:
		url = url_transactions_checking
	case bank.AccountSavings:
		url = url_transactions_savings
	case bank.AccountStocks:
		return []*bank.Transaction{}, nil
	default:
		panic("unrecognized account type")
	}
	url = url_base + fmt.Sprintf(url, account.ID)
	var tr []*bank.Transaction
	return tr, b.getTransactions(&tr, url, "", after)
}

type trError struct {
	id  int
	err error
}

func (t trError) Error() string {
	return fmt.Sprintf("transaction #%d: %v", t.id, t.err)
}

func (b *Boursorama) getTransactions(tr *[]*bank.Transaction, url_base, token string, after time.Time) error {
	trID := len(*tr)

	url := url_base
	if token != "" {
		url += fmt.Sprintf(url_transactions_next, token)
	}
	_, n, err := b.getLoggedIn(url)
	if err != nil {
		return err
	}

	elems, _ := hquery.FindAll("ul.list__movement > li.list-operation-date-line, ul.list__movement > li.list-operation-item", n)
	if len(elems) == 0 || elems[0].FirstChild == nil {
		return trError{trID, fmt.Errorf("could not find first date")}
	}
	date, err := parseDate(elems[0].FirstChild.Data)
	if err != nil {
		return trError{trID, fmt.Errorf("error = %w, first date = %v", err, elems[0].FirstChild.Data)}
	} else if date.Before(after) {
		return nil
	}

	for _, e := range elems[1:] {
		c, _ := hquery.FindAttr("class", e)
		switch c {
		case "list-operation-date-line":
			if e.FirstChild == nil {
				return trError{trID, fmt.Errorf("could not find date")}
			}
			var err error
			date, err = parseDate(e.FirstChild.Data)
			if err != nil {
				return trError{trID, fmt.Errorf("error = %w, date = %v", err, e.FirstChild.Data)}
			} else if date.Before(after) {
				return nil
			}
		case "list-operation-item", "list-operation-item list-operation-item--splitted":
			isSplit, _ := hquery.FindFirst("li>ul.list__movement", e)
			if isSplit != nil {
				continue
			}
			id, ok := hquery.FindAttr("data-id", e)
			if !ok {
				return trError{trID, fmt.Errorf("could not find id")}
			}
			pending, _ := hquery.FindAttr("data-tag-commander-click", e)
			isPending := pending == "pending_authorizations"
			nameNode, _ := hquery.FindFirst("span.list__movement--label-user", e)
			if nameNode == nil || nameNode.FirstChild == nil {
				return trError{trID, fmt.Errorf("could not find name")}
			}
			rawName := nameNode.FirstChild.Data
			categoryNode, _ := hquery.FindFirst("span.list-operation-item__category", e)
			if categoryNode == nil || categoryNode.FirstChild == nil {
				return trError{trID, fmt.Errorf("could not find category")}
			}
			category := categoryNode.FirstChild.Data
			valueNode, _ := hquery.FindFirst("div.list-operation-item__amount", e)
			if valueNode == nil || valueNode.FirstChild == nil {
				return trError{trID, fmt.Errorf("could not find value")}
			}
			value, err := parseValue(valueNode.FirstChild.Data)
			if err != nil {
				return trError{trID, fmt.Errorf("error = %w, value = %v", err, valueNode.FirstChild.Data)}
			}

			// rawName analysis
			var (
				typ      bank.TransactionType
				dateDone time.Time
				card     string
				name     string
			)
			switch {
			case strings.HasPrefix(rawName, "CARTE"):
				typ = bank.TransactionCard
				dateDone, card, name, err = parseNameCard(rawName, "CARTE")
				if err != nil {
					return trError{trID, err}
				}
			case strings.HasPrefix(rawName, "AVOIR"):
				typ = bank.TransactionCredit
				dateDone, card, name, err = parseNameCard(rawName, "AVOIR")
				if err != nil {
					return trError{trID, err}
				}
			case strings.HasPrefix(rawName, "VIR SEPA"):
				typ = bank.TransactionTrsfSEPA
				name = rawName[len("VIR SEPA")+1:]
			case strings.HasPrefix(rawName, "VIR INST"):
				typ = bank.TransactionTrsfINST
				name = rawName[len("VIR INST")+1:]
			case strings.HasPrefix(rawName, "VIR"):
				typ = bank.TransactionTrsf
				name = rawName[len("VIR")+1:]
			case strings.HasPrefix(rawName, "PRLV SEPA"):
				typ = bank.TransactionDDebitSEPA
				name = rawName[len("PRLV SEPA")+1:]
			case isPending:
				idx := strings.Index(category, "CB*")
				if idx < 0 {
					return trError{trID, errors.New("could not find pending iss card of " + category)}
				}
				dateDone = date
				name = rawName
				card = category[idx:]
				category = ""
			default:
				name = rawName
				typ = bank.TransactionNone
			}

			*tr = append(*tr, &bank.Transaction{ID: id, Pending: isPending, DateDebit: date,
				DateDone: dateDone, Name: name, Card: card, Type: typ,
				RawName: rawName, Category: category, Value: value})
			trID++
		default:
			return trError{trID, fmt.Errorf("unable to match class %s", c)}
		}
	}

	// Look for other transactions
	link, _ := hquery.FindFirst("li.list__movement__range-summary", n)
	if link == nil {
		return trError{trID, fmt.Errorf("cannot find link to next page")}
	}
	token, ok := hquery.FindAttr("data-operations-next-pagination", link)
	if !ok {
		return nil // End of list
	}
	return b.getTransactions(tr, url_base, token, after)
}

func parseNameCard(raw string, typ string) (dateDone time.Time, card string, name string, err error) {
	if len(raw) < len(typ)+10 {
		err = fmt.Errorf("error = could not get card date, raw = %v", raw)
		return
	}
	dateStr := raw[len(typ)+1 : len(typ)+9]
	d, err := time.Parse("02/01/06", dateStr)
	if err != nil {
		err = fmt.Errorf("error = %v, rawDate = %v", err, dateStr)
		return
	}
	dateDone = d
	idx := strings.Index(raw, "CB*")
	if idx < 0 {
		err = fmt.Errorf("error = could not find iss card, rawName = %v", raw)
		return
	}
	card = raw[idx:]
	name = raw[len("CARTE")+10 : idx-1] // to remove extra space
	return
}

// parseValue returns money.Money number from s. It must match the following
// format: "12 345,67 €".
func parseValue(s string) (money.Money, error) {
	// This fn must check the presence of the euro sign and comma sign.
	var m money.Money
	m.Code = money.EUR
	var hasComma bool
	var hasEuro bool
	var sb strings.Builder
	sb.Grow(len(s))
	for _, r := range s {
		if r == '€' {
			hasEuro = true
		} else if r == ',' {
			hasComma = true
		} else if r == '\u2212' {
			sb.WriteRune('-')
		} else if !unicode.IsSpace(r) {
			sb.WriteRune(r)
		}
	}
	if !hasComma || !hasEuro {
		return m, fmt.Errorf("incorrect value")
	}
	v, err := strconv.ParseInt(sb.String(), 10, 64)
	m.Amount = v
	return m, err
}

func parseDate(s string) (t time.Time, err error) {
	fields := strings.Fields(s)
	if len(fields) != 4 {
		err = fmt.Errorf("malformed date %s", s)
		return
	}
	day, _ := strconv.Atoi(fields[1])
	year, _ := strconv.Atoi(fields[3])
	month := -1
	for i, m := range months {
		if m == fields[2] {
			month = i
			break
		}
	}
	if month < 0 {
		err = fmt.Errorf("could not find month %s", fields[1])
		return
	}
	t = time.Date(year, time.Month(month+1), day, 0, 0, 0, 0, time.UTC)
	return
}
