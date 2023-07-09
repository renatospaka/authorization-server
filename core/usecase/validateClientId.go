package usecase

import (
	"github.com/renatospaka/authorization-server/core/entity"
	pkgEntity "github.com/renatospaka/authorization-server/utils/entity"
)

func validateClientId(id string) error {
	if id == "" || id == "00000000-0000-0000-0000-000000000000" {
		return entity.ErrClientIDIsRequired
	}

	if _, err := pkgEntity.Parse(id); err != nil {
		return entity.ErrInvalidClientID
	}

	return nil
}
