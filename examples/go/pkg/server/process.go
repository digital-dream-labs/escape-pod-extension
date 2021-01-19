package server

import (
	"context"
	"errors"

	"github.com/digital-dream-labs/escape-pod-extension/proto/lang/go/podextension"
)

// Unary checks the function map for an appropriate processor and either processes it or returns an error
func (s *Server) Unary(ctx context.Context, req *podextension.UnaryReq) (*podextension.UnaryResp, error) {
	r, ok := s.processors[req.IntentName]
	if !ok {
		return nil, errors.New("cannot find processor for request")
	}
	key, params, err := r.Get(ctx, req.IntentName, req.Parameters)
	if err != nil {
		return nil, err
	}

	response := &podextension.UnaryResp{
		IntentName: key,
		Parameters: params,
	}
	return response, nil
}
