# Port Scanner
[![License](https://img.shields.io/github/license/whitehatboy005/Blockchain-Based-Voting-System)](LICENSE.md)

This project is a simple command-line tool written in Go that scans ports on a target domain or IP address. It can scan a range of ports, check which ports are open, and record the results to an output file. It also supports scanning a list of domains from a file.

## Features

- **Port Scanning**: Scan a specific range of ports on a target domain or IP address.
- **Service Detection**: Identifies common services running on open ports such as FTP, HTTP, SSH, etc.
- **Domain List Scanning**: Scan multiple domains from a list stored in a file.
- **Output to File**: Save results to a file in a readable format.
- **Parallel Scanning**: Uses goroutines to scan multiple ports concurrently for faster results.

## Requirements

- Go 1.18 or higher

## Installation

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/whitehatboy005/Port-Scanner.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Port-Scanner
   ```
3. Build the project:
   ```bash
   go build -o portscan portscan.go
   ```
4. Make it executable
   ```bash
    sudo chmod +x portscan
   ```
5. Move the binary to /usr/local/bin for global access
   ```bash
   sudo mv portscan /usr/local/bin/
   ```
## Usage

You can run the port scanner using the following command-line flags:

### Flags

- `-t <target>`: Target domain or IP address to scan.
- `-s <start port>`: Start port number (default: 1).
- `-e <end port>`: End port number (default: 10000).
- `-o <output file>`: Output file to save the results (default: `results.txt`).
- `-l <file>`: File with a list of domains to scan.
- `-h`: Display help menu.
#
![Screenshot 2024-12-29 110115](https://github.com/user-attachments/assets/ae8d472a-c1bb-437e-920c-b06a441acb57)
## Examples

1. **Scan a single target (example: google.com) from port 1 to 1000 and save the results to `results.txt`:**

   ```bash
   portscan -t google.com -s 1 -e 1000 -o results.txt
   ```
2. **Scan a list of domains from domains.txt file from port 1 to 1000 and save the results to results.txt:**
   ```bash
   portscan -l domains.txt -s 1 -e 1000 -o results.txt
   ```
## Example output
  ```
     Results for google.com:
     Port 80 - HTTP
     Port 443 - HTTPS
     Scanning completed in 10.02 seconds
  ```
#
## License

This project is licensed under the terms of the [MIT license](LICENSE.md).
