package main

/*
	[1]
	1,use redis as memche,mysql w/r split
	2.use redis as mq to finish mysql.CURD operater
	[2]use rpc as base operator
	[3]

*/

import (
	"fmt"
	"os"
	"os/signal"
	_ "runtime"

	ps "github.com/aalness/go-redis-pubsub"
	"github.com/cn0512/GoServerFrame/Config"
	"github.com/cn0512/GoServerFrame/Core/DB/Mysql"
	redis "github.com/cn0512/GoServerFrame/Core/DB/Redis"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	"github.com/cn0512/GoServerFrame/Core/Utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	svrCfg "github.com/cn0512/GoServerFrame/Core/Svrs"
)

type User struct {
	Id       int64
	NickName string
	Account  string
	Pwd      string
	RegTime  string
}

var u = &User{}
var svrItem *svrCfg.SvrItem

func test(engine *xorm.Engine) {
	err := engine.CreateTables(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

	size := 100
	queue := make(chan int, size)

	for i := 0; i < size; i++ {
		go func(x int) {
			//x := i
			err := engine.Ping()
			if err != nil {
				fmt.Println(err)
			} else {
				/*err = engine.(u)
				if err != nil {
					fmt.Println("Map user failed")
				} else {*/
				for j := 0; j < 10; j++ {
					if x+j < 2 {
						_, err = engine.Get(u)
					} else if x+j < 4 {
						users := make([]User, 0)
						err = engine.Find(&users)
					} else if x+j < 8 {
						_, err = engine.Count(u)
					} else if x+j < 16 {
						_, err = engine.Insert(&User{NickName: "xlw"})
					} else if x+j < 32 {
						_, err = engine.ID(1).Delete(u)
					}
					if err != nil {
						fmt.Println(err)
						queue <- x
						return
					}
				}
				fmt.Printf("%v success!\n", x)
				//}
			}
			queue <- x
		}(i)
	}

	for i := 0; i < size; i++ {
		<-queue
	}

	//conns := atomic.LoadInt32(&xorm.ConnectionNum)
	//fmt.Println("connection number:", conns)
	fmt.Println("end")
}

func init() {
	svrItem = &svrCfg.SvrItem{
		SvrType: "Storage",
		Uuid:    "",
		State:   false,
		Addr:    "",
	}
}

var redis_cfg = redis.Config{
	Server:    "localhost:6379",
	MaxIdle:   1,
	MaxActive: 0,
}

func RegLocal(reg bool) {
	pool := redis.NewRConnectionPool(redis_cfg)
	con := pool.Get()
	var item = svrCfg.SvrItem{
		SvrType: "Storage",
		State:   true,
		Uuid:    "123456798",
	}
	if reg {
		//redis.Set(&con, "reg:svr:"+item.SvrType, item)
		redis.HSet(&con, "reg:local", item.Uuid, item)
	} else {
		redis.HDel(&con, "reg:local", item.Uuid)
	}
}

func FInit() {
	//注销
	svrItem.Unreg()
	pub.Publish(Config.Topic_Svrs_Unreg, svrItem.Encode())
	RegLocal(false)
}

var pub ps.Publisher

func main() {
	//【1】初始化Mysql
	fmt.Println("Storage is running")
	engine, err := Mysql.GetMysql()
	if err != nil {
		Utils.Logout("Error:%v\n", err)
		return
	}
	defer engine.Close()
	//test(engine)

	//【2】初始化Redis
	go RegLocal(true)

	//【2】注册服务器
	Utils.Logout("try Pub`s redis")
	pub = mq.NewPub()
	defer pub.Shutdown()
	svrItem.Reg()
	pub.Publish(Config.Topic_Svrs_Reg, svrItem.Encode())
	Utils.Logout("finsh Pub`s redis")

	//【3】关闭
	defer FInit()

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt

}
