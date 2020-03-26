package resolvers

import (
	"context"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"strings"
)

func (r *Resolvers) ShowInfo(ctx context.Context, args ShowInfoArgs) (response *ShowInfoResponse,err error)  {
	iban := r.GetIbanById(args.Id)
	response = &ShowInfoResponse{}
	defer func() {
		if err != nil {
			msg := err.Error()
			response.Msg = &msg
		}else{
			response.Status = true
		}
	}()
	if iban.IbanID == 0 {
		err = fmt.Errorf("iban is not exist")
		return
	}
	if strings.TrimSpace(args.Password) == "" {
		err = fmt.Errorf("password is empty")
		return
	}

	if !iban.ComparePassword(args.Password) {
		err = fmt.Errorf("password is not correct")
	}
	return
}




type ShowInfoArgs struct {
	Id graphql.ID
	Password string
}

// Response for checking iban password
type ShowInfoResponse struct {
	Status bool
	Msg    *string
}

// Ok for ShowInfoResponse
func (r *ShowInfoResponse) Ok() bool {
	return r.Status
}

// Error for ShowInfoResponse
func (r *ShowInfoResponse) Error() *string {
	return r.Msg
}