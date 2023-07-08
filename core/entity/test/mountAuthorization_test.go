package test

import (
	"testing"

	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestMountAuthorization(t *testing.T) {
	authData = getTestData()
	auth, err := entity.MountAuthorization(authData)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.NotEmpty(t, auth.GetID())
	assert.Equal(t, c_CLIENT_ID, auth.GetClientID())
	assert.Equal(t, c_TRANSACTION_ID, auth.GetTransactionID())
	assert.Equal(t, float32(200.00), auth.GetValue())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
	assert.True(t, auth.ReceivedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())

	err = auth.Validate()
	assert.Nil(t, err)
}

func TestMountAuthorizationInvalidAuthorizationID(t *testing.T) {
	authData = getTestData()
	authData.ID = "X"+c_AUTHORIZATION_ID
	auth, err := entity.MountAuthorization(authData)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidID.Error())
	
	authData = getTestData()
	authData.ID = ""
	auth2, err2 := entity.MountAuthorization(authData)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidID.Error())
	
	authData = getTestData()
	authData.ID = c_AUTHORIZATION_ID_FAKE
	auth3, err3 := entity.MountAuthorization(authData)
	assert.NotNil(t, err3)
	assert.Nil(t, auth3)
	assert.EqualError(t, err3, entity.ErrIDIsRequired.Error())
}

func TestMountAuthorizationInvalidTransactionID(t *testing.T) {
	authData = getTestData()
	authData.TransactionId = "X"+c_TRANSACTION_ID
	auth, err := entity.MountAuthorization(authData)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidTransactionID.Error())
	
	authData = getTestData()
	authData.TransactionId = ""
	auth2, err2 := entity.MountAuthorization(authData)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidTransactionID.Error())
	
	authData = getTestData()
	authData.TransactionId = c_TRANSACTION_ID_FAKE
	auth3, err3 := entity.MountAuthorization(authData)
	assert.NotNil(t, err3)
	assert.Nil(t, auth3)
	assert.EqualError(t, err3, entity.ErrTransactionIDIsRequired.Error())
}

func TestMountAuthorizationInvalidClientID(t *testing.T) {
	authData = getTestData()
	authData.ClientId = "X"+c_CLIENT_ID
	auth, err := entity.MountAuthorization(authData)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidClientID.Error())
	
	authData = getTestData()
	authData.ClientId = ""
	auth2, err2 := entity.MountAuthorization(authData)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidClientID.Error())
	
	authData = getTestData()
	authData.ClientId = c_CLIENT_ID_FAKE
	auth3, err3 := entity.MountAuthorization(authData)
	assert.NotNil(t, err3)
	assert.Nil(t, auth3)
	assert.EqualError(t, err3, entity.ErrClientIDIsRequired.Error())
}

func TestMountAuthorizationValueIsZero(t *testing.T) {
	authData = getTestData()
	authData.Value = 0
	auth, err := entity.MountAuthorization(authData)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestMountAuthorizationValueIsNegative(t *testing.T) {
	authData = getTestData()
	authData.Value = -200.00
	auth, err := entity.MountAuthorization(authData)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}
