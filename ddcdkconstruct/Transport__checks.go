//go:build !no_runtime_type_checking

package ddcdkconstruct

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (t *jsiiProxy_Transport) validateApplyEnvVarsParameters(lam interface{}) error {
	if lam == nil {
		return fmt.Errorf("parameter lam is required, but nil was provided")
	}
	switch lam.(type) {
	case awslambda.Function:
		// ok
	case awslambda.SingletonFunction:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(lam) {
			return fmt.Errorf("parameter lam must be one of the allowed types: awslambda.Function, awslambda.SingletonFunction; received %#v (a %T)", lam, lam)
		}
	}

	return nil
}

func (j *jsiiProxy_Transport) validateSetFlushMetricsToLogsParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Transport) validateSetSiteParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

