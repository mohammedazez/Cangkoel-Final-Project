package admin

import (
	"backend/entity"
	"backend/helper"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SShowAllAdmin() ([]AdminFormat, error)
	SRegisterAdmin(admin entity.AdminInput) (AdminFormat, error)
	SFindAdminByID(adminID string) (AdminFormat, error)
	SDeleteAdminByID(adminID string) (interface{}, error)
	SUpdateAdminByID(adminID string, input entity.UpdateAdminInput) (AdminFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SShowAllAdmin() ([]AdminFormat, error) {
	admins, err := s.repository.RShowAllAdmin()
	var formatadmins []AdminFormat

	for _, admin := range admins {
		formatadmin := Format(admin)
		formatadmins = append(formatadmins, formatadmin)

	}
	if err != nil {
		return formatadmins, err
	}

	return formatadmins, nil
}

func (s *service) SRegisterAdmin(admin entity.AdminInput) (AdminFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.MinCost)

	if err != nil {
		return AdminFormat{}, err
	}

	var newAdmin = entity.Admin{
		FullName:  admin.FullName,
		Email:     admin.Email,
		Password:  string(genPassword),
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createAdmin, err := s.repository.RRegisterAdmin(newAdmin)
	formatAdmin := Format(createAdmin)

	if err != nil {
		return formatAdmin, err
	}

	return formatAdmin, nil
}

func (s *service) SFindAdminByID(adminID string) (AdminFormat, error) {
	admin, err := s.repository.RFindAdminByID(adminID)

	if err != nil {
		return AdminFormat{}, err
	}

	if admin.ID == 0 {
		newError := fmt.Sprintf("Admin id %s not found", adminID)
		return AdminFormat{}, errors.New(newError)
	}

	formatAdmin := Format(admin)

	return formatAdmin, nil
}

func (s *service) SDeleteAdminByID(adminID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(adminID); err != nil {
		return nil, err
	}

	admin, err := s.repository.RFindAdminByID(adminID)

	if err != nil {
		return nil, err
	}

	if admin.ID == 0 {
		newError := fmt.Sprintf("Admin id %s not found", adminID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.RDeleteAdminByID(adminID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete Admin ID : %s", adminID)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}

func (s *service) SUpdateAdminByID(adminID string, input entity.UpdateAdminInput) (AdminFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(adminID); err != nil {
		return AdminFormat{}, err
	}

	admin, err := s.repository.RFindAdminByID(adminID)

	if err != nil {
		return AdminFormat{}, err
	}

	if admin.ID == 0 {
		newError := fmt.Sprintf("Admin id %s not found", adminID)
		return AdminFormat{}, errors.New(newError)
	}

	if input.FullName != "" || len(input.FullName) != 0 {
		dataUpdate["FullName"] = input.FullName
	}
	if input.Email != "" || len(input.Email) != 0 {
		dataUpdate["Email"] = input.Email
	}
	if input.Password != "" || len(input.Password) != 0 {
		dataUpdate["Passowrd"] = input.Password
	}
	dataUpdate["updated_at"] = time.Now()

	adminUpdated, err := s.repository.RUpdateAdminByID(adminID, dataUpdate)

	if err != nil {
		return AdminFormat{}, err
	}

	formatAdmin := Format(adminUpdated)

	return formatAdmin, nil
}