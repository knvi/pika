# pika go

0 dependencies Golang implementation for Pika.

## Usage

```go
package main

import "github.com/knvi/pika"

func main() {
	prefixes := []pika.PikaPrefixDefinition{
		{
			Prefix:      "test",
			Description: "test",
			Secure:      false,
		},
	}
	
	p := pika.NewPika(prefixes, pika.PikaInitOptions{
		Epoch:            1650153600000,
		NodeID:           622,
		DisableLowercase: true,
	})
	
	id := p.Gen("test")

	println(id)
}
```
