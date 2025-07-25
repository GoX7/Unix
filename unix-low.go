package unix

import (
	"syscall"
	"unsafe"
)

func Mmap(addr uintptr, len uintptr, prots uintptr, flags uintptr, fd uintptr, offset uintptr) (int, error) {
	out, _, err := syscall.Syscall6(
		uintptr(SYS_MMAP),
		uintptr(addr),
		uintptr(len),
		uintptr(prots),
		uintptr(flags),
		uintptr(fd),
		uintptr(offset),
	)
	if err != 0 {
		return 0, err
	}

	return int(out), nil
}

func Munmap(addr uintptr, len uintptr, prots uintptr, flags uintptr, fd uintptr, offset uintptr) error {
	_, _, err := syscall.Syscall6(
		uintptr(SYS_MUNMAP),
		uintptr(addr),
		uintptr(len),
		uintptr(prots),
		uintptr(flags),
		uintptr(fd),
		uintptr(offset),
	)
	if err != 0 {
		return err
	}
	return nil
}

func Pipe(pipefd []int32) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_PIPE),
		uintptr(unsafe.Pointer(&pipefd[0])),
		0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Fork() error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_FORK),
		0, 0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Sync() error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_SYNC),
		0, 0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}
