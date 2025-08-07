package scan

import (
	"encoding/csv"
	"os"
	"strconv"
)

// WriteResultsToCSV writes the Nmap results to a file
func WriteResultsToCSV(filename string, hosts []Host) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	writer.Write([]string{"IP Address", "Hostname", "Port", "Protocol", "Service"})

	for _, host := range hosts {
		if host.Status.State != "up" {
			continue
		}

		ip := "unknown"
		if len(host.Addresses) > 0 {
			ip = host.Addresses[0].Addr
		}

		hostname := "unknown"
		if len(host.Hostnames) > 0 {
			hostname = host.Hostnames[0].Name
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				row := []string{
					ip,
					hostname,
					strconv.Itoa(port.PortID),
					port.Protocol,
					port.Service.Name,
				}
				writer.Write(row)
			}
		}
	}

	return nil
}
