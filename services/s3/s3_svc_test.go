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
