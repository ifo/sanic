package sanic

import (
	"time"
)

type Worker struct {
	ID             int64 // 0 - 2 ^ IDBits
	IDBits         uint64
	IDShift        uint64
	Group          int64 // 0 - 2 ^ GroupBits
	GroupBits      uint64
	GroupShift     uint64
	Sequence       int64 // 0 - 2 ^ SequenceBits
	SequenceBits   uint64
	LastTimeStamp  int64
	TimeStampBits  uint64
	TimeStampShift uint64
	CustomEpoch    int64
	Out            chan<- int64 // if nil, NextID must be called individually
}

func NewWorker(
	id, group, epoch int64,
	idBits, groupBits, sequenceBits, timestampBits uint64,
	output chan<- int64) Worker {

	return Worker{
		ID:             id,
		IDBits:         idBits,
		IDShift:        sequenceBits,
		Group:          group,
		GroupBits:      groupBits,
		GroupShift:     sequenceBits + idBits,
		Sequence:       0,
		SequenceBits:   sequenceBits,
		LastTimeStamp:  TimeMillis(),
		TimeStampBits:  timestampBits,
		TimeStampShift: sequenceBits + idBits + groupBits,
		CustomEpoch:    epoch,
		Out:            output,
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
		w.Group<<w.GroupShift |
		w.ID<<w.IDShift |
		w.Sequence
}

func (w *Worker) GenIDsForever() {
	if w.Out != nil {
		go func() {
			for {
				w.Out <- w.NextID()
			}
		}()
	}
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
