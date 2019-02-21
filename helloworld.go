package main

// Request object with JSON
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	// Request 객체 생성
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	// Content-Type Header added
	req.Header.Add("Content-Type", "application/json")

	// Client object에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response check...
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

// JSON data POST
/*
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response check...
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	str := string(respBody)
	println(str)
}
*/

/*
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if err != nil {
		panic(err)
	}

	// 필요시 헤더 추가 가능
	req.Header.Add("User-Agent", "Crawler")

	// Client 객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) // 바이트를 문자열로...
	fmt.Println(str)
}
*/
/*
import (
	"io"
	"os"
)

func main() {
	fi, err := os.Open("./go/src/github.com/ingubae/hello-world/README.md")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("/home/igbae/go/src/github.com/ingubae/hello-world/file.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buff := make([]byte, 1024)

	for {
		// Read ...
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 {
			break
		}

		// Write ...
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
*/
/*
import "time"

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 done")

		case <-done2:
			println("run2 done")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
*/

/*
import (
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func sum(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	a, _ := strconv.Atoi(v["a"])
	b, _ := strconv.Atoi(v["b"])

	fmt.Fprintf(w, "%d", a+b)
}

func main() {
	fmt.Println("Hello World")

	h := mux.NewRouter()
	h.HandleFunc("/hello", hello).Methods("GET")
	h.HandleFunc("/sum/{a}/{b}", sum).Methods("GET")
	http.Handle("/", h)
	http.ListenAndServe(":3000", nil)
}
*/
