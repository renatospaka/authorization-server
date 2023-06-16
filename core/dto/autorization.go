package dto

type AuthorizationProcessDto struct {
	Value      float32 `json:"value"`
}

type AuthorizationProcessResultDto struct {
	ID         string  `json:"authorization_id"`
	Status     string  `json:"status"`
	Value      float32 `json:"value"`
	DeniedAt   string  `json:"denied_at"`
	ApprovedAt string  `json:"approved_at"`
}
