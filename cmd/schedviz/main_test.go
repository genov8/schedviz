package main

import (
	"testing"

	"github.com/genov8/schedviz/pkg/schedviz"
)

func TestFormatSnapshotWithoutPreviousSnapshot(t *testing.T) {
	snap := &schedviz.Snapshot{
		TimestampMs:        1000,
		Gomaxprocs:         14,
		GlobalRunQueue:     312,
		BusyProcs:          14,
		TotalLocalRunQueue: 221,
		TotalRunnable:      533,
		PressureLevel:      "critical",
	}

	got := formatSnapshot(snap, nil)
	want := "1.0s  Ps 14/14 busy  runnable 533  pressure CRITICAL  GQ 312  LQ 221"
	if got != want {
		t.Fatalf("formatSnapshot() = %q, want %q", got, want)
	}
}

func TestFormatSnapshotWithDeltas(t *testing.T) {
	prev := &schedviz.Snapshot{
		TotalRunnable:      186,
		GlobalRunQueue:     124,
		TotalLocalRunQueue: 62,
	}
	snap := &schedviz.Snapshot{
		TimestampMs:        2000,
		Gomaxprocs:         14,
		GlobalRunQueue:     27,
		BusyProcs:          12,
		TotalLocalRunQueue: 17,
		TotalRunnable:      44,
		PressureLevel:      "medium",
	}

	got := formatSnapshot(snap, prev)
	want := "2.0s  Ps 12/14 busy  runnable 44 ↓142  pressure MEDIUM  GQ 27 ↓97  LQ 17 ↓45"
	if got != want {
		t.Fatalf("formatSnapshot() = %q, want %q", got, want)
	}
}

func TestFormatSnapshotWithUnchangedDeltas(t *testing.T) {
	prev := &schedviz.Snapshot{
		TotalRunnable:      44,
		GlobalRunQueue:     27,
		TotalLocalRunQueue: 17,
	}
	snap := &schedviz.Snapshot{
		TimestampMs:        3000,
		Gomaxprocs:         14,
		GlobalRunQueue:     27,
		BusyProcs:          12,
		TotalLocalRunQueue: 17,
		TotalRunnable:      44,
		PressureLevel:      "medium",
	}

	got := formatSnapshot(snap, prev)
	want := "3.0s  Ps 12/14 busy  runnable 44 ±0  pressure MEDIUM  GQ 27 ±0  LQ 17 ±0"
	if got != want {
		t.Fatalf("formatSnapshot() = %q, want %q", got, want)
	}
}

func TestFormatSnapshotWithIncreasedDeltas(t *testing.T) {
	prev := &schedviz.Snapshot{
		TotalRunnable:      40,
		GlobalRunQueue:     20,
		TotalLocalRunQueue: 20,
	}
	snap := &schedviz.Snapshot{
		TimestampMs:        4000,
		Gomaxprocs:         14,
		GlobalRunQueue:     27,
		BusyProcs:          14,
		TotalLocalRunQueue: 25,
		TotalRunnable:      52,
		PressureLevel:      "medium",
	}

	got := formatSnapshot(snap, prev)
	want := "4.0s  Ps 14/14 busy  runnable 52 ↑12  pressure MEDIUM  GQ 27 ↑7  LQ 25 ↑5"
	if got != want {
		t.Fatalf("formatSnapshot() = %q, want %q", got, want)
	}
}
