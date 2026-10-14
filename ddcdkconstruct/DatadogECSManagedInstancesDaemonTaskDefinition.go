package ddcdkconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v5/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v5/internal"
)

// Deploys the Datadog Agent as an ECS Managed Daemon on ECS Managed Instances.
//
// Unlike `DatadogECSFargateTaskDefinition`, this construct does not extend an
// AWS CDK L2 task definition class - `AWS::ECS::DaemonTaskDefinition` has no
// L2 equivalent, and its container definition is a single value set at
// construction rather than built incrementally via `addContainer`. This is
// inherent to the daemon deployment model, not a limitation relative to the
// Fargate module: the daemon task definition holds only the Datadog Agent
// container. Application containers run in their own, separate task
// definitions and are wired to the daemon over the shared UDS socket volume
// (see `appDdSocketsVolume` / `appDdSocketsMountPoint`) or TCP (see
// `dogstatsdEnvironment` / `apmEnvironment`).
type DatadogECSManagedInstancesDaemonTaskDefinition interface {
	constructs.Construct
	// Environment variables for APM in application containers.
	//
	// Includes the
	// trace agent URL (the Agent socket when UDS is in use, otherwise the daemon
	// bridge address), `DD_PROFILING_ENABLED` when
	// `apm.isProfilingEnabled` is set, and `DD_TRACE_INFERRED_PROXY_SERVICES_ENABLED`
	// when `apm.traceInferredProxyServices` is set. Empty unless APM is enabled.
	ApmEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	// Mount point for the shared UDS socket directory.
	//
	// Add this to an
	// application container's mountPoints. Only defined when UDS is required.
	AppDdSocketsMountPoint() *awsecs.CfnDaemonTaskDefinition_MountPointProperty
	// Volume definition for the shared UDS socket directory.
	//
	// Add this to an
	// application task definition's volumes to enable UDS communication with
	// the Datadog Agent daemon. Only defined when UDS is required.
	AppDdSocketsVolume() *awsecs.CfnDaemonTaskDefinition_VolumeProperty
	// The underlying `AWS::ECS::Daemon` resource.
	//
	// Only defined if `createDaemon` is true.
	Daemon() awsecs.CfnDaemon
	// Environment variables for Data Streams Monitoring in application containers.
	//
	// Empty unless APM and `apm.dataStreams` are both enabled.
	// Data Streams Monitoring is an application/tracer-side feature - this is
	// for the application container that produces or consumes messages, not
	// the Datadog Agent container.
	DataStreamsEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	// Environment variables for DogStatsD in application containers.
	//
	// Empty
	// unless DogStatsD is enabled. Points the application at the Agent socket
	// when UDS is in use, or at the daemon bridge address when it is not.
	DogstatsdEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	// The IAM role used by ECS to pull the Datadog Agent image and access secrets.
	ExecutionRole() awsiam.IRole
	// The tree node.
	Node() constructs.Node
	// The underlying `AWS::ECS::DaemonTaskDefinition` resource.
	TaskDefinition() awsecs.CfnDaemonTaskDefinition
	// The IAM role used by the Datadog Agent container to make AWS API calls.
	TaskRole() awsiam.IRole
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DatadogECSManagedInstancesDaemonTaskDefinition
type jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) ApmEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty {
	var returns *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	_jsii_.Get(
		j,
		"apmEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) AppDdSocketsMountPoint() *awsecs.CfnDaemonTaskDefinition_MountPointProperty {
	var returns *awsecs.CfnDaemonTaskDefinition_MountPointProperty
	_jsii_.Get(
		j,
		"appDdSocketsMountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) AppDdSocketsVolume() *awsecs.CfnDaemonTaskDefinition_VolumeProperty {
	var returns *awsecs.CfnDaemonTaskDefinition_VolumeProperty
	_jsii_.Get(
		j,
		"appDdSocketsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) Daemon() awsecs.CfnDaemon {
	var returns awsecs.CfnDaemon
	_jsii_.Get(
		j,
		"daemon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) DataStreamsEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty {
	var returns *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	_jsii_.Get(
		j,
		"dataStreamsEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) DogstatsdEnvironment() *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty {
	var returns *[]*awsecs.CfnDaemonTaskDefinition_KeyValuePairProperty
	_jsii_.Get(
		j,
		"dogstatsdEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) TaskDefinition() awsecs.CfnDaemonTaskDefinition {
	var returns awsecs.CfnDaemonTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


func NewDatadogECSManagedInstancesDaemonTaskDefinition(scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) DatadogECSManagedInstancesDaemonTaskDefinition {
	_init_.Initialize()

	if err := validateNewDatadogECSManagedInstancesDaemonTaskDefinitionParameters(scope, id, datadogProps); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition{}

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSManagedInstancesDaemonTaskDefinition",
		[]interface{}{scope, id, datadogProps},
		&j,
	)

	return &j
}

func NewDatadogECSManagedInstancesDaemonTaskDefinition_Override(d DatadogECSManagedInstancesDaemonTaskDefinition, scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) {
	_init_.Initialize()

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSManagedInstancesDaemonTaskDefinition",
		[]interface{}{scope, id, datadogProps},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DatadogECSManagedInstancesDaemonTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogECSManagedInstancesDaemonTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSManagedInstancesDaemonTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSManagedInstancesDaemonTaskDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

