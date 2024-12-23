package main

import (
	"encoding/json"
	"fmt"

	"github.com/ales999/TestEasyJson/myttr"
)

func main() {

	var dataWr = []byte(`{"id": 1,"sotr": {"name": "Alexey","age": 30}}`)

	var t myttr.MyTtr
	//best := myttr.NewMyTtr(1, "Alexey", 30)
	//PrintAppJSON(best)

	fmt.Println("Before:", t)

	t.UnmarshalJSON(dataWr)

	tstb, err := t.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("After:", string(tstb))
	fmt.Println("-------")
	PrintAppJSON(t)

}

func PrintAppJSON(structToPrint interface{}) error {

	jsonContent, err := json.MarshalIndent(structToPrint, "", "  ")
	if err != nil {
		fmt.Println("Error Marshal to JSON: ", err)
		return err
	}
	// конвертируем []byte в строку и затем печатаем
	fmt.Printf("%+v\n", string(jsonContent))
	return nil
}
