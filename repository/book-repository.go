package repository

import (
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"gorm.io/gorm"
)

type SongRepository interface {
	InsertSong(song entity.Song) entity.Song
	UpdateSong(song entity.Song) entity.Song
	DeleteSong(song entity.Song)
	AllSongs() []entity.Song
	FindSongById(songId uint) entity.Song
}

type songRepository struct {
	dbConnection *gorm.DB
}

func (songRepo *songRepository) InsertSong(song entity.Song) entity.Song {
	songRepo.dbConnection.Save(&song)
	songRepo.dbConnection.Preload("Song").Find(&song)
	return song
}

func (songRepo *songRepository) UpdateSong(song entity.Song) entity.Song {
	songRepo.dbConnection.Save(&song)
	songRepo.dbConnection.Preload("Song").Find(&song)
	return song
}

func (songRepo *songRepository) DeleteSong(song entity.Song) {
	songRepo.dbConnection.Delete(&song)
}

func (songRepo *songRepository) AllSongs() []entity.Song {
	var allSongs []entity.Song
	songRepo.dbConnection.Preload("Song").Find(&allSongs)
	return allSongs
}

func (songRepo *songRepository) FindSongById(songId uint) entity.Song {
	var foundSong entity.Song
	songRepo.dbConnection.Preload("Song").Find(&foundSong, songId)
	return foundSong
}

func NewSongRepository(dbCon *gorm.DB) SongRepository {
	return &songRepository{
		dbConnection: dbCon,
	}
}
