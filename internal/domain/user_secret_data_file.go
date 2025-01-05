package domain

type UserSecretDataFile struct {
	File string `json:"file"`
}

var _ UserSecretData = &UserSecretDataFile{}

func NewUserSecretFile(file string) *UserSecretDataFile {
	return &UserSecretDataFile{
		File: file,
	}
}

func newUserSecretFileFromData(data []byte) (*UserSecretDataFile, error) {
	secretData := NewUserSecretFile(string(data))

	return secretData, nil
}

func (d *UserSecretDataFile) GetType() UserSecretType {
	return UserSecretFileType
}

func (d *UserSecretDataFile) GetData() ([]byte, error) {
	return []byte(d.File), nil
}
