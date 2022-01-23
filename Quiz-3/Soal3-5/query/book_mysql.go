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
	book_table = "book"
)

func GetAllBook(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + book_table + " ORDER BY created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryId,
		); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookByCategory(ctx context.Context, idCat string) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + book_table + " WHERE category_id = " + idCat + " ORDER BY created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryId,
		); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	return books, nil
}

// func FilterBookCategory(ctx context.Context, title, minYear, maxYear, minPage, maxPage, sortByTitle string) ([]models.Book, error) {
// 	var books []models.Book
// 	db, err := config.MySQL()

// 	var titleCondition, minYearCondition, maxYearCondition, minPageCondition, maxPageCondition, sortByTitleCondition string

// 	if err != nil {
// 		log.Fatal("Cannot Connect to MySQL", err)
// 	}

// 	fmt.Println("title = " + title)
// 	fmt.Println("minYear = " + minYear)

// 	if title != "" {
// 		fmt.Println(`Ketika title bukan ""`)
// 		titleCondition = fmt.Sprintf("LOWER(book.title) LIKE LOWER('%v')", title)
// 	}

// 	if minYear != "" {
// 		fmt.Println(`Ketika minYear bukan ""`)
// 		minYearCondition = fmt.Sprintf("release_year > %v", minYear)
// 	}

// 	if maxYear != "" {
// 		maxYearCondition = fmt.Sprintf("release_year < %v", maxYear)
// 	}

// 	if minPage != "" {
// 		minPageCondition = fmt.Sprintf("total_page > %v", minPage)
// 	}

// 	if maxPage != "" {
// 		maxPageCondition = fmt.Sprintf("total_page < %v", maxPage)
// 	}

// 	if strings.ToLower(sortByTitle) == "asc" {
// 		sortByTitleCondition = "ORDER BY title ASC"
// 	} else if strings.ToLower(sortByTitle) == "desc" {
// 		sortByTitleCondition = "ORDER BY title DESC"
// 	}

// 	queryText := "SELECT * FROM " + book_table + " WHERE " + minYearCondition + " " + maxYearCondition + " " + minPageCondition + " " + maxPageCondition + " " + titleCondition + " " + sortByTitleCondition
// 	// queryText := fmt.Sprintf("SELECT * FROM %v", book_table)
// 	rowQuery, err := db.QueryContext(ctx, queryText)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rowQuery.Next() {
// 		var book models.Book
// 		var createdAt, updatedAt string
// 		if err = rowQuery.Scan(
// 			&book.ID,
// 			&book.Title,
// 			&book.Description,
// 			&book.ImageUrl,
// 			&book.ReleaseYear,
// 			&book.Price,
// 			&book.TotalPage,
// 			&book.Thickness,
// 			&createdAt,
// 			&updatedAt,
// 			&book.CategoryId,
// 		); err != nil {
// 			return nil, err
// 		}

// 		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		books = append(books, book)
// 	}
// 	return books, nil
// }

func InsertBook(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ('%v', '%v', '%v', %v, '%v', %v, '%v', NOW(), NOW(), %v)",
		book_table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryId,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(ctx context.Context, book models.Book, idBook string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET title='%v', description='%v', image_url='%v', release_year=%v, price='%v', total_page=%v, thickness='%v', category_id=%v, updated_at=NOW() WHERE id=%v",
		book_table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryId,
		idBook)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(ctx context.Context, idBook string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id=%v",
		book_table, idBook)

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
