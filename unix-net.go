package unix

import (
	"syscall"
	"unsafe"
)

func Socket(domain int, stype int, protocol int) (int, error) {
	socket, _, err := syscall.Syscall(
		uintptr(SYS_SOCKET),
		uintptr(domain),
		uintptr(stype),
		uintptr(protocol),
	)
	if err != 0 {
		return 0, err
	}
	return int(socket), nil
}

func Bind(fd int, domain int, host [4]int, port int) error {
	var connect [16]byte
	connect[0] = byte(domain)
	connect[2] = byte(port >> 8)
	connect[3] = byte(port & 0xFF)
	connect[4] = byte(host[0])
	connect[5] = byte(host[1])
	connect[6] = byte(host[2])
	connect[7] = byte(host[3])

	_, _, err := syscall.Syscall(
		uintptr(SYS_BIND),
		uintptr(fd),
		uintptr(unsafe.Pointer(&connect[0])),
		uintptr(16),
	)
	if err != 0 {
		return err
	}

	return nil
}

func Listen(fd int, backlog int) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_LISTEN),
		uintptr(fd),
		uintptr(backlog),
		0,
	)
	if err != 0 {
		return err
	}
	return nil
}

func Accept(fd int, buffer [16]byte) (int, error) {
	bufferlen := uintptr(unsafe.Sizeof(buffer))
	ndf, _, err := syscall.Syscall(
		uintptr(SYS_ACCEPT),
		uintptr(fd),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&bufferlen)),
	)
	if err != 0 {
		return 0, err
	}
	return int(ndf), nil
}
