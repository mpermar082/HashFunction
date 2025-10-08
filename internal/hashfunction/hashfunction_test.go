// internal/hashfunction/hashfunction_test.go
package hashfunction

import (
    "testing"
)

// TestNewApp verifies the NewApp function returns a non-nil app with the expected properties.
func TestNewApp(t *testing.T) {
    // Create a new app with verbosity enabled
    app := NewApp(true)
    if app == nil {
        t.Fatal("NewApp returned nil")
    }
    // Verify the app's verbosity is enabled
    if !app.Verbose {
        t.Error("Expected verbose to be true")
    }
    // Verify ProcessedCount is 0
    if app.ProcessedCount != 0 {
        t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
    }
}

// TestProcess tests the Process function's behavior.
func TestProcess(t *testing.T) {
    // Create a new app with verbosity disabled
    app := NewApp(false)
    // Process the input data
    result, err := app.Process("test data")
    
    if err != nil {
        t.Fatalf("Process returned error: %v", err)
    }
    
    // Verify the result is successful
    if !result.Success {
        t.Error("Expected result.Success to be true")
    }
    
    // Verify ProcessedCount is incremented
    if app.ProcessedCount != 1 {
        t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
    }
}

// TestRun tests the Run function's behavior.
func TestRun(t *testing.T) {
    // Create a new app with verbosity disabled
    app := NewApp(false)
    // Run the app with empty input
    err := app.Run("", "")
    
    if err != nil {
        t.Fatalf("Run returned error: %v", err)
    }
}