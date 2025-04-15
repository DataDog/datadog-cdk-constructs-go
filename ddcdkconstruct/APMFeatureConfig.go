package ddcdkconstruct


// APM feature configuration.
type APMFeatureConfig struct {
	// Enables APM.
	IsEnabled *bool `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Enables APM traces traffic over Unix Domain Socket.
	//
	// Falls back to TCP configuration for application containers when disabled.
	IsSocketEnabled *bool `field:"optional" json:"isSocketEnabled" yaml:"isSocketEnabled"`
}

