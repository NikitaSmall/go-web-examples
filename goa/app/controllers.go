//************************************************************************//
// API "goa": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// TasksController is the controller interface for the Tasks actions.
type TasksController interface {
	goa.Muxer
	Create(*CreateTasksContext) error
	Index(*IndexTasksContext) error
}

// MountTasksController "mounts" a Tasks resource controller on the given service.
func MountTasksController(service *goa.Service, ctrl TasksController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateTasksContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateTasksPayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/tasks", ctrl.MuxHandler("Create", h, unmarshalCreateTasksPayload))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Create", "route", "POST /tasks")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewIndexTasksContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Index(rctx)
	}
	service.Mux.Handle("GET", "/tasks", ctrl.MuxHandler("Index", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Index", "route", "GET /tasks")
}

// unmarshalCreateTasksPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTasksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createTasksPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
