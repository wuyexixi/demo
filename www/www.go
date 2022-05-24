package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Println("Received request from " + r.RemoteAddr)

	count, err := rdb.IncrBy("counter",1).Result()

	fmt.Fprintf(w, "pod14: "+hostname+" at your service in "+strconv.FormatInt(count, 10)+" times !\n")
}

func main() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)

	http.HandleFunc("/", handler)
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
