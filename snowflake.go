package pika

import (
	"fmt"
	"strconv"
)

type Snowflake struct {
	Epoch int64
	NodeID int
	Seq int
	LastTS int64
}

type DeconstructedSnowflake struct {
	ID int64
	Timestamp int64
	NodeID int
	Seq int
	Epoch int64
}

func NewSnowflake(epoch int64, node_id int) *Snowflake {
	return &Snowflake{
		Epoch: epoch,
		NodeID: node_id,
		Seq: 0,
		LastTS: 0,
	}
}

// -- Snowflake Struct Functions
func (s *Snowflake) Gen() string {
	return s.GenWithTimestamp(nowTimestamp())
}


func (s *Snowflake) GenWithTimestamp(timestamp int64) string {
	if s.Seq >= 4095 && s.LastTS == timestamp {
		for nowTimestamp() - timestamp < 1 {
			continue
		}
	}

	snowflake := ((timestamp - s.Epoch) << 22) | (int64(s.NodeID) << 12) | int64(s.Seq)

	s.Seq = ((s.Seq + 1) & 4095)

	if s.Seq == 4095 {
		s.LastTS = timestamp
	}

	return fmt.Sprint(snowflake)
}

func (s *Snowflake) Deconstruct(id string) *DeconstructedSnowflake {
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}

	return &DeconstructedSnowflake{
		ID: intId,
		Timestamp: (intId >> 22) + s.Epoch,
		NodeID: int((intId >> 12) & 0b1111111111),
		Seq: int(intId & 0b1111_1111_1111),
		Epoch: s.Epoch,
	}
}