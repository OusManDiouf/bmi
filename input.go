package bmi

import (
	"bufio"
	"os"
)

// Reader Cant be const, because here we have a function call (runtime moment)
// const value is defined at compile time not at runtime !
var reader = bufio.NewReader(os.Stdin)
