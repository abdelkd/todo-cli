package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/abdelkd/todo-cli/internal/relativetime"
	"github.com/abdelkd/todo-cli/internal/todo"
)

type FileModel struct {
	Path string
}

type FileModelSchema struct {
	Todos []todo.Todo `json:"todos"`
}

func openOrInitFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(path)
			if err != nil {
				return nil, err
			}

			file.Write([]byte("{}"))
		} else {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func marshalAndSave(file *os.File, jsonContent any) error {
	jsonBytes, err := json.Marshal(jsonContent)
	if err != nil {
		return err
	}

	file.Seek(0, 0)
	writtenBytes, err := file.Write(jsonBytes)
	if err != nil {
		return err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	if err = file.Truncate(int64(writtenBytes + 1)); err != nil {
		return err
	}

	return nil
}

func getTodos(file *os.File) (FileModelSchema, error) {
	var jsonContent FileModelSchema
	todosContent, err := io.ReadAll(file)
	if err != nil {
		return jsonContent, err
	}

	err = json.Unmarshal(todosContent, &jsonContent)

	return jsonContent, err
}

func (f FileModel) AddItem(content string) error {
	file, err := openOrInitFile(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	todosContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var jsonContent FileModelSchema
	err = json.Unmarshal(todosContent, &jsonContent)
	if err != nil {
		return err
	}

	myTodo := todo.Todo{
		Id:        len(jsonContent.Todos),
		Content:   content,
		CreatedAt: time.Now(),
		IsDone:    false,
	}

	for _, todoEntry := range jsonContent.Todos {
		if strings.Compare(todoEntry.Content, content) == 0 {
			fmt.Println("A todo with the same name already exists")
			// TODO: an AlreadyExists error
			return nil
		}
	}

	jsonContent.Todos = append(jsonContent.Todos, myTodo)
	if err := marshalAndSave(file, jsonContent); err != nil {
		return err
	}

	return nil
}

func (f FileModel) RemoveItem(id int) error {
	file, err := openOrInitFile(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContent, err := getTodos(file)
	if len(jsonContent.Todos) == 0 {
		fmt.Println("The todo list is empty")
		return nil
	}

	found := false
	for i, k := range jsonContent.Todos {
		if k.Id == id {
			jsonContent.Todos = append(jsonContent.Todos[:i], jsonContent.Todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("Item was not found")
	}

	if err := marshalAndSave(file, jsonContent); err != nil {
		return err
	}

	return nil
}

func (f FileModel) ListItems() error {
	file, err := openOrInitFile(f.Path)
	if err != nil {
		return err
	}

	jsonData, err := getTodos(file)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 4, 2, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "ID\tName\tIs Created\tCreated At")
	fmt.Fprintln(w, "\t\t\t")

	for _, todo := range jsonData.Todos {
		fmt.Fprintf(w, "%d\t%s\t%t\t%s\n", todo.Id, todo.Content, todo.IsDone, relativetime.RelativeTime(time.Now()))

	}

	fmt.Fprintln(w, "\t\t\t")

	return w.Flush()
}

func (f FileModel) ToggleItem(id int) error {
	return nil
}

func (f FileModel) EditItem(id int) error {
	return nil
}
