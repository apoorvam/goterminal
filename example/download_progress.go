package main

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	"time"
)

func main() {
	writer := goterminal.New()

	for i := 0; i < 100; i = i + 5 {
		writer.Buf.WriteString(fmt.Sprintf("Downloading (%d/100) bytes...\n", i))
		writer.Write()
		time.Sleep(time.Millisecond * 200)
		writer.Clear()
	}

	fmt.Println("Download finished")
}
