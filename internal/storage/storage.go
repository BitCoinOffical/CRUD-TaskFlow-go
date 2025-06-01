package storage

import (
	"main.go/internal/models"
)

func CreateTask(task *models.Tasks) error {
	insert := `INSERT INTO tasks(description, priority, status, title) VALUES (?, ?, ?, ?)`
	res, err := DB.Exec(insert, task.Description, task.Priority, task.Status, task.Title)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = int(id)
	return nil
}

func GetTask() ([]models.Tasks, error) {
	selectall := `SELECT id, title, description, priority, status FROM tasks ORDER BY id DESC`
	rows, err := DB.Query(selectall)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Tasks
	for rows.Next() {
		var task models.Tasks
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTaskByID(id int, status string) error {
	update := `UPDATE tasks SET status = ? WHERE id = ?`
	_, err := DB.Exec(update, status, id)
	return err
}

func DeleteTask(ReceivedId int) (bool, error) {
	delete := `DELETE FROM tasks WHERE id = ?`
	res, err := DB.Exec(delete, ReceivedId)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil

}
