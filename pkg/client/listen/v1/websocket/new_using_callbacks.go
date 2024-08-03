// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	"context"

	klog "k8s.io/klog/v2"

	live "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket"
	msginterfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket/interfaces"
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
)

/*
NewForDemo creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY

Deprecated: Use NewUsingCallbackForDemo instead
*/
func NewForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*Client, error) {
	return NewUsingCallbackForDemo(ctx, options)
}

/*
NewForDemo creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewUsingCallbackForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*Client, error) {
	return New(ctx, "", &interfaces.ClientOptions{}, options, nil)
}

/*
NewWithDefaults creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The callback handler is set to the default handler which just prints all messages to the console

Deprecated: Use NewUsingCallbackWithDefaults instead
*/
func NewWithDefaults(ctx context.Context, options *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*Client, error) {
	return NewUsingCallback(ctx, "", &interfaces.ClientOptions{}, options, callback)
}

/*
NewWithDefaults creates a new websocket connection with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The callback handler is set to the default handler which just prints all messages to the console
*/
func NewUsingCallbackWithDefaults(ctx context.Context, options *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*Client, error) {
	return New(ctx, "", &interfaces.ClientOptions{}, options, callback)
}

/*
New creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription

Deprecated: Use NewUsingCallback instead
*/
func New(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*Client, error) {
	return NewUsingCallback(ctx, apiKey, cOptions, tOptions, callback)
}

/*
New creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription
*/
func NewUsingCallback(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*Client, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	return NewUsingCallbackWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, callback)
}

/*
New creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- ctxCancel: allow passing in own cancel
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription
*/
func NewUsingCallbackWithCancel(ctx context.Context, ctxCancel context.CancelFunc, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*Client, error) {
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

	if callback == nil {
		klog.V(2).Infof("Using DefaultCallbackHandler.\n")
		callback = live.NewDefaultCallbackHandler()
	}

	// init
	var router msginterfaces.Router
	router = live.NewCallbackRouter(callback)

	conn := Client{
		cOptions:  cOptions,
		tOptions:  tOptions,
		sendBuf:   make(chan []byte, 1),
		callback:  callback,
		router:    &router,
		ctx:       ctx,
		ctxCancel: ctxCancel,
		retry:     true,
	}

	klog.V(3).Infof("NewDeepGramWSClient Succeeded\n")
	klog.V(6).Infof("live.New() LEAVE\n")

	return &conn, nil
}
