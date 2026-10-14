package ddcdkconstruct


// Cloud Network Monitoring configuration.
type NetworkMonitoringFeatureConfig struct {
	// Enables Cloud Network Monitoring.
	//
	// Linux only.
	IsEnabled *bool `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

