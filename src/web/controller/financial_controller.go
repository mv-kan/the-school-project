package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/utils"
)

func NewFinancial(log logger.Logger, financialServ service.IFinancialService) IFinancialController {
	return &financialController{log: log, financialServ: financialServ}
}

type IFinancialController interface {
	GetAllLodgingDebtors(w http.ResponseWriter, r *http.Request)
	CollectedMoneyForMonth(w http.ResponseWriter, r *http.Request)
}

type financialController struct {
	financialServ service.IFinancialService
	log           logger.Logger
}

func (c *financialController) GetAllLodgingDebtors(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)

}

func (c *financialController) CollectedMoneyForMonth(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)

}
