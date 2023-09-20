package routes

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// * Upload File and store in projects uploads directory
func uploadFile(w http.ResponseWriter, r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return file, fileHeader, err
	}

	defer file.Close()

	// err = r.ParseForm()
	// if err != nil {
	// 	fmt.Println("Error Parsing Form")
	// 	return file, fileHeader, err
	// }

	// // * Get Date Input Value
	// dateValue := r.PostFormValue("date")

	// date, err := time.Parse("2006-01-02", dateValue)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return file, fileHeader, err
	// }
	// // * Format Date Input
	// formatDate := date.Format("02-01-2006")

	// fmt.Printf("Hello, %s!", formatDate)

	// * Create the uploads folder if it doesn't
	// * already exist
	err = os.MkdirAll("./configs/uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return file, fileHeader, err
	}

	// * Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return file, fileHeader, err
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return file, fileHeader, err
	}

	fmt.Printf("\n\nUpload successful\n\n")

	return file, fileHeader, err
}

// * Delete uploaded file from projects uploads directory
func deleteFile(w http.ResponseWriter, r *http.Request, file multipart.File, fileHeader *multipart.FileHeader) {
	fmt.Printf("Deleting uploaded file %s \n", fileHeader.Filename)
	err := os.Remove(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted:  %s \n", fileHeader.Filename)
}
