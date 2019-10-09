// go build -o ~/.local/bin/blur_background blur_background.go
package main

import (
	"fmt"
	"os"

	"os/exec"

	"github.com/Difrex/gosway/ipc"
	log "github.com/sirupsen/logrus"
)

const (
	WALLPAPER_PATH string = `/home/difrex/.config/sway/wallpapers/Photo by SpaceX.jpg`
	BLURED_PATH    string = `/tmp/blured.jpg`
)

type blurer struct {
	isSetup bool
}

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

	var blurer blurer
	for {
		event := <-ch
		if event.Change == "focus" {
			blurer.blur(com, event.Container)
		}
	}
}

func (b *blurer) setUnblur(con *ipc.SwayConnection) {
	b.isSetup = true
	log.Info("Setting up unblurred image ", wallpaper())
	setBackground(con, wallpaper())
}

func (b *blurer) setBlur(con *ipc.SwayConnection) {
	b.isSetup = false
	log.Info("Setting up blurred image ", BLURED_PATH)
	if !checkBluredImage() {
		makeBluredImage()
	}
	setBackground(con, BLURED_PATH)
}

func (b *blurer) blur(con *ipc.SwayConnection, container ipc.Container) {
	if !b.checkWindow(con, container) {
		if !b.isSetup {
			b.setUnblur(con)
			return
		}
		return
	}
	if b.isSetup {
		b.setBlur(con)
	}
}

func (b *blurer) checkWindow(con *ipc.SwayConnection, container ipc.Container) bool {
	tree, err := con.GetTree()
	if err != nil {
		log.Error(err)
		return false
	}

	floating := ipc.GetAllFloatingWindows(tree.Nodes)
	for _, f := range floating {
		if f.ID == int64(container.ID) {
			return true
		}
	}
	return false
}

func setBackground(con *ipc.SwayConnection, path string) {
	ws, err := con.GetActiveOutput()
	if err != nil {
		log.Error(err)
		return
	}
	wall := fmt.Sprintf(`output '%s' background "%s" fill`, ws.Name, path)
	log.Warn("Running Sway cmd: ", wall)
	b, err := con.RunSwayCommand(wall)
	if err != nil {
		log.Error(err)
	}
	log.Info("Received data: ", string(b))
}

func checkBluredImage() bool {
	f, err := os.Open(BLURED_PATH)
	defer f.Close()
	return os.IsExist(err)
}

func makeBluredImage() {
	cmd := exec.Command("convert", wallpaper(), "-blur", "0x24", BLURED_PATH)
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
}

// wallpaper returns path of the wallpaper
func wallpaper() string {
	path := os.Getenv("WALLPAPER_PATH")
	if path != "" {
		return path
	}
	return WALLPAPER_PATH
}
