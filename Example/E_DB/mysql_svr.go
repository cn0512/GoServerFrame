package main

import (
	_ "MSvrs/Config"
	"MSvrs/Core/DB/Mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "os"
	"runtime"
)

type User struct {
	Id       int64
	NickName string
	Account  string
	Pwd      string
	RegTime  string
}

var u = &User{}

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

func main() {
	runtime.GOMAXPROCS(2)
	engine,err := Mysql.GetMysql()
	if err != nil {
		return
	}
	defer engine.Close()
	test(engine)
}
