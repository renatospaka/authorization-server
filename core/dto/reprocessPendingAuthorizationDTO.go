package dto

type AuthorizationReprocessPendingDto struct {
	ClientID        string  `json:"client_id"`
	TransactionID   string  `json:"transaction_id"`
	Value           float32 `json:"value"`
}

type AuthorizationReprocessPendingResultDto struct {
	AuthorizationID string  `json:"authorization_id"`
	ClientID        string  `json:"client_id"`
	TransactionID   string  `json:"transaction_id"`
	Status          string  `json:"status"`
	Value           float32 `json:"value"`
	DeniedAt        string  `json:"denied_at"`
	ApprovedAt      string  `json:"approved_at"`
	ErrorMessage    string  `json:"error_message"`
}
