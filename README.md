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
| `Open` / `Openat` / `Stat`      | `getwd`            | `Pipe` / `Fork` / `Sync`  | `Bind`              |
| `Mkdir` / `Rmdir` / `Chdir`     | `getpid`           |                           | `listen`/`connect`  |
| `Pause` / `Close` / `Exit` / `Shutdown` |            |                           | `Accept`            |

---

## 🚀 Installation

```bash
go get github.com/gox7/unix
````

---

## ⚠️ Disclaimer

> Unsafe. Raw syscalls. For advanced use only.

---

## 📄 License

This project is licensed under the MIT License.
See [LICENSE](LICENSE) for details.
