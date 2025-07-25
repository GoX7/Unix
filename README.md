# Unix

[![Go Reference](https://pkg.go.dev/badge/github.com/gox7/unix.svg)](https://pkg.go.dev/github.com/gox7/unix)  
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**Low‑level syscall wrappers for Linux in Go**  
Direct access to `syscall.Syscall`—no abstractions, no dependencies.

---

## 🛠️ Features

| Base Operations                 |System information  | Low‑level Operations      | Networking          |
| --------------------------------|--------------------| ------------------------- | ------------------- |
| `Write` / `Read`                | `getuid`/`geteuid` | `Mmap` / `Munmap`         | `Socket`            |
| `Open` / `Openat` / `Stat`      | `getwd`            | `Pipe` / `Fork` / `Sync`  | `Bind` / `Listen`   |
| `Mkdir` / `Rmdir` / `Chdir`     | `getpid`           |                           | `Accept`            |
| `Pause` / `Close` / `Exit` / `Shutdown` |            |                           |                     |

---

## 🚀 Installation

```bash
go get github.com/gox7/unix
````

---

## 📋 Example

```go
package main

import (
    "log"
    "fmt"
    "github.com/gox7/unix"
)

func main() {
    // Create TCP socket
    fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
    if err != nil {
        log.Fatal(err)
    }
    defer unix.Close(fd)

    // Bind & listen
    unix.Bind(fd, unix.AF_INET, [4]int{127, 0, 0, 1}, 8080)
    unix.Listen(fd, 10)

    // Accept loop
    for {
        clientFD, err := unix.Accept(fd, [16]byte{})
        if err != nil {
            log.Println("accept error:", err)
            continue
        }
        unix.Write(clientFD, []byte("hi\n"))
        fmt.Printf("Sent greeting to fd %d\n", clientFD)
        unix.Close(clientFD)
    }
}
```

---

## ⚠️ Disclaimer

> Unsafe. Raw syscalls. For advanced use only.

---

## 📄 License

This project is licensed under the MIT License.
See [LICENSE](LICENSE) for details.
