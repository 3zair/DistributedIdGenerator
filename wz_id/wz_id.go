package wz_id

import (
	"fmt"
	"time"
	"wz_id_generator/api"
	"wz_id_generator/conf"
)

const (
	MAX_SEQUENCE int64 = 1 << 16
	ID_BITS            = 53
	TIME_BITS          = 32
	WORK_ID_BITS       = 5
	SEQ_BITS           = 16
)

var (
	DEFAULT_T int64 = 0
	sequence  int64 = 0
)

type WzID struct {
	Id int64
}

func (wi *WzID) NewID(in interface{}, id *int64) error {

	now := time.Now().Unix()
	api.EpochSecond = now - DEFAULT_T

	if api.EpochSecond < api.LastSecond {
		fmt.Printf("clock is back: %d from previous: %d\n", api.EpochSecond, api.LastSecond)
	}

	if api.LastSecond != api.EpochSecond {
		api.LastSecond = api.EpochSecond
		sequence = 0
	}

	sequence++

	// TODO ?
	if sequence > MAX_SEQUENCE {
		for {
			api.EpochSecond = time.Now().Unix() - DEFAULT_T
			if api.LastSecond < api.EpochSecond {
				break
			}
		}
	}

	*id = api.EpochSecond<<(WORK_ID_BITS+SEQ_BITS) + int64(conf.C.App.WorkID)<<SEQ_BITS + sequence

	return nil
}
