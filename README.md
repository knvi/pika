# pika go

0 dependencies Golang implementation for [pika.](https://github.com/hopinc/pika)

## Install
To install, run
```sh
go get github.com/knvi/pika@v1.1.0
```

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
