package reporting

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"net/http"
	"time"
)

func SendResultToControlPlane(cert *x509.Certificate) {

	durationDays := int(cert.NotAfter.Sub(cert.NotBefore).Hours() / 24)

	data := map[string]interface{}{
		"cn":        cert.Subject.CommonName,
		"sans":      cert.DNSNames,
		"certUrl":   "https://smallstep.kvcc.edu",
		"not_after": cert.NotAfter.Format(time.RFC3339),
		"duration":  durationDays,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	reporting_url := "https://smallstep.kvcc.edu/records"
	//reporting_url := "http://localhost:4444/records"
	resp, err := http.Post(reporting_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
