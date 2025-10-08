// cmd/hashfunction/main.go
package main

import (
    "flag"
    "log"
    "os"

    "hashfunction/internal/hashfunction"
)

func main() {
    // Define command line flags
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path")
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    // Create a new application instance with logging verbosity
    app := hashfunction.NewApp(*verbose)

    // Run the application with input and output file paths
    if err := app.Run(*input, *output); err != nil {
        // Log the error and exit with a non-zero status code
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}