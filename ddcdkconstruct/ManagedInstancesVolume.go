package ddcdkconstruct


// A host-path volume mounted into the Datadog Agent container.
//
// The volume block on
// `AWS::ECS::DaemonTaskDefinition` only supports host paths - no Docker volume/EFS/FSx
// volume types.
type ManagedInstancesVolume struct {
	// The absolute path on the host to mount.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// The volume name.
	//
	// Cannot be `containerd_sock`, `proc`, `cgroup`, `dd-sockets` or
	// `debug`, which the construct manages. Letters, numbers, underscores and hyphens only.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The absolute path inside the Datadog Agent container.
	// Default: the same as `hostPath`.
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Mounts the volume read-only.
	// Default: true.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

