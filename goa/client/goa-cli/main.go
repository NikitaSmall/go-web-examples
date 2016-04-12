package main

import (
	"fmt"
	"github.com/nikitasmall/frameworks-comparsion/goa/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "goa-cli",
		Short: `CLI client for the goa service`,
	}
	c := client.New(nil)
	c.UserAgent = "goa-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:3000", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `Creates new task`,
	}
	tmp1 := new(CreateTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks "/tasks"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `Deletes some task`,
	}
	tmp2 := new(DeleteTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks "/tasks/:id"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `Shows the task`,
	}
	tmp3 := new(GetTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks "/tasks/:id"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "index",
		Short: `Shows all tasks`,
	}
	tmp4 := new(IndexTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks "/tasks"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Updates some task`,
	}
	tmp5 := new(UpdateTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks "/tasks/:id"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)

}
