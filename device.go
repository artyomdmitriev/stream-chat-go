package stream_chat

import (
	"net/http"
)

const (
	PushProviderAPNS     = pushProvider("apn")
	PushProviderFirebase = pushProvider("firebase")
)

type pushProvider = string

type Device struct {
	ID           string       //The device ID.
	UserID       string       //The user ID for this device.
	PushProvider pushProvider //The push provider for this device. One of constants PushProvider*
}

// Get list of devices for user
func (c *Client) GetDevices(userId string) (devices []Device, err error) {
	params := map[string][]string{
		"user_id": {userId},
	}

	var resp struct {
		Devices []Device `json:"devices"`
	}

	err = c.makeRequest(http.MethodGet, "devices", params, nil, &resp)

	return resp.Devices, err
}

// Add device to a user. Provider should be one of PushProvider* constant
func (c *Client) AddDevice(device Device) error {
	return c.makeRequest(http.MethodPost, "devices", nil, device, nil)
}

// Delete a device for a user
func (c *Client) DeleteDevice(userId string, deviceID string) error {
	params := map[string][]string{
		"id":      {deviceID},
		"user_id": {userId},
	}

	return c.makeRequest(http.MethodDelete, "devices", params, nil, nil)
}
