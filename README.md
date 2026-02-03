# PomoCLI

Yeah, it seems that software is becoming more and more a commodity with the advances in AI. I was studying and wanted a simple Pomodoro timer app. I did what I usually do, opened the browser and started searching for a good app. But I just wanted something really simple, a simple CLI app would suffice, so I ended up opening Claude Code and one-shotting this simple app that does exactly  what I wanted.

## Build

```bash
make build      # Build to ./build/pomocli
make install    # Install to /usr/local/bin
```

## Usage

```bash
pomocli         # 45-minute countdown (default)
pomocli 25      # 25-minute countdown

# Or with make
make run        # Run with default time
make run 25     # Run with 25 minutes
```
