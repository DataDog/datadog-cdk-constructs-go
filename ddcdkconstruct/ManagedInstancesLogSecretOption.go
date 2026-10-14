package ddcdkconstruct


type ManagedInstancesLogSecretOption struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

