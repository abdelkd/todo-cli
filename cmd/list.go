package cmd

import (
	"github.com/abdelkd/todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "View list of todos",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		todoModel := ctx.Value("model").(models.Model)

		if err := todoModel.ListItems(); err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
