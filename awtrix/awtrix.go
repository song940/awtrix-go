package awtrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Awtrix struct {
	api string
}

type AwtrixResponse = []byte
type AwtrixKVResponse = map[string]any

func NewClient(api string) *Awtrix {
	return &Awtrix{api: api}
}

func (a *Awtrix) Call(path string, body interface{}) (out AwtrixResponse, err error) {
	method := "POST"
	url := a.api + path
	data, err := json.Marshal(body)
	if err != nil {
		return
	}
	payload := bytes.NewBuffer(data)
	headers := map[string]string{"Content-Type": "application/json"}
	client := &http.Client{}
	// log.Println(method, url, string(data))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}
	out, err = io.ReadAll(resp.Body)
	// log.Println(string(out))
	return
}

func (c *Awtrix) Notify(text string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"text": text}
	return c.Call("/notify", payload)
}

func (c *Awtrix) CallBasics(payload map[string]interface{}) (AwtrixResponse, error) {
	return c.Call("/basics", payload)
}

func (c *Awtrix) Set(values map[string]interface{}) (AwtrixResponse, error) {
	return c.CallBasics(values)
}

func (c *Awtrix) Get(key string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"get": key}
	return c.CallBasics(payload)
}

func (a *Awtrix) GetSettings() (out AwtrixKVResponse, err error) {
	resp, err := a.Get("settings")
	err = json.Unmarshal(resp, &out)
	return
}

func (a *Awtrix) GetSettingByKey(key string) (out any, err error) {
	resp, err := a.GetSettings()
	if err != nil {
		return
	}
	out = resp[key]
	return
}

func (a *Awtrix) GetVersion() (string, error) {
	resp, err := a.Get("version")
	var kv AwtrixKVResponse
	err = json.Unmarshal(resp, &kv)
	version := kv["version"].(string)
	return version, err
}

func (a *Awtrix) GetUptime() (string, error) {
	resp, err := a.Get("uptime")
	var kv AwtrixKVResponse
	json.Unmarshal(resp, &kv)
	uptime := kv["uptime"].(string)
	return uptime, err
}

func (a *Awtrix) SetBrightness(value int) error {
	payload := map[string]interface{}{"Brightness": value}
	a.Set(payload)
	return nil
}

func (a *Awtrix) GetBrightness() (float64, error) {
	value, err := a.GetSettingByKey("Brightness")
	return value.(float64), err
}

func (a *Awtrix) Power(state bool) error {
	payload := map[string]interface{}{"power": state}
	resp, err := a.CallBasics(payload)
	log.Println(resp)
	return err
}

func (a *Awtrix) PowerOn() error {
	return a.Power(true)
}

func (a *Awtrix) PowerOff() error {
	return a.Power(false)
}

func (c *Awtrix) SetSettings(value map[string]interface{}) ([]byte, error) {
	return c.Call("/settings", value)
}
