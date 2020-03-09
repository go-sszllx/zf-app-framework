package scheduler

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
)

var (
	once sync.Once
	scheduler *Scheduler
)

// DataBlock represents data block which contains message and extra info
type DataBlock struct {
	Extra   interface{}
	Message interface{}
}

// ControlInfo represents task number, if used, Name and CtrNum should not be nil
type ControlInfo struct {
	Name		string  // control name
	CtrlNum		int  // control all task number
}

// Task represents basic unit to be dispatched.
type Task struct {
	CtrlInfo	*ControlInfo  // for control task
	Data   		*DataBlock  // data to do in the task
	DoTask 		func(*DataBlock)  // do task function
}

// dispatcher represents  one dispatcher of scheduler
type dispatcher struct {
	ctrlQueueNum	chan bool  // control the num of input into routine pool
	channel 		chan Task
	pool    		*workerpool.WorkerPool
}

type Scheduler struct {
	name       			string
	dispatcherMap 		sync.Map  // for dispatch all except ctrl task
	dispatcherCtrlMap	sync.Map  // for need control dispatch task number
}

func Init() *Scheduler {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())

		scheduler = newScheduler()
	})

	return scheduler
}

func newScheduler() *Scheduler {

	/*
	TODO:
		dispatcherNumber = beego.AppConfig.DefaultInt("scheduler::channelNum", cm.SchedulerChannelNum)
		taskQueueSize = beego.AppConfig.DefaultInt("scheduler::taskQueueSize", cm.SchedulerTaskQueueSize)
		scheduler = &Scheduler{}
	*/

	return nil
}

func (s *Scheduler) initDispatcher(name string, num int) *dispatcher {
	/*
	d := new(dispatcher)
	d.ctrlQueueNum = make(chan bool, num)  // control the num of input into routine pool
	d.channel = make(chan Task, 1)
	d.pool = workerpool.New(num)
	s.dispatcherCtrlMap.Store(name, d)

	// make dispatcher running
	go func(d *dispatcher) {
		for {
			select {
			case task := <-d.channel:
				d.ctrlDoTask(task)
			}
		}
	}(d)

	log.WithFields(log.Fields{
		"name":		name,
		"poolNum":	num,
	}).Info("dispatcher control init success...")

	return d
	*/

	return nil
}

