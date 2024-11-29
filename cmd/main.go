package main

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"os"
	"path/filepath"

	utilities "background-remover/pkg"
)

const (
	imageFolder = "./input_images"
	outFolder   = "./output_images"
	threshold   = uint8(128)
)

type Response struct {
	Message string `json:"message"`
	Output  string `json:"output,omitempty"`
}

func handleRemoveBackground(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if header.Size == 0 {
		http.Error(w, "Invalid or empty file", http.StatusBadRequest)
		return
	}

	inputPath := filepath.Join(imageFolder, header.Filename)
	outFile, err := os.Create(inputPath)
	if err != nil {
		http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Failed to write file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		http.Error(w, "Failed to stat file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("File size: %d bytes\n", fileInfo.Size())

	baseImage, err := os.Open(inputPath)
	if err != nil {
		http.Error(w, "Failed to open file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer baseImage.Close()

	baseImage.Seek(0, 0)

	_, format, err := image.Decode(baseImage)
	if err != nil {
		http.Error(w, "Unsupported or invalid image format: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Image format: %s\n", format)

	baseImage.Seek(0, 0)

	noBackgroundImage := utilities.Transform(baseImage, threshold, nil)
	err = utilities.SaveImageToFile(header.Filename, noBackgroundImage, outFolder)
	if err != nil {
		http.Error(w, "Failed to save processed image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "Image processed successfully",
		Output:  filepath.Join(outFolder, header.Filename),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	os.MkdirAll(imageFolder, os.ModePerm)
	os.MkdirAll(outFolder, os.ModePerm)

	http.HandleFunc("/remove-background", handleRemoveBackground)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
