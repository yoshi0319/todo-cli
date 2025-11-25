package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/db"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:     "rm [ID]",
	Aliases: []string{"remove", "delete"},
	Short:   "タスクの削除",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		if err := db.DeleteTodo(id); err != nil {
			return err
		}

		fmt.Printf("ID %d のタスクを削除しました。\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
