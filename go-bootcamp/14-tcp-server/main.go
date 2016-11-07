package main

import (
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// for {
		// 	bs := make([]byte, 1024)
		// 	n, err := conn.Read(bs)
		//
		// 	if err != nil {
		// 		break
		// 	}
		// 	_, err = conn.Write(bs[:n])
		// 	if err != nil {
		// 		break
		// 	}
		// }

		// Simple way to echo
		// io.Copy(conn, conn)


                // Multi-thread via go routines
                go func() {
		        io.Copy(conn, conn)
                        conn.Close()
                }()

		// io.WriteString(conn, "Hello World\n")

		// conn.Close()
	}
}
