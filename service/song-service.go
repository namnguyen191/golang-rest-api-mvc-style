package service

import (
	"github.com/namnguyen191/themuzix-golang-rest-api/dto"
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"github.com/namnguyen191/themuzix-golang-rest-api/repository"
)

type SongService interface {
	InsertSong(dto.SongCreateDTO) entity.Song
	UpdateSong(dto.SongUpdateDTO) entity.Song
	Delete(entity.Song)
	All() []entity.Song
	FindById(songId uint) entity.Song
	IsAllowedToEdit(userID string, bookID uint) bool
}

type songService struct {
	songRepo repository.SongRepository
}

func (s *songService) InsertSong(songDto dto.SongCreateDTO) entity.Song {

	songToInsert := entity.Song{
		Title:       songDto.Title,
		Description: songDto.Description,
		Artist:      entity.Artist{ID: songDto.ArtistID},
	}

	res := s.songRepo.InsertSong(songToInsert)
	return res
}

func (s *songService) UpdateSong(songDto dto.SongUpdateDTO) entity.Song {
	songToUpdate := entity.Song{
		ID:          songDto.ID,
		Title:       songDto.Title,
		Description: songDto.Description,
		Artist:      entity.Artist{ID: songDto.ArtistID},
	}

	res := s.songRepo.UpdateSong(songToUpdate)
	return res
}

func (s *songService) Delete(song entity.Song) {
	panic("not implemented") // TODO: Implement
}

func (s *songService) All() []entity.Song {
	panic("not implemented") // TODO: Implement
}

func (s *songService) FindById(songId uint) entity.Song {
	panic("not implemented") // TODO: Implement
}

func (s *songService) IsAllowedToEdit(userID string, bookID uint) bool {
	panic("not implemented") // TODO: Implement
}

func NewSongService(songRepo repository.SongRepository) SongService {
	return &songService{
		songRepo: songRepo,
	}
}
