// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/divoc/registration-api/swagger_gen/models"
)

// NewRegistrationAPIAPI creates a new RegistrationAPI instance
func NewRegistrationAPIAPI(spec *loads.Document) *RegistrationAPIAPI {
	return &RegistrationAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		BookSlotOfFacilityHandler: BookSlotOfFacilityHandlerFunc(func(params BookSlotOfFacilityParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation BookSlotOfFacility has not yet been implemented")
		}),
		DeleteAppointmentHandler: DeleteAppointmentHandlerFunc(func(params DeleteAppointmentParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation DeleteAppointment has not yet been implemented")
		}),
		DeleteRecipientHandler: DeleteRecipientHandlerFunc(func(params DeleteRecipientParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation DeleteRecipient has not yet been implemented")
		}),
		EnrollRecipientHandler: EnrollRecipientHandlerFunc(func(params EnrollRecipientParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation EnrollRecipient has not yet been implemented")
		}),
		GenerateOTPHandler: GenerateOTPHandlerFunc(func(params GenerateOTPParams) middleware.Responder {
			return middleware.NotImplemented("operation GenerateOTP has not yet been implemented")
		}),
		GetRecipientsHandler: GetRecipientsHandlerFunc(func(params GetRecipientsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation GetRecipients has not yet been implemented")
		}),
		GetSlotsForFacilitiesHandler: GetSlotsForFacilitiesHandlerFunc(func(params GetSlotsForFacilitiesParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation GetSlotsForFacilities has not yet been implemented")
		}),
		InitializeFacilitySlotsHandler: InitializeFacilitySlotsHandlerFunc(func(params InitializeFacilitySlotsParams) middleware.Responder {
			return middleware.NotImplemented("operation InitializeFacilitySlots has not yet been implemented")
		}),
		VerifyOTPHandler: VerifyOTPHandlerFunc(func(params VerifyOTPParams) middleware.Responder {
			return middleware.NotImplemented("operation VerifyOTP has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*models.JWTClaimBody, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*RegistrationAPIAPI Registration API */
type RegistrationAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (*models.JWTClaimBody, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// BookSlotOfFacilityHandler sets the operation handler for the book slot of facility operation
	BookSlotOfFacilityHandler BookSlotOfFacilityHandler
	// DeleteAppointmentHandler sets the operation handler for the delete appointment operation
	DeleteAppointmentHandler DeleteAppointmentHandler
	// DeleteRecipientHandler sets the operation handler for the delete recipient operation
	DeleteRecipientHandler DeleteRecipientHandler
	// EnrollRecipientHandler sets the operation handler for the enroll recipient operation
	EnrollRecipientHandler EnrollRecipientHandler
	// GenerateOTPHandler sets the operation handler for the generate o t p operation
	GenerateOTPHandler GenerateOTPHandler
	// GetRecipientsHandler sets the operation handler for the get recipients operation
	GetRecipientsHandler GetRecipientsHandler
	// GetSlotsForFacilitiesHandler sets the operation handler for the get slots for facilities operation
	GetSlotsForFacilitiesHandler GetSlotsForFacilitiesHandler
	// InitializeFacilitySlotsHandler sets the operation handler for the initialize facility slots operation
	InitializeFacilitySlotsHandler InitializeFacilitySlotsHandler
	// VerifyOTPHandler sets the operation handler for the verify o t p operation
	VerifyOTPHandler VerifyOTPHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *RegistrationAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *RegistrationAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *RegistrationAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *RegistrationAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *RegistrationAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *RegistrationAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *RegistrationAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *RegistrationAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *RegistrationAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the RegistrationAPIAPI
func (o *RegistrationAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.BookSlotOfFacilityHandler == nil {
		unregistered = append(unregistered, "BookSlotOfFacilityHandler")
	}
	if o.DeleteAppointmentHandler == nil {
		unregistered = append(unregistered, "DeleteAppointmentHandler")
	}
	if o.DeleteRecipientHandler == nil {
		unregistered = append(unregistered, "DeleteRecipientHandler")
	}
	if o.EnrollRecipientHandler == nil {
		unregistered = append(unregistered, "EnrollRecipientHandler")
	}
	if o.GenerateOTPHandler == nil {
		unregistered = append(unregistered, "GenerateOTPHandler")
	}
	if o.GetRecipientsHandler == nil {
		unregistered = append(unregistered, "GetRecipientsHandler")
	}
	if o.GetSlotsForFacilitiesHandler == nil {
		unregistered = append(unregistered, "GetSlotsForFacilitiesHandler")
	}
	if o.InitializeFacilitySlotsHandler == nil {
		unregistered = append(unregistered, "InitializeFacilitySlotsHandler")
	}
	if o.VerifyOTPHandler == nil {
		unregistered = append(unregistered, "VerifyOTPHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *RegistrationAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *RegistrationAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *RegistrationAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *RegistrationAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *RegistrationAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *RegistrationAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the registration API API
func (o *RegistrationAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *RegistrationAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/appointment"] = NewBookSlotOfFacility(o.context, o.BookSlotOfFacilityHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/appointment"] = NewDeleteAppointment(o.context, o.DeleteAppointmentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/recipients"] = NewDeleteRecipient(o.context, o.DeleteRecipientHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/recipients"] = NewEnrollRecipient(o.context, o.EnrollRecipientHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/generateOTP"] = NewGenerateOTP(o.context, o.GenerateOTPHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/recipients"] = NewGetRecipients(o.context, o.GetRecipientsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/facility/slots"] = NewGetSlotsForFacilities(o.context, o.GetSlotsForFacilitiesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/facility/slots/init"] = NewInitializeFacilitySlots(o.context, o.InitializeFacilitySlotsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/verifyOTP"] = NewVerifyOTP(o.context, o.VerifyOTPHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *RegistrationAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *RegistrationAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *RegistrationAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *RegistrationAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *RegistrationAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
