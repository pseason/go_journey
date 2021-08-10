package call

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 17:38
version 1.0.0
desc:

*/

const (
	ChanSize       = 15
	TimeOut  int64 = 3
)

var futureTaskChan chan string

var futureTaskMap sync.Map

var node *snowflake.Node

type futureTask struct {
	callId   string
	deadline int64
	exec     func(data interface{})
	data     interface{}
}

/*
初始化
*/
func Init() {
	futureTaskChan = make(chan string, ChanSize)
	node, _ = snowflake.NewNode(1)
	// do receive callback doTask
	go func() {
		for {
			select {
			case callId := <-futureTaskChan:
				go doTask(callId)
			}
		}
	}()
	// do timeout range delete
	go func() {
		time.Sleep(time.Millisecond * 100)
		for {
			futureTaskMap.Range(func(key, value interface{}) bool {
				future := value.(futureTask)
				if time.Now().Unix() > future.deadline {
					fmt.Println("Range Delete:", future.callId)
					futureTaskMap.Delete(key)
					return false
				}
				return true
			})
		}
	}()
}

/*
callback
*/
func Call(callId string) {
	futureTaskChan <- callId
}

func AfterFutureTask(exec func(data interface{}), data interface{}) (callId string) {
	callId = node.Generate().String()
	store(callId, exec, data)
	return
}

func store(callId string, exec func(data interface{}), data interface{}) {
	// struct
	future := futureTask{
		callId:   callId,
		deadline: time.Now().Unix() + TimeOut,
		exec:     exec,
		data:     data,
	}
	// 存入map
	futureTaskMap.Store(callId, future)
}

func doTask(callId string) {
	if v, ok := futureTaskMap.LoadAndDelete(callId); ok {
		future := v.(futureTask)
		// 如果未超时，则执行
		if time.Now().Unix() < future.deadline {
			fmt.Println("doTask: ", callId)
			future.exec(future.data)
		} else {
			fmt.Println("unDoTask: ", callId)
		}
	}
}
