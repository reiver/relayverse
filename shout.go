package main

import (
	"fmt"
	"math/rand"
)

func shout() {
	const red     = 31
	const green   = 32
	const yellow  = 33
	const magenta = 35
	const cyan    = 36
	const white   = 37

	var numbers []int = []int{red,red,red, green,green,green, yellow, magenta,magenta,magenta, cyan,cyan,cyan, white,white}

	var random int = rand.Int()
	var index int = random % len(numbers)
	var number int = numbers[index]

	fmt.Printf("\x1b[%dm", number)

	fmt.Print(`
              __
   ________  / /___ ___  ___   _____  _____________
  / ___/ _ \/ / __ `+"`"+`/ / / / | / / _ \/ ___/ ___/ _ \
 / /  /  __/ / /_/ / /_/ /| |/ /  __/ /  (__  )  __/
/_/   \___/_/\__,_/\__, / |___/\___/_/  /____/\___/
                  /____/

`)
	fmt.Print("\x1b[0m")
}
