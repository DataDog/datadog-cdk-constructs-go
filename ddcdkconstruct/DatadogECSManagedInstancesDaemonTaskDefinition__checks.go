//go:build !no_runtime_type_checking

package ddcdkconstruct

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateDatadogECSManagedInstancesDaemonTaskDefinition_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDatadogECSManagedInstancesDaemonTaskDefinitionParameters(scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(datadogProps, func() string { return "parameter datadogProps" }); err != nil {
		return err
	}

	return nil
}

