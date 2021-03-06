package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aweisser/ev3/goev3"
	"github.com/aweisser/ev3/robot"
)

var ev3 = goev3.Create(robot.Specs{WheelDiameter: 3.5})

func main() {
	defer ev3.Handle(goev3.WAIT_FOR_ESCAPE_BUTTON)
	fmt.Printf("Starting EV3 restless command control ...")
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is EV3 mission control. Start a mission!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	ev3.Greet()
}
