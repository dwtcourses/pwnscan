package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	var banner = `
				--------\
				|
				|
				|
				|
				|
	`
	color.Cyan("%s", banner)
	color.Red("          message")
	url := flag.String("url", "u", "-url")
	flag.Parse()
	if *url != "u" {
		color.Green("PWNScan Options :: \n")
		color.Red("[01] Simple Nmap Scan")
		color.Red("[02] Advanced Nmap Scan")
		color.Red("[03] Who-is Lookup")
		color.Red("[04] DNS Lookup")
		color.Red("[05] Reverse DNS Lookup")
		color.Red("[06] ETC")
		color.Red("[07] ETC")
		color.Red("[08] ETC")
		color.Red("[09] ETC")
		color.Red("[10] ETC")

		var ch int
		color.Green("Select an option : ")
		fmt.Scanf("%d", &ch)
		if ch == 1 {
			cmnd := exec.Command("whois.go", "arg")
			cmnd.Start()
			log.Println("log")
		} else if ch == 2 {
			cmnd := exec.Command("whois.go", "arg")
			cmnd.Start()
			log.Println("log")
		}
	} else {
		color.Red("Error : ")
		color.Green("Usage : ")
	}
}
