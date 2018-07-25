package scipipe

import (
	"time"
)

// AuditInfo contains structured audit/provenance logging information for a
// particular task (invocation), to go with all outgoing IPs from that task
type AuditInfo struct {
	ID          string
	ProcessName string
	Command     string
	Params      map[string]string
	Tags        map[string]string
	StartTime   time.Time
	FinishTime  time.Time
	ExecTimeNS  time.Duration
	Upstream    map[string]*AuditInfo
	OutFiles    map[string]string
}

// NewAuditInfo returns a new AuditInfo struct
func NewAuditInfo() *AuditInfo {
	return &AuditInfo{
		ID:          randSeqLC(20),
		ProcessName: "",
		Command:     "",
		Params:      make(map[string]string),
		Tags:        make(map[string]string),
		ExecTimeNS:  -1,
		Upstream:    make(map[string]*AuditInfo),
		OutFiles:    make(map[string]string),
	}
}
