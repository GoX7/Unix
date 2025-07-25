package unix

import (
	"syscall"
	"unsafe"
)

// Create fd socket
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

// Create connect ip4 for bind
func SockAddrInet4(host [4]byte, port int) [16]byte {
	var connect [16]byte
	connect[0] = byte(AF_INET)
	connect[2] = byte(port >> 8)
	connect[3] = byte(port & 0xFF)
	connect[4] = byte(host[0])
	connect[5] = byte(host[1])
	connect[6] = byte(host[2])
	connect[7] = byte(host[3])
	return connect
}

// Create connect ip6 for bind
func SockAddrInet6(host [16]byte, port int) [28]byte {
	var connect [28]byte
	connect[0] = byte(AF_INET6)
	connect[2] = byte(port >> 8)
	connect[3] = byte(port & 0xFF)
	copy(connect[8:24], host[:])
	return connect
}

// Binding connect from connect bytes for socket
func Bind(fd int, addr []byte) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_BIND),
		uintptr(fd),
		uintptr(unsafe.Pointer(&addr[0])),
		uintptr(len(addr)),
	)
	if err != 0 {
		return err
	}

	return nil
}

// Listen connect for socket
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

// Connect to socket
func Connect(fd int, addr []byte) error {
	_, _, err := syscall.Syscall(
		uintptr(SYS_CONNECT),
		uintptr(fd),
		uintptr(unsafe.Pointer(&addr[0])),
		uintptr(len(addr)),
	)
	if err != 0 {
		return err
	}
	return nil
}

// Accept connect for connection
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
