package ddcdkconstruct


// Application language for automatic APM instrumentation.
type TracerLanguage string

const (
	TracerLanguage_JAVA TracerLanguage = "JAVA"
	TracerLanguage_NODEJS TracerLanguage = "NODEJS"
	TracerLanguage_DOTNET TracerLanguage = "DOTNET"
	TracerLanguage_PYTHON TracerLanguage = "PYTHON"
	TracerLanguage_RUBY TracerLanguage = "RUBY"
	TracerLanguage_PHP TracerLanguage = "PHP"
)

