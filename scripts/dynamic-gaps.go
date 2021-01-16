// go build -o ~/.local/bin/dyngaps
package main

import (
	"fmt"
	"log"

	"github.com/Difrex/gosway/ipc"
)

func main() {
	sc, err := ipc.NewSwayConnection()
	if err != nil {
		log.Fatal(err)
	}

	cc, err := ipc.NewSwayConnection()
	if err != nil {
		log.Fatal(err)
	}

	_, err = sc.SendCommand(ipc.IPC_SUBSCRIBE, `["window", "workspace"]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to sway")

	ch := make(chan *ipc.Event)
	go sc.SubscribeListener(ch)

	log.Println("Subscribing to events")

	for {
		<-ch
		setGaps(cc)
	}
}

func setGaps(cc *ipc.SwayConnection) error {
	gaps := horizontalGaps(cc)

	log.Println("setting gaps to", gaps)

	_, err := cc.RunSwayCommand(fmt.Sprintf("gaps horizontal current set %d", gaps))
	return err
}

func horizontalGaps(cc *ipc.SwayConnection) int {
	w, err := cc.GetFocusedWorkspaceWindows()
	if err != nil {
		log.Println("Error:", err)
	}

	switch len(w) {
	case 1:
		return 300
	case 2:
		return 200
	case 3:
		return 150
	default:
		return 50
	}
}
