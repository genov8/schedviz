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

	for scanner.Scan() {
		line := scanner.Text()

		snap, err := schedviz.ParseLine(line)
		if err != nil {
			continue
		}
		if snap == nil {
			continue
		}

		fmt.Println(formatSnapshot(snap))
	}
}

func formatSnapshot(s *schedviz.Snapshot) string {
	return fmt.Sprintf(
		"%.1fs  Ps %d/%d busy  runnable %d  pressure %s  GQ %d  LQ %d",
		float64(s.TimestampMs)/1000,
		s.BusyProcs,
		s.Gomaxprocs,
		s.TotalRunnable,
		strings.ToUpper(s.PressureLevel),
		s.GlobalRunQueue,
		s.TotalLocalRunQueue,
	)
}
