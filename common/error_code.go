package common

type ErrorCodeTempAndArgs struct {
	Template string
	Args     []interface{}
}

type ErrorCode struct {
	error
	Code        int                   `json:"retcode"`
	Message     string                `json:"message"`
	TempAndArgs *ErrorCodeTempAndArgs `json:"-"`
}
