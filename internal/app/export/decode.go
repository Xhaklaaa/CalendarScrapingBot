package export

import (
	"fmt"
	"io"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	jsondata, err := os.Open("out.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsondata.Close()

	data := make([]byte, 64)

	for {
		n, err := jsondata.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}

// func decodingJSON() {
// 	var m Message
// 	err := json.Unmarshal()

// }
