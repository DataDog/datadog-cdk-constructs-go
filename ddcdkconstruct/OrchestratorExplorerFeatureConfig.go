package ddcdkconstruct


// Orchestrator Explorer configuration.
type OrchestratorExplorerFeatureConfig struct {
	// Enables Orchestrator Explorer.
	IsEnabled *bool `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// The URL of the Orchestrator Explorer API.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

