package pika

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var VALID_PREFIX = regexp.MustCompile(`^[a-z0-9_]+$`)
var DEFAULT_EPOCH int64 = 1640995200000 // Jan 1, 2022

type PikaPrefixDefinition struct {
	Prefix string
	Description string
	Secure bool
}

type DecodedPika struct {
	Prefix string
	Tail string
	Snowflake int64
	NodeID int
	Seq int
	Version int
	Epoch int64
	PrefixRecord PikaPrefixDefinition
}

type Pika struct {
	Prefixes map[string]PikaPrefixDefinition
	Epoch int64
	NodeID int
	DisableLowercase bool
	snowflake *Snowflake
}

type PikaInitOptions struct {
	Epoch int64
	NodeID int
	DisableLowercase bool
}

func NewPika(prefixes []PikaPrefixDefinition, opts PikaInitOptions) *Pika {
	var epoch int64
	if opts.Epoch == 0 { epoch = DEFAULT_EPOCH } else { epoch = opts.Epoch }

	var nodeID int
	if opts.NodeID == 0 { nodeID = ComputeNodeID() } else { nodeID = opts.NodeID }

	p := &Pika {
		Prefixes: make(map[string]PikaPrefixDefinition),
		Epoch: epoch,
		NodeID: nodeID,
		DisableLowercase: opts.DisableLowercase,
		snowflake: NewSnowflake(epoch, nodeID),
	}

	for _, prefix := range prefixes {
        if !VALID_PREFIX.MatchString(prefix.Prefix) {
            panic(fmt.Sprintf("invalid prefix; prefixes must be alphanumeric (a-z0-9_) and may include underscores; received: %s", prefix.Prefix))
        }

        p.Prefixes[prefix.Prefix] = prefix
    }

	return p
}

func (p *Pika) Gen(prefix string) string {
	if !VALID_PREFIX.MatchString(prefix) {
        panic(fmt.Sprintf("invalid prefix; prefixes must be alphanumeric (a-z0-9_) and may include underscores; received: %s", prefix))
    }

	snowflake := p.snowflake.Gen()

	var tail string 
	if p.Prefixes[prefix].Secure {
		b := make([]byte, 16)
		rand.Read(b)
		tail = fmt.Sprintf("s_%s_%s", hex.EncodeToString(b), snowflake)
	} else {
		tail = snowflake
	}

	return fmt.Sprintf("%s_%s", prefix, base64.URLEncoding.EncodeToString([]byte(tail)))
}

func (p *Pika) Decode(id string) (DecodedPika, error) {
	s := strings.Split(id, "_")
	tail := s[len(s) - 1]
	prefix := strings.Join(s[:len(s)-1], "_")

	decodedTail, _ := base64.URLEncoding.DecodeString(tail)

	decodedTailStr := string(decodedTail)
	sf := strings.Split(decodedTailStr, "_")
	sfStr := sf[len(sf)-1]

	snowflake := p.snowflake.Deconstruct(sfStr)

	return DecodedPika{
		Prefix: prefix,
		Tail: tail,
		Snowflake: snowflake.ID,
		NodeID: p.NodeID,
		Seq: snowflake.Seq,
		Epoch: p.Epoch,
		Version: 1,
		PrefixRecord: p.Prefixes[prefix],
	}, nil
}

func ComputeNodeID() int {
	interfaces, _ := net.Interfaces()

	for _, iface := range interfaces {
		if iface.HardwareAddr.String() != "00:00:00:00:00:00" {
			mac := iface.HardwareAddr.String()
			macInt, _ := strconv.ParseInt(strings.ReplaceAll(mac, ":", ""), 16, 64)
			return int(macInt % 1024)
		}
	}

	return 0
}