package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/genov8/schedviz/pkg/schedviz"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		snap, err := schedviz.ParseLine(line)
		if err != nil {
			continue
		}
		if snap == nil {
			continue
		}

		fmt.Printf(
			"t=%4dms | P=%2d (idle=%2d) | M=%2d (idle=%2d) | GRQ=%4d | LRQ=%v\n",
			snap.TimestampMs,
			snap.Gomaxprocs,
			snap.Idleprocs,
			snap.Threads,
			snap.IdleThreads,
			snap.GlobalRunQueue,
			snap.LocalRunQueues,
		)
	}
}
