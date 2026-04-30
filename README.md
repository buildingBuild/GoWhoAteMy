# GoWhoAteMy

GoWhoAteMy is a CLI tool that helps answer a simple question:

**what is slowing my computer down right now?**

When your system starts lagging, you usually have to jump between a bunch of tools: `top` for CPU, Activity Monitor for memory, network commands for ports, and logs for clues. Even then, you mostly get dumped raw numbers instead of a clear answer about what changed and which process is responsible.

This tool GOs through it (pun intended) 

---

## Features

* real-time process monitoring for CPU, memory, and networking
* tracks process behavior over time
* detects CPU spikes, memory growth, high swap usage, and network changes
* flags abnormal things like sudden resource jumps and new listening ports
* identifies the main process causing slowdown
* notification system for important alerts
* simple, readable explanations instead of raw stats
* lightweight CLI

---

## The idea

System tools today are good at showing data,
but bad at explaining it.

GoWhoAteMyCPU focuses on:

* what changed
* how quickly it changed
* and what actually caused the slowdown

---

## How it works (high level)

* collect system + process metrics
* keep a short rolling history
* compare current values with recent behavior
* detect patterns like spikes and pressure
* output a simple explanation
* send out important notifications
(Fully supports mac)

---

## Tech stack

* Go
* gopsutil for CPU, memory, disk, process, and network metrics
* macOS `osascript` for system notifications
* Kong for cli management

---

## How to run

```bash
go build -o gowhoatemy .
./gowhoatemy --computer
```

To install it somewhere on your `PATH`, build it into a path folder like `~/bin`:

```bash
go build -o ~/bin/gowhoatemy .
```

Then run it from anywhere:

```bash
gowhoatemy --computer
```

Available flags:

```bash
gowhoatemy --cpu                         # shows CPU metrics
gowhoatemy --memory                      # shows memory metrics
gowhoatemy --network                     # shows network metrics
gowhoatemy --computer                    # shows all computer metrics and sends notifications every 10 seconds
gowhoatemy --computer --interval 20      # shows all computer metrics and checks every 20 seconds
gowhoatemy --help                        # shows available flags
```


This project is being built for the  
[Rebuilding the OS: Core System Utilities Hackathon](https://bitbuilders-code-race-apr-2026.devpost.com/?_gl=1*1lql3uk*_gcl_au*MTA5ODE1NDk0LjE3NzU2Nzk3MTU.*_ga*MTQ2NDUzMTY5NC4xNzc1Njc5NzE1*_ga_0YHJK3Y10M*czE3NzU2ODk1NzckbzMkZzEkdDE3NzU2OTI2NTIkajE2JGwwJGgw)
