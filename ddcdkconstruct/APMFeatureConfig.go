package ddcdkconstruct


// APM feature configuration.
type APMFeatureConfig struct {
	// Enables Data Streams Monitoring.
	//
	// This is an application/tracer-side
	// feature: `DD_DATA_STREAMS_ENABLED` must be set on the application
	// container that produces or consumes messages (Kafka, SQS, Kinesis, SNS,
	// RabbitMQ, etc.), not on the Datadog Agent container.
	DataStreams *bool `field:"optional" json:"dataStreams" yaml:"dataStreams"`
	// Enables APM.
	IsEnabled *bool `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Enables Profile collection.
	//
	// Requires the Datadog tracer in your application containers, installed in the image or added with `apmInstrumentation`.
	IsProfilingEnabled *bool `field:"optional" json:"isProfilingEnabled" yaml:"isProfilingEnabled"`
	// Enables APM traces traffic over Unix Domain Socket.
	//
	// Falls back to TCP configuration for application containers when disabled.
	IsSocketEnabled *bool `field:"optional" json:"isSocketEnabled" yaml:"isSocketEnabled"`
	// Enables inferred spans for proxy services like AWS API Gateway.
	//
	// When enabled, the tracer will create spans for proxy services by using headers
	// passed from the proxy service to the application.
	TraceInferredProxyServices *bool `field:"optional" json:"traceInferredProxyServices" yaml:"traceInferredProxyServices"`
}

