package unix

import (
	"syscall"
	"unsafe"
)

func Getwd(buf []byte) (int, error) {
	cwd, _, err := syscall.Syscall(
		uintptr(SYS_GETCWD),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf),
		0,
	)
	if err != 0 {
		return 0, err
	}
	return int(cwd), nil
}

func Getpid() (int, error) {
	id, _, err := syscall.Syscall(
		uintptr(SYS_GETPID),
		0, 0, 0,
	)
	if err != 0 {
		return 0, err
	}
	return int(id), nil
}

func Getuid() (int, error) {
	id, _, err := syscall.Syscall(
		uintptr(SYS_GETUID),
		0, 0, 0,
	)
	if err != 0 {
		return 0, err
	}
	return int(id), nil
}

func Geteuid() (int, error) {
	id, _, err := syscall.Syscall(
		uintptr(SYS_GETEUID),
		0, 0, 0,
	)
	if err != 0 {
		return 0, err
	}
	return int(id), nil
}

func Getgid() (int, error) {
	id, _, err := syscall.Syscall(
		uintptr(SYS_GETGID),
		0, 0, 0,
	)
	if err != 0 {
		return 0, err
	}
	return int(id), nil
}
