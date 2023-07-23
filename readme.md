

gdriveuploader

gdriveuploader is a GoLang package that provides a simple utility to upload files to Google Drive and obtain public links that can be used as HTML links to access the uploaded files.

## Requirements
    Google Drive API Credentials:

        Before using this package, you need to set up a project on the Google Developers Console.
        Create credentials for your project and download the credentials JSON file.
        Google Drive API Token:

        To access the Google Drive API on behalf of a user, you need an OAuth2 token.
        The token will be generated after the user grants access to the Google Drive API.
        Google Drive Folder ID:

        You need to provide the ID of the Google Drive folder where you want to upload the files.
        The folder ID can be obtained from the URL when you navigate to the folder on Google Drive.
## Installation
To use the gdriveuploader package, you need to install it into your Go workspace:

``` bash
    go get github.com/your-username/gdriveuploader
```
## Usage
    Import the package in your Go code:
    ```go

        import "github.com/your-username/gdriveuploader"
    ```
    Initialize the GoogleDriveUploader:
    ```go
    uploader, err := gdriveuploader.NewUploader("path/to/credentials.json", "path/to/token.json")
    if err != nil {
        // Handle error
    }
    ```
    Upload a file to Google Drive:
    ```go
    filePath := "path/to/upload_file.txt"
    folderID := "your_google_drive_folder_id"

    link, err := uploader.UploadFile(filePath, folderID)
    if err != nil {
        // Handle error
    }

    fmt.Println("File uploaded successfully. Public link:", link)
    ```
    Replace "path/to/credentials.json" and "path/to/token.json" with the actual paths to your Google Drive API credentials and token files, respectively.

    Replace "path/to/upload_file.txt" with the path of the file you want to upload to Google Drive.

    Replace "your_google_drive_folder_id" with the ID of the Google Drive folder where you want to upload the file.

## License
    This package is licensed under the MIT License. See the LICENSE file for more details.