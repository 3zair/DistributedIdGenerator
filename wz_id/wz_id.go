package wz_id

import (
	"time"
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
	*id = 1 << ID_BITS

	now := time.Now().Unix()
	t := now - DEFAULT_T

	if sequence > MAX_SEQUENCE {
		// TODO
	}

	*id = *id + t<<(WORK_ID_BITS+SEQ_BITS) + int64(conf.C.App.WorkID)<<SEQ_BITS + sequence
	sequence++
	return nil
}
