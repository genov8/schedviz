package main

import (
	"testing"

	"github.com/genov8/schedviz/pkg/schedviz"
)

func TestFormatSnapshot(t *testing.T) {
	snap := &schedviz.Snapshot{
		TimestampMs:        1000,
		Gomaxprocs:         14,
		GlobalRunQueue:     312,
		BusyProcs:          14,
		TotalLocalRunQueue: 221,
		TotalRunnable:      533,
		PressureLevel:      "critical",
	}

	got := formatSnapshot(snap)
	want := "1.0s  Ps 14/14 busy  runnable 533  pressure CRITICAL  GQ 312  LQ 221"
	if got != want {
		t.Fatalf("formatSnapshot() = %q, want %q", got, want)
	}
}
