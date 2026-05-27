package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	tsRe = regexp.MustCompile(`^SCHED\s+(\d+)ms:`)

	schedRe = regexp.MustCompile(
		`gomaxprocs=(\d+)\s+` +
			`idleprocs=(\d+)\s+` +
			`threads=(\d+)\s+` +
			`spinningthreads=(\d+)\s+` +
			`needspinning=(\d+)\s+` +
			`idlethreads=(\d+)\s+` +
			`runqueue=(\d+)\s+\[([0-9 ]+)\]\s+` +
			`schedticks=\[([0-9 ]+)\]`,
	)
)

type Snapshot struct {
	TimestampMs     int   // 0, 1000, 2000 ...
	Gomaxprocs      int   // gomaxprocs
	Idleprocs       int   // idleprocs
	Threads         int   // threads
	SpinningThreads int   // spinningthreads
	NeedSpinning    int   // needspinning
	IdleThreads     int   // idlethreads
	GlobalRunQueue  int   // runqueue
	LocalRunQueues  []int // [ ... ]
	SchedTicks      []int // schedticks=[ ... ]

	BusyProcs          int
	TotalLocalRunQueue int
	TotalRunnable      int
	RunnablePerP       float64
	PressureLevel      string
}

func ParseLine(line string) (*Snapshot, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	tsMatch := tsRe.FindStringSubmatch(line)
	if tsMatch == nil {
		return nil, nil
	}

	ts, err := strconv.Atoi(tsMatch[1])
	if err != nil {
		return nil, fmt.Errorf("parse timestamp: %w", err)
	}

	after := strings.TrimSpace(line[len(tsMatch[0]):])

	m := schedRe.FindStringSubmatch(after)
	if m == nil {
		return nil, fmt.Errorf("unrecognized sched format: %q", line)
	}

	s := &Snapshot{TimestampMs: ts}

	intFields := []struct {
		dest *int
		val  string
	}{
		{&s.Gomaxprocs, m[1]},
		{&s.Idleprocs, m[2]},
		{&s.Threads, m[3]},
		{&s.SpinningThreads, m[4]},
		{&s.NeedSpinning, m[5]},
		{&s.IdleThreads, m[6]},
		{&s.GlobalRunQueue, m[7]},
	}

	for _, f := range intFields {
		n, err := strconv.Atoi(f.val)
		if err != nil {
			return nil, fmt.Errorf("parse int field: %w", err)
		}
		*f.dest = n
	}

	localRunq, err := parseIntSlice(m[8])
	if err != nil {
		return nil, fmt.Errorf("parse local runq: %w", err)
	}
	s.LocalRunQueues = localRunq

	schedTicks, err := parseIntSlice(m[9])
	if err != nil {
		return nil, fmt.Errorf("parse schedticks: %w", err)
	}
	s.SchedTicks = schedTicks
	s.deriveMetrics()

	return s, nil
}

func (s *Snapshot) deriveMetrics() {
	s.BusyProcs = s.Gomaxprocs - s.Idleprocs

	for _, n := range s.LocalRunQueues {
		s.TotalLocalRunQueue += n
	}

	s.TotalRunnable = s.GlobalRunQueue + s.TotalLocalRunQueue
	if s.Gomaxprocs > 0 {
		s.RunnablePerP = float64(s.TotalRunnable) / float64(s.Gomaxprocs)
	}

	s.PressureLevel = pressureLevel(s.TotalRunnable, s.RunnablePerP)
}

func pressureLevel(totalRunnable int, runnablePerP float64) string {
	switch {
	case totalRunnable == 0:
		return "idle"
	case runnablePerP < 1:
		return "low"
	case runnablePerP < 5:
		return "medium"
	case runnablePerP < 20:
		return "high"
	default:
		return "critical"
	}
}

func parseIntSlice(raw string) ([]int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}

	parts := strings.Fields(raw)
	out := make([]int, 0, len(parts))

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("parseIntSlice: %w", err)
		}
		out = append(out, n)
	}

	return out, nil
}
