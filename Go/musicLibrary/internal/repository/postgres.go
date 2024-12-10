package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"musicLibrary/internal/config"
	"musicLibrary/internal/models"
)

type Repository interface {
	CreateSong(song *models.Song) error
	GetSongById(id int64) (*models.Song, error)
	UpdateSong(song *models.Song) error
	DeleteSong(id int64) error
	ListSongs(offset, limit int) ([]models.Song, error)
}

type PostgresRepository struct {
	db *pgx.Conn
}

func NewRepository(config *config.Config) *PostgresRepository {
	conn, err := pgx.Connect(context.Background(), config.GetMusicDBConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return &PostgresRepository{db: conn}
}

func (r *PostgresRepository) CreateSong(song *models.Song) error {
	query := `
	INSERT INTO songs (group, name, releaseDate, text, link)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	err := r.db.QueryRow(context.Background(), query, 
		song.Name, 
		song.Group, 
		song.ReleaseDate, 
		song.Text, 
		song.Link,
	).Scan(&song.ID)
	if (err != nil) { 
		return err
	}	

	return nil
}

func (r *PostgresRepository) GetSongById(id int64) (*models.Song, error) {
	query := `
	SELECT id, name, group, releaseDate, text, link 
	FROM song 
	WHERE id = $1;
	`
	song := &models.Song{}

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&song.ID,
        &song.Name,
        &song.Group,
        &song.ReleaseDate,
        &song.Text,
        &song.Link)
	if (err != nil) { 
        return nil, err
    }

	return song, nil
}

func (r *PostgresRepository) UpdateSong(song *models.Song) error {
	query := `
	UPDATE song
	SET name = $1, group = $2, releaseDate = $3, text = $4, link = $5
	WHERE id = $6
	`

	result, err := r.db.Exec(context.Background(), query)
	if err != nil { 
		return err
	}

	rows := result.RowsAffected()
	if rows == 0 {
        return fmt.Errorf("song not found. No rows affected")
    }

	return nil
}

func (r *PostgresRepository) DeleteSong(id int64) error {
	query := `
	DELETE FROM songs WHERE id = $1
	`

	result, err := r.db.Exec(context.Background(), query, id)
	if err!= nil { 
        return err
    }

	rows := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("song not found. No rows affected")
	}

	return nil
}

func (r *PostgresRepository) ListSongs(offset, limit int) ([]models.Song, error) {
	query := `
		SELECT id, name, group, releseDate, text, link
		FROM songs
		ORDER BY id
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() { 
		song := &models.Song{}
		err := rows.Scan(
			&song.ID, 
			&song.Name, 
			&song.Group, 
			&song.ReleaseDate, 
			&song.Text, 
			&song.Link)
		
		if err!= nil {
            return nil, err
        }
		songs = append(songs, *song)
	}


	return songs, nil
}
