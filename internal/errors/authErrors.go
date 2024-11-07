package errors

type SignTokenError struct{}

func (e SignTokenError) Error() string { return "Failed to sign token" }
