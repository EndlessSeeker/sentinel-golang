package go_zero

import (
	"github.com/alibaba/sentinel-golang/core/base"
	"net/http"
)

type (
	Option  func(*options)
	options struct {
		resourceExtract  func(r *http.Request) string
		blockFallback    func(r *http.Request) (int, string)
		sentinelFallback func(*http.ResponseWriter, *http.Request, string, base.BlockType)
	}
)

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, opt := range opts {
		opt(optCopy)
	}

	return optCopy
}

// WithResourceExtractor sets the resource extractor of the web requests.
func WithResourceExtractor(fn func(r *http.Request) string) Option {
	return func(opts *options) {
		opts.resourceExtract = fn
	}
}

// WithBlockFallback sets the fallback handler when requests are blocked.
func WithBlockFallback(fn func(r *http.Request) (int, string)) Option {
	return func(opts *options) {
		opts.blockFallback = fn
	}
}

// WithSentinelFallback use sentinel fallback handler when requests are blocked
func WithSentinelFallback() Option {
	return func(opts *options) {
		opts.sentinelFallback = sentinelFallback
	}
}
