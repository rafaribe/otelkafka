package otelkafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func TestWithTracerProvider(t *testing.T) {
	mtp := mockTracerProvider(nil)
	// Default is to use the global TracerProvider. This will override that.
	c := NewConfig(iName, WithTracerProvider(mtp))
	expected := mtp.Tracer(iName)
	assert.Same(t, expected, c.Tracer)
}

func TestWithAttributes(t *testing.T) {
	attr := attribute.String("key", "value")
	c := NewConfig(iName, WithAttributes([]attribute.KeyValue{attr}))
	ssc := trace.NewSpanStartConfig(c.DefaultStartOpts...)
	assert.Contains(t, ssc.Attributes(), attr)
}

func TestWithPropagator(t *testing.T) {
	p := propagation.NewCompositeTextMapPropagator()
	// Use a non-nil value.
	p = propagation.NewCompositeTextMapPropagator(p)
	assert.Equal(t, p, NewConfig(iName, WithPropagator(p)).Propagator)
}
