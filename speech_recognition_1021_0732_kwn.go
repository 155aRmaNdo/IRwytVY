// 代码生成时间: 2025-10-21 07:32:45
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/envy"
    "log"
    "os"
    "fmt"
    "github.com/googleapis/google-cloud-go/speech/v1p1beta1"
    "golang.org/x/net/context"
    "google.golang.org/api/option"
)

// SpeechRecognitionService represents the speech recognition system service
type SpeechRecognitionService struct {
    Client *speechpb.SpeechClient
}

// NewSpeechRecognitionService initializes a new speech recognition service
func NewSpeechRecognitionService(ctx context.Context) (*SpeechRecognitionService, error) {
    client, err := speechpb.NewSpeechClient(ctx, option.WithCredentialsFile(envy.Get("GOOGLE_APPLICATION_CREDENTIALS")))
    if err != nil {
        return nil, fmt.Errorf("speechClient.NewSpeechClient: %v", err)
    }
    return &SpeechRecognitionService{Client: client}, nil
}

// RecognizeAudio performs speech recognition on the provided audio file
func (s *SpeechRecognitionService) RecognizeAudio(audioFile string) (string, error)
{
    // Open the audio file
    file, err := os.Open(audioFile)
    if err != nil {
        return "", fmt.Errorf("os.Open: %v", err)
    }
    defer file.Close()

    // Create a config for speech recognition
    config := &speechpb.RecognitionConfig{
        Encoding:        speechpb.RecognitionConfig_LINEAR16,
       SampleRateHertz: 16000,
        LanguageCode:    "en-US",
       EnableAutomaticPunctuation: true,
    }

    // Create a request for the speech recognition service
    request := &speechpb.RecognizeRequest{
        Config:  config,
        Audio:   &speechpb.RecognitionAudio{Uri: audioFile},
    }

    // Call the speech recognition service
    response, err := s.Client.Recognize(context.Background(), request)
    if err != nil {
        return "", fmt.Errorf("SpeechClient.Recognize: %v", err)
    }

    // Extract the transcription from the response
    transcription := ""
    for _, result := range response.Results {
        for _, alternative := range result.Alternatives {
            transcription += alternative.Transcript
        }
    }

    return transcription, nil
}

// App is the main application struct
type App struct{
    *buffalo.App
}

// New initializes a new application
func New() *App {
    app := buffalo.New(buffalo.Options{
       PreWares: []buffalo.PreWare{middleware.Logger},
       // Additional middleware stacks here
    })
    return &App{App: app}
}

// Start runs the application
func (app *App) Start(address string) error {
    app.ServeFiles("/assets/", assetsPath)
    app.GET("/", HomeHandler)
    app.POST("/recognize", RecognizeHandler)
    return app.Serve(address)
}

// HomeHandler is the handler for the home page
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("index.html"))
}

// RecognizeHandler handles the POST request for speech recognition
func RecognizeHandler(c buffalo.Context) error {
    // Extract the audio file path from the request
    audioFile := c.Param("audioFile")

    // Initialize the speech recognition service
    service, err := NewSpeechRecognitionService(c.Request.Context())
    if err != nil {
        return c.Error(500, err)
    }

    // Perform speech recognition
    transcription, err := service.RecognizeAudio(audioFile)
    if err != nil {
        return c.Error(500, err)
    }

    // Return the transcription as JSON
    return c.Render(200, r.JSON(map[string]string{"transcription": transcription}))
}

// main entry point for the application
func main() {
    app := New()
    if err := app.Start(":3000"); err != nil {
        log.Fatal(err)
    }
}
