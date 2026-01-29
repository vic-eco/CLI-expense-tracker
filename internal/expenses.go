package internal

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Database struct {
	Expenses []Expense `json:"expenses"`
}
type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ListExpenses() error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "ID\tDescription\tAmount\tCreated At")

	for _, e := range expenses {
		fmt.Fprintf(w, "%d\t%s\t$%.2f\t%s\n",
			e.ID,
			e.Description,
			e.Amount,
			e.CreatedAt.Format("2006-01-02"),
		)
	}

	w.Flush()

	return nil
}
