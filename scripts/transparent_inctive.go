// go build -o ~/.local/bin/tr_in transparent_inctive.go
package main

import (
	"fmt"
	"os"

	"github.com/Difrex/gosway/ipc"
	log "github.com/sirupsen/logrus"
)

const (
	OPACITY = 0.75
)

func main() {
	sub, err := ipc.NewSwayConnection()
	if err != nil {
		panic(err)
	}

	com, err := ipc.NewSwayConnection()
	if err != nil {
		panic(err)
	}

	o, err := sub.SendCommand(ipc.IPC_SUBSCRIBE, "[\"window\"]")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(o))

	ch := make(chan *ipc.Event)
	go sub.SubscribeListener(ch)

	fmt.Println("Waiting for the events")

	for {
		event := <-ch
		if event.Change == "focus" {
			changeTransparency(com, event.Container)
		}
	}
}

func changeTransparency(con *ipc.SwayConnection, focused ipc.Container) {
	o, err := con.RunSwayCommand(fmt.Sprintf("[con_id=%d] opacity %d", focused.ID, 1))
	if err != nil {
		fmt.Println(string(o))
		log.Error(err)
		checkSway()
		return
	}

	windows, err := con.GetFocusedWorkspaceWindows()
	if err != nil {
		log.Error(err)
		checkSway()
		return
	}

	for _, window := range windows {
		if int(window.ID) != focused.ID {
			o, err := window.Command(fmt.Sprintf("opacity %1.1f", OPACITY))
			if err != nil {
				fmt.Println(string(o))
				log.Error(err)
				checkSway()
				return
			}
		}
	}

	tree, err := con.GetTree()
	if err != nil {
		log.Error(err)
		checkSway()
		return
	}
	for _, window := range ipc.GetAllFloatingWindows(tree.Nodes) {
		if int(window.ID) != focused.ID {
			o, err := window.Command(fmt.Sprintf("opacity %1.1f", OPACITY))
			if err != nil {
				fmt.Println(string(o))
				log.Error(err)
				checkSway()
				return
			}
		}
	}
}

func checkSway() {
	if !ipc.IsSwayAvailable() {
		log.Error("Sway not available")
		os.Exit(2)
	}
}
