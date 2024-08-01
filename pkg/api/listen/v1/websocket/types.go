// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket/interfaces"
)

type ResponseChan chan interface{}

// ChanRouter routes events
type ChanRouter struct {
	callback       interfaces.LiveMessageCallback
	debugWebsocket bool

	// call out to channels
	channels map[interfaces.TypeResponse][]*ResponseChan
}

// CallbackRouter routes events
type CallbackRouter struct {
	callback       interfaces.LiveMessageCallback
	debugWebsocket bool
}
type MessageRouter CallbackRouter
