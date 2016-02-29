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
	Frequency      time.Duration
	TotalBits      uint64
	CustomEpoch    int64
}

func NewWorker(
	id, epoch int64, idBits, sequenceBits, timestampBits uint64,
	frequency time.Duration) Worker {

	totalBits := idBits + sequenceBits + timestampBits + 1
	if totalBits%6 != 0 {
		log.Fatal("totalBits + 1 must be evenly divisible by 6")
	}

	w := Worker{
		ID:             id,
		IDBits:         idBits,
		IDShift:        sequenceBits,
		Sequence:       0,
		SequenceBits:   sequenceBits,
		TimeStampBits:  timestampBits,
		TimeStampShift: sequenceBits + idBits,
		Frequency:      frequency,
		TotalBits:      totalBits,
		CustomEpoch:    epoch,
	}
	w.LastTimeStamp = w.Time()
	return w
}

func (w *Worker) NextID() int64 {
	timestamp := w.Time()

	if w.LastTimeStamp > timestamp {
		w.WaitForNextTime()
	}

	if w.LastTimeStamp == timestamp {
		w.Sequence = (w.Sequence + 1) % (1 << w.SequenceBits)
		if w.Sequence == 0 {
			w.WaitForNextTime()
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

func (w *Worker) IDString(id int64) string {
	str, _ := IntToString(id, w.TotalBits)
	return str
}

func (w *Worker) WaitForNextTime() {
	ts := w.Time()
	for ts <= w.LastTimeStamp {
		ts = w.Time()
	}
	w.LastTimeStamp = ts
}

func (w *Worker) Time() int64 {
	return time.Now().UnixNano() / int64(w.Frequency)
}
