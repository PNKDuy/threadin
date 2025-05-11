package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/gin-gonic/gin"
	mp3 "github.com/hajimehoshi/go-mp3"
	oto "github.com/hajimehoshi/oto/v2"
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

// Announce payment using Google Cloud Text-to-Speech
func announcePayment(payment PaymentNotification) {
	// Create the announcement text in Vietnamese
	text := fmt.Sprintf("Đã nhận thanh toán: %.2f %s từ %s",
		payment.Amount, payment.Currency, payment.Payer)

	// Set up Google Cloud TTS client
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create TTS client: %v", err)
		fmt.Println(text) // Fallback to console
		return
	}
	defer client.Close()

	// Configure the TTS request
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "vi-VN",            // Vietnamese
			Name:         "vi-VN-Standard-A", // Female voice (Standard-A), you can also use "vi-VN-Wavenet-A" for higher quality
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	// Send the request to synthesize speech
	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		log.Printf("Failed to synthesize speech: %v", err)
		fmt.Println(text) // Fallback to console
		return
	}

	// Decode the MP3 audio
	mp3Decoder, err := mp3.NewDecoder(io.NopCloser(io.Reader(bytes.NewReader(resp.AudioContent))))
	if err != nil {
		log.Printf("Failed to decode MP3: %v", err)
		fmt.Println(text)
		return
	}

	// Set up audio context
	otoCtx, readyChan, err := oto.NewContext(22050, 1, oto.FormatSignedInt16LE)
	if err != nil {
		log.Printf("Failed to initialize audio context: %v", err)
		fmt.Println(text)
		return
	}
	<-readyChan // Wait for the context to be ready

	// Play the audio
	player := otoCtx.NewPlayer(mp3Decoder)
	defer player.Close()
	player.Play()

	// Wait for playback to finish (calculate duration based on sample rate and bytes)
	for player.IsPlaying() {
		time.Sleep(time.Millisecond * 100)
	}
}
