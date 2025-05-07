package services

import (
	"errors"
	"sync"
	"time"
	"todo-app/models"
)

type TodoService struct {
	todos      map[uint]*models.TodoList
	items      map[uint]*models.TodoItem
	lastID     uint
	lastItemID uint
	mutex      sync.RWMutex
}

var todoService *TodoService
var once sync.Once

func GetTodoService() *TodoService {
	once.Do(func() {
		todoService = &TodoService{
			todos:      make(map[uint]*models.TodoList),
			items:      make(map[uint]*models.TodoItem),
			lastID:     0,
			lastItemID: 0,
		}
	})
	return todoService
}

func (s *TodoService) CreateTodo(todo *models.TodoList) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lastID++
	todo.ID = s.lastID
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	todo.Completion = 0
	todo.Items = make([]models.TodoItem, 0)
	s.todos[todo.ID] = todo
	return nil
}

func (s *TodoService) GetTodo(id uint) (*models.TodoList, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	todo, exists := s.todos[id]
	if !exists || todo.DeletedAt != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *TodoService) GetAllTodos(userID uint, isAdmin bool) []*models.TodoList {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var result []*models.TodoList
	for _, todo := range s.todos {
		if todo.DeletedAt == nil && (isAdmin || todo.UserID == userID) {
			result = append(result, todo)
		}
	}
	return result
}

func (s *TodoService) UpdateTodo(todo *models.TodoList) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existing, exists := s.todos[todo.ID]
	if !exists || existing.DeletedAt != nil {
		return errors.New("todo not found")
	}

	todo.CreatedAt = existing.CreatedAt
	todo.UpdatedAt = time.Now()
	todo.Items = existing.Items

	s.todos[todo.ID] = todo
	return nil
}

func (s *TodoService) DeleteTodo(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	todo, exists := s.todos[id]
	if !exists || todo.DeletedAt != nil {
		return errors.New("todo not found")
	}

	now := time.Now()
	todo.DeletedAt = &now

	// Todo'ya ait tüm itemları da sil
	for _, item := range todo.Items {
		if item.DeletedAt == nil {
			item.DeletedAt = &now
			s.items[item.ID] = &item
		}
	}

	return nil
}

func (s *TodoService) CreateTodoItem(todoID uint, item *models.TodoItem) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	todo, exists := s.todos[todoID]
	if !exists || todo.DeletedAt != nil {
		return errors.New("todo not found")
	}

	s.lastItemID++
	item.ID = s.lastItemID
	item.TodoListID = todoID
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	s.items[item.ID] = item
	todo.Items = append(todo.Items, *item)

	s.updateTodoCompletion(todo)
	return nil
}

func (s *TodoService) UpdateTodoItem(item *models.TodoItem) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existing, exists := s.items[item.ID]
	if !exists || existing.DeletedAt != nil {
		return errors.New("todo item not found")
	}

	item.CreatedAt = existing.CreatedAt
	item.UpdatedAt = time.Now()
	s.items[item.ID] = item

	// Todo listesini güncelle
	if todo, exists := s.todos[item.TodoListID]; exists {
		for i, existingItem := range todo.Items {
			if existingItem.ID == item.ID {
				todo.Items[i] = *item
				break
			}
		}
		s.updateTodoCompletion(todo)
	}

	return nil
}

func (s *TodoService) DeleteTodoItem(todoID, itemID uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	item, exists := s.items[itemID]
	if !exists || item.DeletedAt != nil || item.TodoListID != todoID {
		return errors.New("todo item not found")
	}

	now := time.Now()
	item.DeletedAt = &now
	s.items[itemID] = item

	// Todo listesini güncelle
	if todo, exists := s.todos[todoID]; exists {
		for i, existingItem := range todo.Items {
			if existingItem.ID == itemID {
				todo.Items[i].DeletedAt = &now
				break
			}
		}
		s.updateTodoCompletion(todo)
	}

	return nil
}

func (s *TodoService) updateTodoCompletion(todo *models.TodoList) {
	total := 0
	completed := 0
	for _, item := range todo.Items {
		if item.DeletedAt == nil {
			total++
			if item.IsCompleted {
				completed++
			}
		}
	}
	if total > 0 {
		todo.Completion = float32(completed) / float32(total) * 100
	} else {
		todo.Completion = 0
	}
}
