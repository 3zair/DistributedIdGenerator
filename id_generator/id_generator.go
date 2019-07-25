package id_generator

import (
	"DistributedIdGenerator/conf"
	"sync"
	"time"
)

const (
	workerIdBits uint8 = 5
	seqBits      uint8 = 16
)

type IdGenerator struct {
	epoch     int64
	seq       int64
	seqMask   int64
	lastTime  int64
	workerId  int64
	timeShift uint8
	m         sync.Mutex
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{
		epoch:     conf.C.Epoch,
		seq:       0,
		seqMask:   2<<seqBits - 1,
		lastTime:  0,
		workerId:  conf.C.WorkId << seqBits,
		timeShift: workerIdBits + seqBits,
	}
}

func (ig *IdGenerator) NextId(in interface{}, id *int64) error {
	ig.m.Lock()
	defer ig.m.Unlock()

	now := time.Now().Unix()

	if now > ig.lastTime {
		ig.seq = 0
	} else {
		ig.seq = (ig.seq + 1) & ig.seqMask

		if ig.seq == 0 {
			for now <= ig.lastTime {
				time.Sleep(time.Millisecond * 100)
				now = time.Now().Unix()
			}
		}
	}

	ig.lastTime = now
	*id = ((now - ig.epoch) << ig.timeShift) | ig.workerId | ig.seq
	return nil
}
