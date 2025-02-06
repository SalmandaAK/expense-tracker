package view

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/SalmandaAK/expense-tracker/internal/expense/domain"
)

func DisplayError(err error) {
	fmt.Printf("Error: %v", err)
}

func DisplayMessage(message string) {
	fmt.Printf("%v", message)
}

func DisplayExpenseList(expenses []*domain.Expense) {
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)

	fmt.Fprintf(tabwriter, "\tID\t\tDate\t\tDescription\t\tAmount\t\n")
	for _, e := range expenses {
		fmt.Fprintf(tabwriter, "\t%d\t\t%s\t\t%s\t\t%d\t\n", e.Id, e.CreatedAt.Format(time.DateOnly), e.Description, e.Amount)
	}
	tabwriter.Flush()
}

func DisplaySummary(summary int) {
	fmt.Printf("Total expenses: %v", summary)
}
