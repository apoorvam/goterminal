package main

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
	"time"
)

func main() {
	writer := goterminal.New()
	for i := 0; i < 5; i++ {
		msg := "This is a golang library to re-write text on Terminal, works on Windows too!.\n" +
			"Color of this text should change to Green after some processing."

		ct.Foreground(ct.Yellow, false)
		fmt.Fprintln(writer, msg)
		writer.Print()
		ct.ResetColor()

		time.Sleep(time.Second)

		fmt.Fprintln(writer, "Some text here.\nThis text will change later.")
		writer.Print()
		time.Sleep(time.Second)

		// processing done here, after which color should change or text should be over-written.
		writer.Clear()

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, msg)
		writer.Print()
		ct.ResetColor()

		fmt.Fprintln(writer, "This is new text.\nThis is re-written text.")
		writer.Print()
		time.Sleep(time.Second)

		writer.Reset()
	}

}
