package entities

type GenderType string

const (
	GenderMale   GenderType = "male"
	GenderFemale GenderType = "female"
	GenderOther  GenderType = "other"
)

type TicketOwner struct {
	TicketID    int64      `json:"ticket_id"`    // ID của vé
	FirstName   string     `json:"first_name"`   // Tên
	LastName    string     `json:"last_name"`    // Họ
	PhoneNumber string     `json:"phone_number"` // Số điện thoại
	Gender      GenderType `json:"gender"`       // Giới tính
}
