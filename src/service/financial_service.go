package service

import (
	"time"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/shopspring/decimal"
)

const SERVICE_TYPE_LODGING = 1

type IFinancialService interface {
	FindAllLodgingDebtors() ([]entity.Pupil, error)
	CollectedMoneyForMonth(time.Time) (decimal.Decimal, error)
}

type finacialService struct {
	invoiceRepo repo.IRepository[entity.Invoice]
	pupilRepo   repo.IRepository[entity.Pupil]
}

func (s finacialService) FindAllDebtors() ([]entity.Pupil, error) {
	// find all not debtors
	// get all invoices
	invoices, err := s.invoiceRepo.FindAll()
	if err != nil {
		return nil, err
	}
	// check for type of invoice and payment due date to get all fresh invoices
	notDebtors := make([]int, 0)
	for _, invoice := range invoices {
		if time.Now().Before(invoice.PaymentDue) {
			// extract pupil ids from invoices
			notDebtors = append(notDebtors, invoice.PupilID)
		}
	}
	// get all pupils
	pupils, err := s.pupilRepo.FindAll()
	if err != nil {
		return nil, err
	}
	// if pupil id is not in extracted pupil ids and pupil with that id has room id
	// then it is debtor and save him
	debtors := make([]entity.Pupil, 0)
	for _, pupil := range pupils {
		in := func(x int, set []int) bool {
			for _, value := range set {
				if x == value {
					return true
				}
			}
			return false
		}
		if !in(pupil.ID, notDebtors) && pupil.RoomID.Valid {
			debtors = append(debtors, pupil)
		}
	}
	return debtors, nil
}

func (s finacialService) CollectedMoneyForMonth(thisMonth time.Time) (decimal.Decimal, error) {
	// get all invoices
	invoices, err := s.invoiceRepo.FindAll()
	if err != nil {
		return decimal.Decimal{}, err
	}
	// compare date and add up amount of money
	sum := decimal.Decimal{}
	for _, invoice := range invoices {
		if invoice.DateOfPayment.Month() == thisMonth.Month() && invoice.DateOfPayment.Year() == thisMonth.Year() {
			sum = decimal.Sum(invoice.AmountOfMoney, sum)
		}
	}
	// return
	return sum, nil
}
