package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/models"
	"time"
)

const (
	category_table = "category"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllCategory(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + category_table + " ORDER BY created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category models.Category
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&category.ID,
			&category.Name,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func InsertCategory(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) VALUES ('%v', NOW(), NOW())",
		category_table,
		category.Name)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(ctx context.Context, category models.Category, idCat string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET name='%v', updated_at=NOW() WHERE id=%v",
		category_table,
		category.Name,
		idCat)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(ctx context.Context, idCat string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id=%v",
		category_table, idCat)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
