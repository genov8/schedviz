package schedviz

import (
	"github.com/genov8/schedviz/internal/parser"
)

type Snapshot = parser.Snapshot

func ParseLine(line string) (*Snapshot, error) {
	return parser.ParseLine(line)
}
