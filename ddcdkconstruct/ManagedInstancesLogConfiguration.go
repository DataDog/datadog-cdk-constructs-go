package ddcdkconstruct


// Log configuration for the Datadog Agent container's own logs.
type ManagedInstancesLogConfiguration struct {
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	SecretOptions *[]*ManagedInstancesLogSecretOption `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

