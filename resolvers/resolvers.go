package resolvers

import (
	"github.com/monocash/iban.im/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
