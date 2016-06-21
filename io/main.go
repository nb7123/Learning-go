package main

import "fmt"
import "bufio"
import "os"
import "reflect"
import "io"
import "strings"

func main() {
//	ReadSth()
//	ReadFile()
//	Read()
	PrintArgs()
}

func PrintArgs() {
	if len(os.Args) > 0 {
		ret := strings.Join(os.Args[1:], "|")
		fmt.Println(ret)
	}
}

func Read() {
	inFile, err := os.Open("buffer.go")
	if nil != err {
		fmt.Println("Open file error:", err)
		return
	}
	defer inFile.Close()

	buffer := make([]byte, 10)
	for {
	_, e := inFile.Read(buffer)
	if e == io.EOF {
		return
	}
	fmt.Println(buffer)
	}
}

func ReadFile() {
	inFile, err := os.Open("buffer.go")
	f := func(){ fmt.Println("Defer"); inFile.Close()}
	defer f()
	if err != nil {
		fmt.Println("some error:", err)
		return
	}

	reader := bufio.NewReader(inFile)

	for {
	read, ee := reader.ReadString('\n')
	if io.EOF == ee {
		fmt.Println("Read end", ee)
		return
	}

	fmt.Println("Read:", read)
	}
}

func Defer(obj interface{}) {
	fmt.Println("Defer")
	v := reflect.ValueOf(obj)
	v.Close()
}

func ReadSth() {
	buffer := bufio.NewReader(os.Stdin)

	input, err := buffer.ReadString('S')

	if err != nil {
		fmt.Println("Some error:", err)
		return
	}

	var count int = 0
	for i, c := range input {
		count++
		fmt.Println(i, c)
	}

	fmt.Println(count)
}
