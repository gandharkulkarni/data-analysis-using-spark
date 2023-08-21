package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
)

const (
	DirectoryPath = "/bigdata/mmalensek/nam/3hr_sample/sampled_2019/part-00198-20f67c5c-ec6c-4ad1-b1f8-3303ac1542e2-c000.tdv.gz"
	Host          = "localhost"
	Port          = "9999"
)

func streamFile(conn net.Conn, filePath string) {
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading file: %s", err)
			}
			break
		}
		conn.Write(buffer[:n])
	}
}

func startServer() {
	address := fmt.Sprintf("%s:%s", Host, Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
	defer listener.Close()

	log.Printf("Server started on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}

		fileList, err := filepath.Glob(filepath.Join(DirectoryPath, "*"))
		if err != nil {
			log.Printf("Failed to read file list: %s", err)
			conn.Close()
			continue
		}

		var wg sync.WaitGroup
		wg.Add(len(fileList))

		for _, filePath := range fileList {
			go func(path string) {
				defer wg.Done()
				streamFile(conn, path)
			}(filePath)
		}

		wg.Wait()
		conn.Close()
	}
}

func main() {
	startServer()
}
