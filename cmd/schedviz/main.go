package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/genov8/schedviz/pkg/schedviz"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var prev *schedviz.Snapshot

	for scanner.Scan() {
		line := scanner.Text()

		snap, err := schedviz.ParseLine(line)
		if err != nil {
			continue
		}
		if snap == nil {
			continue
		}

		fmt.Println(formatSnapshot(snap, prev))
		prev = snap
	}
}

func formatSnapshot(s, prev *schedviz.Snapshot) string {
	return fmt.Sprintf(
		"%.1fs  Ps %d/%d busy  runnable %d%s  pressure %s  GQ %d%s  LQ %d%s",
		float64(s.TimestampMs)/1000,
		s.BusyProcs,
		s.Gomaxprocs,
		s.TotalRunnable,
		formatDelta(s.TotalRunnable, prev, func(p *schedviz.Snapshot) int {
			return p.TotalRunnable
		}),
		strings.ToUpper(s.PressureLevel),
		s.GlobalRunQueue,
		formatDelta(s.GlobalRunQueue, prev, func(p *schedviz.Snapshot) int {
			return p.GlobalRunQueue
		}),
		s.TotalLocalRunQueue,
		formatDelta(s.TotalLocalRunQueue, prev, func(p *schedviz.Snapshot) int {
			return p.TotalLocalRunQueue
		}),
	)
}

func formatDelta(current int, prev *schedviz.Snapshot, previousValue func(*schedviz.Snapshot) int) string {
	if prev == nil {
		return ""
	}

	previous := previousValue(prev)
	delta := current - previous
	switch {
	case delta > 0:
		return fmt.Sprintf(" ↑%d", delta)
	case delta < 0:
		return fmt.Sprintf(" ↓%d", -delta)
	default:
		return " ±0"
	}
}
