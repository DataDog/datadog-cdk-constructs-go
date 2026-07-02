package ddcdkconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct/v4/jsii"
)

type DatadogDefaultLayerVersions interface {
}

// The jsii proxy struct for DatadogDefaultLayerVersions
type jsiiProxy_DatadogDefaultLayerVersions struct {
	_ byte // padding
}

func NewDatadogDefaultLayerVersions() DatadogDefaultLayerVersions {
	_init_.Initialize()

	j := jsiiProxy_DatadogDefaultLayerVersions{}

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDatadogDefaultLayerVersions_Override(d DatadogDefaultLayerVersions) {
	_init_.Initialize()

	_jsii_.Create(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		nil, // no parameters
		d,
	)
}

func DatadogDefaultLayerVersions_DOTNET() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"DOTNET",
		&returns,
	)
	return returns
}

func DatadogDefaultLayerVersions_EXTENSION() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"EXTENSION",
		&returns,
	)
	return returns
}

func DatadogDefaultLayerVersions_JAVA() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"JAVA",
		&returns,
	)
	return returns
}

func DatadogDefaultLayerVersions_NODE() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"NODE",
		&returns,
	)
	return returns
}

func DatadogDefaultLayerVersions_PYTHON() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"PYTHON",
		&returns,
	)
	return returns
}

func DatadogDefaultLayerVersions_RUBY() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"datadog-cdk-constructs-v2.DatadogDefaultLayerVersions",
		"RUBY",
		&returns,
	)
	return returns
}

