package authentication

type TestSigner struct{}

func NewTestSigner() (Signer, error) {
	return &TestSigner{}, nil
}

func (s *TestSigner) DefaultAlgorithm() string {
	return ""
}

func (s *TestSigner) KeyFingerprint() string {
	return ""
}

func (s *TestSigner) Sign(dateHeader string) (string, error) {
	return "", nil
}

func (s *TestSigner) SignRaw(toSign string) (string, string, error) {
	return "", "", nil
}
