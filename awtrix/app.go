package awtrix

import "encoding/json"

func (c *Awtrix) ListApps() (apps []string, err error) {
	resp, err := c.Get("installedApps")
	err = json.Unmarshal(resp, &apps)
	return
}

func (c *Awtrix) EnableApp(array []interface{}) (AwtrixResponse, error) {
	payload := map[string]interface{}{"enable": array}
	return c.Set(payload)
}

func (c *Awtrix) DisableApp(appName string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"disable": appName}
	return c.Set(payload)
}

func (c *Awtrix) SwitchTo(appName string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"switchTo": appName}
	return c.Call("/switchTo", payload)
}

func (c *Awtrix) AppControl(cmd string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"app": cmd}
	return c.Set(payload)
}

func (c *Awtrix) Pause() (AwtrixResponse, error) {
	return c.AppControl("pause")
}

func (c *Awtrix) Hold() (AwtrixResponse, error) {
	return c.AppControl("hold")
}

func (c *Awtrix) NextApp() (AwtrixResponse, error) {
	return c.AppControl("next")
}

func (c *Awtrix) PrevApp() (AwtrixResponse, error) {
	return c.AppControl("back")
}

func (c *Awtrix) SetOrders(apps []string) (AwtrixResponse, error) {
	payload := map[string]interface{}{"appList": apps}
	return c.Set(payload)
}
