package parser

import "testing"

func TestParseLineDerivesMetrics(t *testing.T) {
	line := "SCHED 1000ms: gomaxprocs=14 idleprocs=0 threads=15 spinningthreads=0 needspinning=0 idlethreads=0 runqueue=312 [21 18 25 14 20 23 19 22 17 16 15 14 19 18] schedticks=[1 2 3 4 5 6 7 8 9 10 11 12 13 14]"

	snap, err := ParseLine(line)
	if err != nil {
		t.Fatalf("ParseLine() error = %v", err)
	}
	if snap == nil {
		t.Fatal("ParseLine() returned nil snapshot")
	}

	if snap.BusyProcs != 14 {
		t.Fatalf("BusyProcs = %d, want 14", snap.BusyProcs)
	}
	if snap.TotalLocalRunQueue != 261 {
		t.Fatalf("TotalLocalRunQueue = %d, want 261", snap.TotalLocalRunQueue)
	}
	if snap.TotalRunnable != 573 {
		t.Fatalf("TotalRunnable = %d, want 573", snap.TotalRunnable)
	}
	if snap.RunnablePerP != float64(573)/14 {
		t.Fatalf("RunnablePerP = %f, want %f", snap.RunnablePerP, float64(573)/14)
	}
	if snap.PressureLevel != "critical" {
		t.Fatalf("PressureLevel = %q, want critical", snap.PressureLevel)
	}
}

func TestParseLinePressureLevels(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		pressure string
	}{
		{
			name:     "idle",
			line:     "SCHED 1000ms: gomaxprocs=4 idleprocs=4 threads=5 spinningthreads=0 needspinning=0 idlethreads=4 runqueue=0 [0 0 0 0] schedticks=[1 1 1 1]",
			pressure: "idle",
		},
		{
			name:     "low",
			line:     "SCHED 1000ms: gomaxprocs=4 idleprocs=3 threads=5 spinningthreads=0 needspinning=0 idlethreads=4 runqueue=0 [1 0 0 0] schedticks=[1 1 1 1]",
			pressure: "low",
		},
		{
			name:     "medium",
			line:     "SCHED 1000ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 needspinning=0 idlethreads=0 runqueue=2 [1 1 1 0] schedticks=[1 1 1 1]",
			pressure: "medium",
		},
		{
			name:     "high",
			line:     "SCHED 1000ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 needspinning=0 idlethreads=0 runqueue=10 [3 3 2 2] schedticks=[1 1 1 1]",
			pressure: "high",
		},
		{
			name:     "critical",
			line:     "SCHED 1000ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 needspinning=0 idlethreads=0 runqueue=50 [10 10 5 5] schedticks=[1 1 1 1]",
			pressure: "critical",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snap, err := ParseLine(tt.line)
			if err != nil {
				t.Fatalf("ParseLine() error = %v", err)
			}
			if snap.PressureLevel != tt.pressure {
				t.Fatalf("PressureLevel = %q, want %q", snap.PressureLevel, tt.pressure)
			}
		})
	}
}
