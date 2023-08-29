package awtrix

func (c *Awtrix) EnableApp(array []interface{}) ([]byte, error) {
	payload := map[string]interface{}{"enable": array}
	return c.Set(payload)
}

func (c *Awtrix) DisableApp(appName string) ([]byte, error) {
	payload := map[string]interface{}{"disable": appName}
	return c.Set(payload)
}

func (c *Awtrix) SwitchTo(appName string) ([]byte, error) {
	payload := map[string]interface{}{"switchTo": appName}
	return c.Call("/switchTo", payload)
}

func (c *Awtrix) AppControl(cmd string) ([]byte, error) {
	payload := map[string]interface{}{"app": cmd}
	return c.Set(payload)
}

func (c *Awtrix) Pause() ([]byte, error) {
	return c.AppControl("pause")
}

func (c *Awtrix) Hold() ([]byte, error) {
	return c.AppControl("hold")
}

func (c *Awtrix) NextApp() ([]byte, error) {
	return c.AppControl("next")
}

func (c *Awtrix) PrevApp() ([]byte, error) {
	return c.AppControl("back")
}

func (c *Awtrix) SetOrders(apps []string) ([]byte, error) {
	payload := map[string]interface{}{"appList": apps}
	return c.Set(payload)
}
