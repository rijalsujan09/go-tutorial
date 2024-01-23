package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang..")
	content := "This is the text need to go insid file.."

	file, err := os.Create("./mytext.txt")

	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)

	checkNilerror(err)

	fmt.Println("length  is : ", length)

	readFile("./mytext.txt")

	defer file.Close()
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilerror(err)
	fmt.Println("text date in the file is :\n", string(dataByte))
}

func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}

}
