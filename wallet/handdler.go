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
	var req DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not parse id",
		})
		return
	}

	wal, err := GetWalletById(accID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("can not get wallet %w", err),
		})
		return
	}

	if wal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "wallet not found",
		})
		return
	}

	newBalance := req.Amount + wal.Balance
	if err := updateWalletBalance(accID, newBalance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("can not update balance cause %w", err),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":      accID,
		"balance": newBalance,
	})
	return

}

func WithdrawByID(c *gin.Context) {
	pathID := c.Param("id")
	accID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("can not parse ID cause %w", err),
		})
	}

	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("can not parse request cause %w", err),
		})
	}

	wal, err := GetWalletById(accID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("can not get wallet by id: %w", err),
		})
		return
	}

	if wal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "wallet not found",
		})
		return
	}

	if wal.Balance-req.Amount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "balance is not enough to withdraw",
		})
		return
	}

	newBalance := wal.Balance - req.Amount
	if err := updateWalletBalance(accID, newBalance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("can not update balance cause %w", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      accID,
		"balance": newBalance,
	})
	return
}
