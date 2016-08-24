// Code generated by thriftrw-plugin-envelope
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package servicegenerator

import (
	"github.com/thriftrw/thriftrw-go/internal/envelope"
	"github.com/thriftrw/thriftrw-go/wire"
	"github.com/thriftrw/thriftrw-go/plugin/api"
)

// Client implements a ServiceGenerator client.
type client struct {
	client envelope.Client
}

// NewClient builds a new ServiceGenerator client.
func NewClient(c envelope.Client) api.ServiceGenerator {
	return &client{
		client: c,
	}
}

func (c *client) Generate(
	_Request *api.GenerateServiceRequest,
) (success *api.GenerateServiceResponse, err error) {
	args := GenerateHelper.Args(_Request)

	var body wire.Value
	body, err = args.ToWire()
	if err != nil {
		return
	}

	body, err = c.client.Send("generate", body)
	if err != nil {
		return
	}

	var result GenerateResult
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = GenerateHelper.UnwrapResponse(&result)
	return
}

// Handler serves an implementation of the ServiceGenerator service.
type Handler struct {
	impl api.ServiceGenerator
}

// NewHandler builds a new ServiceGenerator handler.
func NewHandler(service api.ServiceGenerator) Handler {
	return Handler{
		impl: service,
	}
}

// Handle receives and handles a request for the ServiceGenerator service.
func (h Handler) Handle(name string, reqValue wire.Value) (wire.Value, error) {
	switch name {

	case "generate":
		var args GenerateArgs
		if err := args.FromWire(reqValue); err != nil {
			return wire.Value{}, err
		}

		result, err := GenerateHelper.WrapResponse(
			h.impl.Generate(args.Request),
		)
		if err != nil {
			return wire.Value{}, err
		}

		return result.ToWire()

	default:

		return wire.Value{}, envelope.ErrUnknownMethod(name)

	}
}
