package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type todo_item struct {
	Task         string
	Done         bool
	Created_at   time.Time
	Completed_at time.Time
}

type Todos []todo_item

func (t *Todos) AddTodo(task string) {
	todo := todo_item{
		Task:         task,
		Done:         false,
		Created_at:   time.Now(),
		Completed_at: time.Time{},
	}
	*t = append(*t, todo)
}
func (t *Todos) CompleteTodo(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid")
	}
	ls[index-1].Completed_at = time.Now()
	ls[index-1].Done = true

	return nil

}
func (t *Todos) DeleteTodo(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}
func (t *Todos) Load(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}
func (t *Todos) Store(fileName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}
func (t *Todos) PrintTodo() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "UpdatedAt"},
		},
	}
	var cells [][]*simpletable.Cell
	for i, item := range *t {
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i+1)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.Created_at.Format(time.RFC822)},
			{Text: item.Completed_at.Format(time.RFC822)},
		})
	}
	table.Body = &simpletable.Body{
		Cells: cells,
	}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "All todo are here."},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
