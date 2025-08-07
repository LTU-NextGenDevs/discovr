# Discovr – Asset Discovery Tool

Discovr is a lightweight, portable asset discovery tool designed to rapidly identify and map digital assets across traditional and cloud infrastructure (AWS and Azure). Built to assist cybersecurity teams in environments with limited visibility, it enables fast deployment of security tooling by revealing all reachable systems—particularly those suitable for agent deployment.

## Purpose

Developed in collaboration with Triskele Labs as part of a university capstone project, Discovr fills a market gap where existing tools are either too complex, incomplete, or commercially restrictive.

## Key Features

- **Portable**: Runs as a single binary on Windows and Linux with minimal setup.
- **Comprehensive Scanning**: Active and passive discovery using tools like Nmap.
- **Cloud-Aware**: Detects cloud-hosted VMs (AWS, Azure) using runtime credentials.
- **Host Details**: Exports asset details such as IP, hostname, OS, and role.
- **Open Source**: Built using open-source tools and libraries—no commercial dependencies.
- **User-Friendly**: Simple CLI (with optional minimal GUI or web-based report).

## Project Structure

```bash
├── cmd/
│   ├── root.go
│   └── scan.go
├── internal/
│   └── scan/
│       ├── csv.go
│       ├── nmap.go
│       └── parser.go
 |── go.mod
├── go.sum
├── main.go
└── README.md
```

## Installation

Discovr is written in Go. To build from source:

```bash
git clone https://github.com/LTU-NextGenDevs/discovr.git
cd discovr
go build -o discovr
```

To run:
```bash
go run main.go scan
```

## Usage

```bash
# Scan local network with default settings
discovr scan

# Specify subnet to scan
discovr scan --target 10.0.0.0/24
