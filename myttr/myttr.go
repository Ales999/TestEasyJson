package myttr

//go:generate go run github.com/mailru/easyjson/...@latest .

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//easyjson:json
type MyTtr struct {
	ID   int    `json:"id"`
	Sotr Person `json:"sotr"`
}

func NewMyTtr(id int, name string, age int) *MyTtr {
	pers := Person{
		Name: name,
		Age:  age,
	}
	return &MyTtr{
		ID:   id,
		Sotr: pers,
	}
}
