package scan

import (
	"encoding/xml"
)

// NmapRun is the root of the XML tree
type NmapRun struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	Status    Status     `xml:"status"`
	Addresses []Address  `xml:"address"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Ports     []Port     `xml:"ports>port"`
}

type Status struct {
	State string `xml:"state,attr"`
}

type Address struct {
	Addr string `xml:"addr,attr"`
}

type Hostname struct {
	Name string `xml:"name,attr"`
}

type Port struct {
	Protocol string    `xml:"protocol,attr"`
	PortID   int       `xml:"portid,attr"`
	State    PortState `xml:"state"`
	Service  Service   `xml:"service"`
}

type PortState struct {
	State string `xml:"state,attr"`
}

type Service struct {
	Name string `xml:"name,attr"`
}

// ✅ Exported function (capitalized)
func ParseNmapXML(data []byte) (*NmapRun, error) {
	var result NmapRun
	err := xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
