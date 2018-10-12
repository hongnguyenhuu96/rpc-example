package types

// Make a new ToDo type that is a typed collection of fields
// (Title and Status), both of which are of type string

type ToDo struct {
	Title, Status string
}

type EditToDo struct {
	Title, NewTitle, NewStatus string
}
