package service

import (
	"fmt"
	"time"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func NewEnroll(db *gorm.DB) IEnrollService {
	pupilRepo := repo.New[entity.Pupil](db)
	invoiceRepo := repo.New[entity.Invoice](db)
	return enrollService{db: db, pupilRepo: pupilRepo, invoiceRepo: invoiceRepo}
}

type IEnrollService interface {
	// creates pupil and first invoice for education
	Enroll(pupil entity.Pupil) error
}

type enrollService struct {
	db          *gorm.DB
	pupilRepo   repo.IRepository[entity.Pupil]
	invoiceRepo repo.IRepository[entity.Invoice]
}

func (s enrollService) Enroll(pupil entity.Pupil) error {

	tx := s.db.Begin()
	pupil, err := s.pupilRepo.WithTx(tx).Save(pupil)
	if err != nil {
		tx.Rollback()
		return err
	}
	// get school end date for payment due
	today := time.Now()
	var start_date, end_date time.Time

	toDate := func(year int, month int, day int) (time.Time, error) {
		layoutISO := "2006-1-2"
		date, err := time.Parse(layoutISO, fmt.Sprintf("%d-%d-%d", year, month, day))
		return date, err
	}
	if today.Month() >= START_OF_SCHOOL_YEAR_MONTH {
		start_date, err = toDate(today.Year()+1, START_OF_SCHOOL_YEAR_MONTH, START_OF_SCHOOL_YEAR_DAY)
		if err != nil {
			tx.Rollback()
			return err
		}
		end_date, err = toDate(today.Year()+2, END_OF_SCHOOL_YEAR_MONTH, END_OF_SCHOOL_YEAR_DAY)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if today.Month() < START_OF_SCHOOL_YEAR_MONTH {
		start_date, err = toDate(today.Year(), START_OF_SCHOOL_YEAR_MONTH, START_OF_SCHOOL_YEAR_DAY)
		if err != nil {
			tx.Rollback()
			return err
		}
		end_date, err = toDate(today.Year()+1, END_OF_SCHOOL_YEAR_MONTH, END_OF_SCHOOL_YEAR_DAY)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// create invoice with the pupil's data
	invoice := entity.Invoice{
		PaymentStart:    start_date,
		PaymentDue:      end_date,
		DateOfPayment:   today,
		PupilID:         pupil.ID,
		TypeOfServiceID: SERVICE_TYPE_EDUCATION,
		AmountOfMoney:   decimal.NewFromInt(CHARGE_FOR_EDUCATION),
	}
	_, err = s.invoiceRepo.WithTx(tx).Save(invoice)
	if err != nil {
		tx.Rollback()
		return err
	}
	// save it and commit to database
	tx.Commit()
	return nil
}
