// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	"encoding/json"
	"os"
	"strings"

	prettyjson "github.com/hokaccha/go-prettyjson"
	klog "k8s.io/klog/v2"

	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/websocket/interfaces"
)

// NewWithDefault creates a ChanRouter with the default callback handler
func NewChanWithDefault() *ChanRouter {
	return NewChanRouter(NewDefaultCallbackHandler())
}

// New creates a ChanRouter with a user-defined callback
func NewChanRouter(callback interfaces.LiveMessageCallback) *ChanRouter {
	debugStr := os.Getenv("DEEPGRAM_DEBUG_WEBSOCKET")
	return &ChanRouter{
		debugWebsocket: strings.EqualFold(strings.ToLower(debugStr), "true"),
		channels:       make(map[interfaces.TypeResponse][]*ResponseChan),
	}
}

// AddChanCallback adds a channel to the router for a given message type
func (r *ChanRouter) AddChanCallback(msgType interfaces.TypeResponse, ch *ResponseChan) error {
	if ch == nil {
		return ErrUserChanNotDefined
	}
	switch msgType {
	case interfaces.TypeMessageResponse, interfaces.TypeMetadataResponse, interfaces.TypeSpeechStartedResponse, interfaces.TypeUtteranceEndResponse, interfaces.TypeErrorResponse:
		break
	default:
		return ErrInvalidMessageType
	}
	r.channels[msgType] = append(r.channels[msgType], ch)
	return nil
}

// OpenHelper handles the OpenResponse message
func (r *ChanRouter) OpenHelper(or *interfaces.OpenResponse) error {
	byMsg, err := json.Marshal(or)
	if err != nil {
		klog.V(1).Infof("json.Marshal(or) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeMessageResponse, byMsg, or)
}

// OpenResponse handles the OpenResponse message
func (r *ChanRouter) CloseHelper(or *interfaces.CloseResponse) error {
	byMsg, err := json.Marshal(or)
	if err != nil {
		klog.V(1).Infof("json.Marshal(or) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeMessageResponse, byMsg, or)
}

// ErrorResponse handles the OpenResponse message
func (r *ChanRouter) ErrorHelper(er *interfaces.ErrorResponse) error {
	byMsg, err := json.Marshal(er)
	if err != nil {
		klog.V(1).Infof("json.Marshal(er) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeMessageResponse, byMsg, er)
}

// processGeneric generalizes the handling of all message types
func (r *ChanRouter) processGeneric(msgType interfaces.TypeResponse, byMsg []byte, data interface{}) error {
	klog.V(6).Infof("router.%s ENTER\n", msgType)

	r.printDebugMessages(5, msgType, byMsg)

	chans := r.channels[msgType]
	if chans != nil {
		for _, ch := range chans {
			klog.V(5).Infof("callback.%s succeeded\n", msgType)
			*ch <- data
		}
	} else {
		klog.V(5).Infof("callback.%s is empty\n", msgType)
	}

	klog.V(6).Infof("router.%s LEAVE\n", msgType)
	return nil
}

func (r *ChanRouter) processMessage(byMsg []byte) error {
	var msg interfaces.MessageResponse
	if err := json.Unmarshal(byMsg, &msg); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageResponse) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeMessageResponse, byMsg, msg)
}

func (r *ChanRouter) processMetadata(byMsg []byte) error {
	var msg interfaces.MetadataResponse
	if err := json.Unmarshal(byMsg, &msg); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageResponse) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeMetadataResponse, byMsg, msg)
}

func (r *ChanRouter) processSpeechStartedResponse(byMsg []byte) error {
	var msg interfaces.SpeechStartedResponse
	if err := json.Unmarshal(byMsg, &msg); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageResponse) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeSpeechStartedResponse, byMsg, msg)
}

func (r *ChanRouter) processUtteranceEndResponse(byMsg []byte) error {
	var msg interfaces.UtteranceEndResponse
	if err := json.Unmarshal(byMsg, &msg); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageResponse) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeUtteranceEndResponse, byMsg, msg)
}

func (r *ChanRouter) processErrorResponse(byMsg []byte) error {
	var msg interfaces.ErrorResponse
	if err := json.Unmarshal(byMsg, &msg); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageResponse) failed. Err: %v\n", err)
		return err
	}

	return r.processGeneric(interfaces.TypeErrorResponse, byMsg, msg)
}

// Message handles platform messages and routes them appropriately based on the MessageType
func (r *ChanRouter) Message(byMsg []byte) error {
	klog.V(6).Infof("router.Message ENTER\n")

	if r.debugWebsocket {
		klog.V(5).Infof("Raw Message:\n%s\n", string(byMsg))
	}

	var mt interfaces.MessageType
	if err := json.Unmarshal(byMsg, &mt); err != nil {
		klog.V(1).Infof("json.Unmarshal(MessageType) failed. Err: %v\n", err)
		klog.V(6).Infof("router.Message LEAVE\n")
		return err
	}

	var err error
	switch interfaces.TypeResponse(mt.Type) {
	case interfaces.TypeMessageResponse:
		err = r.processMessage(byMsg)
	case interfaces.TypeMetadataResponse:
		err = r.processMetadata(byMsg)
	case interfaces.TypeSpeechStartedResponse:
		err = r.processSpeechStartedResponse(byMsg)
	case interfaces.TypeUtteranceEndResponse:
		err = r.processUtteranceEndResponse(byMsg)
	case interfaces.TypeErrorResponse:
		err = r.processErrorResponse(byMsg)
	default:
		err = r.UnhandledMessage(byMsg)
	}

	if err == nil {
		klog.V(6).Infof("MessageType(%s) after - Result: succeeded\n", mt.Type)
	} else {
		klog.V(5).Infof("MessageType(%s) after - Result: %v\n", mt.Type, err)
	}
	klog.V(6).Infof("router.Message LEAVE\n")
	return err
}

// UnhandledMessage logs and handles any unexpected message types
func (r *ChanRouter) UnhandledMessage(byMsg []byte) error {
	klog.V(6).Infof("router.UnhandledMessage ENTER\n")
	r.printDebugMessages(3, "UnhandledMessage", byMsg)
	klog.V(1).Infof("Unknown Event was received\n")
	klog.V(6).Infof("router.UnhandledMessage LEAVE\n")
	return ErrInvalidMessageType
}

// printDebugMessages formats and logs debugging messages
func (r *ChanRouter) printDebugMessages(level klog.Level, function interfaces.TypeResponse, byMsg []byte) {
	prettyJSON, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Format failed. Err: %v\n", err)
		return
	}
	klog.V(level).Infof("\n\n-----------------------------------------------\n")
	klog.V(level).Infof("%s RAW:\n%s\n", function, prettyJSON)
	klog.V(level).Infof("-----------------------------------------------\n\n\n")
}
