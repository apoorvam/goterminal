package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apoorvam/goterminal"
)

func main() {
	// get an instance of writer
	writer := goterminal.New(os.Stdout)

	for i := 0; i < 100; i = i + 10 {
		// add your text to writer's buffer
		fmt.Fprintf(writer, "Downloading (%d/100) bytes...\n", i)
		// write to terminal
		writer.Print()
		time.Sleep(time.Millisecond * 200)

		// clear the text written by previous write, so that it can be re-written.
		writer.Clear()
	}

	// reset the writer
	writer.Reset()
	fmt.Println("Download finished!")
}
