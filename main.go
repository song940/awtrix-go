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
	fmt.Printf("  node-awtrix %s\n", version)
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
	fmt.Println("  https://github.com/song940/node-awtrix")
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
		res, _ := client.Set(map[string]interface{}{key: value})
		fmt.Println(res)
	case "get":
		key := args[1]
		value, _ := client.Get(key)
		fmt.Println(value)
	case "get_version":
		v, _ := client.GetVersion()
		fmt.Println(v)
	case "uptime":
		uptime, _ := client.GetUptime()
		fmt.Println(uptime)
	case "power":
		state := args[1]
		res, _ := client.Power(state == "on")
		fmt.Println(res)
	case "brightness":
		brightnessValue, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Brightness value should be a valid integer.")
			return
		}
		res, err := client.Brightness(brightnessValue)
		fmt.Println(res)
	case "notify":
		message := args[1]
		res, _ := client.Notify(message)
		fmt.Println(res)
	default:
		fmt.Println("Unknown command:", command)
	}
}
