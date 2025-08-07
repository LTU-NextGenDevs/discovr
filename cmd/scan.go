package cmd

import (
	"capstone/internal/scan"
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

var target string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a target network for assets",
	Run: func(cmd *cobra.Command, args []string) {
		if target == "" {
			detected, err := getLocalSubnet()
			if err != nil {
				fmt.Println("Could not detect local subnet:", err)
				return
			}
			fmt.Println("No target provided. Using detected subnet:", detected)
			target = detected
		} else {
			fmt.Println("Scanning user-defined target:", target)
		}

		xmlOutput, err := scan.RunNmapScan(target)
		if err != nil {
			fmt.Println("Scan failed:", err)
			return
		}

		result, err := scan.ParseNmapXML(xmlOutput)
		if err != nil {
			fmt.Println("Failed to parse XML:", err)
			return
		}

		for _, host := range result.Hosts {
			if host.Status.State != "up" {
				continue
			}

			ip := "unknown"
			if len(host.Addresses) > 0 {
				ip = host.Addresses[0].Addr
			}

			name := "unknown"
			if len(host.Hostnames) > 0 {
				name = host.Hostnames[0].Name
			}

			fmt.Printf("Host: %s (%s)\n", ip, name)

			for _, port := range host.Ports {
				if port.State.State == "open" {
					fmt.Printf("   - Port %d (%s): %s\n", port.PortID, port.Protocol, port.Service.Name)
				}
			}

			err = scan.WriteResultsToCSV("scan_results1.csv", result.Hosts)
			if err != nil {
				fmt.Println("Failed to write CSV:", err)
				return
			}

			fmt.Println("Results saved to scan_results.csv")
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&target, "target", "t", "", "Target IP range or subnet (e.g., 192.168.1.0/24)")
	//scanCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(scanCmd)
}

// getLocalSubnet returns the detected local subnet in CIDR form (e.g. 192.168.1.0/24)
func getLocalSubnet() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Skip down or loopback interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			ipParts := strings.Split(ip.String(), ".")
			if len(ipParts) != 4 {
				continue
			}

			// Construct /24 subnet from first 3 octets
			return fmt.Sprintf("%s.%s.%s.0/24", ipParts[0], ipParts[1], ipParts[2]), nil
		}
	}

	return "", fmt.Errorf("unable to detect local subnet")
}

//
