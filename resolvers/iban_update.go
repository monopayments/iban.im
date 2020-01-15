package resolvers

import (
	"context"

	"github.com/monocash/iban.im/handler"
	"github.com/monocash/iban.im/model"
	"fmt"
)

// IbanUpdate mutation change profile
func (r *Resolvers) IbanUpdate(ctx context.Context, args IbanUpdateMutationArgs) (*IbanUpdateResponse, error) {
	userID := ctx.Value(handler.ContextKey("UserID"))
	// ibanID:=1
	fmt.Println("ibanid: ",userID)
	fmt.Printf("ctx: %+v, args: %+v\n",ctx,args)



	if userID == nil {
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
	userid,_:= userID.(int)
	ibans:=r.FindIbanByOwner(userid)
	
	fmt.Printf("ibans: %+v\n",ibans)

	// if err := r.DB.First(&iban, ibanID).Error; err != nil {
	// 	msg := "Not existing iban"
	// 	return &IbanUpdateResponse{Status: false, Msg: &msg, Iban: nil}, nil
	// }

	// if args.Password != "" {
	// 	iban.Password = args.Password
	// }

	// r.DB.Save(&iban)
	return &IbanUpdateResponse{Status: true, Msg: nil, Iban: &IbanResponse{i: &iban}}, nil
}

type IbanUpdateMutationArgs struct {
	Text     string
	Password string
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

func (r *Resolvers)FindIbanByOwner(userID int)[]model.Iban{
	ibans:=[]model.Iban{}
	// Get all matched records
	r.DB.Where("owner_id = ?", userID).Find(&ibans)
	return ibans
}

