# Unix

Low-level syscall wrappers for Linux in Go.

Direct access to `syscall.Syscall`, no abstractions, no dependencies.

## Features

### Base operation:
- `write`/`read`
- `open`/`openat`/`stat`/`rename`
- `mkdir`/`rmdir`/`chdir`/`fchdir`
- `pause`/`close`/`exit`/`shutdown`
### Low operation:
- `mmap`/`munmap`
- `pipe`/`fork`/`sync`
### Net operation:
- `socket`
- `bind`/`listen`
- `accept`

## Installation

```bash
go get github.com/gox7/unix
````

## Example

```go
fd, _ := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
unix.Bind(fd, unix.AF_INET, [4]int{127, 0, 0, 1}, 8080)
unix.Listen(fd, 10)

for {
	client, _ := unix.Accept(fd, [16]byte{})
	unix.Write(client, []byte("hi\n"))
	unix.Close(client)
}
```

## Disclaimer

⚠️ Unsafe. Raw syscalls. For advanced use only.
