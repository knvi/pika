package pika

import (
	"testing"
)

func TestInitPika(t *testing.T) {
    prefixes := []PikaPrefixDefinition{
        {
            Prefix:      "test",
            Description: "test",
            Secure:      false,
        },
        {
            Prefix:      "s_test",
            Description: "test",
            Secure:      true,
        },
    }

    p := NewPika(prefixes, PikaInitOptions{
        Epoch:            1650153600000,
        NodeID:           ComputeNodeID(),
        DisableLowercase: true,
    })

    id := p.Gen("test")
    deconstructed, _ := p.Decode(id)

    sID := p.Gen("s_test")
    sDeconstructed, _ := p.Decode(sID)

    if deconstructed.NodeID != p.NodeID {
        t.Errorf("Expected NodeID to be %d, got %d", p.NodeID, deconstructed.NodeID)
    }
    if deconstructed.Seq != 0 {
        t.Errorf("Expected Seq to be 0, got %d", deconstructed.Seq)
    }
    if deconstructed.Version != 1 {
        t.Errorf("Expected Version to be 1, got %d", deconstructed.Version)
    }
    if deconstructed.Epoch != 1650153600000 {
        t.Errorf("Expected Epoch to be 1650153600000, got %d", deconstructed.Epoch)
    }

    if sDeconstructed.NodeID != p.NodeID {
        t.Errorf("Expected NodeID to be %d, got %d", p.NodeID, sDeconstructed.NodeID)
    }
    if sDeconstructed.Seq != 1 {
        t.Errorf("Expected Seq to be 1, got %d", sDeconstructed.Seq)
    }
    if sDeconstructed.Version != 1 {
        t.Errorf("Expected Version to be 1, got %d", sDeconstructed.Version)
    }
    if sDeconstructed.Epoch != 1650153600000 {
        t.Errorf("Expected Epoch to be 1650153600000, got %d", sDeconstructed.Epoch)
    }
}

func TestInitPikaWithNodeID(t *testing.T) {
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
    deconstructed, _ := p.Decode(id)

    if deconstructed.NodeID != 622 {
        t.Errorf("Expected NodeID to be 622, got %d", deconstructed.NodeID)
    }
    if deconstructed.Seq != 0 {
        t.Errorf("Expected Seq to be 0, got %d", deconstructed.Seq)
    }
    if deconstructed.Version != 1 {
        t.Errorf("Expected Version to be 1, got %d", deconstructed.Version)
    }
    if deconstructed.Epoch != 1650153600000 {
        t.Errorf("Expected Epoch to be 1650153600000, got %d", deconstructed.Epoch)
    }
}