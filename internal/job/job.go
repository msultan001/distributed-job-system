package job

import "github.com/google/uuid"

type Job struct {
	ID      uuid.UUID
	Type    string
	Payload interface{}
}
