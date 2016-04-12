//************************************************************************//
// API "goa": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nikitasmall/frameworks-comparsion/goa
// --design=github.com/nikitasmall/frameworks-comparsion/goa/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateTasksContext provides the tasks create action context.
type CreateTasksContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *CreateTasksPayload
}

// NewCreateTasksContext parses the incoming request URL and body, performs validations and creates the
// context used by the tasks controller create action.
func NewCreateTasksContext(ctx context.Context, service *goa.Service) (*CreateTasksContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateTasksContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// createTasksPayload is the tasks create action payload.
type createTasksPayload struct {
	// task text
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

// Publicize creates CreateTasksPayload from createTasksPayload
func (payload *createTasksPayload) Publicize() *CreateTasksPayload {
	var pub CreateTasksPayload
	if payload.Text != nil {
		pub.Text = payload.Text
	}
	return &pub
}

// CreateTasksPayload is the tasks create action payload.
type CreateTasksPayload struct {
	// task text
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

// DeleteTasksContext provides the tasks delete action context.
type DeleteTasksContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	ID      int
}

// NewDeleteTasksContext parses the incoming request URL and body, performs validations and creates the
// context used by the tasks controller delete action.
func NewDeleteTasksContext(ctx context.Context, service *goa.Service) (*DeleteTasksContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteTasksContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// GetTasksContext provides the tasks get action context.
type GetTasksContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	ID      int
}

// NewGetTasksContext parses the incoming request URL and body, performs validations and creates the
// context used by the tasks controller get action.
func NewGetTasksContext(ctx context.Context, service *goa.Service) (*GetTasksContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := GetTasksContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetTasksContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// IndexTasksContext provides the tasks index action context.
type IndexTasksContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewIndexTasksContext parses the incoming request URL and body, performs validations and creates the
// context used by the tasks controller index action.
func NewIndexTasksContext(ctx context.Context, service *goa.Service) (*IndexTasksContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := IndexTasksContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *IndexTasksContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// UpdateTasksContext provides the tasks update action context.
type UpdateTasksContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	ID      int
	Payload *UpdateTasksPayload
}

// NewUpdateTasksContext parses the incoming request URL and body, performs validations and creates the
// context used by the tasks controller update action.
func NewUpdateTasksContext(ctx context.Context, service *goa.Service) (*UpdateTasksContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateTasksContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// updateTasksPayload is the tasks update action payload.
type updateTasksPayload struct {
	// task completion
	Complete *bool `json:"complete,omitempty" xml:"complete,omitempty"`
	// task text
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

// Publicize creates UpdateTasksPayload from updateTasksPayload
func (payload *updateTasksPayload) Publicize() *UpdateTasksPayload {
	var pub UpdateTasksPayload
	if payload.Complete != nil {
		pub.Complete = payload.Complete
	}
	if payload.Text != nil {
		pub.Text = payload.Text
	}
	return &pub
}

// UpdateTasksPayload is the tasks update action payload.
type UpdateTasksPayload struct {
	// task completion
	Complete *bool `json:"complete,omitempty" xml:"complete,omitempty"`
	// task text
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}
