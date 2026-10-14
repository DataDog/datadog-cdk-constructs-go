package ddcdkconstruct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

type DatadogECSManagedInstancesProps struct {
	// The Datadog API key string.
	//
	// Must define at least 1 source for the API key.
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The Datadog API key secret.
	//
	// Must define at least 1 source for the API key.
	ApiKeySecret awssecretsmanager.ISecret `field:"optional" json:"apiKeySecret" yaml:"apiKeySecret"`
	// The ARN of the Datadog API key secret.
	//
	// Must define at least 1 source for the API key.
	ApiKeySecretArn *string `field:"optional" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	// The JSON field to read from the Datadog API key secret in Secrets Manager.
	//
	// Applies to `apiKeySecret` and `apiKeySecretArn`.
	ApiKeySecretField *string `field:"optional" json:"apiKeySecretField" yaml:"apiKeySecretField"`
	// The ARN of the Datadog API key in SSM Parameter Store.
	//
	// Must define at least 1 source for the API key.
	ApiKeySsmArn *string `field:"optional" json:"apiKeySsmArn" yaml:"apiKeySsmArn"`
	// APM feature configuration.
	Apm *APMFeatureConfig `field:"optional" json:"apm" yaml:"apm"`
	// The Datadog Agent checks tag cardinality.
	ChecksCardinality Cardinality `field:"optional" json:"checksCardinality" yaml:"checksCardinality"`
	// The cluster name to use for tagging.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The minimum number of CPU units to reserve for the Datadog Agent container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Configure health check for the Datadog Agent container.
	DatadogHealthCheck *awsecs.HealthCheck `field:"optional" json:"datadogHealthCheck" yaml:"datadogHealthCheck"`
	// DogStatsD feature configuration.
	Dogstatsd *DogstatsdFeatureConfig `field:"optional" json:"dogstatsd" yaml:"dogstatsd"`
	// The task environment name.
	//
	// Used for tagging (UST).
	Env *string `field:"optional" json:"env" yaml:"env"`
	// Datadog Agent environment variables.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Global tags to apply to all data sent by the Agent.
	//
	// Overrides any DD_TAGS values in environmentVariables.
	GlobalTags *string `field:"optional" json:"globalTags" yaml:"globalTags"`
	// The version of the Datadog Agent container image to use.
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
	// Configure added containers to have container dependency on the Datadog Agent container.
	IsDatadogDependencyEnabled *bool `field:"optional" json:"isDatadogDependencyEnabled" yaml:"isDatadogDependencyEnabled"`
	// Configure Datadog Agent container to be essential for the task.
	IsDatadogEssential *bool `field:"optional" json:"isDatadogEssential" yaml:"isDatadogEssential"`
	// The amount (in MiB) of memory to present to the Datadog Agent container.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// Orchestrator Explorer feature configuration.
	OrchestratorExplorer *OrchestratorExplorerFeatureConfig `field:"optional" json:"orchestratorExplorer" yaml:"orchestratorExplorer"`
	// Configure Datadog Agent container to run with read-only root filesystem enabled.
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The registry to pull the Datadog Agent container image from.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// The task service name.
	//
	// Used for tagging (UST).
	Service *string `field:"optional" json:"service" yaml:"service"`
	// The Datadog site to send data to.
	Site *string `field:"optional" json:"site" yaml:"site"`
	// The task version.
	//
	// Used for tagging (UST).
	Version *string `field:"optional" json:"version" yaml:"version"`
	// A unique name for the daemon task definition family.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Log configuration for the Datadog Agent container's OWN logs (not application container logs, which the agent cannot collect in daemon mode on ECS Managed Instances), e.g. routing to CloudWatch via the awslogs driver.
	AgentLogConfiguration *ManagedInstancesLogConfiguration `field:"optional" json:"agentLogConfiguration" yaml:"agentLogConfiguration"`
	// ARNs of ECS Managed Instances capacity providers the daemon should run on.
	//
	// Required if `createDaemon` is true. This construct does not create the
	// capacity provider, infrastructure role, or instance profile - those must
	// already exist.
	CapacityProviderArns *[]*string `field:"optional" json:"capacityProviderArns" yaml:"capacityProviderArns"`
	// Path to the cgroup directory on the host.
	// Default: "/sys/fs/cgroup/".
	//
	CgroupPath *string `field:"optional" json:"cgroupPath" yaml:"cgroupPath"`
	// ARN of the ECS cluster where the Datadog Agent daemon will run.
	//
	// Required if `createDaemon` is true.
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Whether to create the `AWS::ECS::Daemon` resource.
	//
	// If false, only the
	// daemon task definition is created.
	// Default: true.
	//
	CreateDaemon *bool `field:"optional" json:"createDaemon" yaml:"createDaemon"`
	// Path to the containerd socket on the host.
	//
	// ECS Managed Instances uses
	// containerd, not Docker. The socket is always mounted at
	// `/var/run/containerd/containerd.sock` inside the Datadog Agent container.
	// Default: "/var/run/containerd/containerd.sock"
	//
	CriSocketPath *string `field:"optional" json:"criSocketPath" yaml:"criSocketPath"`
	// Name of the ECS daemon.
	// Default: `${family}-datadog-agent`.
	//
	DaemonName *string `field:"optional" json:"daemonName" yaml:"daemonName"`
	// Controls the daemon's rolling deployment behavior across instances.
	//
	// Any
	// change to the daemon task definition - including something as small as a
	// tag - forces a full replacement, and AWS's daemon deployment model drains
	// and replaces every EC2 instance in the attached capacity provider(s). This
	// block is write-only on the AWS side (not readable back from the API), so
	// `cdk diff` will always show a diff here - that is expected.
	DeploymentConfiguration *DeploymentConfiguration `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// Enable ECS managed tags for the daemon.
	// Default: true.
	//
	EnableEcsManagedTags *bool `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Enable ECS Exec for daemon tasks.
	// Default: false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The IAM role that grants the Amazon ECS container agent permission to make AWS API calls.
	//
	// Created automatically if not provided.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The IPC namespace mode for the daemon task.
	//
	// AWS only accepts `shared` or
	// `none` for a daemon task definition (`host` is rejected).
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// Log collection is not supported for the Datadog Agent running as an ECS Managed Daemon on ECS Managed Instances.
	//
	// Defining this property with
	// `isEnabled: true` will cause construct validation to fail.
	LogCollection *LogCollectionFeatureConfig `field:"optional" json:"logCollection" yaml:"logCollection"`
	// Configuration for Datadog Cloud Network Monitoring.
	//
	// Linux only.
	NetworkMonitoring *NetworkMonitoringFeatureConfig `field:"optional" json:"networkMonitoring" yaml:"networkMonitoring"`
	// The PID namespace mode for the daemon task.
	//
	// AWS only accepts `shared` or
	// `none` for a daemon task definition (`host` is rejected). DogStatsD/APM
	// origin detection does not require this to be set: Datadog client
	// libraries embed the container ID directly in the packet, which works the
	// same way over UDS and TCP regardless of PID namespace configuration.
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// Configuration for Datadog Live Process collection.
	ProcessCollection *ProcessCollectionFeatureConfig `field:"optional" json:"processCollection" yaml:"processCollection"`
	// Path to the /proc directory on the host.
	// Default: "/proc/".
	//
	ProcPath *string `field:"optional" json:"procPath" yaml:"procPath"`
	// Propagate tags to daemon tasks.
	// Default: "DAEMON".
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A map of additional tags to add to the daemon task definition/daemon created.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Number of cpu units used by the daemon task, as a string (e.g. "256").
	TaskCpu *string `field:"optional" json:"taskCpu" yaml:"taskCpu"`
	// Amount (in MiB) of memory used by the daemon task, as a string (e.g. "512").
	TaskMemory *string `field:"optional" json:"taskMemory" yaml:"taskMemory"`
	// The IAM role that allows the Datadog Agent container to make calls to other AWS services.
	//
	// Created automatically if not provided.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// Additional host-path volumes to mount into the Datadog Agent container, beyond the ones this construct manages (containerd socket, /proc, /sys/fs/cgroup, UDS sockets, network monitoring debug mount).
	Volumes *[]*ManagedInstancesVolume `field:"optional" json:"volumes" yaml:"volumes"`
}

