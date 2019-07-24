package wz_id

import (
	"fmt"
	"sync"
	"time"
	"wz_id_generator/conf"
)

const (
	MAX_SEQUENCE        int64 = 1 << 16
	TIME_OFFSET_BITS          = 21
	WORK_ID_OFFSET_BITS       = 16
)

var (
	StartTime   int64 = 0
	sequence    int64 = 0
	epochSecond int64 = 0
	lastSecond  int64 = 0
	workID            = conf.C.App.WorkID
	m                 = &sync.Mutex{}

	count int64 = 0
)

type WzID struct {
	Id int64
}

func (wi *WzID) NewID(in interface{}, id *int64) error {
	m.Lock()
	epochSecond = getEpochSecond()
	if epochSecond < lastSecond {
		// clock is back
		fmt.Printf("clock is back: %d from previous: %d\n", epochSecond, lastSecond)
	}

	if lastSecond != epochSecond {
		lastSecond = epochSecond
		sequence = 0
		*id = getNextID()

		m.Unlock()
		return nil
	}

	CheckSequence()
	*id = getNextID()

	m.Unlock()
	return nil
}

func getNextID() int64 {
	return epochSecond<<TIME_OFFSET_BITS + workID<<WORK_ID_OFFSET_BITS + sequence
}

func getEpochSecond() int64 {
	return time.Now().Unix() - StartTime
}

func CheckSequence() {
	if sequence > MAX_SEQUENCE {
		tickerChan := time.NewTicker(time.Microsecond * 10).C

		select {
		case <-tickerChan:
			epochSecond = getEpochSecond()
			if lastSecond < epochSecond {
				sequence = 0
				break
			}
		}
	} else {
		sequence++
	}
}
