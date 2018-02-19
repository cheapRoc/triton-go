package authentication

import "time"

type ManualSigner struct {
	dateHeader string
	authHeader string
}

func (s *ManualSigner) SetDate(dateHeader string) {
	s.dateHeader = dateHeader
}

func (s *ManualSigner) SetSign(authHeader string) {
	s.authHeader = authHeader
}

func (s *ManualSigner) Date() string {
	if s.dateHeader == "" {
		s.dateHeader = time.Now().UTC().Format(time.RFC1123)
	}
	return s.dateHeader
}

// Sign simply passes through the manually set authentication signature header.
func (s *ManualSigner) Sign() (string, error) {
	return s.authHeader, nil
}
