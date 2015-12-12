# goterminal
A cross-platform Go-library for updating progress in terminal.

## Installation

````
  go get -u github.com/apoorvam/goterminal
````

## Example

````
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
````
Examples can be found [here](https://github.com/apoorvam/goterminal/tree/master/examples).

### Usage

* Create a `Writer` instance.
* Add your text to writer's buffer and call `Write()` to print text on Terminal. This can be called any number of times.
* Call `Clear()` to move the cursor to position where first `Write()` started so that it can be over-written.
* `Reset()` writer, so it clears its state for next Write.

