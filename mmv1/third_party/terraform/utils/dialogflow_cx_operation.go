package google

import (
	"encoding/json"
	"fmt"
	"time"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

type DialogflowCXOperationWaiter struct {
	Config    *transport_tpg.Config
	UserAgent string
	CommonOperationWaiter
	Location string
}

func (w *DialogflowCXOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://%s-dialogflow.googleapis.com/v3/%s", w.Location, w.CommonOperationWaiter.Op.Name)

	return transport_tpg.SendRequest(w.Config, "GET", "", url, w.UserAgent, nil)
}

func createDialogflowCXWaiter(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent, location string) (*DialogflowCXOperationWaiter, error) {
	w := &DialogflowCXOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
		Location:  location,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func DialogflowCXOperationWaitTimeWithResponse(config *transport_tpg.Config, op map[string]interface{}, response *map[string]interface{}, activity, userAgent, location string, timeout time.Duration) error {
	w, err := createDialogflowCXWaiter(config, op, activity, userAgent, location)
	if err != nil {
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func DialogflowCXOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent, location string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createDialogflowCXWaiter(config, op, activity, userAgent, location)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
