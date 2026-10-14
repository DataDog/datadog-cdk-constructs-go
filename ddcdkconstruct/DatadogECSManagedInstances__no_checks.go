//go:build no_runtime_type_checking

package ddcdkconstruct

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatadogECSManagedInstances) validateDaemonTaskDefinitionParameters(scope constructs.Construct, id *string, datadogProps *DatadogECSManagedInstancesProps) error {
	return nil
}

func validateNewDatadogECSManagedInstancesParameters(datadogProps *DatadogECSManagedInstancesProps) error {
	return nil
}

