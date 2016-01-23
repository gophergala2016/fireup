package main

import (
	"net"
	"io"
	"os"
	"fmt"
)

func acquireReader(name string) (io.Reader, error) {
	if len(name) == 0 {
		return os.Stdin, nil
	}else {
		file, err := os.Open(name)
		return file, err
	}
}

func PushFile(host string, port int, name string) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return err
	}

	reader, err := acquireReader(name)

	if err != nil {
		return err
	}

	fmt.Fprintln(conn, name)
	
	_, err = io.Copy(conn, reader)

	if err != nil {
		return err
	}

	return nil
}
