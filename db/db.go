package db

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type store struct {
	LastID int    `json:"lastid"`
	Todos  []Todo `json:"todos"`
}

const dbFileName = ".todo.json"

func loadStore() (*store, error) {
	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		return &store{LastID: 0, Todos: []Todo{}}, nil
	}

	file, err := os.ReadFile(dbFileName)
	if err != nil {
		return nil, fmt.Errorf("ファイル読み込み失敗: %w", err)
	}

	var s store
	if len(file) == 0 {
		return &store{LastID: 0, Todos: []Todo{}}, nil
	}

	if err := json.Unmarshal(file, &s); err != nil {
		return nil, fmt.Errorf("JSONデータが破損しているか形式が古いです。%s を削除して初期化することを検討してください: %w", dbFileName, err)
	}
	return &s, nil
}

func saveStore(s *store) error {
	data, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		return fmt.Errorf("JSON変換失敗: %w", err)
	}

	if err := os.WriteFile(dbFileName, data, 0644); err != nil {
		return fmt.Errorf("ファイル書き込み失敗: %w", err)
	}
	return nil
}

func GetTodos() ([]Todo, error) {
	s, err := loadStore()
	if err != nil {
		return nil, err
	}

	return s.Todos, nil
}

func AddTodo(task string) (*Todo, error) {
	s, err := loadStore()
	if err != nil {
		return nil, err
	}

	s.LastID++
	newID := s.LastID

	newTodo := Todo{
		ID:        newID,
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	s.Todos = append(s.Todos, newTodo)
	if err := saveStore(s); err != nil {
		return nil, err
	}
	return &newTodo, nil
}

func DeleteTodo(id int) error {
	s, err := loadStore()
	if err != nil {
		return err
	}

	newTodos := []Todo{}
	found := false
	for _, todo := range s.Todos {
		if todo.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, todo)
	}

	if !found {
		return fmt.Errorf("ID %d のタスクは見つかりませんでした。", id)
	}

	s.Todos = newTodos
	return saveStore(s)
}

func CompleteTodo(id int) (*Todo, error) {
	s, err := loadStore()
	if err != nil {
		return nil, err
	}

	var targetTodo *Todo
	for i := range s.Todos {
		if s.Todos[i].ID == id {
			s.Todos[i].Done = true
			targetTodo = &s.Todos[i]
			break
		}
	}

	if targetTodo == nil {
		return nil, fmt.Errorf("ID %d のタスクは見つかりませんでした。", id)
	}

	if err := saveStore(s); err != nil {
		return nil, err
	}

	return targetTodo, nil
}
