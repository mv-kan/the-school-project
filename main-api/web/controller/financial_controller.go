package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/utils"
	"github.com/shopspring/decimal"
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
	debtors, err := c.financialServ.FindAllLodgingDebtors()
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, debtors)
}

func (c *financialController) CollectedMoneyForMonth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	year, err := strconv.Atoi(params["year"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	month, err := strconv.Atoi(params["month"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}

	money, err := c.financialServ.CollectedMoneyForMonth(year, month)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]decimal.Decimal{"amountOfMoney": money})
}
