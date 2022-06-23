package main


import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {

    path := os.Args[1]

    fmt.Println(path)

    // open the file
    f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    Info, err := os.Stat(path)
    fsize := Info.Size() + 1

    fmt.Print("filename:")
    fmt.Println(f)

    fmt.Print("filesize:")
    fmt.Println(fsize)


    sum := 0
    for sum < 3 {
	fmt.Println("loop")
	buff := make([]byte, fsize)
	ioutil.WriteFile(path, buff, 0666)
	sum += 1
    }
}
