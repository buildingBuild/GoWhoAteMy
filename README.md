# GoWhoAteMyCPU

GoWhoAteMyCPU is a CLI tool that helps answer a simple question:

**who is eating my CPU right now?**

When your system starts lagging, tools like `top` or Activity Monitor give you raw metrics  but they don’t tell you what actually caused the problem. CPU stats, logs, and system metrics get annoying when its just numbers

This tool GOs through it (pun intended) 

---

## Features

* real-time process monitoring (CPU + memory + networking)
* tracks process behavior over time 
* detects CPU spikes and abnormal usage
* flags processes with steady memory growth
* identifies the main process causing slowdown
* sends notifications when abnormal usage is detected
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


This project is being built for the  
[Rebuilding the OS: Core System Utilities Hackathon](https://bitbuilders-code-race-apr-2026.devpost.com/?_gl=1*1lql3uk*_gcl_au*MTA5ODE1NDk0LjE3NzU2Nzk3MTU.*_ga*MTQ2NDUzMTY5NC4xNzc1Njc5NzE1*_ga_0YHJK3Y10M*czE3NzU2ODk1NzckbzMkZzEkdDE3NzU2OTI2NTIkajE2JGwwJGgw)
