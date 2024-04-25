// Copyright 2023 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package replay

import (
	"os"
	"sync"

	_ "azul3d.org/audio/flac.v0" // Add FLAC decoding support.

	flac "github.com/mewkiz/flac"
	wav "github.com/youpy/go-wav"
)

// ReplayOpts defines options for this device
type ReplayOptions struct {
	FullFilename string
}

// Client is a replay device. In this case, an audio stream.
type WavClient struct {
	options ReplayOptions

	// wav
	file    *os.File
	decoder *wav.Reader

	// operational stuff
	stopChan chan struct{}
	mute     sync.Mutex
	muted    bool
}
type FlacClient struct {
	options ReplayOptions

	// wav
	file    *os.File
	decoder *flac.Stream

	// operational stuff
	stopChan chan struct{}
	mute     sync.Mutex
	muted    bool
}
