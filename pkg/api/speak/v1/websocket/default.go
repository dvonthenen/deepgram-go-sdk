// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	prettyjson "github.com/hokaccha/go-prettyjson"
	klog "k8s.io/klog/v2"

	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/api/speak/v1/websocket/interfaces"
)

// DefaultCallbackHandler is a default callback handler for text-to-speech connections
type DefaultCallbackHandler struct{}

// NewDefaultCallbackHandler creates a new DefaultCallbackHandler
func NewDefaultCallbackHandler() DefaultCallbackHandler {
	return DefaultCallbackHandler{}
}

// Open is the callback for when the connection opens
func (dch DefaultCallbackHandler) Open(or *interfaces.OpenResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(or)
		if err != nil {
			klog.V(1).Infof("Open json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nOpen Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n\n[OpenResponse]\n\n")

	return nil
}

// Metadata is the callback for information about the connection
func (dch DefaultCallbackHandler) Metadata(md *interfaces.MetadataResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(md)
		if err != nil {
			klog.V(1).Infof("Metadata json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nMetadata Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n\nMetadata.RequestID: %s\n", strings.TrimSpace(md.RequestID))

	return nil
}

// Flushed is the callback for when the connection flushes
func (dch DefaultCallbackHandler) Flush(fr *interfaces.FlushedResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(fr)
		if err != nil {
			klog.V(1).Infof("Flush json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nFlush Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n\nFlushed.SequenceID: %d\n", fr.SequenceID)

	return nil
}

// Binary is the callback for when the connection receives binary data
func (dch DefaultCallbackHandler) Binary(br []byte) error {
	klog.V(3).Infof("Received binary data: %d bytes", len(br))
	return nil
}

// Close is the callback for when the connection closes
func (dch DefaultCallbackHandler) Close(or *interfaces.CloseResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(or)
		if err != nil {
			klog.V(1).Infof("Close json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nClose Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n\n[CloseResponse]\n\n")

	return nil
}

// Warning is the callback for error messages
func (dch DefaultCallbackHandler) Warning(wr *interfaces.WarningResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(wr)
		if err != nil {
			klog.V(1).Infof("Error json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nWarning Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n[WarningResponse]\n")
	fmt.Printf("\nError.Code: %s\n", wr.WarnCode)
	fmt.Printf("Error.Message: %s\n", wr.WarnMsg)

	return nil
}

// Error is the callback for error messages
func (dch DefaultCallbackHandler) Error(er *interfaces.ErrorResponse) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		data, err := json.Marshal(er)
		if err != nil {
			klog.V(1).Infof("Error json.Marshal failed. Err: %v\n", err)
			return err
		}

		prettyJSON, err := prettyjson.Format(data)
		if err != nil {
			klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
			return err
		}
		klog.V(2).Infof("\n\nError Object:\n%s\n\n", prettyJSON)

		return nil
	}

	// handle the message
	fmt.Printf("\n[ErrorResponse]\n")
	fmt.Printf("\nError.Type: %s\n", er.ErrCode)
	fmt.Printf("Error.Message: %s\n", er.ErrMsg)
	fmt.Printf("Error.Description: %s\n\n", er.Description)
	fmt.Printf("Error.Variant: %s\n\n", er.Variant)

	return nil
}

// UnhandledEvent is the callback for unknown messages
func (dch DefaultCallbackHandler) UnhandledEvent(byData []byte) error {
	var debugStr string
	if v := os.Getenv("DEEPGRAM_DEBUG"); v != "" {
		klog.V(4).Infof("DEEPGRAM_DEBUG found")
		debugStr = v
	}

	if strings.EqualFold(debugStr, "true") {
		prettyJSON, err := prettyjson.Format(byData)
		if err != nil {
			klog.V(2).Infof("\n\nRaw Data:\n%s\n\n", string(byData))
		} else {
			klog.V(2).Infof("\n\nError Object:\n%s\n\n", prettyJSON)
		}

		return nil
	}

	// handle the message
	fmt.Printf("\n[UnhandledEvent]")
	fmt.Printf("Dump:\n%s\n\n", string(byData))

	return nil
}
