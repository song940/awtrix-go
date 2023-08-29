package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/song940/awtrix-go/awtrix"
)

const version = "0.0.1" // Update this to your desired version

var awtrixAPI = os.Getenv("AWTRIX_API")
var client = awtrix.NewClient(awtrixAPI)

func help() {
	fmt.Println()
	fmt.Printf("  awtrix %s\n", version)
	fmt.Println()
	fmt.Println("  Usage: awtrix <command>")
	fmt.Println()
	fmt.Println("  Commands:")
	fmt.Println()
	fmt.Println("  get_settings")
	fmt.Println("  set <key> <value>")
	fmt.Println("  brightness <value>")
	fmt.Println("  notify <text>")
	fmt.Println("  draw <array>")
	fmt.Println("  version")
	fmt.Println()
	fmt.Println("  Examples:")
	fmt.Println()
	fmt.Println("  awtrix get_settings")
	fmt.Println("  awtrix set Brightness 50")
	fmt.Println("  awtrix brightness 50")
	fmt.Println("  awtrix notify \"Hello World\"")
	fmt.Println()
	fmt.Println("  https://github.com/song940/awtrix-go")
	fmt.Println()
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		help()
		return
	}
	command := args[0]
	switch command {
	case "help":
		help()
	case "set":
		key := args[1]
		value := args[2]
		client.Set(map[string]interface{}{key: value})
	case "get":
		key := args[1]
		value, _ := client.Get(key)
		fmt.Println(value)
	case "get_version":
		v, _ := client.GetVersion()
		fmt.Println(v)
	case "get_settings":
		value, _ := client.GetSettings()
		fmt.Println(value)
	case "list_apps":
		apps, _ := client.ListApps()
		fmt.Println(apps)
	case "uptime":
		uptime, _ := client.GetUptime()
		fmt.Println(uptime)
	case "power":
		state := args[1]
		client.Power(state == "on")
	case "set_brightness":
		brightnessValue, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Brightness value should be a valid integer.")
			return
		}
		client.SetBrightness(brightnessValue)
	case "get_brightness":
		brightness, err := client.GetBrightness()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(brightness)
	case "notify":
		message := args[1]
		res, _ := client.Notify(message)
		fmt.Println(res)
	default:
		fmt.Println("Unknown command:", command)
	}
}
