package entity

import (
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
	id     pkgEntity.ID
	status string
	value  float32
	valid  bool
}

// Create a new authorization
func NewAuthorization(value float32) (*Authorization, error) {
	authorization := &Authorization{
		id:         pkgEntity.NewID(),
		value:      value,
		status:     TR_PENDING,
		deniedAt:   time.Time{},
		approvedAt: time.Time{},
		TrailDate:  &pkgEntity.TrailDate{},
		valid:      false,
	}
	authorization.TrailDate.SetCreationToToday()

	// deliver the new authorization validated
	err := authorization.Validate()
	if err != nil {
		return nil, err
	}
	return authorization, nil
}

// Get the ID of the authorization request
func (t *Authorization) GetID() string {
	return t.id.String()
}

// Approve the authorization request
func (t *Authorization) ApprovedIt() {
	t.approvedAt = time.Now()
	t.deniedAt = time.Time{}
	t.TrailDate.SetAlterationToToday()
	t.status = TR_APPROVED
}

// Deny the authorization request
func (t *Authorization) DenyIt() {
	t.approvedAt = time.Time{}
	t.deniedAt = time.Now()
	t.TrailDate.SetAlterationToToday()
	t.status = TR_DENIED
}

// Get when the authorization was denied (if it was)
func (t *Authorization) DeniedAt() time.Time {
	return t.deniedAt
}

// Delete the authorization request
func (t *Authorization) DeleteIt() {
	t.approvedAt = time.Time{}
	t.deniedAt = time.Time{}
	t.TrailDate.SetDeletionToToday()
	t.status = TR_DELETED
}

// Get when the authorization was approved (if it was)
func (t *Authorization) ApprovedAt() time.Time {
	return t.approvedAt
}

// Get the current status of the authorization request
func (t *Authorization) GetStatus() string {
	return t.status
}

// Get the value of the authorization
func (t *Authorization) GetValue() float32 {
	return t.value
}

// Validates all business rules to authorize this
func (t *Authorization) Validate() error {
	t.valid = false
	if t.id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(t.id.String()); err != nil {
		return ErrInvalidID
	}

	if t.value < 0 {
		return ErrValueIsNegative
	}

	if t.value == 0 {
		return ErrValueIsZero
	}

	if t.status != TR_APPROVED &&
		t.status != TR_DELETED &&
		t.status != TR_DENIED &&
		t.status != TR_PENDING {
		return ErrInvalidStatus
	}

	t.valid = true
	return nil
}

// Return whether the structure is valid or not
func (t *Authorization) IsValid() bool {
	return t.valid
}
