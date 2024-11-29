# Image Background Removal API

This is a Go-based API that removes the background from images using a threshold value. The service supports PNG images, processes the images, and returns them with the background removed.

## Features

- **Upload PNG images**: Upload `.png` images to the server via a POST request.
- **Background Removal**: The background of the uploaded image will be removed based on the luminance threshold.
- **Save Processed Image**: The processed image will be saved on the server and available for download.

## Prerequisites

Before running this project, make sure you have the following installed:

- **Go (Golang)**: [Download Go](https://golang.org/dl/)
- **Go Modules**: Ensure your Go environment is set up to use modules (`go mod`).

## Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/alrescha79-cmd/image-background-removal.git
    ```

2. move to `image-background-removal`

    ```bash
    cd image-background-removal
    ```

3. Install dependencies:

    The project uses Go's `image` package for image decoding and processing, and it doesn't have external dependencies other than Go’s standard library.

    If needed, you can run:

    ```bash
    go mod tidy
    ```

## Running the Server

1. Navigate to the project directory.

2. Run the server with the following command:

    ```bash
    go run ./cmd/main.go
    ```

3. The server will start and listen on port `8080` (default). You can change the port in the code if needed.

4. The server will be available at: `http://localhost:8080/`

## API Endpoint

### `POST /remove-background`

This endpoint accepts a PNG image, processes it to remove the background, and returns the processed image.

#### Request Body

- **Method**: `POST`
- **Content-Type**: `multipart/form-data`
- **Key**: `image` (the name of the file input)
- **File Type**: PNG image

#### Example Request (Postman or curl)

**Postman**:

- Set the request method to `POST`.
- Set the URL to `http://localhost:8080/remove-background`.
- In the body, select `form-data` and upload a PNG file with the key `image`.

**curl**:

```bash
curl -X POST -F "image=@path/to/your/image.png" http://localhost:8080/remove-background
```

#### Response

- **Success**: If the image is processed successfully, you will receive a response with the file path to the processed image.

    Example:

    ```json
    {
        "message": "Image processed successfully",
        "output": "output_folder/your-image-no-bg.png"
    }
    ```

- **Error**: If there’s an issue with the image (e.g., invalid format or file error), an error response with an appropriate HTTP status code will be returned.

    Example:

    ```json
    {
        "error": "Unsupported or invalid image format"
    }
    ```

## Image Processing Details

- **Threshold-based Background Removal**: The service processes the image by converting it to grayscale and adjusting the transparency of each pixel based on its luminance.
- **Supported Formats**: Currently, only PNG images are supported.

## File Uploads

The images uploaded to the server will be saved in the `input_images` folder. After processing, the result will be saved in the specified `output_folder`.

## Folder Structure

The server expects the following folder structure:

```root
    background-remover/
    ├── cmd/
    │   └── main.go
    ├── pkg/
    │   └── utils.go
    ├── input_images/
    ├── output_images/
    ├── go.mod

```

Make sure the `input_images` and `output_images` folders exist or are created during the file processing.

## Error Handling

If the uploaded file is empty, in an unsupported format, or the image cannot be processed, the API will return an appropriate error message.

## Troubleshooting

- **Issue**: The image is not processed correctly, or the file is not accepted.
- **Solution**: Ensure that the uploaded file is a valid PNG image. You can check the file size and format in the server logs.

## Example

- **Input**
    ![input](input_images/golang.png)
- **Output**
    ![output](output_images/golang-no-bg.png)

## Author

- **Anggun Caksono** - [alrescha79-cmd](https://github.com/alrescha79-cmd)
