# Events
Simple golang events library

## Example
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/night-codes/events"
)

var (
	myEvent = events.New()
)

func main() {
	for i := 0; i < 10; i++ {
		l := i
		myEvent.On(func(data interface{}) {
			fmt.Println("Listener"+strconv.Itoa(l), data)
		})
	}
	for j := 0; j < 10; j++ {
		fmt.Println("====== Emit: " + strconv.Itoa(j) + " =====")
		myEvent.Emit("Event" + strconv.Itoa(j))
	}
}
```

## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2018 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
