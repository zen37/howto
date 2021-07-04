package main

//docker       run image <cmd> <params>
//go run main,go run 		<cmd> <params>

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//docker

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("unrecognized command")
	}

}

func run() {

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	fmt.Println("going to run command ["+os.Args[2]+"]"+" params", os.Args[3:])
	//cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderrr
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err.Error())
	}
}
