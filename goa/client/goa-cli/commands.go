package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/nikitasmall/frameworks-comparsion/goa/app"
	"github.com/nikitasmall/frameworks-comparsion/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// CreateTasksCommand is the command line data structure for the create action of tasks
	CreateTasksCommand struct {
		Payload string
	}
	// IndexTasksCommand is the command line data structure for the index action of tasks
	IndexTasksCommand struct {
	}
)

// Run makes the HTTP request corresponding to the CreateTasksCommand command.
func (cmd *CreateTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/tasks"
	}
	var payload app.CreateTasksPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	resp, err := c.CreateTasks(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateTasksCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the IndexTasksCommand command.
func (cmd *IndexTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/tasks"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	resp, err := c.IndexTasks(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *IndexTasksCommand) RegisterFlags(cc *cobra.Command) {
}
