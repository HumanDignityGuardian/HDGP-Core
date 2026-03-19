package engine

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"
)

// simple in-memory audit log with optional append-only file persistence (TD-A1).
// 当 HDGP_AUDIT_LOG_PATH 非空时，每条审计记录以 NDJSON 行追加写入文件。

type auditEntry struct {
	Timestamp       time.Time        `json:"timestamp"`
	Request         EvaluateRequest  `json:"request"`
	Response        EvaluateResponse `json:"response"`
	IntegrityEvents []IntegrityEvent `json:"integrity_events,omitempty"`
}

var (
	auditMu       sync.Mutex
	auditLogFileMu sync.Mutex // protects auditLogFile writes
	auditLog      []auditEntry
	maxAuditN     = 100
	auditLogPath  = os.Getenv("HDGP_AUDIT_LOG_PATH") // 空则仅内存
	auditLogFile  *os.File
	auditLogInit  sync.Once
)

func openAuditLogFile() (*os.File, error) {
	if auditLogPath == "" {
		return nil, nil
	}
	return os.OpenFile(auditLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}

// recordAudit appends a new audit entry in a bounded in-memory buffer.
// When HDGP_AUDIT_LOG_PATH is set, also appends one NDJSON line to the file (append-only).
func recordAudit(req EvaluateRequest, resp EvaluateResponse) {
	entry := auditEntry{
		Timestamp:       time.Now().UTC(),
		Request:         req,
		Response:        resp,
		IntegrityEvents: resp.IntegrityEvents,
	}

	auditMu.Lock()
	auditLog = append(auditLog, entry)
	if len(auditLog) > maxAuditN {
		auditLog = auditLog[len(auditLog)-maxAuditN:]
	}
	auditMu.Unlock()

	// Append-only file write (TD-A1); failure does not block in-memory recording
	if auditLogPath != "" {
		auditLogInit.Do(func() {
			f, err := openAuditLogFile()
			if err != nil {
				return
			}
			auditLogFile = f
		})
		if auditLogFile != nil {
			line, _ := json.Marshal(entry)
			auditLogFileMu.Lock()
			auditLogFile.Write(line)
			auditLogFile.Write([]byte("\n"))
			auditLogFileMu.Unlock()
		}
	}
}

// NewAuditHandler exposes a read-only view of recent audit entries.
// GET /hdgp/v1/audit?limit=N
func NewAuditHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		limit := clampLimit(r.URL.Query().Get("limit"), maxAuditN, 20)

		auditMu.Lock()
		defer auditMu.Unlock()

		entries := lastEntries(auditLog, limit)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(entries); err != nil {
			http.Error(w, "failed to encode audit log", http.StatusInternalServerError)
			return
		}
	})
}

