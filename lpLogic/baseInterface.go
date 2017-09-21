/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"sync"
)

const (
	task_cmd_exit = 0x3001 // 退出
)

// 通讯结构
type CommMessage struct {
	Cmd      int
	Text     string
	ExInfo   interface{}
	Callback func(interface{})
}

// 接口
type BaseInterface interface {
	ProcessTask(*CommMessage)
}

// 基类
type BaseClass struct {
	BaseInterface
	chTask  chan *CommMessage
	msgPool *sync.Pool
}

// 启动
func (this *BaseClass) Start() {
	this.chTask = make(chan *CommMessage, 100)
	this.msgPool = &sync.Pool{New: func() interface{} { return new(CommMessage) }}
	go this.Routine_Task()
}

// 处理线程
func (this *BaseClass) Routine_Task() {
	for {
		select {
		case msg := <-this.chTask:
			// 结束
			if msg.Cmd == task_cmd_exit {
				this.msgPool.Put(msg)
				return
			}
			this.ProcessTask(msg)
			this.msgPool.Put(msg)
		}
	}
}

// 创建任务
func (this *BaseClass) CreateTask(cmd int, text string, ex interface{}, cb func(interface{})) {
	msg := this.msgPool.Get().(*CommMessage)
	msg.Cmd = cmd
	msg.Text = text
	msg.ExInfo = ex
	msg.Callback = cb
	this.chTask <- msg
}
