package ddcdkconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v2/internal"
)

// The Datadog ECS Fargate Task Definition automatically instruments the ECS Fargate task and containers with configured Datadog features.
type DatadogECSFargateTaskDefinition interface {
	awsecs.FargateTaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() awsecs.Compatibility
	// The container definitions.
	Containers() *[]awsecs.ContainerDefinition
	CwsContainer() awsecs.ContainerDefinition
	DatadogContainer() awsecs.ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() awsecs.ContainerDefinition
	SetDefaultContainer(val awsecs.ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*awsecs.InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	LogContainer() awsecs.ContainerDefinition
	// The Docker networking mode to use for the containers in the task.
	//
	// Fargate tasks require the awsvpc network mode.
	NetworkMode() awsecs.NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The process namespace to use for the containers in the task.
	//
	// Only supported for tasks that are hosted on AWS Fargate if the tasks
	// are using platform version 1.4.0 or later (Linux). Not supported in
	// Windows containers. If pidMode is specified for a Fargate task,
	// then runtimePlatform.operatingSystemFamily must also be specified.  For more
	// information, see [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_definition_pidmode).
	PidMode() awsecs.PidMode
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	//
	// Modifies properties of container to support specified agent configuration in task.
	AddContainer(id *string, containerProps *awsecs.ContainerDefinitionOptions) awsecs.ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension awsecs.ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *awsecs.FirelensLogRouterDefinitionOptions) awsecs.FirelensLogRouter
	// Adds an inference accelerator to the task definition.
	AddInferenceAccelerator(inferenceAccelerator *awsecs.InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint awsecs.PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *awsecs.Volume)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) awsecs.ContainerDefinition
	// Determine the existing port mapping for the provided name.
	//
	// Returns: PortMapping for the provided name, if it exists.
	FindPortMappingByName(name *string) *awsecs.PortMapping
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants permissions to run this task definition.
	//
	// This will grant the following permissions:
	//
	//   - ecs:RunTask
	// - iam:PassRole.
	GrantRun(grantee awsiam.IGrantable) awsiam.Grant
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DatadogECSFargateTaskDefinition
type jsiiProxy_DatadogECSFargateTaskDefinition struct {
	internal.Type__awsecsFargateTaskDefinition
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Compatibility() awsecs.Compatibility {
	var returns awsecs.Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Containers() *[]awsecs.ContainerDefinition {
	var returns *[]awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) CwsContainer() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"cwsContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) DatadogContainer() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"datadogContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) DefaultContainer() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) InferenceAccelerators() *[]*awsecs.InferenceAccelerator {
	var returns *[]*awsecs.InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) LogContainer() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"logContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) NetworkMode() awsecs.NetworkMode {
	var returns awsecs.NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) PidMode() awsecs.PidMode {
	var returns awsecs.PidMode
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


func NewDatadogECSFargateTaskDefinition(scope constructs.Construct, id *string, props *awsecs.FargateTaskDefinitionProps, datadogProps *DatadogECSFargateProps) DatadogECSFargateTaskDefinition {
	_init_.Initialize()

	if err := validateNewDatadogECSFargateTaskDefinitionParameters(scope, id, props, datadogProps); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogECSFargateTaskDefinition{}

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		[]interface{}{scope, id, props, datadogProps},
		&j,
	)

	return &j
}

func NewDatadogECSFargateTaskDefinition_Override(d DatadogECSFargateTaskDefinition, scope constructs.Construct, id *string, props *awsecs.FargateTaskDefinitionProps, datadogProps *DatadogECSFargateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		[]interface{}{scope, id, props, datadogProps},
		d,
	)
}

func (j *jsiiProxy_DatadogECSFargateTaskDefinition)SetDefaultContainer(val awsecs.ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
func DatadogECSFargateTaskDefinition_FromFargateTaskDefinitionArn(scope constructs.Construct, id *string, fargateTaskDefinitionArn *string) awsecs.IFargateTaskDefinition {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_FromFargateTaskDefinitionArnParameters(scope, id, fargateTaskDefinitionArn); err != nil {
		panic(err)
	}
	var returns awsecs.IFargateTaskDefinition

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"fromFargateTaskDefinitionArn",
		[]interface{}{scope, id, fargateTaskDefinitionArn},
		&returns,
	)

	return returns
}

// Import an existing Fargate task definition from its attributes.
func DatadogECSFargateTaskDefinition_FromFargateTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *awsecs.FargateTaskDefinitionAttributes) awsecs.IFargateTaskDefinition {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_FromFargateTaskDefinitionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awsecs.IFargateTaskDefinition

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"fromFargateTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func DatadogECSFargateTaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) awsecs.ITaskDefinition {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_FromTaskDefinitionArnParameters(scope, id, taskDefinitionArn); err != nil {
		panic(err)
	}
	var returns awsecs.ITaskDefinition

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func DatadogECSFargateTaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *awsecs.TaskDefinitionAttributes) awsecs.ITaskDefinition {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_FromTaskDefinitionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awsecs.ITaskDefinition

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"fromTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DatadogECSFargateTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DatadogECSFargateTaskDefinition_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatadogECSFargateTaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatadogECSFargateTaskDefinition_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogECSFargateTaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddContainer(id *string, containerProps *awsecs.ContainerDefinitionOptions) awsecs.ContainerDefinition {
	if err := d.validateAddContainerParameters(id, containerProps); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{id, containerProps},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddExtension(extension awsecs.ITaskDefinitionExtension) {
	if err := d.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addExtension",
		[]interface{}{extension},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddFirelensLogRouter(id *string, props *awsecs.FirelensLogRouterDefinitionOptions) awsecs.FirelensLogRouter {
	if err := d.validateAddFirelensLogRouterParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsecs.FirelensLogRouter

	_jsii_.Invoke(
		d,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddInferenceAccelerator(inferenceAccelerator *awsecs.InferenceAccelerator) {
	if err := d.validateAddInferenceAcceleratorParameters(inferenceAccelerator); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddPlacementConstraint(constraint awsecs.PlacementConstraint) {
	if err := d.validateAddPlacementConstraintParameters(constraint); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	if err := d.validateAddToExecutionRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	if err := d.validateAddToTaskRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) AddVolume(volume *awsecs.Volume) {
	if err := d.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{volume},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) FindContainer(containerName *string) awsecs.ContainerDefinition {
	if err := d.validateFindContainerParameters(containerName); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		d,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) FindPortMappingByName(name *string) *awsecs.PortMapping {
	if err := d.validateFindPortMappingByNameParameters(name); err != nil {
		panic(err)
	}
	var returns *awsecs.PortMapping

	_jsii_.Invoke(
		d,
		"findPortMappingByName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) GrantRun(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantRunParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRun",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		d,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogECSFargateTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

