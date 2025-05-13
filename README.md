# Forensic Timeline and Event Correlation Dashboard

This Go project builds a dashboard that aggregates logs from multiple sources (endpoint, network, cloud, and application), normalizes timestamps, and correlates events into a chronological timeline for forensic analysis. It also provides an automated JSON report outlining key forensic findings.

## Key Features
- **Log Aggregation:** Loads logs from a CSV file.
- **Timestamp Normalization:** Converts and sorts timestamps.
- **Timeline Correlation:** Uses simple rule-based logic for ordering events.
- **Interactive Dashboard:** Displays events on a Plotly.js timeline.
- **Automated Reporting:** Outputs a JSON report summarizing the timeline.

## How to Run
1. Install [Go](https://golang.org/dl/) (version 1.16+ recommended).
2. In the repository directory, run: go mod tidy go run main.go
3. Open your browser and navigate to [http://localhost:8080/dashboard](http://localhost:8080/dashboard) to view the dashboard.
4. Visit [http://localhost:8080/api/logs](http://localhost:8080/api/logs) for JSON log data or [http://localhost:8080/api/report](http://localhost:8080/api/report) for the incident report.

## Repository Structure
```
forensic-timeline-dashboard/
├── README.md
├── go.mod
├── data/
│   └── logs.csv
├── templates/
│   └── dashboard.html
└── main.go
```
