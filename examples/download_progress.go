package main

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	"time"
)

func main() {
	// get an instance of writer
	writer := goterminal.New()

	for i := 0; i < 100; i = i + 10 {
		// add your text to writer's buffer
		writer.Buf.WriteString(fmt.Sprintf("Downloading (%d/100) bytes...\n", i))
		// write to terminal
		writer.Write()
		time.Sleep(time.Millisecond * 200)

		// clear the text written by previous write, so that it can be re-written.
		writer.Clear()
	}

	// reset the writer
	writer.Reset()
	fmt.Println("Download finished!")
}
