package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "sync"
    "time"
    "net" // used for scanning ports
)

// Define port names for known services
var portNames = map[int]string{
    21:   "FTP",
    22:   "SSH",
    23:   "Telnet",
    25:   "SMTP",
    53:   "DNS",
    80:   "HTTP",
    443:  "HTTPS",
    3306: "MySQL",
    5432: "PostgreSQL",
    8080: "HTTP Proxy",
    3389: "RDP",
}

// Function to scan a specific port
func scanPort(target string, port int, wg *sync.WaitGroup, resultChan chan string) {
    defer wg.Done()

    // Format the address for the target with the port
    address := fmt.Sprintf("%s:%d", target, port)

    // Attempt to connect to the target on the specified port
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err != nil {
        // If there's an error, it means the port is closed or unreachable
        return
    }
    defer conn.Close()

    // If successful, the port is open, so send the result to the channel
    portName, exists := portNames[port]
    if !exists {
        portName = "Unknown"
    }
    resultChan <- fmt.Sprintf("Port %d - %s", port, portName)
}

// Function to scan ports in a given range and output results to a file
func scanPorts(target string, startPort int, endPort int, outputFile string) {
    var wg sync.WaitGroup
    resultChan := make(chan string)

    // Open output file for writing (overwrite mode)
    file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Write the results header to the file
    writer := bufio.NewWriter(file)
    writer.WriteString(fmt.Sprintf("Results for %s:\n", target))

    // Start scanning
    startTime := time.Now() // Capture start time

    // Perform port scanning
    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(target, port, &wg, resultChan)
    }

    // Wait for all scanning goroutines to complete
    go func() {
        wg.Wait()
        close(resultChan)
    }()

    // Collect open ports
    var openPorts []string
    for result := range resultChan {
        openPorts = append(openPorts, result)
        writer.WriteString(result + "\n") // Write to the file as well
    }

    // Calculate the time taken for scanning
    elapsedTime := time.Since(startTime)

    // Display the scanning completion time and open ports in the terminal
    fmt.Printf("Scanning target: %s\n", target)
    if len(openPorts) > 0 {
        fmt.Println("Open Ports:")
        for _, result := range openPorts {
            fmt.Println(result)
        }
    } else {
        fmt.Println("No open ports found.")
        writer.WriteString("No open ports found.\n") // Write this message to the file as well
    }
    fmt.Printf("Scanning completed in %s\n", elapsedTime)

    // Add a blank line between results of different scans
    writer.WriteString("\n") // Blank line after each scan

    // Flush the buffered writer to the file
    writer.Flush()
}

// Function to scan a list of domains from a file
func scanDomainList(fileName string, startPort int, endPort int, outputFile string) {
    // Open domain list file
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening domain list file:", err)
        return
    }
    defer file.Close()

    // Read domains from the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        domain := scanner.Text()
        if domain != "" {
            // Scan each domain
            fmt.Printf("Scanning target: %s\n", domain)
            scanPorts(domain, startPort, endPort, outputFile)
        }
    }

    // Check for errors reading the file
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading domain list:", err)
    }
}

// Main function to handle command-line arguments and start the scan
func main() {
    // Define flags for command-line arguments
    targetPtr := flag.String("t", "", "Target domain or IP address to scan")
    startPortPtr := flag.Int("s", 1, "Start port (default: 1)")
    endPortPtr := flag.Int("e", 10000, "End port (default: 10000)")
    outputFilePtr := flag.String("o", "results.txt", "Output file to save results (default: results.txt)")
    listFilePtr := flag.String("l", "", "File with list of domains to scan")
    helpPtr := flag.Bool("h", false, "Display help menu")

    // Parse the command-line arguments
    flag.Parse()

    // If help flag is set, show the usage
if *helpPtr {
    fmt.Println("Usage:")
    fmt.Println("  -t <target>   : Target domain or IP address to scan")
    fmt.Println("  -s <start port>: Start port number (default: 1)")
    fmt.Println("  -e <end port> : End port number (default: 10000)")
    fmt.Println("  -o <output>   : Output file to save the results (default: results.txt)")
    fmt.Println("  -l <file>     : File with list of domains to scan")
    fmt.Println("Examples:")
    fmt.Println("  portscan -t google.com -s 1 -e 1000 -o results.txt")
    fmt.Println("  portscan -l domains.txt -s 1 -e 1000 -o results.txt")
    return
}

    // If -l flag is provided, scan the list of domains
    if *listFilePtr != "" {
        scanDomainList(*listFilePtr, *startPortPtr, *endPortPtr, *outputFilePtr)
    } else if *targetPtr != "" {
        // Otherwise, scan the single target
        scanPorts(*targetPtr, *startPortPtr, *endPortPtr, *outputFilePtr)
    } else {
        fmt.Println("Please provide a target domain or file using -t or -l flag.")
    }

    fmt.Println("All scans complete!")
}
