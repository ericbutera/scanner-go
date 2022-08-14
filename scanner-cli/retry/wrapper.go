package retry

// What I want:
// Rate limit
// Span metrics around outbound
// Timeout
// Non-http specific implementation; should work with anything that can fail
//
// Solid implementation: https://github.com/aws/aws-sdk-go-v2/blob/v1.16.11/aws/retry/standard.go

import (
	"context"
	"log"
	"time"

	// "github.com/afex/hystrix-go/hystrix" this might be useful
	"github.com/cenkalti/backoff"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Retryable func() error

func Run(operationName string, fn Retryable) error {
	ctx := context.Background()

	// operation that may fail
	operation := func() error {
		log.Printf("Running operation %s", operationName)
		return fn()
	}

	expo := backoff.NewExponentialBackOff()
	expo.MaxElapsedTime = 10 * time.Second
	b := backoff.WithContext(expo, ctx)

	span := tracer.StartSpan(operationName)
	defer span.Finish()
	err := backoff.Retry(operation, b)
	span.Finish(tracer.WithError(err))

	if err != nil {
		log.Printf("Operation failure %+v", err)
		return err
	}

	return err
}
