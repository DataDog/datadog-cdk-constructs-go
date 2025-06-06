package ddcdkconstruct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

type DatadogLambdaStrictProps struct {
	AddLayers *bool `field:"required" json:"addLayers" yaml:"addLayers"`
	CaptureCloudServicePayload *bool `field:"required" json:"captureCloudServicePayload" yaml:"captureCloudServicePayload"`
	CaptureLambdaPayload *bool `field:"required" json:"captureLambdaPayload" yaml:"captureLambdaPayload"`
	EnableDatadogASM *bool `field:"required" json:"enableDatadogASM" yaml:"enableDatadogASM"`
	EnableDatadogLogs *bool `field:"required" json:"enableDatadogLogs" yaml:"enableDatadogLogs"`
	EnableDatadogTracing *bool `field:"required" json:"enableDatadogTracing" yaml:"enableDatadogTracing"`
	EnableMergeXrayTraces *bool `field:"required" json:"enableMergeXrayTraces" yaml:"enableMergeXrayTraces"`
	GrantSecretReadAccess *bool `field:"required" json:"grantSecretReadAccess" yaml:"grantSecretReadAccess"`
	InjectLogContext *bool `field:"required" json:"injectLogContext" yaml:"injectLogContext"`
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	ApiKeySecret awssecretsmanager.ISecret `field:"optional" json:"apiKeySecret" yaml:"apiKeySecret"`
	ApiKeySecretArn *string `field:"optional" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	ApiKmsKey *string `field:"optional" json:"apiKmsKey" yaml:"apiKmsKey"`
	ExtensionLayerArn *string `field:"optional" json:"extensionLayerArn" yaml:"extensionLayerArn"`
	ExtensionLayerVersion *float64 `field:"optional" json:"extensionLayerVersion" yaml:"extensionLayerVersion"`
	FlushMetricsToLogs *bool `field:"optional" json:"flushMetricsToLogs" yaml:"flushMetricsToLogs"`
	ForwarderArn *string `field:"optional" json:"forwarderArn" yaml:"forwarderArn"`
	JavaLayerArn *string `field:"optional" json:"javaLayerArn" yaml:"javaLayerArn"`
	JavaLayerVersion *float64 `field:"optional" json:"javaLayerVersion" yaml:"javaLayerVersion"`
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	NodeLayerArn *string `field:"optional" json:"nodeLayerArn" yaml:"nodeLayerArn"`
	NodeLayerVersion *float64 `field:"optional" json:"nodeLayerVersion" yaml:"nodeLayerVersion"`
	PythonLayerArn *string `field:"optional" json:"pythonLayerArn" yaml:"pythonLayerArn"`
	PythonLayerVersion *float64 `field:"optional" json:"pythonLayerVersion" yaml:"pythonLayerVersion"`
	RedirectHandler *bool `field:"optional" json:"redirectHandler" yaml:"redirectHandler"`
	Site *string `field:"optional" json:"site" yaml:"site"`
	SourceCodeIntegration *bool `field:"optional" json:"sourceCodeIntegration" yaml:"sourceCodeIntegration"`
}

