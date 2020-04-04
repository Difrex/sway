// go build -o ~/.local/bin/key_layout key_layout.go
package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/Difrex/gosway/ipc"
	log "github.com/sirupsen/logrus"
)

const (
	EnLangEmoji    string          = "🇺🇸"
	RuLangEmoji    string          = "🇷🇺"
	EnLang         string          = "En"
	RuLang         string          = "Ru"
	XkbLayoutEvent ipc.ChangeEvent = "xkb_layout"

	UseEmojiEnvVar string = "SWAY_KEY_LAYOUT_USE_EMOJIS"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered")
		}
	}()

	ch := make(chan *ipc.Event)
	go subscribe(ch, "input")

	con, err := ipc.NewSwayConnection()
	if err != nil {
		log.Fatal(err)
	}

	// First launch
	printActiveLayout(con)

	for {
		event := <-ch
		if event.Change == XkbLayoutEvent {
			printActiveLayout(con)
		}
	}
}

func printActiveLayout(con *ipc.SwayConnection) {
	inputs, err := con.GetInputs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getActiveLayout(inputs))
}

func isEmojisEnabled() bool {
	useEmoji := strings.ToLower(os.Getenv(UseEmojiEnvVar))
	if useEmoji == "true" {
		return true
	}
	return false
}

func getActiveLayout(inputs []ipc.Input) string {
	primaryLayout := EnLang
	secondaryLayout := RuLang
	if isEmojisEnabled() {
		primaryLayout = EnLang
		secondaryLayout = RuLang
	}

	for _, input := range inputs {
		if input.Type == "keyboard" {
			switch input.XkbActiveLayoutName {
			case "Russian":
				return secondaryLayout
			default:
				return primaryLayout
			}
		}
	}
	return EnLang
}

func subscribe(ch chan *ipc.Event, eventType string) {
	con, err := ipc.NewSwayConnection()
	if err != nil {
		log.Fatal(err)
	}

	_, err = con.SendCommand(ipc.IPC_SUBSCRIBE, fmt.Sprintf(`["%s"]`, eventType))
	if err != nil {
		log.Fatal(err)
	}

	con.SubscribeListener(ch)
}
