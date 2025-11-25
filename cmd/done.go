package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/db"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [ID]",
	Short: "タスクの完了",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("無効なIDです: %s", idStr)
		}

		todo, err := db.CompleteTodo(id)
		if err != nil {
			return err
		}

		fmt.Printf("タスク \"%s\" (ID: %d) を完了しました。\n", todo.Task, id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
