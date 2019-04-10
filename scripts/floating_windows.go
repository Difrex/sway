package main

// Switch to the scratchpad window

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Difrex/gosway/ipc"
)

const (
	rofi = "rofi"
)

func main() {
	sc, err := ipc.NewSwayConnection()
	if err != nil {
		panic(err)
	}

	tree, err := sc.GetTree()
	if err != nil {
		panic(err)
	}

	windows := ipc.GetAllFloatingWindows(tree.Nodes)
	var windowsListForRofi []string
	for _, w := range windows {
		windowsListForRofi = append(windowsListForRofi, w.Name)
	}

	var id int64
	selectedWindow := strings.Trim(execRofi(strings.Join(windowsListForRofi, "\n")), "\n")
	for _, w := range windows {
		if w.Name == selectedWindow {
			id = w.ID
			break
		}
	}

	resp, err := sc.RunSwayCommand(fmt.Sprintf("[con_id=%d] focus", id))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func execRofi(windows string) string {
	var stdout bytes.Buffer

	cmd := exec.Command(rofi, "-dmenu", "-p", "select floating window")

	reader := strings.NewReader(windows)

	cmd.Stdin = reader
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd.Wait()

	return stdout.String()
}
