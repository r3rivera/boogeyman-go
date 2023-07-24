package s3

import "testing"

const BUCKET_NAME = "amdg-r2app"

func Test_S3PresignedUrl(t *testing.T) {

	url, err := GeneratePresignedUrl("File001.txt", BUCKET_NAME)
	if err != nil {
		t.Fail()
	}
	t.Logf("URL :: \n\n%s \n\n", url)
}

func Test_S3PutObject(t *testing.T) {

	s := NewS3FileInfo("files/", "test001.txt", BUCKET_NAME)
	err := s.UploadFile()
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
}
