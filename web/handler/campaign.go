package handler

import (
	"go_startup_api/campaign"
	"go_startup_api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	userService     user.Service
}

func NewCampaignHandler(campaignService campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{campaignService, userService}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.GetCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}

func (h *campaignHandler) New(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	var input campaign.FormCreateCampaignInput
	input.Users = users

	c.HTML(http.StatusOK, "campaign_new.html", input)
}

func (h *campaignHandler) Create(c *gin.Context) {
	var input campaign.FormCreateCampaignInput
	users, err := h.userService.GetAllUsers()
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "campaign_new.html", input)
		return
	}
	input.Users = users

	err = c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "campaign_new.html", input)
		return
	}

	user, err := h.userService.GetUserByID(input.UserID)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "campaign_new.html", input)
		return
	}

	campaignInput := campaign.CreateCampaignInput{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		User:             user,
	}

	_, err = h.campaignService.CreateCampaign(campaignInput)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "campaign_new.html", input)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}
