//go:build !windows
// +build !windows

package tts

import "io"

// NewDecSpeech uses DECtalk to generate a speech (wav)
func NewDecSpeech(text string) (io.ReadCloser, error) {
	panic("unimplemented")
}
