package service

import (
	"errors"
	"go-url-shortener/internal/model"
	"go-url-shortener/internal/repository"
	"math/rand"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// Declare the interface
type URLService interface {
	Shorten(original string, expiredAtISO *string) (*model.URL, error)
	GetOriginal(code string) (*model.URL, error)
	Increment(code string) error
}

type urlService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{repo}
}

// Check whether the string is a valid HTTP or HTTPS URL and returns
// Error when the format is invalid
func validateURL(raw string) error {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return errors.New("invalid URL format")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("only HTTP and HTTPS allowed")
	}
	return nil
}

// Generate random short code len 6 char
func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Implement the urlService interface and Shorten method
func (s *urlService) Shorten(original string, expiredAtISO *string) (*model.URL, error) {
	if err := validateURL(original); err != nil {
		return nil, err
	}

	var exp time.Time
	if expiredAtISO != nil {
		t, err := time.Parse(time.RFC3339, *expiredAtISO)
		if err != nil {
			return nil, errors.New("invalid expired_at datetime format")
		}
		exp = t
	} else {
		// default 24 jam dari sekarang
		exp = time.Now().Add(24 * time.Hour)
	}

	url := &model.URL{
		ID:          uuid.New(),
		OriginalURL: original,
		ShortCode:   generateShortCode(),
		ClickCount:  0,
		ExpiresAt:   &exp,
	}

	return url, s.repo.Create(url)
}

func (s *urlService) GetOriginal(code string) (*model.URL, error) {
	return s.repo.FindByCode(code)
}

func (s *urlService) Increment(code string) error {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return err
	}

	return s.repo.IncrementClick(url.ID)
}
