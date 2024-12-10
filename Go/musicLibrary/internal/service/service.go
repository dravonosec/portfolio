package service

import (
	"encoding/json"
	"fmt"
	"musicLibrary/internal/models"
	"musicLibrary/internal/repository"
	"net/http"
	"net/url"
	"time"
)

type Service interface {
	CreateSong(song *models.Song) error
	GetSongById(id int64) (*models.Song, error)
	UpdateSong(song *models.Song) error
	DeleteSongById(id int64) error
	ListSongs(offset, limit int) ([]models.Song, error)
}

type MusicService struct {
	repo repository.Repository
}

func NewMusicService(repo repository.Repository) *MusicService {
	return &MusicService{repo: repo}
}

func (s *MusicService) CreateSong(song *models.Song) error {
	if err := validateSong(song); err != nil {
		return err
	}

	err := enrichSongData(song)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return s.repo.CreateSong(song)
}

func (s *MusicService) GetSongById(id int64) (*models.Song, error) {
	song, err := s.repo.GetSongById(id)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return song, nil
}

func (s *MusicService) UpdateSong(song *models.Song) error {
	if err:=validateSong(song); err != nil {
		return err;
	}
	return s.repo.UpdateSong(song)
}

func (s *MusicService) DeleteSongById(id int64) error {
	return s.repo.DeleteSong(id)
}

func (s *MusicService) ListSongs(offset int, limit int) ([]models.Song, error) {
	return s.repo.ListSongs(offset, limit)
}

func validateSong(song *models.Song) error {
	if song.Group == "" || song.Name == "" {
		return fmt.Errorf("song must contain group and name")
	}
	return nil
}

func enrichSongData(song *models.Song) error {
    client := &http.Client {
		Timeout: time.Second * 10,
	}
	
    baseURL := "http://api.example.com/info"
	params := url.Values {}
	params.Add("group", song.Group)
	params.Add("song", song.Name)

	req, err := http.NewRequest("GET", baseURL + "?" + params.Encode(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err);
	}

	resp, err := client.Do(req)
	if err!= nil {
        return fmt.Errorf("failed to execute request: %v", err);
    }

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get song info: %s", resp.StatusCode)
	}

	var SongDetails struct {
		ReleaseDate string `json:"release"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&SongDetails); err != nil {
		return fmt.Errorf("failed to decode response: %v", err);
	}

	song.ReleaseDate = SongDetails.ReleaseDate
	song.Text = SongDetails.Text
	song.Link = SongDetails.Link

	return nil
}

