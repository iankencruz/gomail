package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./configs/web/index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// truncated for brevity

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = r.ParseForm()
	if err != nil {
		fmt.Println("Error Parsing Form")
		return
	}

	dateValue := r.PostFormValue("date")

	date, err := time.Parse("2006-01-02", dateValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatDate := date.Format("02-01-2006")

	fmt.Printf("Hello, %s!", formatDate)

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./configs/uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("\n\nUpload successful\n\n")
	fmt.Printf("\n\nDo stuff Here \n\n")

	fmt.Printf("\n\nNow delete the same file \n\n")

	// e := os.Remove(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
	// if e != nil {
	// 	log.Fatal(e)
	// }

	fmt.Printf("\n\nFile has been deleted \n\n")

	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./configs/web/success.html")

}

func closeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "exit.html")
	fmt.Println("Server Closed")
	os.Exit(1)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

}
