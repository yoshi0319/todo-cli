package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "簡単なCLI-ToDoアプリ",
	Long:  "コマンドライン入力によりToDoリストを管理できるアプリケーション",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
