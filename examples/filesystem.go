package examples

import (
	"fmt"

	"github.com/gox7/unix"
)

func Filesystem() {
	// open file
	file, err := unix.Openat("file.txt", unix.O_CREATE|unix.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	// close file
	unix.Close(file)

	// stat file
	stat, _ := unix.Stat("file.txt")
	fmt.Printf("Size - %d\n", stat.Size) // get size from stat

	// write on file
	unix.Write(file, []byte("Hello, world!")) // file.txt - "Hello, world!"

	// read file
	buf := make([]byte, 1024)
	unix.Read(file, buf) // buf - "Hello, world!"

	// create dir (mkdir)
	unix.Mkdir("dir", 0755)

	// delete dir (rmdir)
	unix.Rmdir("dir")

	// rename file/dir
	unix.Rename("old path", "new path")

	// get curret work dir
	buf = make([]byte, 100)
	unix.Getwd(buf) // buf - work dir
}
