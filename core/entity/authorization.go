package entity

import (
	"math/rand"
	"time"

	pkgEntity "github.com/renatospaka/authorization-server/utils/entity"
)

const (
	TR_APPROVED = "approved"
	TR_DENIED   = "denied"
	TR_PENDING  = "pending"
	TR_DELETED  = "deleted"
)

type Authorization struct {
	receivedAt time.Time
	deniedAt   time.Time
	approvedAt time.Time
	*pkgEntity.TrailDate
	id            pkgEntity.ID
	clientId      pkgEntity.ID
	transactionId pkgEntity.ID
	status        string
	value         float32
	valid         bool
}

// Create a new authorization
func NewAuthorization(clientId string, transactionId string, value float32) (*Authorization, error) {
	clientUuid, err := pkgEntity.Parse(clientId)
	if err != nil {
		return nil, ErrInvalidClientID
	}

	transactionUuid, err := pkgEntity.Parse(transactionId)
	if err != nil {
		return nil, ErrInvalidTransactionID
	}

	authorization := &Authorization{
		id:            pkgEntity.NewID(),
		clientId:      clientUuid,
		transactionId: transactionUuid,
		value:         value,
		status:        TR_PENDING,
		deniedAt:      time.Time{},
		approvedAt:    time.Time{},
		TrailDate:     &pkgEntity.TrailDate{},
		valid:         false,
	}
	authorization.TrailDate.SetCreationToToday()

	// deliver the new authorization validated
	err = authorization.Validate()
	if err != nil {
		return nil, err
	}
	return authorization, nil
}

// Randomically define what is the response for the authorization proceess
// 21% of the authorizations should be denied
func (a *Authorization) Process() string {
	min, max := 0, 100
	random := rand.Intn(max-min) + min
	if random <= 21 {
		a.Deny()
	} else {
		a.Approve()
	}
	return a.status
}

// Use the same algorithm of Process
// but validates status before processing -  
// only pending authorizations can be reprocessed
func (a *Authorization) Reprocess() string {
	if a.status != TR_PENDING {
		return a.status
	}
	return a.Process()
}

// Get the ID of the authorization request
func (a *Authorization) GetID() string {
	uuid := a.id.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Get the Client ID of the transaction
func (a *Authorization) GetClientID() string {
	uuid := a.clientId.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Get the Transaction ID of the transaction
func (a *Authorization) GetTransactionID() string {
	uuid := a.transactionId.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Approve the authorization request
func (a *Authorization) Approve() {
	a.approvedAt = time.Now()
	a.deniedAt = time.Time{}
	a.receivedAt = time.Now()
	a.TrailDate.SetAlterationToToday()
	a.status = TR_APPROVED
}

// Deny the authorization request
func (a *Authorization) Deny() {
	a.approvedAt = time.Time{}
	a.deniedAt = time.Now()
	a.receivedAt = time.Now()
	a.TrailDate.SetAlterationToToday()
	a.status = TR_DENIED
}

// Get when the authorization was denied (if it was)
func (a *Authorization) DeniedAt() time.Time {
	return a.deniedAt
}

// Delete the authorization request
func (a *Authorization) DeleteIt() {
	a.approvedAt = time.Time{}
	a.deniedAt = time.Time{}
	a.TrailDate.SetDeletionToToday()
	a.status = TR_DELETED
}

// Get when the authorization was approved (if it was)
func (a *Authorization) ApprovedAt() time.Time {
	return a.approvedAt
}

// Get when the authorization was receeived (if it was)
func (a *Authorization) ReceivedAt() time.Time {
	return a.receivedAt
}

// Get the current status of the authorization request
func (a *Authorization) GetStatus() string {
	return a.status
}

// Get the value of the authorization
func (a *Authorization) GetValue() float32 {
	return a.value
}

// Validates all business rules to authorize this
func (a *Authorization) Validate() error {
	a.valid = false
	if a.id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(a.id.String()); err != nil {
		return ErrInvalidID
	}

	if a.value < 0 {
		return ErrValueIsNegative
	}

	if a.value == 0 {
		return ErrValueIsZero
	}

	if a.status != TR_APPROVED &&
		a.status != TR_DELETED &&
		a.status != TR_DENIED &&
		a.status != TR_PENDING {
		return ErrInvalidStatus
	}

	a.valid = true
	return nil
}

// Return whether the structure is valid or not
func (a *Authorization) IsValid() bool {
	return a.valid
}
