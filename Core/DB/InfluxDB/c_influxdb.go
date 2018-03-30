package InfluxDB

import (
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/influxdata/influxdb/client"
)

type InfluxClient struct {
	Con *client.Client
}

func New(uri string) *InfluxClient {
	//fmt.Sprintf("http://%s:%d", "localhost", 8086)
	host, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: this assumes you've setup a user and have setup shell env variables,
	// namely INFLUX_USER/INFLUX_PWD. If not just omit Username/Password below.
	conf := client.Config{
		URL:      *host,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}
	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection", con)

	return &InfluxClient{
		Con: con,
	}
}

func (c *InfluxClient) Ping() {
	dur, ver, err := c.Con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Happy as a hippo! %v, %s", dur, ver)
}

func (c *InfluxClient) Query() {
	q := client.Query{
		Command:  "select count(value) from shapes",
		Database: "square_holes",
	}
	if response, err := c.Con.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	} else {
		log.Println(err)
	}
}

func (c *InfluxClient) Wirte() {
	var (
		shapes     = []string{"circle", "rectangle", "square", "triangle"}
		colors     = []string{"red", "blue", "green"}
		sampleSize = 1000
		pts        = make([]client.Point, sampleSize)
	)

	rand.Seed(42)
	for i := 0; i < sampleSize; i++ {
		pts[i] = client.Point{
			Measurement: "shapes",
			Tags: map[string]string{
				"color": strconv.Itoa(rand.Intn(len(colors))),
				"shape": strconv.Itoa(rand.Intn(len(shapes))),
			},
			Fields: map[string]interface{}{
				"value": rand.Intn(sampleSize),
			},
			Time:      time.Now(),
			Precision: "s",
		}
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "BumbeBeeTuna",
		RetentionPolicy: "default",
	}
	_, err := c.Con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}
