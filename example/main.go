package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":4150")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println(text)
		}
	}()

	conn.Write([]byte("  V2"))
	for i := 0; i < 10; i++ {
		Pub(conn, "order", fmt.Sprintf("hello world %d", i))
	}
	Pub(conn, "task", "task one")

	time.Sleep(time.Second * 3)
}

func Pub(conn net.Conn, topic, message string)  {
	var temp = make([]byte, 4)
	binary.BigEndian.PutUint32(temp, uint32(len(message)))
	conn.Write([]byte("PUB " + topic + "\n"))
	conn.Write(temp)
	conn.Write([]byte(message))

}
