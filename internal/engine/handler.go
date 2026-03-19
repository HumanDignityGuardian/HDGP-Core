package engine

import (
	"encoding/json"
	"io"
	"net/http"
)

const engineVersion = "0.1.0-mvp"

// maxEvaluateRequestBytes limits request body size (SH-A3: prevent DoS via huge JSON).
const maxEvaluateRequestBytes = 1 << 20 // 1 MiB

// NewEvaluateHandler returns an HTTP handler for /hdgp/v1/evaluate.
func NewEvaluateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body := io.LimitReader(r.Body, maxEvaluateRequestBytes)
		decoder := json.NewDecoder(body)
		decoder.DisallowUnknownFields()

		var req EvaluateRequest
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		applyMetaDefaults(&req)

		resp := Evaluate(req)
		recordAudit(req, resp)

		if r.URL.Query().Get("include_report") == "true" || r.URL.Query().Get("include_report") == "1" {
			resp.Report = GenerateReport(&req, &resp)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}



