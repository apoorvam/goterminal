# goterminal
A cross-platform Go-library for updating progress in terminal.

## Installation

````
    go get -u github.com/apoorvam/goterminal
````

## Example

````
    writer := goterminal.New()      // get an instance of writer

	for i := 0; i < 100; i = i + 5 {
		writer.Buf.WriteString(fmt.Sprintf("Downloading (%d/100) bytes...\n", i))   // add your text to writer's buffer
		writer.Write()              // write to terminal
		time.Sleep(time.Millisecond * 200)
		writer.Clear()              // clear the text written by previous write, so that it can be re-written.
	}

    writer.Reset()                  // reset the writer
	fmt.Println("Download finished")
````

### Usage

* Create a `Writer` instance
* Add your text to writer's buffer and call Write() to print text on Terminal. This can be called any number of times.
* Call Clear() to move the cursor to position where first Write started so that it can be over-written.
* Reset() writer, so it clears its state for next Write.

