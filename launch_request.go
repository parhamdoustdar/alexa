package alexa

import "time"

type LaunchRequest struct {
	RequestID string    `json:"requestId"`
	Time      time.Time `json:"timestamp"`
	Locale    string    `json:"locale"`
	Session   Session
	Context   Context
}

func newLaunchRequestFromRequest(request Request) (*LaunchRequest, error) {
	requestMap := request.RequestMap

	time, err := time.Parse("2006-01-02T15:04:05Z", requestMap["timestamp"].(string))
	if err != nil {
		return nil, err
	}

	requestId := requestMap["requestId"].(string)
	locale := requestMap["locale"].(string)

	launchRequest := &LaunchRequest{
		Time:      time,
		RequestID: requestId,
		Locale:    locale,
		Session:   request.Session,
		Context:   request.Context,
	}

	return launchRequest, nil
}
