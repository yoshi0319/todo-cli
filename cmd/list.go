package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todo-cli/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "タスク一覧の表示",
	RunE: func(cmd *cobra.Command, args []string) error {
		todos, err := db.GetTodos()
		if err != nil {
			return err
		}

		if len(todos) == 0 {
			fmt.Println("タスクはありません。")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tStatus\tTask\tCreated At")
		fmt.Fprintln(w, "--\t------\t----\t----------")

		for _, todo := range todos {
			status := "[ ]"
			if todo.Done {
				status = "[x]"
			}

			dateStr := todo.CreatedAt.Format("2006-01-02 15:04")
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", todo.ID, status, todo.Task, dateStr)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
