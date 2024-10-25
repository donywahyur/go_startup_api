package transaction

import "go_startup_api/user"

type GetTransactionCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
