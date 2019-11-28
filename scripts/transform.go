// go build -o ~/.local/bin/transform_output transform.go
package main

import (
	"fmt"

	"strconv"

	"github.com/Difrex/gosway/ipc"
)

func main() {
	conn, err := ipc.NewSwayConnection()
	if err != nil {
		panic(err)
	}

	out, err := conn.GetActiveOutput()
	if err != nil {
		panic(err)
	}

	if out.Active {
		rotate(out, conn)
	}
}

func rotate(out *ipc.Output, conn *ipc.SwayConnection) {
	var transform int
	switch out.Transform {
	case "normal":
		transform = 0
	default:
		t, err := strconv.Atoi(out.Transform)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		transform = t
	}

	fmt.Println("Current transform:", transform)

	if transform < 270 {
		run(transform+90, out, conn)
		return
	}
	run(0, out, conn)
}

func run(t int, out *ipc.Output, conn *ipc.SwayConnection) {
	cmd := `output %s transform %d`
	msg, err := conn.RunSwayCommand(fmt.Sprintf(cmd, out.Name, t))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))
}
