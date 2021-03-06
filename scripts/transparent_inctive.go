// go build -o ~/.local/bin/tr_in transparent_inctive.go
package main

import (
	"fmt"
	"os"

	"strconv"

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

	_, err = sub.SendCommand(ipc.IPC_SUBSCRIBE, "[\"window\"]")
	if err != nil {
		panic(err)
	}

	ch := make(chan *ipc.Event)
	go sub.SubscribeListener(ch)

	log.Info("Waiting for the events")

	for {
		event := <-ch
		if event.Change == "focus" {
			changeTransparency(com, event.Container)
		}
	}
}

func changeTransparency(con *ipc.SwayConnection, focused ipc.Container) {
	_, err := con.RunSwayCommand(fmt.Sprintf("[con_id=%d] opacity %d", focused.ID, 1))
	if err != nil {
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
			o, err := window.Command(fmt.Sprintf("opacity %1.1f", getOpacity()))
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
			_, err := window.Command(fmt.Sprintf("opacity %1.1f", getOpacity()))
			if err != nil {
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

func getOpacity() float64 {
	o := os.Getenv("TR_IN_OPACITY")
	if o == "" {
		return OPACITY
	}

	opacity, err := strconv.ParseFloat(o, 1)
	if err != nil {
		log.Error(err)
		return OPACITY
	}
	return opacity
}
