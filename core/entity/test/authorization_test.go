package test

import (
	"testing"

	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthorization(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.ObjectsAreEqual(&entity.Authorization{}, auth)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.NotEmpty(t, auth.GetID())
	assert.Equal(t, c_CLIENT_ID, auth.GetClientID())
	assert.Equal(t, c_TRANSACTION_ID, auth.GetTransactionID())
	assert.Equal(t, float32(200.00), auth.GetValue())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())

	err = auth.Validate()
	assert.Nil(t, err)
}

func TestNewAuthorizationInvalidClientID(t *testing.T) {
	auth, err := entity.NewAuthorization("X"+c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidClientID.Error())

	auth2, err2 := entity.NewAuthorization("", c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidClientID.Error())

	auth3, err3 := entity.NewAuthorization(c_CLIENT_ID_FAKE, c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, err3)
	assert.Nil(t, auth3)
	assert.EqualError(t, err3, entity.ErrClientIDIsRequired.Error())
}

func TestNewAuthorizationInvalidTransactionID(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, "X"+c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidTransactionID.Error())

	auth2, err2 := entity.NewAuthorization(c_CLIENT_ID, "", 200.00)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidTransactionID.Error())

	auth3, err3 := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID_FAKE, 200.00)
	assert.NotNil(t, err3)
	assert.Nil(t, auth3)
	assert.EqualError(t, err3, entity.ErrTransactionIDIsRequired.Error())
}

func TestNewAuthorizationValueIsZero(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 0)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestNewAuthorizationValueIsNegative(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, -200.00)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}

func TestApprove(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	err = auth.Approve()
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_APPROVED, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.False(t, auth.ApprovedAt().IsZero())
}

func TestDeny(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	err = auth.Deny()
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_DENIED, auth.GetStatus())
	assert.False(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
}

func TestDeleteIt(t *testing.T) {
	auth, err := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	auth.DeleteIt()
	assert.Equal(t, entity.TR_DELETED, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
	assert.False(t, auth.TrailDate.DeletedAt().IsZero())
}

func TestProcess(t *testing.T) {
	auth, _ := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, auth)
	
	deny, approve, other := 0, 0, 0
	for i := 0; i < 10000; i++ {
		auth.Process()
		status := auth.GetStatus()
		if status == entity.TR_APPROVED {
			approve++
		} else if status == entity.TR_DENIED {
			deny++
		} else {
			other++
		}
	}
	denyP := float32(deny) / 10000.00 * 100.00
	approveP := float32(approve) / 10000.00 * 100.00

	assert.EqualValues(t, 0, other)
	assert.GreaterOrEqual(t, approveP, float32(75.0))
	assert.LessOrEqual(t, denyP, float32(25.0))
}

func TestReprocess(t *testing.T) {
	auth, _ := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, auth)

	status := auth.GetStatus()
	assert.Equal(t, entity.TR_PENDING, status)

	newStatus, err := auth.Reprocess()
	assert.Nil(t, err)
	assert.NotEmpty(t, newStatus)
	assert.Condition(t, func() bool {
		if newStatus == entity.TR_DENIED {return true}
		if newStatus == entity.TR_APPROVED {return true}
		return false
	})
}

func TestReprocessCannotReprocess(t *testing.T) {
	auth, _ := entity.NewAuthorization(c_CLIENT_ID, c_TRANSACTION_ID, 200.00)
	assert.NotNil(t, auth)

	status := auth.Process()
	assert.NotEmpty(t, status)
	assert.Condition(t, func() bool {
		if status == entity.TR_DENIED {return true}
		if status == entity.TR_APPROVED {return true}
		return false
	})

	status, err := auth.Reprocess()
	assert.NotNil(t, err)
	assert.EqualError(t, err, entity.ErrCannotReprocess.Error())
}
