package main

import (
    "encoding/csv"
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "os"
    "sort"
    "strconv"
    "time"
)

// Event represents a single log entry.
type Event struct {
    EventID     int       `json:"EventID"`
    Timestamp   time.Time `json:"Timestamp"`
    Source      string    `json:"Source"`
    Description string    `json:"Description"`
}

// loadLogs loads and parses log data from a CSV file.
func loadLogs(filePath string) ([]Event, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var events []Event
    // Skip the header row
    for i, record := range records {
        if i == 0 {
            continue
        }
        if len(record) < 4 {
            continue
        }
        id, err := strconv.Atoi(record[0])
        if err != nil {
            continue
        }
        ts, err := time.Parse(time.RFC3339, record[1])
        if err != nil {
            continue
        }
        event := Event{
            EventID:     id,
            Timestamp:   ts,
            Source:      record[2],
            Description: record[3],
        }
        events = append(events, event)
    }
    // Sort events by timestamp
    sort.SliceStable(events, func(i, j int) bool {
        return events[i].Timestamp.Before(events[j].Timestamp)
    })
    return events, nil
}

// logsHandler serves the log data as JSON on /api/logs.
func logsHandler(w http.ResponseWriter, r *http.Request) {
    events, err := loadLogs("data/logs.csv")
    if err != nil {
        http.Error(w, "Unable to load logs", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(events)
}

// reportHandler generates a basic incident report on /api/report.
func reportHandler(w http.ResponseWriter, r *http.Request) {
    events, err := loadLogs("data/logs.csv")
    if err != nil {
        http.Error(w, "Unable to load logs", http.StatusInternalServerError)
        return
    }
    report := map[string]interface{}{
        "report_generated": time.Now().Format(time.RFC3339),
        "total_events":     len(events),
        "first_event":      events[0],
        "last_event":       events[len(events)-1],
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}

// dashboardHandler serves the HTML dashboard on /dashboard.
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/dashboard.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func main() {
    http.HandleFunc("/api/logs", logsHandler)
    http.HandleFunc("/api/report", reportHandler)
    http.HandleFunc("/dashboard", dashboardHandler)

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
