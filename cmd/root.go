/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"path"
	"strconv"

	"github.com/abdelkd/todo-cli/internal/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A Simple CLI to manage todo lists",
	Long: `
todo is a CLI application made for working with todos on the terminal
It has ability to store todo lists in plain text or SQLite database`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli.yaml)")
	rootCmd.PersistentFlags().Bool("sqlite", false, "Use SQLite instead of json file")

	useSqliteFlag := rootCmd.Flag("sqlite")
	useSqlite, err := strconv.ParseBool(useSqliteFlag.Value.String())
	if err != nil {
		useSqlite = false
	}

	var todoModel models.Model
	if useSqlite {
		return
	} else {
		todosFilePath := path.Join(os.Getenv("HOME"), ".todos.json")
		todoModel = models.FileModel{
			Path: todosFilePath,
		}
	}

	ctx := context.WithValue(context.Background(), "model", todoModel)
	rootCmd.SetContext(ctx)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
