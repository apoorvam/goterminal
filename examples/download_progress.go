package main

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	"time"
)

func main() {
	writer := goterminal.New() // get an instance of writer

	for i := 0; i < 100; i = i + 5 {
		writer.Buf.WriteString(fmt.Sprintf("Downloading (%d/100) bytes...\n", i)) // add your text to writer's buffer
		writer.Write()                                                            // write to terminal
		time.Sleep(time.Millisecond * 200)
		writer.Clear() // clear the text written by previous write, so that it can be re-written.
	}

	writer.Reset() // reset the writer
	fmt.Println("Download finished!")
}
