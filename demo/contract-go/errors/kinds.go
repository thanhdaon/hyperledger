package errors

type Kind struct {
	slug string
}

var (
	KUnknown  = Kind{}
	KBadInput = Kind{"bad input"}
	KNotFound = Kind{"not found"}
)
