package internal

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ListExpenses() error {
	result, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}
	if result.FileCreated {
		fmt.Println("File was just created, no past history")
		return nil
	} else if result.FileEmpty {
		fmt.Println("No past history available")
		return nil
	}

	expenses := result.Expenses

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "ID\tDescription\tAmount\tCreated At")

	for _, e := range expenses {
		fmt.Fprintf(w, "%d\t%s\t$%.2f\t%s\n",
			e.ID,
			e.Description,
			e.Amount,
			e.CreatedAt.Format(time.DateOnly),
		)
	}

	w.Flush()

	return nil
}

func AddExpenses(desc string, amount float64) error {
	result, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	expenses := result.Expenses
	id := expenses[len(expenses)-1].ID

	newExpense := Expense{ID: id + 1, Description: desc, Amount: amount, CreatedAt: time.Now()}

	expenses = append(expenses, newExpense)

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	return nil
}
