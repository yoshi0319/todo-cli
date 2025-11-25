package cmd

import (
	"fmt"
	"strings"
	"todo-cli/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [タスク名]",
	Short: "新しいタスクの追加",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName := strings.Join(args, " ")

		newTodo, err := db.AddTodo(taskName)
		if err != nil {
			return err
		}

		fmt.Printf("タスクを追加しました。: [%d] %s\n", newTodo.ID, newTodo.Task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
