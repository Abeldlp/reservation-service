package models

type Appointment struct {
	ID              int64  `json:"id"`
	Priority        int64  `json:"priority"`
	AppointmentDate string `json:"appointment_date"`
	Description     string `json:"description"`
	DoctorID        int64  `json:"doctor_id"`
	Doctor          Doctor `json:"doctor" gorm:"foreignKey:DoctorID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewAppointment() Appointment {
	return Appointment{
		Priority:        0,
		AppointmentDate: "",
		Description:     "",
	}
}
