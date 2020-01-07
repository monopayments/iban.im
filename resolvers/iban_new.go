package resolvers

import (
	"github.com/monocash/iban.im/model"
	"context"
	"github.com/monocash/iban.im/handler"
	"fmt"
	
	
)

// IbanNew mutation creates iban
func (r *Resolvers) IbanNew(ctx context.Context,args IbanNewMutationArgs) (*IbanNewResponse, error) {
	UserID := ctx.Value(handler.ContextKey("UserID"))
	if UserID == nil {
		msg := "Not Authorized"
		return &IbanNewResponse{Status: false, Msg: &msg, Iban: nil}, nil
	}
	// userid,_:= strconv.Atoi(UserID.(string))
	userid,_:= UserID.(int)
	fmt.Printf("UserID: %+v, userid: %i\n",UserID,userid)
	IbanNew := model.Iban{Text: args.Text, Password: args.Password, Handle: args.Handle, OwnerID:uint(userid)}

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
