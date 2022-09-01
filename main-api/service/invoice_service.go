package service

import (
	"errors"
	"time"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"gorm.io/gorm"
)

// use this contructor because this one uses upgraded version of service struct specifically for invoice case
func NewInvoice(db *gorm.DB) IService[entity.Invoice] {
	invoiceRepo := repo.New[entity.Invoice](db)
	invoiceNoteRepo := repo.New[entity.InvoiceNote](db)
	invoiceService := invoiceService{db: db, invoiceRepo: invoiceRepo, invoiceNoteRepo: invoiceNoteRepo}
	return &invoiceService
}

// the main difference between service[Invoice] and invoiceService is
// that invoiceService gets invoices with notes and service[Invoice] doesn't
type invoiceService struct {
	db              *gorm.DB
	invoiceNoteRepo repo.IRepository[entity.InvoiceNote]
	invoiceRepo     repo.IRepository[entity.Invoice]
}

func (s invoiceService) FindAll() ([]entity.Invoice, error) {
	// get all invoices
	invoices, err := s.invoiceRepo.FindAll()
	if err != nil {
		return nil, err
	}
	// get all invoice notes for each invoice
	for i, invoice := range invoices {
		note, err := s.invoiceNoteRepo.Find(invoice.ID)
		if errors.Is(err, repo.ErrRecordNotFound) {
			invoices[i].Note = nil
		} else if err != nil {
			return nil, err
		} else {
			invoices[i].Note = &note
		}
	}
	// return invoices with invoice notes
	return invoices, nil
}

func (s invoiceService) Find(id int) (entity.Invoice, error) {
	invoice, err := s.invoiceRepo.Find(id)
	if err != nil {
		return invoice, err
	}
	note, err := s.invoiceNoteRepo.Find(invoice.ID)
	if errors.Is(err, repo.ErrRecordNotFound) {
		invoice.Note = nil
	} else if err != nil {
		return invoice, err
	} else {
		invoice.Note = &note
	}
	return invoice, nil
}

func (s invoiceService) Delete(id int) error {
	tx := s.db.Begin()
	err := s.invoiceNoteRepo.WithTx(tx).Delete(id)
	if err != nil && !errors.Is(err, repo.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}
	err = s.invoiceRepo.WithTx(tx).Delete(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// TODO set payment date and write test for this functionality
func (s invoiceService) Create(invoice entity.Invoice) (entity.Invoice, error) {
	tx := s.db.Begin()
	invoice.DateOfPayment = time.Now()
	// Note: if note is in invoice already gorm db will create automatically note in the table
	invoice, err := s.invoiceRepo.WithTx(tx).Create(invoice)
	if err != nil {
		tx.Rollback()
		return invoice, err
	}
	return invoice, tx.Commit().Error
}

func (s invoiceService) Update(invoice entity.Invoice) error {
	tx := s.db.Begin()
	err := s.invoiceRepo.WithTx(tx).Update(invoice)
	if err != nil {
		tx.Rollback()
		return err
	}
	if invoice.Note != nil {
		err = s.invoiceNoteRepo.WithTx(tx).Update(*invoice.Note)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		err = s.invoiceNoteRepo.WithTx(tx).Delete(invoice.ID)
	}
	return tx.Commit().Error
}
