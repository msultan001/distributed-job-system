package job

type JobDTO struct {
	ID      int         `json:"id"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
