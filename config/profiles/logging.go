package profiles

type Logging struct {
	Level                 string `json:"level"`
	WithCaller            bool   `json:"with_caller"`
	WithStacktraceOnError bool   `json:"with_stacktrace_on_error"`
}
