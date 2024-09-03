package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// type LoggerWrite struct{}

// func main() {
// res, err := http.Get("http://google.com")

// if err != nil {
// 	fmt.Println("Error:", err)
// 	os.Exit(1)
// }
// buf := make([]byte, 99999)
// n, _ := res.Body.Read(buf)

// if err != nil {

// 	fmt.Println("Error reading response body:", err)
// 	os.Exit(1)
// }

// println(string(buf[:n]))
// fmt.Println(n)

// p := make([]byte, 99999)
// new := io.LimitedReader{R: res.Body, N: 30000}

// io.Copy(os.Stdout, &new)
// io.Copy(os.Stdout, res.Body)

// lw := LoggerWrite{}
// io.Copy(lw, res.Body)
// }

// func (LoggerWrite) Write(b []byte) (int, error) {

// 	fmt.Println(string(b))
// 	fmt.Println("Writing all value:", len(b))

// 	return len(b), nil
// }

// exercise

// type Triangle struct {
// 	base   float32
// 	height float32
// }

// type Square struct {
// 	side float32
// }

// type LoggerArea interface {
// 	getArea() float32
// }

// func main() {

// 	triangle, square := initTest()

// 	printArea(triangle)
// 	printArea(square)

// }

// func initTest() (Triangle, Square) {

// 	triangleInit := Triangle{}

// 	triangleInit.base = float32(rand.Intn(1000))

// 	for {
// 		triangleInit.height = float32(rand.Intn(1000))

// 		if triangleInit.base > triangleInit.height {
// 			break
// 		}
// 	}

// 	return triangleInit, Square{side: float32(rand.Intn(1000))}

// }

// func (t Triangle) getArea() float32 {

// 	return t.base * t.height * 0.5
// }

// func (s Square) getArea() float32 {

// 	return s.side * s.side
// }

// func printArea(a LoggerArea) {
// 	fmt.Println("Area my houses", a.getArea())
// }

type FileCSV struct {
	name string
}

func main() {
	listArgs := os.Args

	if len(listArgs) == 1 {
		panic("Name file not empty")
	}

	nameFile := listArgs[1]

	fileData := FileCSV{}

	fileData.name = nameFile

	io.Copy(fileData, fileData.openFile())
}

func (f FileCSV) openFile() *os.File {
	fmt.Println("File name is:", f.name)

	filePointer, err := os.Open(f.name)

	if err != nil {
		fmt.Println("Oh no error:", err)
		os.Exit((1))
	}

	return filePointer

}

func (f FileCSV) Write(bs []byte) (int, error) {
	list := strings.Split(string(bs), ",")

	for k, s := range list {
		fmt.Println("Value is", s, "at index:", k)
	}

	fmt.Println("This size file is", len(bs), "bytes")
	return len(bs), nil
}
