package ddcdkconstruct


// Automatic APM instrumentation configuration.
type APMInstrumentationConfig struct {
	// The application language, which selects the tracer to add.
	Language TracerLanguage `field:"required" json:"language" yaml:"language"`
	// The name of the application container that loads the tracer.
	//
	// Required when the task definition has more than one application container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The C standard library that the application image uses.
	//
	// Defaults to `GLIBC`.
	TracerLibc TracerLibc `field:"optional" json:"tracerLibc" yaml:"tracerLibc"`
	// The version of the tracer to add.
	//
	// Defaults to `latest`.
	TracerVersion *string `field:"optional" json:"tracerVersion" yaml:"tracerVersion"`
}

