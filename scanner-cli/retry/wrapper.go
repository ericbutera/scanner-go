package retry

// What I want:
// Rate limit
// Span metrics around outbound
// Timeout
// Non-http specific implementation; should work with anything that can fail

import (
	"context"
	"log"
	"time"

	// "github.com/afex/hystrix-go/hystrix" this might be useful
	"github.com/cenkalti/backoff"
)

type Retryable func() error

func Run(fn Retryable) error {
	// A context
	ctx := context.Background()

	// An operation that may fail.
	operation := func() error {
		log.Print("Running operation")
		return fn()
	}

	expo := backoff.NewExponentialBackOff()
	expo.MaxElapsedTime = 10 * time.Second
	b := backoff.WithContext(expo, ctx)

	err := backoff.Retry(operation, b)
	if err != nil {
		log.Printf("Operation failure %+v", err)
		return err
	}

	// Operation is successful.
	return nil
}
