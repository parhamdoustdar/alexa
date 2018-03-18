package alexa

import "context"

// LaunchRequestHandler is called when the skill is launched, which
// means when the user says something like: "alexa, launch <skill>",
// "alexa, open <skill>" and so on.
type LaunchRequestHandler func(context.Context, *LaunchRequest) (*SpeechletResponse, error)

// SessionEndedRequestHandler is called when the request ends, either
// manually by the user, or when what the user says doesn't match the
// specified intents, or when there's an error.
type SessionEndedRequestHandler func(context.Context, *SessionEndedRequest) (*SpeechletResponse, error)

// IntentRequestHandler is the handler for a specific intent.
type IntentRequestHandler func(context.Context, *IntentRequest) (*SpeechletResponse, error)
