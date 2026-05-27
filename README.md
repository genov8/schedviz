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

Local example using the included load generator:

```bash
GODEBUG=schedtrace=1000 ./bin/load 2>&1 | ./bin/schedviz
```

### Example output:

```
1.0s  Ps 14/14 busy  runnable 533  pressure CRITICAL  GQ 312  LQ 221
2.0s  Ps 14/14 busy  runnable 485 ↓48  pressure CRITICAL  GQ 287 ↓25  LQ 198 ↓23
3.0s  Ps 13/14 busy  runnable 367 ↓118  pressure HIGH  GQ 198 ↓89  LQ 169 ↓29
```

### Reading the output

- **Ps busy** shows how many processors (`P`s) are currently doing Go work out of `GOMAXPROCS`.
- **runnable** is the total number of goroutines ready to run but not currently running.
- **pressure** summarizes runnable goroutines per `P`: `IDLE`, `LOW`, `MEDIUM`, `HIGH`, or `CRITICAL`.
- **GQ** is the global run queue: runnable goroutines waiting in the shared scheduler queue.
- **LQ** is the total local run queue size across all `P`s.
- **↑**, **↓**, and **±0** show the change since the previous scheduler snapshot: increased, decreased, or unchanged.

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
