package main

import "fmt"
import "net/http"
//import "log"
import "os"
import "io/ioutil"
//import "in"

var urls = []string {
	"http://www.baidu.com",
	"http://www.sina.com",
	"http://www.163.com",
	"http://www.google.com",
}

func main() {
//	http.HandlerFunc("/", HelloServer)
//	http.HandleFunc("/", HelloServer)
//	http.HandleFunc("/hello/Name", Name)
//	err := http.ListenAndServe("127.0.0.1:8888", nil)

//	if err != nil {
//		log.Fatal("ListenAndServe error: ", err.Error())
//	}

	for _, url := range urls {
		ReadPages(url)	
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Receive request:", req.URL.Path)
	fmt.Fprintf(w, "Hello: " + req.URL.Path)
}

func Name(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request server name")
	fmt.Fprintf(w, "%v", os.Environ())
}

func ReadPages(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Read error: ", url, err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("Read error: ", err.Error())
		return
	}

	fmt.Println(url, ": ", string(data))
}
