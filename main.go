package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "time"
)

type Config struct {
    URL string
}

func parseFlags() (*Config, error) {
    config := &Config{}
    flag.StringVar(&config.URL, "url", "", "URL to send request to")
    flag.Parse() 
    
    if config.URL == "" {
        return nil, fmt.Errorf("Error: URL is required")
    }
    return config, nil
}

func makeRequest(config *Config) error {
    currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
    fmt.Printf("Current Date and Time (UTC): %s\n", currentTime)
    
    currentUser := os.Getenv("USER")
    fmt.Printf("Current User: %s\n\n", currentUser)

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    req, err := http.NewRequest("GET", config.URL, nil)
    if err != nil {
        return fmt.Errorf("Error creating request: %v", err)
    }

    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Error sending request: %v", err)
    }
    defer resp.Body.Close()

    fmt.Printf("Response Status: %s\n", resp.Status)
    fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
    return nil
}

func main() {
    config, err := parseFlags()
    if err != nil {
        fmt.Println(err)
        flag.Usage()
        os.Exit(1)
    }

    fmt.Printf("Making request to: %s\n\n", config.URL)

    if err := makeRequest(config); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}