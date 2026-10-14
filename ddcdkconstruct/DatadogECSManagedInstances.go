package ddcdkconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// The Datadog ECS Managed Instances construct manages the Datadog configuration for Datadog Agent daemons running on ECS Managed Instances.
type DatadogECSManagedInstances interface {
	// Creates a new Datadog Agent daemon task definition (and, unless `createDaemon` is false, the daemon itself).
	//
	// Merges the provided datadogProps with the class's datadogProps.
	DaemonTaskDefinition(scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) DatadogECSManagedInstancesDaemonTaskDefinition
}

// The jsii proxy struct for DatadogECSManagedInstances
type jsiiProxy_DatadogECSManagedInstances struct {
	_ byte // padding
}

func NewDatadogECSManagedInstances(datadogProps *DatadogECSManagedInstancesProps) DatadogECSManagedInstances {
	_init_.Initialize()

	if err := validateNewDatadogECSManagedInstancesParameters(datadogProps); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogECSManagedInstances{}

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSManagedInstances",
		[]interface{}{datadogProps},
		&j,
	)

	return &j
}

func NewDatadogECSManagedInstances_Override(d DatadogECSManagedInstances, datadogProps *DatadogECSManagedInstancesProps) {
	_init_.Initialize()

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogECSManagedInstances",
		[]interface{}{datadogProps},
		d,
	)
}

func (d *jsiiProxy_DatadogECSManagedInstances) DaemonTaskDefinition(scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) DatadogECSManagedInstancesDaemonTaskDefinition {
	if err := d.validateDaemonTaskDefinitionParameters(scope, id, datadogProps); err != nil {
		panic(err)
	}
	var returns DatadogECSManagedInstancesDaemonTaskDefinition

	_jsii_.Invoke(
		d,
		"daemonTaskDefinition",
		[]interface{}{scope, id, datadogProps},
		&returns,
	)

	return returns
}

