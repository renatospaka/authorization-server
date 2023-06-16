package test

import (
	"testing"

	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthorization(t *testing.T) {
	tr, err := entity.NewAuthorization(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.NotEmpty(t, tr.GetID())
	assert.Equal(t, float32(200.00), tr.GetValue())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())

	err = tr.Validate()
	assert.Nil(t, err)
}

func TestNewAuthorizationValueIsZero(t *testing.T) {
	tr, err := entity.NewAuthorization(0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestNewAuthorizationValueIsNegative(t *testing.T) {
	tr, err := entity.NewAuthorization(-200.00)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}

func TestApproveIt(t *testing.T) {
	tr, err := entity.NewAuthorization(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.ApprovedIt()
	assert.Equal(t, entity.TR_APPROVED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.False(t, tr.ApprovedAt().IsZero())
}

func TestDenyIt(t *testing.T) {
	tr, err := entity.NewAuthorization(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.DenyIt()
	assert.Equal(t, entity.TR_DENIED, tr.GetStatus())
	assert.False(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
}

func TestDeleteIt(t *testing.T) {
	tr, err := entity.NewAuthorization(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.DeleteIt()
	assert.Equal(t, entity.TR_DELETED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.False(t, tr.TrailDate.DeletedAt().IsZero())
}
