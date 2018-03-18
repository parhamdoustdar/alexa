# alexa

This project is created to make it easier to create Alexa skills written in Go.

Very recently, Amazon announced [support for Go]() in AWS Lambda. However, up until this point, there was no way of easily creating an Alexa skill, while using the power of Go's static typing system.

For example, follow the tutorial for [creating an Alexa skill in 5 minutes](). Using this library, the code for the Lambda function driving your skill would be as simple as the snippet below:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/parhamdoustdar/alexa"
)

func main() {
	router := alexa.NewRouter()
	router.LaunchRequestHandler = func(context context.Context, request *alexa.LaunchRequest) (*alexa.SpeechletResponse, error) {
		response := alexa.NewSpeechletResponse()
		response.AskPlainText("Welcome to the color picker! I can do the extraordinary job of remembering your favorite color for you. You can tell me what your favorite color is, and then ask me what it is!")
		response.RepromptWithPlainText("I didn't catch that. You can either tell me what your favorite color is, or if you've already told me, ask me what it is.")

		return response, nil
	}

	router.SessionEndedRequestHandler = func(context context.Context, sessionEndedRequest *alexa.SessionEndedRequest) (*alexa.SpeechletResponse, error) {
		response := alexa.NewSpeechletResponse()

		if sessionEndedRequest.Error.Type != "" {
			log.Printf("Error type %s: %s", sessionEndedRequest.Error.Type, sessionEndedRequest.Error.Message)
		}

		return response, nil
	}

	router.AddIntentHandler("WhatsMyColorIntent", func(context context.Context, request *alexa.IntentRequest) (*alexa.SpeechletResponse, error) {
		response := alexa.NewSpeechletResponse()
		colorInterface, ok := request.Session.Attributes["color"]
		if ok {
			color := colorInterface.(string)
			response.TellSSML(fmt.Sprintf(`<speak>Your favorite color is, <amazon:effect name="whispered">%s</amazon:effect>! Got it? <emphasis level="strong">%s</emphasis>!</speak>`, color, color))
			response.ShowSimpleCard("Your Favorite Color Is...", color)
		} else {
			response.TellPlainText("You haven't told me your favorite color yet.")
		}

		return response, nil
	})

	router.AddIntentHandler("MyColorIsIntent", func(context context.Context, request *alexa.IntentRequest) (*alexa.SpeechletResponse, error) {
		response := alexa.NewSpeechletResponse()

		color := request.Intent.Slots["Color"].Value
		response.AskPlainText(fmt.Sprintf("Your favorite color is %s, got it. You can now ask me what your favorite color is.", color))
		response.SessionAttributes["color"] = color

		return response, nil
	})

	router.Start()
}

```

## Usage

What you see in the example above is three things:

- First, you have the `LaunchRequestHandler`, which is called when the user says something like `tell <skill name> <something>`, or `Open <skill name>`.
- Then, you have the `SessionEndedRequest`, which is called when the user's interaction with your skill ends. This usually happens when the user says `Cancel`, or when you call one of the `Tell*` functions on the `Response` object. An example of this is in the handler for the `WhatsMyColorIntent:
  ```go
  response.TellSSML(fmt.Sprintf(`<speak>Your favorite color is, <amazon:effect name="whispered">%s</amazon:effect>! Got it? <emphasis level="strong">%s</emphasis>!</speak>`, color, color))
  ```
- Finally, you have the intent handlers. This is where you can define a handler per intent, which is called when the user requests an intent. Intents are defined in the configuration of your skill, and Alexa translates what the user says into the intent name before calling your Lambda function with it.

## The Request Object

The request object encapsulates the request that Alexa sends to your lambda skill.

- [LaunchRequest]() is sent to the `LaunchRequestHandler`.
- [SessionEndedRequest]() is sent to `SessionEndedRequestHandler`.
- And finally, [IntentRequest]() is sent to the various number of `IntentRequestHandler`s you have added to the router.

## Response

The [SpeechletResponse]() is used by Alexa to understand how to react to what the user just said. You can provide text to say (using `AskSSML`, `TellSSML`, `AskPlainText`, and `TellPlainText`), you can tell it how to prompt the user when what he says doesn't match the intents you've defined (using `RepromptWithSSML` and `RepromptWithPlainText`), and show cards (with `ShowSimpleCard()`, `ShowStandardCard()`, and `ShowLinkAccountCard()`).

[support for Go]: https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/
[creating an Alexa skill in 5 minutes]: https://developer.amazon.com/alexa-skills-kit/alexa-skill-quick-start-tutorial
[LaunchRequest]: https://godoc.org/github.com/parhamdoustdar/alexa#LaunchRequest
[SessionEndedRequest]: https://godoc.org/github.com/parhamdoustdar/alexa#SessionEndedRequest
[IntentRequest]: https://godoc.org/github.com/parhamdoustdar/alexa#IntentRequest
[SpeechletResponse]: https://godoc.org/github.com/parhamdoustdar/alexa#SpeechletResponse
