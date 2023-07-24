package filestore

type Uploader interface {
	UploadFile() error
}
