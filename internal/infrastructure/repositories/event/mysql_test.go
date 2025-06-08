package event

import (
	"context"
	"database/sql"
	"errors"
	"goevents/internal/domain/entities"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *EventRepositoryMySQL) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := NewEventRepositoryMySQL(db).(*EventRepositoryMySQL)
	return db, mock, repo
}

func TestEventRepositoryMySQL_Create(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	event := &entities.Event{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        time.Now(),
	}

	mock.ExpectExec("INSERT INTO events").
		WithArgs(event.Title, event.Description, event.Date).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.TODO(), event)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Create_Error(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	event := &entities.Event{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        time.Now(),
	}

	mock.ExpectExec("INSERT INTO events").
		WithArgs(event.Title, event.Description, event.Date).
		WillReturnError(errors.New("error"))

	err := repo.Create(context.TODO(), event)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_FindAll(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "date"}).
		AddRow("1", "Event 1", "Desc 1", time.Now()).
		AddRow("2", "Event 2", "Desc 2", time.Now())

	mock.ExpectQuery("SELECT \\* FROM events").
		WillReturnRows(rows)

	events, err := repo.FindAll(context.TODO())

	assert.NoError(t, err)
	assert.Len(t, events, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_FindAll_Error(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM events").
		WillReturnError(errors.New("error"))

	events, err := repo.FindAll(context.TODO())
	assert.Error(t, err)
	assert.Nil(t, events)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_FindAll_Error_TimeParse(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "date"}).
		AddRow("1", "Event 1", "Desc 1", "not valid datetime")

	mock.ExpectQuery("SELECT \\* FROM events").
		WillReturnRows(rows)

	events, err := repo.FindAll(context.TODO())
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestEventRepositoryMySQL_Find(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "date"}).
		AddRow("1", "Event 1", "Desc 1", time.Now())

	mock.ExpectQuery("SELECT \\* FROM events WHERE id = \\?").
		WithArgs("1").
		WillReturnRows(rows)

	event, err := repo.Find(context.TODO(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Find_Error(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM events WHERE id = \\?").
		WithArgs("1").
		WillReturnError(errors.New("error"))

	event, err := repo.Find(context.TODO(), "1")
	assert.Error(t, err)
	assert.Nil(t, event)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Update(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	event := &entities.Event{
		ID:          "1",
		Title:       "Updated Event",
		Description: "Updated Description",
		Date:        time.Now(),
	}

	mock.ExpectExec("UPDATE events").
		WithArgs(event.Title, event.Description, event.Date, event.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(context.TODO(), event)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Update_Error(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	event := &entities.Event{
		ID:          "1",
		Title:       "Updated Event",
		Description: "Updated Description",
		Date:        time.Now(),
	}

	mock.ExpectExec("UPDATE events").
		WithArgs(event.Title, event.Description, event.Date, event.ID).
		WillReturnError(errors.New("error"))

	err := repo.Update(context.TODO(), event)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Delete(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM events").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(context.TODO(), "1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEventRepositoryMySQL_Delete_Error(t *testing.T) {
	db, mock, repo := setupTest(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM events").
		WithArgs("1").
		WillReturnError(errors.New("error"))

	err := repo.Delete(context.TODO(), "1")
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
