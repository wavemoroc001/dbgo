package wallet

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SaveWallet(c *gin.Context) {
	var req SaveWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "can not parse request",
		})
	}
	id, err := InsertWallet(req.Owner, req.Balance)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("can not save data: %w", err).Error(),
		})
		return
	}
	resp := GetWalletResponse{
		ID:      id,
		Balance: req.Balance,
		Owner:   req.Owner,
	}
	c.JSON(http.StatusCreated, resp)
}

func GetWalletByID(c *gin.Context) {
	pathID := c.Param("id")
	accID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not parse ID",
		})
		return
	}
	wallet, err := GetWalletById(accID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	if wallet.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func GetBalanceByID(c *gin.Context) {
	pathID := c.Param("id")
	accID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not parse id",
		})
		return
	}
	wallet, err := GetWalletById(accID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	resp := BalanceResponse{wallet.Balance}
	c.JSON(http.StatusOK, resp)
}

func DepositByID(c *gin.Context) {
	pathID := c.Param("id")
	accID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not parse id",
		})
		return
	}
	var req BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not parse id",
		})
		return
	}

	//if err := updateWalletBalance(accID, req.Balance); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "can not update balance" + err.Error(),
	//	})
	//	return
	//}

}

func WithdrawByID(c *gin.Context) {

}
