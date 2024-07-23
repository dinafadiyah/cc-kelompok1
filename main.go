package main

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Instantiates a client.
	ctx := context.Background()

	textToSpeechClient, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	http.HandleFunc("/synthesize", synthesizeText())

	http.ListenAndServe(":8080", nil)

}

// /synthesize?text="hello, kami dari kelompok 1"

func synthesizeText() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		keys, ok := request.URL.Query()["text"]
		if !ok {
			log.Fatal("text not found")
		}

		if !ok || len(keys[0]) < 1 {
			log.Panicln("Url param 'key' is missing")
			return
		}

		// Query()["key"] will retuen an array of items,
		// we only want the single item.
		text := keys[0]
	}
}

func main() {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: "Halo, kami dari kelompok 1!"},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "id-ID",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)
}
