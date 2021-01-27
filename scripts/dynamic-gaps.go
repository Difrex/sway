// go build -o ~/.local/bin/dyngaps
package main

import (
	"fmt"
	"log"

	"github.com/Difrex/gosway/ipc"
)

const (
	largeMax    = 250
	largeMid    = 150
	largeMin    = 100
	smallMax    = 100
	smallMid    = 50
	smallMin    = 25
	scaleFactor = 1.1
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

	o, err := cc.GetActiveOutput()
	if err != nil {
		return 50
	}

	max := largeMax
	mid := largeMid
	min := largeMin
	if o.Scale < 1.1 {
		max = smallMax
		mid = smallMid
		min = smallMin
	}

	switch len(w) {
	case 1:
		return max
	case 2:
		return mid
	case 3:
		return min
	default:
		return smallMin
	}
}
