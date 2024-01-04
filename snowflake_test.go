package pika

import (
	"testing"
)

func TestGenerateSnowflake(t *testing.T) {
    s := NewSnowflake(650153600000, 1023)
    id := s.Gen()

    deconstructed := s.Deconstruct(id)
    if deconstructed.Epoch != 650153600000 {
        t.Errorf("Expected Epoch to be 650153600000, got %d", deconstructed.Epoch)
    }
    if deconstructed.NodeID != 1023 {
        t.Errorf("Expected NodeID to be 1023, got %d", deconstructed.NodeID)
    }
}

func TestGenerateSnowflakes(t *testing.T) {
    s := NewSnowflake(650153600000, 1023)

    var snowflakes []string
    for i := 0; i < 4096; i++ {
        snowflakes = append(snowflakes, s.Gen())
    }
    lastSnowflake := s.Gen()

    for sequence, snowflake := range snowflakes {
        deconstructed := s.Deconstruct(snowflake)
        if deconstructed.Seq != sequence {
            t.Errorf("Expected Seq to be %d, got %d", sequence, deconstructed.Seq)
        }
    }

    deconstructed := s.Deconstruct(lastSnowflake)
    if deconstructed.Seq != 0 {
        t.Errorf("Expected Seq to be 0, got %d", deconstructed.Seq)
    }
}