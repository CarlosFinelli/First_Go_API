package service

import (
	"fmt"
	"hello_world/classes"
	"hello_world/databases"
	"hello_world/repository"
)

func ReturnText() string {
	return repository.ReturnValues()
}

func ReturnArtist(name string) ([]classes.Album, error) {
	db, err := databases.ReturnFromDB()
	if err != nil {
		return nil, fmt.Errorf("error fetching albuns: %f", err)
	}
	repo := repository.NewArtistRepository(db)
	return repo.GetAlbum()
	//return repository.AlbunsByArtist(name)
}

func ReturnById(id int) (classes.Album, error) {
	db, err := databases.ReturnFromDB()
	if err != nil {
		return classes.Album{}, fmt.Errorf("error fetching albuns: %f", err)
	}
	repo := repository.NewArtistRepository(db)
	return repo.GetAlbumById(id)
}

func RegisterAlbum(album classes.Album) (classes.Album, error) {
	db, err := databases.ReturnFromDB()
	if err != nil {
		return classes.Album{}, fmt.Errorf("error registering album: %v", err)
	}
	repo := repository.NewArtistRepository(db)
	return repo.RegisterAlbum(album)
}

func UpdateAlbum(id int, album classes.Album) (classes.Album, error) {
	db, err := databases.ReturnFromDB()
	if err != nil {
		return classes.Album{}, fmt.Errorf("error connecting to database: %v", err)
	}
	repo := repository.NewArtistRepository(db)
	return repo.UpdateAlbum(id, album)
}

func DeleteAlbum(id int) (classes.Album, error) {
	fmt.Printf("Id: %d", id)
	db, err := databases.ReturnFromDB()
	if err != nil {
		return classes.Album{}, fmt.Errorf("error deleting album: %f", err)
	}
	repo := repository.NewArtistRepository(db)
	album, err := repo.DeleteAlbum(id)
	if err != nil {
		return classes.Album{}, fmt.Errorf("error deleting album: %f", err)
	}
	return album, nil
}
