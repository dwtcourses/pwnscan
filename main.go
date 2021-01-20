package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
)

func main() {
	menu := wmenu.NewMenu("Select an option : ")
	//	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Who-is Look up", nil, true, func(opt wmenu.Opt) error {
		// whois()
		// return nil
		//	cmd := exec.Command("whois.go")
		//	return nil
	})
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

}

//func whois() {
//	cmnd := exec.Command("whois.go", "arg")
//	//cmnd.Run() // and wait
//	cmnd.Start()
//}
