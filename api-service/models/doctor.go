package models

type Doctor struct {
	ID           int64      `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	PhoneNumber  string     `json:"phone_number"`
	SpecialityID int64      `json:"speciality_id"`
	Speciality   Speciality `json:"speciality" gorm:"foreignKey:SpecialityID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewDoctor() Doctor {
	return Doctor{}
}
