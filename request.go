package alexa

// Request encapsulates a request sent to the skill by Amazon
type Request struct {
	// version of the alexa request, supplied by Amazon
	Version string `json:"version"`
	// session information
	Session Session `json:"session"`
	// information about the current state of the Alexa service and device at the time the request is sent
	Context Context `json:"context"`
	// the key-value map of the `request` object.
	// See: https://developer.amazon.com/docs/custom-skills/request-types-reference.html
	RequestMap map[string]interface{} `json:"request"`
}

// Session is the information about the current user session. It
// contains arbitrary Attributes, and part of the information also
// present in the Context object
type Session struct {
	// whether this request starts a new session
	New bool `json:"new"`
	// the session ID
	SessionID string `json:"sessionId"`
	// information about the application
	Application Application `json:"application"`
	// a key-value pair for arbitrary values
	Attributes SessionAttributes `json:"attributes"`
	// information about the user and permissions
	User User `json:"user"`
}

// SessionAttributes is a key-value map for storing session
// information. This is typically information that you have returned
// in the Session of a response. Alexa then sends the same map back to
// you so that you can retrieve previously stored values from it.
type SessionAttributes map[string]interface{}

// Application stores information about the application sent by Alexa
type Application struct {
	// the ID of the application
	ApplicationID string `json:"applicationId"`
}

// User encapsulates the information about an Alexa user
type User struct {
	// ID of the current user
	UserID string `json:"userId"`
	// a token identifying the user in another system
	//
	// See: https://developer.amazon.com/docs/custom-skills/link-an-alexa-user-with-a-user-in-your-system.html
	AccessToken string `json:"accessToken"`
}

// Context encapsulates information about the current state of the
// Alexa service and device at the time the request is sent. When the
// request is LaunchRequest and IntentRequest, the information
// available in this object is the same as the information in the
// Session object.
type Context struct {
	// the state of the Alexa system and the device interacting with the skill
	System System
	// information about the current state of the audio player
	AudioPlayer AudioPlayer
}

// System encapsulates information about the Alexa system and the device interacting with the skill.
type System struct {
	// a token that can be used to access Alexa-specific
	// APIs. This token encapsulates: any permissions the user has
	// consented to, and access to other Alexa-specific APIs, such
	// as the Progressive Response API
	APIAccessToken string `json:"apiAccessToken"`
	// the correct base URI to refer to by region, for use with APIs such as the Device Location API and Progressive Response API
	APIEndpoint string `json:"apiEndpoint"`
	// information about the application
	Application Application `json:"application"`
	// information about the device interacting with this skill
	Device Device `json:"device"`
	User   User   `json:"user"`
}

// Device encapsulates the information about the device and its capabilities.
type Device struct {
	// the ID of the device interacting with this skill
	DeviceID string `json:"deviceId"`
	// a map of supported interfaces. If a key is set in this map, it means that the device interacting with this skill has the specified capability
	SupportedInterfaces map[string]interface{} `json:"supportedInterfaces"`
}

// AudioPlayer encapsulates the information about the current state of the audio player interface.
type AudioPlayer struct {
	// the token set by the skill when starting a playback. Only
	// available when this skill is the most recent skill which
	// has started a playback
	Token string `json:"token"`
	// offset in milliseconds from the start of the track, 0 if
	// the track is in the beginning. Only available when this
	// skill is the last skill that started a playback
	OffsetInMilliseconds int64 `json:"offsetInMilliseconds"`
	// the last known state of audio playback
	PlayerActivity PlayerActivity `json:"playerActivity"`
}

// PlayerActivity is the last known state of audio playback.
type PlayerActivity string

// These are the different types of audio playback states.
const (
	PlayerActivityIdle           PlayerActivity = "IDLE"
	PlayerActivityPaused         PlayerActivity = "PAUSED"
	PlayerActivityPlaying        PlayerActivity = "PLAYING"
	PlayerActivityBufferUnderrun PlayerActivity = "BUFFER_UNDERRUN"
	PlayerActivityFinished       PlayerActivity = "FINISHED"
	PlayerActivityStopped        PlayerActivity = "STOPPED"
)
