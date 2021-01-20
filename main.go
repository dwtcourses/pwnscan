package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/likexian/whois-go"
)

func main() {
	// clears the terminal
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	// start
	var banner = `
 ______        ___   _ ____   ____    _    _   _
|  _ \ \      / | \ | / ___| / ___|  / \  | \ | |
| |_) \ \ /\ / /|  \| \___ \| |     / _ \ |  \| |
|  __/ \ V  V / | |\  |___) | |___ / ___ \| |\  |
|_|     \_/\_/  |_| \_|____/ \____/_/   \_|_| \_|
				
	`

	color.Cyan("%s", banner)
	color.Red("	          made with <3 by @nolimitcarter")
	color.White("                   ")
	url := flag.String("url", "u", "-url")
	flag.Parse()
	if *url != "u" {
		color.Green("[-] Entered URL: ")
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Printf(*url)
		color.White("                   ")
		color.White("                   ")

		color.Green("[-] PWNScan Options : ")
		color.White("                   ")
		color.Yellow("[01] Simple Nmap Scan")
		color.Yellow("[02] Advanced Nmap Scan")
		color.Yellow("[03] Who-is Lookup")
		color.Yellow("[04] DNS Lookup")
		color.Yellow("[05] Reverse DNS Lookup")
		color.Yellow("[06] ETC")
		color.Yellow("[07] ETC")
		color.Yellow("[08] ETC")
		color.Yellow("[09] ETC")
		color.Yellow("[10] ETC")
		color.White("                   ")
		color.Yellow("[00] About")
		//color.White("                   ")
		color.Yellow("[99] Exit")
		color.White("                   ")

		var ch int
		color.Green("[-] Select an option : ")
		fmt.Scanf("%d", &ch)
		if ch == 1 {
			//cmnd := exec.Command("whois.go", "arg")
			//cmnd.Start()
			simplenmap(*url)
		} else if ch == 2 {
			advancednmap(*url)
		} else if ch == 3 {
			whoislookup(*url)
		} else if ch == 99 {
			exit()
		} else if ch == 00 {
			about()
		}
	} else {
		foo()
		if err := foo(); err != nil {
		}

		//color.Red("Error : Missing --url")
		//color.Green("Usage : ")
	}

}

//func simplenmap() {
//	var wg sync.WaitGroup
//	for i := 1; i <= 1024; i++ {
//		wg.Add(1)
//		go func(j int) {
//			defer wg.Done()
//			address :=
//		}
//	}
//}

func simplenmap(name string) {
	url := "https://api.hackertarget.com/nmap/?q=" + name
	resp, err := http.Get(url)
	if err != nil {
		foo()
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
}

func advancednmap(name string) {
	url := "https://api.hackertarget.com/nmap/?q=" + name
	resp, err := http.Get(url)
	if err != nil {
		foo()
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
}

func whoislookup(name string) {
	result, err := whois.Whois(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func about() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	var banner = `
 ______        ___   _ ____   ____    _    _   _
|  _ \ \      / | \ | / ___| / ___|  / \  | \ | |
| |_) \ \ /\ / /|  \| \___ \| |     / _ \ |  \| |
|  __/ \ V  V / | |\  |___) | |___ / ___ \| |\  |
|_|     \_/\_/  |_| \_|____/ \____/_/   \_|_| \_|
				
	`

	color.Cyan("%s", banner)
	color.Red("			 made with <3 by @nolimitcarter")
	color.White("                   ")
	color.Green("[-] Tool Created by @nolimitcarter [github.com/nolimitcarter]")
	color.White("                   ")
	color.Cyan("Author  :  @nolimitcarter")
	color.Cyan("Github  :  https://github.com/nolimitcarter")
	color.Cyan("Find me :  https://cartert.dev")
	color.Cyan("Version 1.1")
	color.White("                   ")
	color.Yellow("[00] Exit")
	color.White("                   ")
	var ch int
	color.Green("[-] Select an option : ")
	fmt.Scanf("%d", &ch)
	if ch == 00 {
		color.Yellow("Exiting... Thanks for using this tool!")
		exit()
	} else {
		foo()
	}
}

func exit() {
	color.Red("Exiting... Thanks for using this tool!")
	defer fmt.Println("!")
	os.Exit(0)
}

func foo() error {
	return errors.New("Some error occured")
}
