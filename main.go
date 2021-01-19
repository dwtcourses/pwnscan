package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/dixonwille/wmenu/v5"
)

func main() {
	menu := wmenu.NewMenu("Select an option : ")
	//	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Who-is Look up", nil, true, nil)
	whois()
	menu.Option("Look up", nil, true, nil)
	menu.Option("ETC", nil, true, nil)
	menu.Option("Reverse DNS Look up", nil, false, nil)
	menu.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
		fmt.Printf("Tacos are great")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
	if menu.Option("Who-is Look up", nil, true, nil) {
		fmt.Println("burgers")
	}

}

func whois() {
	cmnd := exec.Command("whois.go", "arg")
	//cmnd.Run() // and wait
	cmnd.Start()
}
