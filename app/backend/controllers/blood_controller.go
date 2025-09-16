package controllers

import (
	"blood-type-compatibility/helpers"
	"blood-type-compatibility/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DonorResponse struct {
	Username  string `json:"username"`
	BloodType string `json:"blood_type,omitempty"`
}

func GetCompatibleDonors(ctx *gin.Context) {
	var userBloodType struct {
		ABO string `form:"abo" binding:"required"`
		Rh  bool   `form:"rh" binding:"required"`
	}

	if ctx.ShouldBindQuery(&userBloodType) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	users, err := services.FindCompatibleDonors(userBloodType.ABO, userBloodType.Rh)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch compatible donors"})
		return
	}

	logrus.Info("Donors found successfully")

	var response []DonorResponse
	for _, u := range users {
		response = append(response, DonorResponse{
			Username:  u.Username,
			BloodType: fmt.Sprintf("%s%s", u.BloodType.ABO, helpers.Ternary(u.BloodType.Rh, "+", "-")),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"compatible_donors": response})
}
