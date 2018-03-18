package alexa

type SessionEndedRequest struct {
	*LaunchRequest
	Reason Reason
	Error  Error
}

type Reason string

const (
	ReasonUserInitiated        Reason = "USER_INITIATED"
	ReasonError                Reason = "ERROR"
	ReasonExceededMaxReprompts Reason = "EXCEEDED_MAX_REPROMPTS"
)

type Error struct {
	Type    ErrorType
	Message string
}

type ErrorType string

const (
	ErrorTypeInvalidResponse          ErrorType = "INVALID_RESPONSE"
	ErrorTypeDeviceCommunicationError ErrorType = "DEVICE_COMMUNICATION_ERROR"
	ErrorTypeInternalError            ErrorType = "INTERNAL_ERROR"
)

func NewSessionEndedRequestFromRequest(request Request) (*SessionEndedRequest, error) {
	launchRequest, err := newLaunchRequestFromRequest(request)
	if err != nil {
		return nil, err
	}

	requestMap := request.RequestMap

	error := Error{}
	if errorMapInterface, ok := requestMap["error"]; ok {
		errorMap := errorMapInterface.(map[string]interface{})

		error = Error{
			Type:    ErrorType(errorMap["type"].(string)),
			Message: errorMap["message"].(string),
		}
	}

	sessionEndedRequest := &SessionEndedRequest{
		LaunchRequest: launchRequest,
		Reason:        Reason(requestMap["reason"].(string)),
		Error:         error,
	}

	return sessionEndedRequest, nil
}
