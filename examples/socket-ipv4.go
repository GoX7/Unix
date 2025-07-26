package examples

import (
	"fmt"

	"github.com/gox7/unix"
)

func Server() {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	addr := unix.SockAddrInet4([4]byte{127, 0, 0, 1}, 8080)
	unix.Bind(fd, addr[:])

	fmt.Println("Listen...")
	unix.Listen(fd, 1)
	for {
		var addrbuf [16]byte
		nfd, _ := unix.Accept(fd, addrbuf)
		unix.Write(nfd, []byte("Hello from socket!"))
		unix.Close(nfd)
	}
}

func Cliet() {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	addr := unix.SockAddrInet4([4]byte{127, 0, 0, 1}, 8080)
	unix.Connect(fd, addr[:])

	buffer := make([]byte, 128)
	unix.Read(fd, buffer)
	fmt.Println(string(buffer))
}
