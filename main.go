// Command fish embellishes a message.
//
// The message is passed with a command-line flag.
// "Hello, World!" is used as the default message if no flag is passed.
package main

import (
	"flag"
	"fmt"

	"github.com/kytrinyx/fish/bubbles"
)

var msg = flag.String("msg", "Hello, World!", "The text that you want embellished.")

func main() {
	flag.Parse()

	fmt.Println(bubbles.Embellish(*msg))
}
