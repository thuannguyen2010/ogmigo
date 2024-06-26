// Copyright 2021 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ogmigo

import "time"

// Options available to ogmios client
type Options struct {
	endpoint         string
	logger           Logger
	pipeline         int
	saveInterval     uint64
	handshakeTimeout time.Duration
}

// Option to cardano client
type Option func(*Options)

// WithEndpoint allows ogmios endpoint to set; defaults to ws://127.0.0.1:1337
func WithEndpoint(endpoint string) Option {
	return func(opts *Options) {
		opts.endpoint = endpoint
	}
}

// WithTimeout allows set timeout for handshaking
func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.handshakeTimeout = timeout
	}
}

// WithInterval specifies how frequently to save checkpoints when reading
func WithInterval(n int) Option {
	return func(options *Options) {
		options.saveInterval = uint64(n)
	}
}

// WithLogger allows custom logger to be specified
func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.logger = logger
	}
}

// WithPipeline allows number of pipelined ogmios requests to be provided
func WithPipeline(n int) Option {
	return func(opts *Options) {
		opts.pipeline = n
	}
}

func buildOptions(opts ...Option) Options {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}
	if options.endpoint == "" {
		options.endpoint = "ws://127.0.0.1:1337"
	}
	if options.logger == nil {
		options.logger = DefaultLogger
	}
	if options.pipeline <= 0 {
		options.pipeline = 50
	}
	if options.saveInterval <= 0 {
		options.saveInterval = 2160
	}
	return options
}
