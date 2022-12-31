package models

type Speciality struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewSpeciality() Speciality {
	return Speciality{}
}
