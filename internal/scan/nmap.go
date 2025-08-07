package scan

import (
	"fmt"
	"os/exec"
)

// RunNmapScan runs an Nmap scan on the target subnet/IP
// and returns the XML output as a byte slice
func RunNmapScan(target string) ([]byte, error) {
	// Use -sT instead of -sS and REMOVE -O
	cmd := exec.Command("nmap", "-sT", "-T4", target, "-oX", "-")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("nmap error: %w\nOutput: %s", err, string(output))
	}

	return output, nil
}

//
