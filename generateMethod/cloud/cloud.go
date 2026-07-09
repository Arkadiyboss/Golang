package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(name string) *CloudDb {
	return &CloudDb{
		url: name,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return nil, nil
}

func (db *CloudDb) Write(content []byte) string {
	return ""
}