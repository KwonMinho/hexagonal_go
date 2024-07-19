package vo

type IDName struct {
	ID string
	Name string
}


func CreateIDName(id, name string) IDName {
	return IDName{ID: id, Name: name}
}

