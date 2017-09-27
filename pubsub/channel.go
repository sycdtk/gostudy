package pubsub

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

//wg.Wrap(GetName, "test", &my)
//    wg.Wrap(GetSex, &my)
func (w *WaitGroupWrapper) Wrap(cb func(argvs ...interface{}), argvs ...interface{}) {
	w.Add(1)
	go func() {
		cb(argvs...)
		w.Done()
	}()
}

////WaitGroup封装结构
//type WaitGroupWrapper struct {
//    sync.WaitGroup
//}

//func (w *WaitGroupWrapper) Wrap(f func()) {
//    w.Add(1)
//    go func() {
//        f()
//        w.Done()
//    }()
//}

type Channel struct {
	Name    string
	clients map[int]*Client
	//  exitChan   chan int
	sync.RWMutex
	waitGroup    WaitGroupWrapper
	messageCount uint64
	exitFlag     int32
}

func NewChannel(channelName string) *Channel {
	return &Channel{
		Name: channelName,
		//  exitChan:       make(chan int),
		clients: make(map[int]*Client),
	}
}

func (ch *Channel) AddClient(client *Client) bool {
	ch.RLock()
	_, found := ch.clients[client.Id]
	ch.RUnlock()

	ch.Lock()
	if !found {
		ch.clients[client.Id] = client
	}
	ch.Unlock()
	return found
}

func (ch *Channel) DeleteClient(client *Client) int {
	var ret int
	ch.ReplyMsg(
		fmt.Sprintf("从channel:%s 中删除client:%d ", ch.Name, client.Id))

	ch.Lock()
	delete(ch.clients, client.Id)
	ch.Unlock()

	ch.RLock()
	ret = len(ch.clients)
	ch.RUnlock()

	return ret
}

func (ch *Channel) Notify(message string) bool {

	ch.RLock()
	defer ch.RUnlock()

	for cid, _ := range ch.clients {
		ch.ReplyMsg(
			fmt.Sprintf("channel:%s client:%d message:%s", ch.Name, cid, message))
	}
	return true
}

func (ch *Channel) ReplyMsg(message string) {
	ch.waitGroup.Wrap(func(msg ...interface{}) { fmt.Println(msg) }, message)
}

func (ch *Channel) Wait() {
	ch.waitGroup.Wait()
}

func (ch *Channel) Exiting() bool {
	return atomic.LoadInt32(&ch.exitFlag) == 1
}

func (ch *Channel) Exit() {
	if !atomic.CompareAndSwapInt32(&ch.exitFlag, 0, 1) {
		return
	}
	//close(ch.exitChan)
	ch.Wait()
}

func (ch *Channel) PutMessage(clientID int, message string) {
	ch.RLock()
	defer ch.RUnlock()

	if ch.Exiting() {
		return
	}

	//select {
	// case <-t.exitChan:
	// return
	//}
	fmt.Println(ch.Name, ":", message)

	atomic.AddUint64(&ch.messageCount, 1)
	return
}
