package phonecard

type Repository interface {
	AddCard(Phonecard) error
	ActiveCard(code string, phoneNumber string) error
}
