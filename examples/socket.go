package examples

import "github.com/gox7/unix"

func Socket() {
	// Creating socket fd
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	// Binding socket fd
	err = unix.Bind(fd, unix.AF_INET, [4]int{127, 0, 0, 1}, 8080) // 127.0.0.1 = [4]int{127, 0, 0, 1}
	if err != nil {
		panic(err)
	}

	// Listen connect on socket (127.0.0.1:8080)
	err = unix.Listen(fd, 1)
	if err != nil {
		panic(err)
	}

	for {
		// create buffer for addrbuf
		var buf [16]byte
		// accept connection
		client, _ := unix.Accept(fd, buf)
		// write message on client
		unix.Write(client, []byte("Hello from socket!"))
		// Close connection
		unix.Close(client)
	}
}
