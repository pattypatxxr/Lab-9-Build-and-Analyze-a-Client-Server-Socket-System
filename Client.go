package main

import (
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			conn, _ := net.Dial("tcp", "127.0.0.1:8080")
			conn.Write([]byte("Hello"))
			conn.Close()
		}()
	}

	wg.Wait()
}
