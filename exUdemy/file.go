package ex

import (
	"fmt"
	"io"
	"os"
)

func PrintContentFile(path string) {
	fsys, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var content []byte
	temp := make([]byte, 99999)
	for {
		n, e := fsys.Read(temp)
		if e == io.EOF {
			break
		}
		content = append(content, temp[:n]...)
	}
	defer fsys.Close()
	fmt.Printf("%v", string(content))
}
