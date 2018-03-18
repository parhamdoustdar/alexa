package alexa

import "encoding/json"

type SpeechletResponse struct {
	Version           string
	SessionAttributes SessionAttributes
	response          map[string]interface{}
}

func NewSpeechletResponse() *SpeechletResponse {
	speechletResponse := &SpeechletResponse{
		Version:           "1.0",
		SessionAttributes: SessionAttributes{},
		response:          map[string]interface{}{},
	}

	speechletResponse.response["shouldEndSession"] = false

	return speechletResponse
}

func (speechletResponse *SpeechletResponse) AskPlainText(text string) {
	speechletResponse.response["outputSpeech"] = speechletResponse.createPlainTextOutputSpeech(text)
}

func (speechletResponse *SpeechletResponse) TellPlainText(text string) {
	speechletResponse.response["outputSpeech"] = speechletResponse.createPlainTextOutputSpeech(text)
	speechletResponse.response["shouldEndSession"] = true
}

func (speechletResponse *SpeechletResponse) AskSSML(text string) {
	speechletResponse.response["outputSpeech"] = speechletResponse.createSSMLOutputSpeech(text)
}

func (speechletResponse *SpeechletResponse) TellSSML(text string) {
	speechletResponse.response["outputSpeech"] = speechletResponse.createSSMLOutputSpeech(text)
	speechletResponse.response["shouldEndSession"] = true
}

func (speechletResponse *SpeechletResponse) RepromptWithPlainText(text string) {
	speechletResponse.response["reprompt"] = map[string]map[string]string{
		"outputSpeech": speechletResponse.createPlainTextOutputSpeech(text),
	}
}

func (speechletResponse *SpeechletResponse) RepromptWithSSML(text string) {
	speechletResponse.response["reprompt"] = map[string]map[string]string{
		"outputSpeech": speechletResponse.createSSMLOutputSpeech(text),
	}
}

func (speechletResponse *SpeechletResponse) ShowSimpleCard(title, content string) {
	card := map[string]string{
		"type":    "Simple",
		"title":   title,
		"content": content,
	}

	speechletResponse.response["card"] = card
}

func (speechletResponse *SpeechletResponse) ShowStandardCard(title, smallImageURL, largeImageUrl, text string) {
	card := map[string]interface{}{
		"type":  "Standard",
		"title": title,
		"text":  text,
		"image": map[string]string{
			"smallImageUrl": smallImageURL,
			"largeImageUrl": largeImageUrl,
		},
	}

	speechletResponse.response["card"] = card
}

func (speechletResponse *SpeechletResponse) ShowLinkAccountResponse() {
	card := map[string]string{
		"type": "LinkAccount",
	}

	speechletResponse.response["card"] = card
}

func (speechletResponse *SpeechletResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Version           string                 `json:"version"`
		SessionAttributes SessionAttributes      `json:"sessionAttributes"`
		Response          map[string]interface{} `json:"response"`
	}{
		Version:           speechletResponse.Version,
		SessionAttributes: speechletResponse.SessionAttributes,
		Response:          speechletResponse.response,
	})
}

func (speechletResponse *SpeechletResponse) createPlainTextOutputSpeech(text string) map[string]string {
	return map[string]string{
		"type": "PlainText",
		"text": text,
	}
}

func (speechletResponse *SpeechletResponse) createSSMLOutputSpeech(text string) map[string]string {
	return map[string]string{
		"type": "SSML",
		"ssml": text,
	}
}
