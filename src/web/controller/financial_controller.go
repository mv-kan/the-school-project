package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/service"
)

func NewFinancial(financialServ service.IFinancialService) IFinancialController {
	return &financialController{financialServ: financialServ}
}

type IFinancialController interface {
	GetAllLodgingDebtors(w http.ResponseWriter, r *http.Request)
	CollectedMoneyForMonth(w http.ResponseWriter, r *http.Request)
}

type financialController struct {
	financialServ service.IFinancialService
}

func (c *financialController) GetAllLodgingDebtors(w http.ResponseWriter, r *http.Request) {

}

func (c *financialController) CollectedMoneyForMonth(w http.ResponseWriter, r *http.Request) {

}
