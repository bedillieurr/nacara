package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("Starting Nacara Scanner...")

    url := "http://example.com" // Ganti dengan target yang sesuai
    autoTemplateDiscovery(url)
    wafBypass(url)
    apiVulnerabilityScan(url)
    headlessBrowserScan(url)

    fmt.Println("\nScan Complete.")
    fmt.Println("nacara v1 by bedillieurr")
}

