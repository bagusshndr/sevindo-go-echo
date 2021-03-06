package models

import "time"

type UserAdmin struct {
	Id                   string     `json:"id" validate:"required"`
	CreatedBy            string     `json:"created_by" validate:"required"`
	CreatedDate          time.Time  `json:"created_date" validate:"required"`
	ModifiedBy           *string    `json:"modified_by"`
	ModifiedDate         *time.Time `json:"modified_date"`
	DeletedBy            *string    `json:"deleted_by"`
	DeletedDate          *time.Time `json:"deleted_date"`
	IsDeleted            int        `json:"is_deleted" validate:"required"`
	IsActive             int        `json:"is_active" validate:"required"`
	Email 				string `json:"email"`
	FullName             string     `json:"full_name"`
	BranchId 			 *string `json:"branch_id"`
}

type UserAdminDto struct {
	Id                   string     `json:"id" validate:"required"`
	Email 				string `json:"email"`
	FullName             string     `json:"full_name"`
	BranchId 			 *string `json:"branch_id"`
}
type UserAdminInfoDto struct {
	Id                   string     `json:"id" validate:"required"`
	Email 				string `json:"email"`
	FullName             string     `json:"full_name"`
	BranchId 			 *string `json:"branch_id"`
	Roles 				string `json:"roles"`
}
type NewCommandUserAdmin struct {
	Id                   string     `json:"id" validate:"required"`
	Email 				string `json:"email"`
	Password 			string `json:"password"`
	FullName             string     `json:"full_name"`
	BranchId 			 *string `json:"branch_id"`
}

type UserAdminWithPagination struct {
	Data []*UserAdminDto  `json:"data"`
	Meta *MetaPagination `json:"meta"`
}