package cmd

import (
	"fmt"

	"github.com/abdelkd/todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add [todo-name]",
	Short: "Add a new todo to the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		var todoModel models.Model
		ctx := cmd.Context()
		todoModel = ctx.Value("model").(models.Model)

		err := todoModel.AddItem(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Println("Todo has been added to the list")
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
