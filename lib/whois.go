// git: github.com/nolimitcarter/pwnscan
// A modular framework designed to automate security tests.
// Licensed under the GPL

package lib

import (
	"fmt"
)

func whois(name string) {
	result, err := whois.Whois(name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
