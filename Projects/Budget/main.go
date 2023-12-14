package main

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	datafile := flag.String("data", "./transactions.gob", "Datafile with the transactions")
	action := flag.String("action", "resume", "Action to realize (resume, expense, income, balance, history, remove)")

	flag.Parse()

	file, err := os.Open(*datafile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data Data

	dec := gob.NewDecoder(file)
	dec.Decode(&data)

	file.Close()

	switch *action {
	case "resume":
		fmt.Println("Resume of transactions")

		lastBalance, money := data.Balance()

		fmt.Println("\nÚltimo balance:")
		fmt.Println(data.Transactions[lastBalance])

		fmt.Println("\nDinero Total:", money)

	case "expense":
		fmt.Println("Adding expense:")

		var quantiy uint64

		fmt.Print("Quantity: $")
		fmt.Scanf("%d\n", &quantiy)

		fmt.Print("Notes: ")
		r := bufio.NewReader(os.Stdin)
		notes, _ := r.ReadString('\n')
		data.AddTransaction(quantiy, Expense, time.Now(), notes)

		fmt.Println("\nAdded data:")
		fmt.Println(data.Transactions[len(data.Transactions)-1])

		file, err := os.Create(*datafile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		enc := gob.NewEncoder(file)
		enc.Encode(data)

		file.Close()

	case "income":
		fmt.Println("Adding income:")

		var quantiy uint64

		fmt.Print("Quantity: $")
		fmt.Scanf("%d\n", &quantiy)

		fmt.Print("Notes: ")
		r := bufio.NewReader(os.Stdin)
		notes, _ := r.ReadString('\n')

		data.AddTransaction(quantiy, Income, time.Now(), notes)

		fmt.Println("\nAdded data:")
		fmt.Println(data.Transactions[len(data.Transactions)-1])

		file, err := os.Create(*datafile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		enc := gob.NewEncoder(file)
		enc.Encode(data)

		file.Close()

	case "balance":
		fmt.Println("Adding balance:")

		var quantiy uint64

		fmt.Print("Quantity: $")
		fmt.Scanf("%d\n", &quantiy)

		fmt.Print("Notes: ")
		r := bufio.NewReader(os.Stdin)
		notes, _ := r.ReadString('\n')

		data.AddTransaction(quantiy, Balance, time.Now(), notes)

		fmt.Println("\nAdded data:")
		fmt.Println(data.Transactions[len(data.Transactions)-1])

		file, err := os.Create(*datafile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		enc := gob.NewEncoder(file)
		enc.Encode(data)

		file.Close()

	case "history":
		fmt.Println("History:")

		for i := 0; i < len(data.Transactions); i++ {
			fmt.Println("-------------------------------")
			fmt.Println(data.Transactions[i])
		}

	case "remove":
		fmt.Println("Removiendo última transacción:")
		fmt.Println()

		fmt.Println(data.Transactions[len(data.Transactions)-1])

		data.Transactions = data.Transactions[:len(data.Transactions)-1]

		file, err := os.Create(*datafile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		enc := gob.NewEncoder(file)
		enc.Encode(data)

		file.Close()

	default:
		fmt.Println("Invalid action.")
		os.Exit(1)
	}
}

type Data struct {
	Transactions []Transaction
}

func (d *Data) AddTransaction(quantity uint64, t Type, date time.Time, notes string) {
	d.Transactions = append(d.Transactions, Transaction{
		Quantity: quantity,
		Type:     t,
		Time:     date,
		Notes:    notes,
	})
}

func (d Data) Balance() (lastBalance int, avaible int) {
	sort.Sort(d)

	lastBalance = -1

	for i := 0; i < len(d.Transactions); i++ {
		if d.Transactions[i].Type == Balance {
			lastBalance = i
		}
	}

	if lastBalance == -1 {
		return -1, 0
	}

	money := d.Transactions[lastBalance].Quantity
	for i := lastBalance + 1; i < len(d.Transactions); i++ {
		q := d.Transactions[i].Quantity

		switch d.Transactions[i].Type {
		case Income:
			money += q
		case Expense:
			money -= q
		}
	}

	return lastBalance, int(money)
}

func (d Data) Len() int {
	return len(d.Transactions)
}

func (d Data) Less(i, j int) bool {
	if d.Transactions[i].Time.Before(d.Transactions[j].Time) {
		return true
	}

	return false
}

func (d Data) Swap(i, j int) {
	d.Transactions[i], d.Transactions[j] = d.Transactions[j], d.Transactions[i]
}

type Transaction struct {
	Quantity uint64
	Time     time.Time
	Type     Type
	Notes    string
}

func (t Transaction) String() string {
	return fmt.Sprintf(`Transacción:
 - Cantidad: %d
 - Tipo: %s
 - Fecha: %s
 - Notas:
 		%s
`, t.Quantity, t.Type, t.Time.Format(time.ANSIC), t.Notes)
}

type Type uint8

const (
	Balance Type = iota
	Expense
	Income
)

func (t Type) String() string {
	switch t {
	case Balance:
		return "Balance"
	case Expense:
		return "Expense"
	case Income:
		return "Income"
	default:
		return "Error"
	}
}
