package awtrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Awtrix struct {
	api string
}

func NewClient(api string) *Awtrix {
	return &Awtrix{api: api}
}

func (a *Awtrix) Call(path string, body interface{}) ([]byte, error) {
	method := "POST"
	url := a.api + path
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewBuffer(data)
	headers := map[string]string{"Content-Type": "application/json"}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *Awtrix) CallBasics(payload map[string]interface{}) ([]byte, error) {
	return c.Call("/basics", payload)
}

func (c *Awtrix) Set(values map[string]interface{}) ([]byte, error) {
	return c.CallBasics(values)
}

func (c *Awtrix) Get(key string) ([]byte, error) {
	payload := map[string]interface{}{"key": key}
	return c.CallBasics(payload)
}

func (c *Awtrix) SetSettings(value map[string]interface{}) ([]byte, error) {
	return c.Call("/settings", value)
}

func (a *Awtrix) GetSettings() ([]byte, error) {
	return a.Get("/settings")
}

func (a *Awtrix) GetVersion() ([]byte, error) {
	return a.Get("version")
}

func (a *Awtrix) GetUptime() ([]byte, error) {
	return a.Get("uptime")
}

func (c *Awtrix) Notify(text string) ([]byte, error) {
	payload := map[string]interface{}{"text": text}
	return c.Call("/notify", payload)
}

func (a *Awtrix) Brightness(value int) ([]byte, error) {
	payload := map[string]interface{}{"Brightness": value}
	return a.Set(payload)
}

func (a *Awtrix) Power(state bool) ([]byte, error) {
	payload := map[string]interface{}{"power": state}
	return a.CallBasics(payload)
}

func (a *Awtrix) PowerOn() ([]byte, error) {
	return a.Power(true)
}

func (a *Awtrix) PowerOff() ([]byte, error) {
	return a.Power(false)
}
