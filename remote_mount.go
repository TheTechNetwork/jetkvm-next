package kvm

import (
	"context"
	"errors"
)

type RemoteImageReader interface {
	Read(ctx context.Context, offset int64, size int64) ([]byte, error)
}

type WebRTCDiskReader struct {
}

var webRTCDiskReader WebRTCDiskReader

func (w *WebRTCDiskReader) Read(ctx context.Context, offset int64, size int64) ([]byte, error) {
	return nil, errors.New("not implemented in jetkvm-next-multisession")
}
