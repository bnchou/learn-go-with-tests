package main

import (
	"os"
	"time"

	clockface "learngowithtests/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
