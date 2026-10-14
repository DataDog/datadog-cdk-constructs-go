package ddcdkconstruct


type DeploymentAlarmConfiguration struct {
	// The CloudWatch alarm names to monitor during a daemon deployment.
	// Default: [].
	//
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Determines whether to use the CloudWatch alarm option in the daemon deployment process.
	// Default: false.
	//
	Enable *bool `field:"optional" json:"enable" yaml:"enable"`
}

