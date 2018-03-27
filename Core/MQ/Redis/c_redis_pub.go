package Redis

import (
	"log"
	_ "strconv"
	"sync"
	"time"

	ps "github.com/aalness/go-redis-pubsub"
	"github.com/cn0512/GoServerFrame/Config"
	"github.com/garyburd/redigo/redis"
)

type testPubHandler struct {
	mutex       sync.Mutex
	connections int
}

func (h *testPubHandler) OnPublishConnect(conn redis.Conn, address string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.connections++
}

func (h *testPubHandler) OnPublishConnectError(err error, nextTime time.Duration) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	log.Fatal(err)
}

func (h *testPubHandler) OnPublishError(err error, channel string, data []byte) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	log.Fatal(err)
}

var _ ps.PublicationHandler = new(testPubHandler)

func NewPub() ps.Publisher {
	ph := &testPubHandler{}
	p := ps.NewRedisPublisher(Config.Redis_addr, ph, 0, 0)
	//defer p.Shutdown()
	return p
}
