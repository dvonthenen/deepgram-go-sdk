// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	"context"

	klog "k8s.io/klog/v2"

	websocket "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket"
	msginterfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket/interfaces"
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
)

/*
NewForDemo creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewUsingChanForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*Client, error) {
	return NewUsingChan(ctx, "", &interfaces.ClientOptions{}, options, nil)
}

/*
NewWithDefaults creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The chans handler is set to the default handler which just prints all messages to the console
*/
func NewUsingChanWithDefaults(ctx context.Context, options *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*Client, error) {
	return NewUsingChan(ctx, "", &interfaces.ClientOptions{}, options, chans)
}

/*
New creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- chans: LiveMessageCallback which is a chans that allows you to perform actions based on the transcription
*/
func NewUsingChan(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*Client, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	return NewUsingChanWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, chans)
}

/*
New creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- ctxCancel: allow passing in own cancel
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- chans: LiveMessageCallback which is a chans that allows you to perform actions based on the transcription
*/
func NewUsingChanWithCancel(ctx context.Context, ctxCancel context.CancelFunc, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*Client, error) {
	klog.V(6).Infof("live.New() ENTER\n")

	if apiKey != "" {
		cOptions.APIKey = apiKey
	}
	err := cOptions.Parse()
	if err != nil {
		klog.V(1).Infof("ClientOptions.Parse() failed. Err: %v\n", err)
		return nil, err
	}
	err = tOptions.Check()
	if err != nil {
		klog.V(1).Infof("TranscribeOptions.Check() failed. Err: %v\n", err)
		return nil, err
	}

	if chans == nil {
		klog.V(2).Infof("Using DefaultCallbackHandler.\n")
		defaultHandler := websocket.NewDefaultChanHandler()
		var handler msginterfaces.LiveMessageChan
		handler = defaultHandler
		chans = &handler
	}

	// init
	var router msginterfaces.Router
	router = websocket.NewChanRouter(chans)

	conn := Client{
		cOptions:  cOptions,
		tOptions:  tOptions,
		sendBuf:   make(chan []byte, 1),
		chans:     make([]*msginterfaces.LiveMessageChan, 0),
		router:    &router,
		ctx:       ctx,
		ctxCancel: ctxCancel,
		retry:     true,
	}

	klog.V(3).Infof("NewDeepGramWSClient Succeeded\n")
	klog.V(6).Infof("live.New() LEAVE\n")

	return &conn, nil
}
