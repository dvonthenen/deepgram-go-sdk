// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

/*
This package provides the prerecorded client implementation for the Deepgram API
*/
package listen

import (
	"context"

	msginterfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket/interfaces"
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
	listenv1rest "github.com/deepgram/deepgram-go-sdk/pkg/client/listen/v1/rest"
	listenv1ws "github.com/deepgram/deepgram-go-sdk/pkg/client/listen/v1/websocket"
)

/***********************************/
// REST Client
/***********************************/
const (
	RESTPackageVersion = listenv1rest.PackageVersion
)

// RestClient is an alias for listenv1rest.Client
type RestClient = listenv1rest.Client

/*
NewRESTWithDefaults creates a new analyze/read client with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewRESTWithDefaults() *listenv1rest.Client {
	return listenv1rest.NewWithDefaults()
}

/*
NewREST creates a new prerecorded client with the specified options

Input parameters:
- apiKey: string containing the Deepgram API key
- options: ClientOptions which allows overriding things like hostname, version of the API, etc.
*/
func NewREST(apiKey string, options *interfaces.ClientOptions) *listenv1rest.Client {
	return listenv1rest.New(apiKey, options)
}

/***********************************/
// WebSocket / Streaming / Live
/***********************************/
const (
	WebSocketPackageVersion = listenv1ws.PackageVersion
)

// WebSocketClient is an alias for listenv1ws.Client
type WebSocketClient = listenv1ws.Client

/*
	Using Callbacks
*/
/*
NewWebSocketUsingCallbackForDemo creates a new websocket connection with all default options

Input parameters:
- ctx: context.Context object
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewWebSocketUsingCallbackForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingCallbackForDemo(ctx, options)
}

/*
NewWebSocketUsingCallbackWithDefaults creates a new websocket connection with all default options

Input parameters:
- ctx: context.Context object
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The callback handler is set to the default handler which just prints all messages to the console
*/
func NewWebSocketUsingCallbackWithDefaults(ctx context.Context, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingCallbackWithDefaults(ctx, tOptions, callback)
}

/*
NewWebSocket creates a new websocket connection with the specified options

Input parameters:
- ctx: context.Context object
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription

Notes:
  - If apiKey is an empty string, the Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The callback handler is set to the default handler which just prints all messages to the console
*/
func NewWebSocketUsingCallback(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	return listenv1ws.NewUsingCallbackWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, callback)
}

/*
NewWebSocketWithCancel creates a new websocket connection but has facilities to BYOC (Bring Your Own Cancel)

Input parameters:
- ctx: context.Context object
- ctxCancel: allow passing in own cancel
- apiKey: string containing the Deepgram API key
- cOptions: ClientOptions which allows overriding things like hostname, version of the API, etc.
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.
- callback: LiveMessageCallback which is a callback that allows you to perform actions based on the transcription

Notes:
  - If apiKey is an empty string, the Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The callback handler is set to the default handler which just prints all messages to the console
*/
func NewWebSocketUsingCallbackWithCancel(ctx context.Context, ctxCancel context.CancelFunc, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingCallbackWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, callback)
}

/*
	Using Channels
*/
/*
NewWebSocketUsingChanForDemo creates a new websocket connection for demo purposes only

Input parameters:
- ctx: context.Context object
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewWebSocketUsingChanForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingChanForDemo(ctx, options)
}

/*
NewWebSocketUsingChanWithDefaults creates a new websocket connection with all default options

Input parameters:
- ctx: context.Context object
- tOptions: LiveTranscriptionOptions which allows overriding things like language, model, etc.

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
  - The chans handler is set to the default handler which just prints all messages to the console
*/
func NewWebSocketUsingChanWithDefaults(ctx context.Context, options *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingChanWithDefaults(ctx, options, chans)
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
func NewWebSocketUsingChan(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*listenv1ws.Client, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	return listenv1ws.NewUsingChanWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, chans)
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
func NewWebSocketUsingChanWithCancel(ctx context.Context, ctxCancel context.CancelFunc, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, chans *msginterfaces.LiveMessageChan) (*listenv1ws.Client, error) {
	return listenv1ws.NewUsingChanWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, chans)
}

/***********************************/
// Deprecated (THESE WILL STILL WORK,
// BUT WILL BE REMOVED IN A FUTURE RELEASE)
/***********************************/

/***********************************/
// REST Client
/***********************************/
// PreRecordedClient is an alias for listenv1rest.Client
//
// Deprecated: This package is deprecated. Use RestClient instead. This will be removed in a future release.
type PreRecordedClient = listenv1rest.Client

// NewPreRecordedWithDefaults is an alias for NewRESTWithDefaults
//
// Deprecated: This package is deprecated. Use NewRESTWithDefaults instead. This will be removed in a future release.
var NewPreRecordedWithDefaults = NewRESTWithDefaults

// NewPreRecorded is an alias for NewREST
//
// Deprecated: This package is deprecated. Use NewREST instead. This will be removed in a future release.
var NewPreRecorded = NewREST

/***********************************/
// WebSocket / Streaming / Live
/***********************************/
// LiveClient is an alias for listenv1rest.Client
//
// Deprecated: This package is deprecated. Use WebSocketClient instead. This will be removed in a future release.
type LiveClient = listenv1ws.Client

/*
NewWebSocketForDemo creates a new websocket connection with all default options

Please see NewWebSocketUsingCallbackForDemo for more information.

TODO: Deprecate this function later
*/
func NewWebSocketForDemo(ctx context.Context, options *interfaces.LiveTranscriptionOptions) (*listenv1ws.Client, error) {
	return NewWebSocketForDemo(ctx, options)
}

// NewLiveForDemo is an alias for NewWebSocketForDemo
//
// Deprecated: This package is deprecated. Use NewWebSocketForDemo instead. This will be removed in a future release.
var NewLiveForDemo = NewWebSocketForDemo

/*
NewWebSocketWithDefaults creates a new websocket connection with all default options

Please see NewWebSocketUsingCallbackWithDefaults for more information.

TODO: Deprecate this function later
*/
func NewWebSocketWithDefaults(ctx context.Context, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	return NewWebSocketUsingCallbackWithDefaults(ctx, tOptions, callback)
}

// NewLiveWithDefaults is an alias for NewWebSocketWithDefaults
//
// Deprecated: This package is deprecated. Use NewWebSocketWithDefaults instead. This will be removed in a future release.
var NewLiveWithDefaults = NewWebSocketWithDefaults

/*
NewWebSocket creates a new websocket connection with the specified options

Please see NewWebSocketUsingCallback for more information.

TODO: Deprecate this function later
*/
func NewWebSocket(ctx context.Context, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	return NewWebSocketUsingCallback(ctx, apiKey, cOptions, tOptions, callback)
}

// NewLive is an alias for NewWebSocket
//
// Deprecated: This package is deprecated. Use NewWebSocket instead. This will be removed in a future release.
var NewLive = NewWebSocket

/*
NewWebSocketWithCancel creates a new websocket connection but has facilities to BYOC (Bring Your Own Cancel)

Please see NewWebSocketUsingCallbackWithCancel for more information.

TODO: Deprecate this function later
*/
func NewWebSocketWithCancel(ctx context.Context, ctxCancel context.CancelFunc, apiKey string, cOptions *interfaces.ClientOptions, tOptions *interfaces.LiveTranscriptionOptions, callback msginterfaces.LiveMessageCallback) (*listenv1ws.Client, error) {
	return NewWebSocketUsingCallbackWithCancel(ctx, ctxCancel, apiKey, cOptions, tOptions, callback)
}

// NewLiveWithCancel is an alias for NewWebSocketWithCancel
//
// Deprecated: This package is deprecated. Use NewWebSocketWithCancel instead. This will be removed in a future release.
var NewLiveWithCancel = NewWebSocketWithCancel
