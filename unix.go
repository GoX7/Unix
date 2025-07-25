package unix

import (
	"syscall"
	"unsafe"
)

func Write(fd int, data []byte) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_WRITE),
		uintptr(fd),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
	)
	return err
}

func Read(fd int, buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}

	n, _, err := syscall.Syscall(
		uintptr(SYS_READ),
		uintptr(fd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if err != 0 {
		return 0, err
	}

	return int(n), nil
}

func Open(path string, flags int, mode int32) (int, error) {
	ptr, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}

	fd, _, errno := syscall.Syscall(
		uintptr(SYS_OPEN),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(flags),
		uintptr(mode),
	)
	if errno != 0 {
		return 0, errno
	}

	return int(fd), nil
}

func Openat(path string, flags int, mode int32) (int, error) {
	ptr, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}

	fd, _, errno := syscall.Syscall6(
		uintptr(SYS_OPENAT),
		uintptr(int(AT_FDCWD)),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(flags),
		uintptr(mode),
		0, 0,
	)
	if errno != 0 {
		return 0, errno
	}

	return int(fd), nil
}

func Mkdir(path string, mode int) error {
	bpath := append([]byte(path), 0)
	_, _, err := syscall.Syscall(
		uintptr(SYS_MKDIR),
		uintptr(unsafe.Pointer(&bpath[0])),
		uintptr(mode),
		0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Rmdir(path string) error {
	bpath := append([]byte(path), 0)
	_, _, err := syscall.Syscall(
		uintptr(SYS_RMDIR),
		uintptr(unsafe.Pointer(&bpath[0])),
		0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Chdir(path string) error {
	bpath := append([]byte(path), 0)
	_, _, err := syscall.Syscall(
		uintptr(SYS_CHDIR),
		uintptr(unsafe.Pointer(&bpath[0])),
		0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Fchdir(path string) error {
	bpath := append([]byte(path), 0)
	_, _, err := syscall.Syscall(
		uintptr(SYS_FCHDIR),
		uintptr(unsafe.Pointer(&bpath[0])),
		0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Stat(path string) (*syscall.Stat_t, error) {
	ptr, err := syscall.BytePtrFromString(path)
	if err != nil {
		return nil, err
	}

	var stat syscall.Stat_t
	_, _, errno := syscall.Syscall6(
		uintptr(SYS_NEWFSTAT),
		uintptr(int(AT_FDCWD)),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(unsafe.Pointer(&stat)),
		uintptr(0), 0, 0,
	)
	if errno != 0 {
		return nil, errno
	}

	return &stat, nil
}

func Rename(old_path string, new_path string) error {
	bop := append([]byte(old_path), 0)
	nop := append([]byte(new_path), 0)
	_, _, err := syscall.Syscall(
		uintptr(SYS_RENAME),
		uintptr(unsafe.Pointer(&bop[0])),
		uintptr(unsafe.Pointer(&nop[0])),
		0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Pause() error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_PAUSE),
		0, 0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Exit(code int) error {
	_, _, err := syscall.RawSyscall(
		uintptr(SYS_EXIT),
		uintptr(code),
		0, 0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Close(fd int) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_CLOSE),
		uintptr(fd),
		0, 0,
	)
	return err
}

func Shutdown(sockfd int, how int) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_SHUTDOWN),
		uintptr(sockfd),
		uintptr(how),
		0,
	)
	if err != 0 {
		return err
	}
	return nil
}
