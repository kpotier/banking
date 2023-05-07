package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/kpotier/banking/pkg/bank"
	_ "github.com/kpotier/banking/pkg/boursorama"
	"golang.org/x/term"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	banks := bank.Banks()
	fmt.Println("Which bank do you want to connect?")
	fmt.Println("List of available banks:", strings.Join(banks, ", "))
	fmt.Print("> ")
	bStr, err := rd.ReadBytes('\n')
	if err != nil {
		fatal(err)
	}

	bNew, ok := bank.GetBank(string(bStr[:len(bStr)-1]))
	if !ok {
		fatal(err)
	}

	fmt.Print("Username: ")
	username, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		fatal(err)
	}

	fmt.Print("Password: ")
	password, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		fatal(err)
	}

	questions := make(chan string)
	defer close(questions)
	answers := make(chan string)
	go func() {
		for {
			q, ok := <-questions
			if !ok {
				break
			}
			fmt.Println(q)
			fmt.Print("Answer: ")
			answer, err := term.ReadPassword(syscall.Stdin)
			fmt.Println()
			if err != nil {
				fatal(err)
			}
			answers <- string(answer[:len(answer)-1])
		}
	}()

	b := bNew()
	err = b.Login(username, password, context.Background(), questions, answers)
	if err != nil {
		fatal(err)
	}

	acc, err := b.Accounts()
	if err != nil {
		fatal(err)
	}
	fmt.Println("Retrieve the transactions made in the last 30 days for the account number:")
	for i, a := range acc {
		// We can safely divide by 100 because currency is either EUR or GBP.
		fmt.Printf("%d: %s (%.2f %s)\n", i, a.Name, float32(a.Balance.Amount)/100, a.Balance.Code)
	}
	fmt.Print("> ")
	idStr, err := rd.ReadBytes('\n')
	if err != nil {
		fatal(err)
	}
	id, err := strconv.Atoi(string(idStr[:len(idStr)-1]))
	if err != nil {
		fatal(err)
	}
	if id < 0 || id >= len(acc) {
		fatal(fmt.Errorf("bad id"))
	}

	tr, err := b.Transactions(acc[id], time.Now().Truncate(30*24*time.Hour))
	if err != nil {
		fatal(err)
	}
	for _, t := range tr {
		fmt.Printf("%s (%s) -> %s (%s): %.2f %s\n", t.Name, t.Card, t.Category, t.DateDebit, float32(t.Value.Amount)/100, t.Value.Code)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
