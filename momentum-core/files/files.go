package files

import (
	"encoding/base64"
	"momentum-core/utils"
)

func fileToBase64(path string) (string, error) {

	f, err := utils.FileOpen(path, utils.FILE_ALLOW_READ_WRITE_ALL)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileAsString := utils.FileAsString(f)

	return base64.RawStdEncoding.EncodeToString([]byte(fileAsString)), nil
}

func FileToRaw(base64Encoded string) (string, error) {

	bytes, err := base64.RawStdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
