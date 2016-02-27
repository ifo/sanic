package sanic

import (
	"log"
	"time"
)

type Worker struct {
	ID             int64 // 0 - 2 ^ IDBits
	IDBits         uint64
	IDShift        uint64
	Sequence       int64 // 0 - 2 ^ SequenceBits
	SequenceBits   uint64
	LastTimeStamp  int64
	TimeStampBits  uint64
	TimeStampShift uint64
	TotalBits      uint64
	CustomEpoch    int64
}

func NewWorker(
	id, epoch int64,
	idBits, sequenceBits, timestampBits uint64) Worker {

	totalBits := idBits + sequenceBits + timestampBits + 1
	if totalBits%6 != 0 {
		log.Fatal("totalBits + 1 must be evenly divisible by 6")
	}

	return Worker{
		ID:             id,
		IDBits:         idBits,
		IDShift:        sequenceBits,
		Sequence:       0,
		SequenceBits:   sequenceBits,
		LastTimeStamp:  TimeMillis(),
		TimeStampBits:  timestampBits,
		TimeStampShift: sequenceBits + idBits,
		TotalBits:      totalBits,
		CustomEpoch:    epoch,
	}
}

func (w *Worker) NextID() int64 {
	timestamp := TimeMillis()

	if w.LastTimeStamp > timestamp {
		w.WaitForNextMilli()
	}

	if w.LastTimeStamp == timestamp {
		w.Sequence = (w.Sequence + 1) % (1 << w.SequenceBits)
		if w.Sequence == 0 {
			w.WaitForNextMilli()
			timestamp = w.LastTimeStamp
		}
	} else {
		w.Sequence = 0
	}

	w.LastTimeStamp = timestamp

	return (timestamp-w.CustomEpoch)<<w.TimeStampShift |
		w.ID<<w.IDShift |
		w.Sequence
}

func (w *Worker) WaitForNextMilli() {
	ts := TimeMillis()
	for ts <= w.LastTimeStamp {
		ts = TimeMillis()
	}
	w.LastTimeStamp = ts
}

func TimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
