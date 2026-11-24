package ddcdkconstruct


type DatadogAppSecMode string

const (
	// Disable App and API Protection.
	DatadogAppSecMode_OFF DatadogAppSecMode = "OFF"
	// Enable App and API Protection.
	DatadogAppSecMode_ON DatadogAppSecMode = "ON"
	// Enable App and API Protection using the Datadog Lambda Extension implementation.
	DatadogAppSecMode_EXTENSION DatadogAppSecMode = "EXTENSION"
	// Enable App and API Protection using the Datadog Lambda Library implementation.
	//
	// **Warning**: This option forces tracer enablement for cases where the Datadog CDK Constructs
	// cannot safely detect that you are using a compatible library. Ensure that you are using the
	// Datadog Lambda Library for Python version "8.114.0" or above.
	// For Java, ensure that you are using Java layer version 31 or later and extension version
	// 102 or later. TRACER does not check versions; use ON for automatic selection and fallback.
	DatadogAppSecMode_TRACER DatadogAppSecMode = "TRACER"
)

