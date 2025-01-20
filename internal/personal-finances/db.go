package personalfinances

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func NewExpenses(description string, value float64, dueDate time.Time, payDate time.Time, expensestype string, subType string, status bool) *Expenses {
	return &Expenses{
		ID:          uuid.New().String(),
		Description: description,
		Value:       value,
		DueDate:     dueDate,
		PayDate:     payDate,
		Type:        expensestype,
		SubType:     subType,
		Status:      status,
	}
}

func InsertExpenses(db *sql.DB, expenses Expenses) error {
	stmt, err := db.Prepare("INSERT INTO expenses (id, description, value, due_date, pay_date, type,  sub_type, status) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		expenses.ID,
		expenses.Description,
		expenses.Value,
		expenses.DueDate,
		expenses.PayDate,
		expenses.Type,
		expenses.SubType,
		expenses.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateExpenses(db *sql.DB, expenses *Expenses) error {
	stmt, err := db.Prepare("UPDATE expenses SET description = ?, value = ?, due_date = ?, pay_date = ?, type = ?,  sub_type = ?, status = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		expenses.ID,
		expenses.Description,
		expenses.Value,
		expenses.DueDate,
		expenses.PayDate,
		expenses.Type,
		expenses.SubType,
		expenses.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllExpenses(db *sql.DB) ([]Expenses, error) {
	rows, err := db.Query("SELECT id, description, value, due_date, pay_date, type,  sub_type, status FROM expenses ")
	if err != nil {
		return []Expenses{}, err
	}
	defer rows.Close()

	var expensesList []Expenses

	for rows.Next() {
		var expenses Expenses
		err = rows.Scan(
			&expenses.ID,
			&expenses.Description,
			&expenses.Value,
			&expenses.DueDate,
			&expenses.PayDate,
			&expenses.Type,
			&expenses.SubType,
			&expenses.Status,
		)

		if err != nil {
			return []Expenses{}, err
		}

		expensesList = append(expensesList, expenses)
	}

	return expensesList, nil
}

func SelecOneExpenses(id string, db *sql.DB) (*Expenses, error) {
	stmt, err := db.Prepare("SELECT id, description, value, due_date, pay_date, type,  sub_type, status FROM expenses WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var e Expenses
	err = stmt.QueryRow(id).Scan(
		&e.ID,
		&e.Description,
		&e.Value,
		&e.DueDate,
		&e.PayDate,
		&e.Type,
		&e.SubType,
		&e.Status,
	)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func DeleteExpenses(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM expenses WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
