package Redis

import (
	"log"
	"sync"
	"time"

	"MSvrs/Config"
	"MSvrs/Core/Utils"
	_ "fmt"
	ps "github.com/aalness/go-redis-pubsub"
	"github.com/garyburd/redigo/redis"
)

type Callback interface{
	Call(msg string)
}
var cb Callback

type testSubHandler struct {
	mutex              sync.Mutex
	connections        int
	unsubscribeErrors  int
	receiveErrors      int
	disconnectedErrors int
	subscribeCount     int
	unsubscribeCount   int
	messages           map[string]map[string]struct{}
	messageChan        chan struct{}
	unsubscribeChan    chan struct{}
}

func (h *testSubHandler) OnSubscriberConnect(s ps.Subscriber,
	conn redis.Conn, address string, slot int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.connections++
}

func (h *testSubHandler) OnSubscriberConnectError(err error, nextTime time.Duration) {
	log.Fatal(err)
}

func (h *testSubHandler) OnSubscribe(channel string, count int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.subscribeCount++
}

func (h *testSubHandler) OnUnsubscribe(channel string, count int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.unsubscribeCount++
	h.unsubscribeChan <- struct{}{}
}

func (h *testSubHandler) OnMessage(channel string, data []byte) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	line, _ := redis.String(data, nil)
	messages, ok := h.messages[channel]
	if !ok {
		messages = make(map[string]struct{})
		h.messages[channel] = messages
	}
	messages[line] = struct{}{}
	h.messageChan <- struct{}{}
	//fmt.Println(message, messages)
	cb.Call(line)
}

func (h *testSubHandler) OnUnsubscribeError(channel string, err error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.unsubscribeErrors++
}

func (h *testSubHandler) OnReceiveError(err error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.receiveErrors++
}

func (h *testSubHandler) OnDisconnected(err error, slot int, channels []string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.disconnectedErrors++
}

func (h *testSubHandler) GetUnsubscribeTimeout() time.Duration {
	return 1 * time.Millisecond
}

func (h *testSubHandler) waitForMessages(count int) {
	seen := 0
	for {
		select {
		case <-h.messageChan:
			if seen++; seen == count {
				return
			}
		case <-time.After(30 * time.Second):
			log.Fatal("Timed out waiting for messages")
		}
	}
}

func (h *testSubHandler) waitForUnsubscribes(count int) {
	seen := 0
	for {
		select {
		case <-h.unsubscribeChan:
			if seen++; seen == count {
				return
			}
		case <-time.After(30 * time.Second):
			log.Fatal("Timed out waiting for unsubscribes")
		}
	}
}

func init() {

}

func newSubHandle() *testSubHandler {
	return &testSubHandler{
		messages:        make(map[string]map[string]struct{}),
		messageChan:     make(chan struct{}, 10000),
		unsubscribeChan: make(chan struct{}, 10000),
	}
}

func NewSub(topic string,sub_cb Callback) (ps.Subscriber, error) {
	h := newSubHandle()
	sub := ps.NewRedisSubscriber(Config.Redis_addr, h, 0)
	//defer sub.Shutdown()
	if err := <-sub.Subscribe(topic); err != nil {
		Utils.Logout("%v\n",err)
		return nil, err
	}
	cb = sub_cb
	return sub, nil
}
