// Copyright 2023 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

// streaming
import (
	"context"
	"fmt"
	"os"
	"time"

	microphone "github.com/deepgram/deepgram-go-sdk/pkg/audio/microphone"
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
	client "github.com/deepgram/deepgram-go-sdk/pkg/client/live"
)

func main() {
	// init library
	microphone.Initialize()

	// print instructions
	fmt.Print("\n\nPress ENTER to exit!\n\n")

	/*
		DG Streaming API
	*/
	// init library
	client.InitWithDefault()

	// Go context
	ctx := context.Background()

	// client options
	cOptions := interfaces.ClientOptions{
		EnableKeepAlive: true,
	}

	// set the Transcription options
	tOptions := interfaces.LiveTranscriptionOptions{
		Model:       "nova-2",
		Language:    "en-US",
		Punctuate:   true,
		Encoding:    "linear16",
		Channels:    1,
		SampleRate:  16000,
		SmartFormat: true,
		// To get UtteranceEnd, the following must be set:
		InterimResults: true,
		UtteranceEndMs: "1000",
		VadEvents:      true,
	}

	// create a Deepgram client

	dgClient, err := client.New(ctx, "", cOptions, tOptions, nil)
	if err != nil {
		fmt.Println("ERROR creating LiveTranscription connection:", err)
		return
	}

	for i := 0; i < 10; i++ {
		if i > 0 {
			time.Sleep(5 * time.Second)
		}

		if i == 0 {
			fmt.Printf("Starting connection #%d...", i)
		} else {
			fmt.Printf("Restarting connection #%d...", i)
		}

		// connect the websocket to Deepgram
		wsconn := dgClient.Connect()
		if wsconn == nil {
			fmt.Println("Client.Connect failed")
			os.Exit(1)
		}

		/*
			Microphone package
		*/
		// mic stuf
		mic, err := microphone.New(microphone.AudioConfig{
			InputChannels: 1,
			SamplingRate:  16000,
		})
		if err != nil {
			fmt.Printf("Initialize failed. Err: %v\n", err)
			continue
		}

		// start the mic
		err = mic.Start()
		if err != nil {
			fmt.Printf("mic.Start failed. Err: %v\n", err)
			continue
		}

		go func() {
			// feed the microphone stream to the Deepgram client (this is a blocking call)
			mic.Stream(dgClient)
		}()

		// capture audio for 15 seconds
		time.Sleep(15 * time.Second)

		// close mic stream
		err = mic.Stop()
		if err != nil {
			fmt.Printf("mic.Stop failed. Err: %v\n", err)
			os.Exit(1)
		}

		// teardown library
		microphone.Teardown()

		// close DG client
		dgClient.Stop()
	}

	fmt.Printf("Program exiting...\n")

}
