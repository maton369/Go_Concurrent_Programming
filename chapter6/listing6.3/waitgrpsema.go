package listing6_3

import (
	listing5_16 "Go_Concurrent_Programming/chapter5/listing5.16"
)

type WaitGrp struct {
    sema *listing5_16.Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
    return &WaitGrp{sema: listing5_16.NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
    wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
    wg.sema.Release()
}