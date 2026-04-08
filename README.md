# GoWhoAteMyCPU

GoWhoAteMyCPU is a CLI tool that helps answer a simple question:

**who is eating my CPU right now?**

When your system starts lagging, tools like `top` or Activity Monitor give you raw metrics — but they don’t tell you what actually caused the problem. CPU stats, logs, and system metrics get annoying when its just numbers

This tool GOs through it (pun intended) 

---

## Features

* real-time process monitoring (CPU + memory)
* tracks process behavior over time (not just snapshots)
* detects CPU spikes and abnormal usage
* flags processes with steady memory growth
* identifies the main process causing slowdown
* simple, readable explanations instead of raw stats
* lightweight CLI (no heavy UI, no noise)

---

## Example output

```bash
⚠ slowdown detected

main suspect:
- Chrome (pid 1234)
- cpu jumped from 12% → 78% in 3 seconds

also noticed:
- node process memory steadily increasing

suggestion:
- close some chrome tabs or restart the process
```

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


