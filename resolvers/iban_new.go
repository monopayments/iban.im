package resolvers

import (
	"github.com/monocash/iban.im/model"
)

// IbanNew mutation creates iban
func (r *Resolvers) IbanNew(args IbanNewMutationArgs) (*IbanNewResponse, error) {

	IbanNew := model.Iban{Text: args.Text, Password: args.Password, Handle: args.Handle}

	r.DB.Create(&IbanNew)

	return &IbanNewResponse{Status: true, Msg: nil, Iban: &IbanResponse{i: &IbanNew}}, nil
}

type IbanNewMutationArgs struct {
	Text     string
	Password string
	Handle   string
}

// IbanNewResponse is the response type
type IbanNewResponse struct {
	Status bool
	Msg    *string
	Iban   *IbanResponse
}

// Ok for IbanNewResponse
func (r *IbanNewResponse) Ok() bool {
	return r.Status
}

// Error for IbanNewResponse
func (r *IbanNewResponse) Error() *string {
	return r.Msg
}
