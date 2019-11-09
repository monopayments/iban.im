package resolvers

import (
	"context"

	"github.com/monocash/iban.im/handler"
	"github.com/monocash/iban.im/model"
)

// IbanUpdate mutation change profile
func (r *Resolvers) IbanUpdate(ctx context.Context, args IbanUpdateMutationArgs) (*IbanUpdateResponse, error) {
	ibanID := ctx.Value(handler.ContextKey("ibanID"))

	if ibanID == nil {
		msg := "Not Authorized"
		return &IbanUpdateResponse{Status: false, Msg: &msg, Iban: nil}, nil
	}
	if args.Text == "" {
		msg := "You have to provide IBAN"
		return &IbanUpdateResponse{Status: false, Msg: &msg, Iban: nil}, nil
	}
	if args.Handle == "" {
		msg := "You have to provide handle"
		return &IbanUpdateResponse{Status: false, Msg: &msg, Iban: nil}, nil
	}
	iban := model.Iban{}

	if err := r.DB.First(&iban, ibanID).Error; err != nil {
		msg := "Not existing iban"
		return &IbanUpdateResponse{Status: false, Msg: &msg, Iban: nil}, nil
	}

	if args.Password != nil {
		iban.Password = *args.Password
	}

	r.DB.Save(&iban)
	return &IbanUpdateResponse{Status: true, Msg: nil, Iban: &IbanResponse{i: &iban}}, nil
}

type IbanUpdateMutationArgs struct {
	Text     string
	Password *string
	Handle   string
}

// IbanUpdateResponse is the response type
type IbanUpdateResponse struct {
	Status bool
	Msg    *string
	Iban   *IbanResponse
}

// Ok for IbanUpdateResponse
func (r *IbanUpdateResponse) Ok() bool {
	return r.Status
}

// Error for IbanUpdateResponse
func (r *IbanUpdateResponse) Error() *string {
	return r.Msg
}
