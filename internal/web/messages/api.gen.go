package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Message defines model for Message.
type Message struct {
	Id *int64 `json:"id,omitempty"`

	// Message The content of the message
	Message *string `json:"message,omitempty"`
}

// PostMessagesJSONRequestBody defines body for PostMessages for application/json ContentType.
type PostMessagesJSONRequestBody = Message

// PatchMessagesIdJSONRequestBody defines body for PatchMessagesId for application/json ContentType.
type PatchMessagesIdJSONRequestBody = Message

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all messages
	// (GET /messages)
	GetMessages(ctx echo.Context) error
	// Create a new message
	// (POST /messages)
	PostMessages(ctx echo.Context) error
	// Delete a message by ID
	// (DELETE /messages/{id})
	DeleteMessagesId(ctx echo.Context, id int) error
	// Get a message by ID
	// (GET /messages/{id})
	GetMessagesId(ctx echo.Context, id int) error
	// Update a message by ID
	// (PATCH /messages/{id})
	PatchMessagesId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {

	err := w.Handler.GetMessages(ctx)
	return err
}

// PostMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PostMessages(ctx echo.Context) error {

	err := w.Handler.PostMessages(ctx)
	return err
}

// DeleteMessagesId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMessagesId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteMessagesId(ctx, id)
	return err
}

// GetMessagesId converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessagesId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMessagesId(ctx, id)
	return err
}

// PatchMessagesId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchMessagesId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchMessagesId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.POST(baseURL+"/messages", wrapper.PostMessages)
	router.DELETE(baseURL+"/messages/:id", wrapper.DeleteMessagesId)
	router.GET(baseURL+"/messages/:id", wrapper.GetMessagesId)
	router.PATCH(baseURL+"/messages/:id", wrapper.PatchMessagesId)

}

type GetMessagesRequestObject struct {
}

type GetMessagesResponseObject interface {
	VisitGetMessagesResponse(w http.ResponseWriter) error
}

type GetMessages200JSONResponse []Message

func (response GetMessages200JSONResponse) VisitGetMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostMessagesRequestObject struct {
	Body *PostMessagesJSONRequestBody
}

type PostMessagesResponseObject interface {
	VisitPostMessagesResponse(w http.ResponseWriter) error
}

type PostMessages201JSONResponse Message

func (response PostMessages201JSONResponse) VisitPostMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteMessagesIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteMessagesIdResponseObject interface {
	VisitDeleteMessagesIdResponse(w http.ResponseWriter) error
}

type DeleteMessagesId204Response struct {
}

func (response DeleteMessagesId204Response) VisitDeleteMessagesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteMessagesId404Response struct {
}

func (response DeleteMessagesId404Response) VisitDeleteMessagesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetMessagesIdRequestObject struct {
	Id int `json:"id"`
}

type GetMessagesIdResponseObject interface {
	VisitGetMessagesIdResponse(w http.ResponseWriter) error
}

type GetMessagesId200JSONResponse Message

func (response GetMessagesId200JSONResponse) VisitGetMessagesIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetMessagesId404Response struct {
}

func (response GetMessagesId404Response) VisitGetMessagesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchMessagesIdRequestObject struct {
	Id   int `json:"id"`
	Body *PatchMessagesIdJSONRequestBody
}

type PatchMessagesIdResponseObject interface {
	VisitPatchMessagesIdResponse(w http.ResponseWriter) error
}

type PatchMessagesId200JSONResponse Message

func (response PatchMessagesId200JSONResponse) VisitPatchMessagesIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchMessagesId404Response struct {
}

func (response PatchMessagesId404Response) VisitPatchMessagesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	GetMessages(ctx context.Context, request GetMessagesRequestObject) (GetMessagesResponseObject, error)

	PostMessages(ctx context.Context, request PostMessagesRequestObject) (PostMessagesResponseObject, error)

	DeleteMessagesId(ctx context.Context, request DeleteMessagesIdRequestObject) (DeleteMessagesIdResponseObject, error)

	GetMessagesId(ctx context.Context, request GetMessagesIdRequestObject) (GetMessagesIdResponseObject, error)

	PatchMessagesId(ctx context.Context, request PatchMessagesIdRequestObject) (PatchMessagesIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetMessages operation middleware
func (sh *strictHandler) GetMessages(ctx echo.Context) error {
	var request GetMessagesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMessages(ctx.Request().Context(), request.(GetMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMessagesResponseObject); ok {
		return validResponse.VisitGetMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostMessages operation middleware
func (sh *strictHandler) PostMessages(ctx echo.Context) error {
	var request PostMessagesRequestObject

	var body PostMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostMessages(ctx.Request().Context(), request.(PostMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostMessagesResponseObject); ok {
		return validResponse.VisitPostMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteMessagesId operation middleware
func (sh *strictHandler) DeleteMessagesId(ctx echo.Context, id int) error {
	var request DeleteMessagesIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteMessagesId(ctx.Request().Context(), request.(DeleteMessagesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteMessagesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteMessagesIdResponseObject); ok {
		return validResponse.VisitDeleteMessagesIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetMessagesId operation middleware
func (sh *strictHandler) GetMessagesId(ctx echo.Context, id int) error {
	var request GetMessagesIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMessagesId(ctx.Request().Context(), request.(GetMessagesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMessagesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMessagesIdResponseObject); ok {
		return validResponse.VisitGetMessagesIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchMessagesId operation middleware
func (sh *strictHandler) PatchMessagesId(ctx echo.Context, id int) error {
	var request PatchMessagesIdRequestObject

	request.Id = id

	var body PatchMessagesIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchMessagesId(ctx.Request().Context(), request.(PatchMessagesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchMessagesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchMessagesIdResponseObject); ok {
		return validResponse.VisitPatchMessagesIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
