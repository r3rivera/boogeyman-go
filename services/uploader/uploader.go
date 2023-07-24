package uploader

type Uploader interface {
	UploadFile(filePath, fileName, bucket string) error
}
