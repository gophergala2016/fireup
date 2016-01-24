package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)


func PushFile(host string, port int, fname string) {
       var scanner *bufio.Scanner

        if len(fname) == 0 {
                scanner = bufio.NewScanner(os.Stdin)
        } else {
                file, err := os.Open(fname)
                defer file.Close()

                if err != nil {
                        fmt.Fprintln(os.Stderr, "ERROR", err)
                        return
                }

                scanner = bufio.NewScanner(file)
        }

        addr := fmt.Sprintf("%s:%d", host, port)
        conn, err := net.Dial("tcp", addr)

        if err != nil {
                fmt.Fprintln(os.Stderr, "ERROR:", err)
                return 
        }

        for scanner.Scan() {
                fmt.Fprintln(conn, scanner.Text())
        }
}
