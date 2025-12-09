# schedviz — Go Scheduler Trace Visualizer

`schedviz` is a lightweight CLI tool and Go library for parsing and visualizing  
Go runtime scheduler traces produced by:

```
GODEBUG=schedtrace=1000
```

It helps you understand what the Go scheduler is doing in real time:  
how many P's are busy, how many M's exist, how many goroutines are runnable,  
and how the local/global run queues evolve under load.

---

## ✨ Features

- Reads Go's `schedtrace` output from **stdin**
- Parses scheduler snapshots (`SCHED ...`)
- Displays key runtime metrics:
    - busy/idle P count
    - OS threads (M)
    - global runqueue size
    - local runqueue per P
    - total runnable goroutines
- Works both as:
    - a **CLI tool**, and
    - a **Go library** (`pkg/schedviz`)

---

## 🚀 Installation (CLI)

```bash
go install github.com/genov8/schedviz/cmd/schedviz@latest
```

The binary will be installed into:

```
~/go/bin/schedviz
```

---

## 🧪 Usage

### 🔥 1. Live visualization (recommended)

```bash
GODEBUG=schedtrace=1000 ./your_app 2>&1 | schedviz
```

### Example output:

```
t=1005ms | P=14 (idle=0) | M=15 (idle=0) | GRQ=312 | LRQ=[21 18 25 14 20 23 19 22 17 16 15 14 19 18]
t=2008ms | P=14 (idle=0) | M=15 (idle=0) | GRQ=287 | LRQ=[19 22 17 16 13 20 18 21 15 14 12 19 17 20]
t=3012ms | P=14 (idle=1) | M=15 (idle=2) | GRQ=198 | LRQ=[12 15 14 11 10 16 13 14  9 12  8 11 10 13]
```

---

## 🧩 Use as a Go Library

```go
import "github.com/genov8/schedviz/pkg/schedviz"

func main() {
    line := "SCHED 1003ms: gomaxprocs=14 idleprocs=0 ..."

    snap, err := schedviz.ParseLine(line)
    if err != nil {
        panic(err)
    }
    if snap != nil {
        fmt.Println("Global runqueue:", snap.GlobalRunQueue)
    }
}
```

---

## 📦 Project Structure

```
schedviz/
 ├── cmd/schedviz/       # CLI tool
 ├── pkg/schedviz/       # Public API (ParseLine, Snapshot)
 └── internal/parser/    # Internal SCHED trace parser
```

---

## 🔧 Notes

- Go's `schedtrace` output is sent to **stderr**
- If you want to pipe it into schedviz, redirect stderr → stdout:

```bash
GODEBUG=schedtrace=1000 ./your_app 2>&1 | schedviz
```

---

## 📅 Roadmap

The project is in active development.
I plan to continue improving schedviz, including:

- better visualization of scheduler activity
- more readable and informative output
- improvements to both the CLI and the parsing engine
- general enhancements based on real usageё

I’m glad to develop this tool further and appreciate any ideas or feedback.

---

## 📄 License

MIT — free to use and modify.
