package ddcdkconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v3/internal"
)

type DatadogStepFunctions interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	Props() *DatadogStepFunctionsProps
	SetProps(val *DatadogStepFunctionsProps)
	Scope() constructs.Construct
	SetScope(val constructs.Construct)
	Stack() awscdk.Stack
	SetStack(val awscdk.Stack)
	AddStateMachines(stateMachines *[]awsstepfunctions.StateMachine, construct constructs.Construct)
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

// The jsii proxy struct for DatadogStepFunctions
type jsiiProxy_DatadogStepFunctions struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DatadogStepFunctions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogStepFunctions) Props() *DatadogStepFunctionsProps {
	var returns *DatadogStepFunctionsProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogStepFunctions) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogStepFunctions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatadogStepFunctions(scope constructs.Construct, id *string, props *DatadogStepFunctionsProps) DatadogStepFunctions {
	_init_.Initialize()

	if err := validateNewDatadogStepFunctionsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogStepFunctions{}

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogStepFunctions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatadogStepFunctions_Override(d DatadogStepFunctions, scope constructs.Construct, id *string, props *DatadogStepFunctionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogStepFunctions",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatadogStepFunctions)SetProps(val *DatadogStepFunctionsProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_DatadogStepFunctions)SetScope(val constructs.Construct) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_DatadogStepFunctions)SetStack(val awscdk.Stack) {
	if err := j.validateSetStackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stack",
		val,
	)
}

func DatadogStepFunctions_BuildLambdaPayloadToMergeTraces(payload *map[string]interface{}) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogStepFunctions",
		"buildLambdaPayloadToMergeTraces",
		[]interface{}{payload},
		&returns,
	)

	return returns
}

func DatadogStepFunctions_BuildStepFunctionTaskInputToMergeTraces(input *map[string]interface{}) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogStepFunctions",
		"buildStepFunctionTaskInputToMergeTraces",
		[]interface{}{input},
		&returns,
	)

	return returns
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
func DatadogStepFunctions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogStepFunctions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"datadog-cdk-constructs-v2.DatadogStepFunctions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogStepFunctions) AddStateMachines(stateMachines *[]awsstepfunctions.StateMachine, construct constructs.Construct) {
	if err := d.validateAddStateMachinesParameters(stateMachines); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addStateMachines",
		[]interface{}{stateMachines, construct},
	)
}

func (d *jsiiProxy_DatadogStepFunctions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogStepFunctions) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

