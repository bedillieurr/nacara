package main

import (
    "fmt"
    "github.com/playwright-community/playwright-go"
)

func headlessBrowserScan(url string) {
    pw, err := playwright.Run()
    if err != nil {
        fmt.Println("Failed to start Playwright:", err)
        return
    }
    defer pw.Stop()

    browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(true),
    })
    if err != nil {
        fmt.Println("Failed to launch browser:", err)
        return
    }
    defer browser.Close()

    page, err := browser.NewPage()
    if err != nil {
        fmt.Println("Failed to create page:", err)
        return
    }

    if err := page.Goto(url); err != nil {
        fmt.Println("Failed to visit the page:", err)
        return
    }

    // Deteksi XSS
    page.Fill("input[name='username']", "<script>alert('XSS')</script>")
    page.Click("button[type='submit']")
    if err := page.WaitForSelector("text='XSS'"); err == nil {
        fmt.Println("[HIGH] Potential XSS vulnerability detected!")
    }
}

