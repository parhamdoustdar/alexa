package alexa

import (
	"fmt"

	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	RequestTypeLaunchRequest       = "LaunchRequest"
	RequestTypeSessionEndedRequest = "SessionEndedRequest"
	RequestTypeIntentRequest       = "IntentRequest"
)

type Router struct {
	LaunchRequestHandler       LaunchRequestHandler
	SessionEndedRequestHandler SessionEndedRequestHandler
	intentRequestHandlers      map[string]IntentRequestHandler
}

func NewRouter() *Router {
	return &Router{
		intentRequestHandlers: make(map[string]IntentRequestHandler),
	}
}

func (router *Router) Start() {
	lambda.Start(router.handle)
}

func (router *Router) handle(context context.Context, request Request) (interface{}, error) {
	requestMap := request.RequestMap

	switch requestMap["type"].(string) {
	case RequestTypeLaunchRequest:
		launchRequest, err := newLaunchRequestFromRequest(request)
		if err != nil {
			return nil, err
		}

		return router.LaunchRequestHandler(context, launchRequest)
	case RequestTypeSessionEndedRequest:
		sessionEndedRequest, err := NewSessionEndedRequestFromRequest(request)
		if err != nil {
			return nil, err
		}

		return router.SessionEndedRequestHandler(context, sessionEndedRequest)
	case RequestTypeIntentRequest:
		intentHandler, err := router.getIntentHandler(requestMap["intent"])
		if err != nil {
			return nil, err
		}

		intentRequest, err := newIntentRequestFromRequest(request)
		if err != nil {
			return nil, err
		}

		return intentHandler(context, intentRequest)
	default:
		return nil, fmt.Errorf("Unsupported request type: %s", requestMap["type"].(string))
	}
}

func (router *Router) AddIntentHandler(intentName string, intentHandler IntentRequestHandler) error {
	if _, ok := router.intentRequestHandlers[intentName]; ok {
		return fmt.Errorf("A handler is already registered to handle the %s intent.", intentName)
	}

	router.intentRequestHandlers[intentName] = intentHandler

	return nil
}

func (router *Router) getIntentHandler(intentMapInterface interface{}) (IntentRequestHandler, error) {
	intentMap := intentMapInterface.(map[string]interface{})

	name := intentMap["name"].(string)

	var handler IntentRequestHandler
	var ok bool
	if handler, ok = router.intentRequestHandlers[name]; !ok {
		return nil, fmt.Errorf("No intent handler found for %s", name)
	}

	return handler, nil
}
