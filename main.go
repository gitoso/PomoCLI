package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

const (
	defaultMinutes = 45

	// ANSI colors
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorBold   = "\033[1m"

	// ANSI controls
	clearScreen = "\033[2J"
	moveCursor  = "\033[H"
	hideCursor  = "\033[?25l"
	showCursor  = "\033[?25h"
)

func main() {
	minutes := defaultMinutes

	if len(os.Args) > 1 {
		m, err := strconv.Atoi(os.Args[1])
		if err != nil || m <= 0 {
			fmt.Printf("%sError: Please provide a valid positive number of minutes%s\n", colorRed, colorReset)
			fmt.Printf("Usage: pomocli [minutes]\n")
			os.Exit(1)
		}
		minutes = m
	}

	runCountdown(minutes)
}

func runCountdown(minutes int) {
	totalSeconds := minutes * 60
	remaining := totalSeconds

	// Hide cursor during countdown
	fmt.Print(hideCursor)
	defer fmt.Print(showCursor)

	for remaining >= 0 {
		displayTimer(remaining, totalSeconds)
		if remaining == 0 {
			break
		}
		time.Sleep(1 * time.Second)
		remaining--
	}

	// Timer complete
	displayComplete()
	playNotificationSound()
}

func displayTimer(remaining, total int) {
	mins := remaining / 60
	secs := remaining % 60

	progress := float64(total-remaining) / float64(total)
	barWidth := 40
	filled := int(progress * float64(barWidth))

	// Build progress bar
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	// Choose color based on remaining time
	var timeColor string
	percentRemaining := float64(remaining) / float64(total) * 100
	switch {
	case percentRemaining > 50:
		timeColor = colorGreen
	case percentRemaining > 20:
		timeColor = colorYellow
	default:
		timeColor = colorRed
	}

	fmt.Print(clearScreen + moveCursor)

	// Display
	fmt.Println()
	fmt.Printf("  %s%s╔══════════════════════════════════════════════╗%s\n", colorBold, colorCyan, colorReset)
	fmt.Printf("  %s%s║            🍅  POMOCLI  🍅                   ║%s\n", colorBold, colorCyan, colorReset)
	fmt.Printf("  %s%s╚══════════════════════════════════════════════╝%s\n", colorBold, colorCyan, colorReset)
	fmt.Println()
	fmt.Println()

	// Big time display
	fmt.Printf("              %s%s%02d:%02d%s\n", colorBold, timeColor, mins, secs, colorReset)
	fmt.Println()

	// Progress bar
	fmt.Printf("      %s[%s]%s\n", colorWhite, bar, colorReset)
	fmt.Printf("      %s%.1f%% complete%s\n", colorPurple, progress*100, colorReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("  %sPress Ctrl+C to cancel%s\n", colorBlue, colorReset)
}

func displayComplete() {
	fmt.Print(clearScreen + moveCursor)
	fmt.Println()
	fmt.Printf("  %s%s╔══════════════════════════════════════════════╗%s\n", colorBold, colorGreen, colorReset)
	fmt.Printf("  %s%s║            🍅  POMOCLI  🍅                   ║%s\n", colorBold, colorGreen, colorReset)
	fmt.Printf("  %s%s╚══════════════════════════════════════════════╝%s\n", colorBold, colorGreen, colorReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("        %s%s✓ TIME'S UP! ✓%s\n", colorBold, colorGreen, colorReset)
	fmt.Println()
	fmt.Printf("    %sTake a break, you've earned it!%s\n", colorYellow, colorReset)
	fmt.Println()
	fmt.Println()
}

func playNotificationSound() {
	// Terminal bell
	fmt.Print("\a")

	// Platform-specific sound
	switch runtime.GOOS {
	case "darwin":
		// macOS: play system sound
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
	case "linux":
		// Linux: try paplay or aplay
		if err := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga").Run(); err != nil {
			exec.Command("aplay", "/usr/share/sounds/sound-icons/bell.wav").Run()
		}
	case "windows":
		// Windows: use PowerShell to play system sound
		exec.Command("powershell", "-c", "[System.Media.SystemSounds]::Asterisk.Play()").Run()
	}
}
