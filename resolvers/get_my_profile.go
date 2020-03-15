package resolvers

import (
	"context"
	"github.com/monocash/iban.im/config"

	"github.com/monocash/iban.im/handler"
	"github.com/monocash/iban.im/model"
	// "github.com/monocash/iban.im/tools"
	
)



// GetMyProfile resolver
func (r *Resolvers) GetMyProfile(ctx context.Context) (*GetMyProfileResponse, error) {
	UserID := ctx.Value(handler.ContextKey("UserID"))
	// tools.GetContextDetails(ctx)
	if UserID == nil {
		msg := "Not Authorized"
		return &GetMyProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	user := model.User{}
	if err := config.DB.First(&user, UserID).Error; err != nil {
		msg := "Not found"
		return &GetMyProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	return &GetMyProfileResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

// GetMyProfileResponse is the response type
type GetMyProfileResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for GetMyProfileResponse
func (r *GetMyProfileResponse) Ok() bool {
	return r.Status
}

// Error for GetMyProfileResponse
func (r *GetMyProfileResponse) Error() *string {
	return r.Msg
}
