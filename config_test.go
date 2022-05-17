package otelkafka

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type testTextMapProp struct{}

var _ propagation.TextMapPropagator = (*testTextMapProp)(nil)

func (*testTextMapProp) Inject(context.Context, propagation.TextMapCarrier) {}

func (*testTextMapProp) Extract(ctx context.Context, _ propagation.TextMapCarrier) context.Context {
	return ctx
}

func (*testTextMapProp) Fields() []string { return nil }

func TestConfigDefaultPropagator(t *testing.T) {
	c := newConfig()
	expected := otel.GetTextMapPropagator()
	assert.Same(t, expected, c.Propagator)
}

func TestConfigUserPropagator(t *testing.T) {
	prop := new(testTextMapProp)
	c := newConfig(WithPropagator(prop))
	assert.Same(t, prop, c.Propagator)
}
