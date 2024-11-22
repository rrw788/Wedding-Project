package models

type Event struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ClientID  uint   `json:"client_id"`
	TypeEvent string `json:"type_event"`
	Venue     string `json:"venue"`
	Budget    int    `json:"budget"`
	Status    string `json:"status"`
}
