package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/likexian/whois-go"
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
			//cmnd := exec.Command("whois.go", "arg")
			//cmnd.Start()
			simplenmap(*url)
		} else if ch == 2 {
			advancednmap(*url)
		} else if ch == 3 {
			whoislookup(*url)
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

func foo() error {
	return errors.New("Some error occured")
}
