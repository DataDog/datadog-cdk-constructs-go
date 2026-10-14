package ddcdkconstruct


// Controls the daemon's rolling deployment behavior across instances.
type DeploymentConfiguration struct {
	// CloudWatch alarm configuration for the daemon deployment.
	Alarms *DeploymentAlarmConfiguration `field:"optional" json:"alarms" yaml:"alarms"`
	// The amount of time (in minutes) to wait after a successful deployment step before proceeding.
	// Default: 0.
	//
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// The percentage of instances to drain simultaneously during a daemon deployment.
	// Default: 25.
	//
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}

