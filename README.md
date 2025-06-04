# Goroutine Scheduler 🚦

This project implements a simple job scheduler in Go using goroutines, simulating the core logic of operating system schedulers (e.g., Round-Robin). Each job is run concurrently with a fixed time slot, helping you understand how context switching and cooperative scheduling can be modeled in Go.

## Features
- ✅ Round-robin scheduling
- ✅ Time-slice control
- 🧩 Easily extendable to priority or preemptive scheduling

## Getting Started
```bash
# Clone the repository
git clone https://github.com/fuzayil-baxtiyoriy/goroutine-scheduler.git
cd goroutine-scheduler

# Run the project
go run main.go
```

## File Structure
```
main.go               # Entry point
scheduler/
  ├── job.go          # Job definition
  └── scheduler.go    # Core scheduler logic
go.mod               # Go module file
```

## Example Output
```
Starting Round-Robin Scheduler...
Running job #1
Executing work for job #1
Finished job #1
...
```

## Future Ideas
- [ ] Add priority scheduling
- [ ] Add preemptive scheduling using context cancel
- [ ] Track job stats and lifecycle
- [ ] CLI dashboard with tview

