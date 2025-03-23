package handler

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

const ThreadsAPIURL = "https://www.threads.net/api/graphql"

type PublicAPI struct {
	APPToken string
}

func prepareHeaders(c *gin.Context, token string) {
	for k, v := range defaultHeaders(token) {
		c.Request.Header.Set(k, v)
	}
}

func defaultHeaders(token string) map[string]string {
	return map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.9",
		"Cache-Control":   "no-cache",
		"Content-Type":    "application/x-www-form-urlencoded",
	}
}

// Generate and announce payment using TTS
func announcePayment(payment PaymentNotification) {
	// Create the announcement text
	text := fmt.Sprintf("Payment received: %.2f %s",
		payment.Amount, payment.Currency)

	// Choose TTS command based on OS
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("say", text)
	case "linux": // Linux
		cmd = exec.Command("espeak", text)
	case "windows": // Windows
		// Note: Windows doesn't have a built-in TTS command like this.
		// You could use 'powershell -Command "Add-Type –AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak(\"" + text + "\")"'
		cmd = exec.Command("powershell", "-Command",
			fmt.Sprintf("Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s')", text))
	default:
		log.Printf("Unsupported OS for TTS: %s, falling back to console output", runtime.GOOS)
		fmt.Println(text)
		return
	}

	// Execute the TTS command
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run TTS command: %v", err)
		// Fallback to console output
		fmt.Println(text)
		return
	}
}
