package cmd

import (
	"strconv"

	"github.com/abdelkd/todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [todo-id]",
	Short: "Removes an item from todo list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		ctx := cmd.Context()
		todoModel := ctx.Value("model").(models.Model)
		err = todoModel.RemoveItem(id)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
