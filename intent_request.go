package alexa

// IntentRequest encapsulates the information for a request to an intent defined for the skill.
type IntentRequest struct {
	*LaunchRequest
	// status of the multi-turn dialog. This is only inclued if
	// the skill meets the requirements to use the dialog
	// directives
	//
	// See: https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#dialog-reqs
	DialogState dialogState `json:"dialogState"`
	// what the user wants
	Intent intent `json:"intent"`
}

// DialogState is the current status of the multi-turn dialog, which
// is set if this skill meets the requirements for using the
// dialog directives.
//
// See: https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#dialog-reqs
type dialogState string

// different dialog states
const (
	DialogStateStarted    dialogState = "STARTED"
	DialogStateInProgress dialogState = "IN_PROGRESS"
	DialogStateCompleted  dialogState = "COMPLETED"
)

// Intent encapsulates the information on what the user wants.
type intent struct {
	// name of the intent the user wants to access
	Name string `json:"name"`
	// whether the user has explicitly confirmed or denied the entire intent
	ConfirmationStatus confirmationStatus `json:"confirmationStatus"`
	// a key-value map of slots, defined in the skill's configuration
	Slots map[string]slot `json:"slots"`
}

// ConfirmationStatus defines if the user has accepted or denied the whole intent.
type confirmationStatus string

// enumiration of confirmation statuses
const (
	ConfirmationStatusNone      confirmationStatus = "NONE"
	ConfirmationStatusConfirmed confirmationStatus = "CONFIRMED"
	ConfirmationStatusDenied    confirmationStatus = "DENIED"
)

// Slot encapsulates the information about the slot defined in the skill's configuration, and the user-specified value for it.
type slot struct {
	// name of the slot
	Name string `json:"name"`
	// value of the slot
	Value string `json:"value"`
}

func newIntentRequestFromRequest(request Request) (*IntentRequest, error) {
	requestMap := request.RequestMap

	launchRequest, err := newLaunchRequestFromRequest(request)
	if err != nil {
		return nil, err
	}

	var dialogStateValue dialogState
	if value, ok := requestMap["dialogState"]; ok {
		dialogStateValue = dialogState(value.(string))
	} else {
		dialogStateValue = ""
	}

	intent := createIntent(requestMap["intent"])

	intentRequest := &IntentRequest{
		LaunchRequest: launchRequest,
		DialogState:   dialogStateValue,
		Intent:        intent,
	}

	return intentRequest, nil
}

func createIntent(intentMapInterface interface{}) intent {
	intentMap := intentMapInterface.(map[string]interface{})
	slots := createSlots(intentMap["slots"])

	return intent{
		Name:               intentMap["name"].(string),
		ConfirmationStatus: confirmationStatus(intentMap["confirmationStatus"].(string)),
		Slots:              slots,
	}
}

func createSlots(slotsMapInterface interface{}) map[string]slot {
	slots := map[string]slot{}

	slotsMap, ok := slotsMapInterface.(map[string]interface{})
	if !ok {
		return slots
	}

	if len(slotsMap) == 0 {
		return slots
	}

	for slotName, slotMapInterface := range slotsMap {
		slotMap := slotMapInterface.(map[string]interface{})
		slot := slot{
			Name:  slotMap["name"].(string),
			Value: slotMap["value"].(string),
		}
		slots[slotName] = slot
	}

	return slots
}
