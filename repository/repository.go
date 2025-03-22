package repository

import (
	"database/sql"
	"fmt"
	"hello_world/classes"
	"hello_world/databases"

	_ "github.com/lib/pq"
)

func ReturnValues() string {
	return "Hello, World!"
}

func AlbunsByArtist(name string) ([]classes.Album, error) {
	db, err := databases.ReturnFromDB()
	if err != nil {
		fmt.Println("Error connecting!")
		return nil, fmt.Errorf("failed to connect to database")
	}
	fmt.Println("Successfully connected!")
	var albuns []classes.Album

	rows, err := db.Query("SELECT * FROM album WHERE artist LIKE %?%", name)
	if err != nil {
		fmt.Println("My errors:")

		fmt.Printf("albumsByArtist %q: %v", name, err)
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb classes.Album
		if err := rows.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price); err != nil {

			fmt.Println("Current Errors:")

			fmt.Printf("albumsByArtist %q: %v", name, err)
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albuns = append(albuns, alb)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Last errors:")

		fmt.Printf("albumsByArtist %q: %v", name, err)
		return nil, fmt.Errorf("albunsByArtist %q: %v", name, err)
	}
	fmt.Println(albuns)
	return albuns, nil
}

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{
		db: db,
	}
}

func (repo *ArtistRepository) GetAlbum() ([]classes.Album, error) {

	query := "SELECT * FROM album"
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	var albuns []classes.Album

	for rows.Next() {
		var alb classes.Album
		err := rows.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		albuns = append(albuns, alb)
	}

	return albuns, nil
}

func (repo *ArtistRepository) GetAlbumById(id int) (classes.Album, error) {
	queryGet := "SELECT * FROM album WHERE id = $1"
	result, err := repo.db.Query(queryGet, id)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error executing query: %v", err)
	}
	var album classes.Album
	for result.Next() {
		var alb classes.Album
		myerr := result.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price)
		if myerr != nil {
			fmt.Printf("My error: %v", myerr)
			return classes.Album{}, fmt.Errorf("error scanning row: %v", err)
		}
		album = alb
	}
	return album, nil
}

func (repo *ArtistRepository) RegisterAlbum(album classes.Album) (classes.Album, error) {
	query := "INSERT INTO album (title, artist, price) VALUES($1, $2, $3) RETURNING *"
	result, err := repo.db.Query(query, album.Title, album.Artist, album.Price)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error inserting album: %v", err)
	}
	var mAlbum classes.Album
	for result.Next() {
		var alb classes.Album
		myerr := result.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price)
		if myerr != nil {
			fmt.Printf("My error: %v", myerr)
			return classes.Album{}, fmt.Errorf("error scanning row: %v", err)
		}
		mAlbum = alb
	}
	return mAlbum, nil
}

func (repo *ArtistRepository) UpdateAlbum(id int, album classes.Album) (classes.Album, error) {
	query := "UPDATE album SET title = $1, artist = $2, price = $3 WHERE id = $4 RETURNING *"
	result, err := repo.db.Query(query, album.Title, album.Artist, album.Price, id)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error updating album: %v", err)
	}
	var mAlbum classes.Album
	for result.Next() {
		var alb classes.Album
		myErr := result.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price)
		if myErr != nil {
			return classes.Album{}, fmt.Errorf("error scanning row: %v", myErr)
		}
		mAlbum = alb
	}
	return mAlbum, nil
}

func (repo *ArtistRepository) DeleteAlbum(id int) (classes.Album, error) {
	album, err := repo.GetAlbumById(id)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error getting the album for id %d: %s", id, err)
	}

	query := "DELETE FROM album WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error deleting the album: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return classes.Album{}, fmt.Errorf("no register encountered: %s", err)
	}
	if rowsAffected == 0 {
		return classes.Album{}, fmt.Errorf("no album encountered for id %d: %s", id, err)
	}
	return album, nil
}
