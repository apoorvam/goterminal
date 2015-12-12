package main

import (
	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
	"time"
)

func main() {
	writer := goterminal.New()
	for i := 0; i < 5; i++ {
		msg := "This is a golang library to re-write text on Terminal, works on Windows too!.\n" +
			"Color of this text should change to Green after some processing.\n"

		ct.Foreground(ct.Yellow, false)
		writer.Buf.WriteString(msg)
		writer.Write()
		ct.ResetColor()

		time.Sleep(time.Second)

		writer.Buf.WriteString("Some text here.\nThis text may change later.\n")
		writer.Write()
		time.Sleep(time.Second)

		writer.Clear()
		// processing done here, after which color should change.

		ct.Foreground(ct.Green, false)
		writer.Buf.WriteString(msg)
		writer.Write()
		ct.ResetColor()

		writer.Buf.WriteString("This is new text.\nThis re-writes the old text.\n")
		writer.Write()
		time.Sleep(time.Second)

		writer.Reset()
	}

}
