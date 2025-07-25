package examples

import "github.com/gox7/unix"

func Console() {
	// Write bytes on output(1)
	unix.Write(unix.CODE_STDOUT, []byte("Name: "))

	// Create buffer for reading
	buffer := make([]byte, 1024)

	// Reading
	_, err := unix.Read(unix.CODE_STDIN, buffer)
	if err != nil {
		// Write error
		unix.Write(unix.CODE_STDOUT, []byte(err.Error()+"\n"))
		// Exit with code 1
		unix.Exit(1)
	}

	// Write name (buffer) on stdout(1)
	unix.Write(unix.CODE_STDOUT, buffer)
}
