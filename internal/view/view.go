package view

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/SalmandaAK/expense-tracker/internal/expense/domain"
)

func DisplayError(err error) {
	fmt.Printf("Error: %v\n", err)
}

func DisplayMessage(message string) {
	fmt.Printf("%v\n", message)
}

func DisplayExpenseList(expenses []*domain.Expense, currency string) {
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)

	fmt.Fprintf(tabwriter, "\tID\t\tDate\t\tDescription\t\tAmount\t\n")
	for _, e := range expenses {
		fmt.Fprintf(tabwriter, "\t%d\t\t%s\t\t%s\t\t%s %d\t\n", e.Id, e.CreatedAt.Format(time.DateOnly), e.Description, currency, e.Amount)
	}
	tabwriter.Flush()
}

func DisplaySummary(summary, month int, currency string) {
	if month == 0 {
		fmt.Printf("Total expenses: %v %v\n", currency, summary)
		return
	}
	fmt.Printf("Total expenses for %v: %v %v\n", time.Month(month).String(), currency, summary)
}
