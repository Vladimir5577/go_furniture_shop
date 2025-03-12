package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadImageFromHTTPRequest(r *http.Request) (string, error) {
	var fileName string
	r.ParseMultipartForm(10 << 20) //10 MB
	file, fileHeader, err := r.FormFile("image")

	if file == nil {
		return fileName, nil
	}

	if err != nil {
		return fileName, err
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		return fileName, err
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return fileName, fmt.Errorf("the provided file format is not allowed. please upload a JPEG or PNG image")
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fileName, err
	}

	// Create the uploads folder if it doesn't
	// already exist
	// err = os.MkdirAll("./uploads", os.ModePerm)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	// Create a new file in the uploads directory
	// after saving record to database
	dst, err := os.Create("./uploads/" + fileName)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	fmt.Println("Upload successful, filename = ", fileName)

	return fileName, nil
}
