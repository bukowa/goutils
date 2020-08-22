package storage

import "fmt"

type ErrorBucketDoesNotExists string

func (e ErrorBucketDoesNotExists) Error() string {
	return fmt.Sprintf("bucket for %s does not exist.", string(e))
}

type ErrorEmptyKey string

func (e ErrorEmptyKey) Error() string {
	return fmt.Sprintf("value of Key() for model %s is empty", string(e))
}
