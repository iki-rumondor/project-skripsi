package response

import "time"

type InstrumenType struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IndikatorType struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Instrument  string    `json:"instrument"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Indikator struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	TypeID      uint      `json:"type_id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
