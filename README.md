# pika go

0 dependencies Golang implementation for Pika.

## Usage

```go
prefixes := []PikaPrefixDefinition{
    {
        Prefix:      "test",
        Description: "test",
        Secure:      false,
    },
}

p := NewPika(prefixes, PikaInitOptions{
    Epoch:            1650153600000,
    NodeID:           622,
    DisableLowercase: true,
})

id := p.Gen("test")
```