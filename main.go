package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	conn, err := net.Dial("tcp", "gopher.floodgap.com:70")
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Fprint(conn, "\r\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
