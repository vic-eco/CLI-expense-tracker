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

	var id int
	if len(expenses) == 0 {
		id = 0
	} else {
		id = expenses[len(expenses)-1].ID
	}
	
	newExpense := Expense{ID: id + 1, Description: desc, Amount: amount, CreatedAt: time.Now()}

	expenses = append(expenses, newExpense)

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	return nil
}

func DeleteExpense(id int) error {
	result, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if result.FileCreated || result.FileEmpty {
		fmt.Println("Expenses file is empty")
		return nil
	}

	expenses := result.Expenses
	var found bool
	for i, e := range expenses {
		if e.ID == id {
			expenses = deleteAtIndex(expenses, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("There is no expense with id:", id)
		return nil
	}

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	return nil

}

func deleteAtIndex(slice []Expense, index int) []Expense {
	return append(slice[:index], slice[index+1:]...)
}
