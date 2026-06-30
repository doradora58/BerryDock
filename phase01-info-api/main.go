package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	CPUModel     string `json:"cpu_model"`
	CPUCount     int    `json:"cpu_count"`
}

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

func getCPUModel() string {
	cpuInfo := readFile("/proc/cpuinfo")
	for _, line := range strings.Split(cpuInfo, "\n") {
		if strings.HasPrefix(line, "model name") || strings.HasPrefix(line, "Hardware") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

func getCPUCount() int {
	return runtime.NumCPU()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := SystemInfo{
		Hostname:     readFile("/etc/hostname"),
		OS:           readFile("/etc/os-release"),
		Architecture: runtime.GOARCH,
		CPUModel:     getCPUModel(),
		CPUCount:     getCPUCount(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
