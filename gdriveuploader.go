package gdriveuploader

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/api/drive/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleDriveUploader struct {
	service *drive.Service
}

func NewUploader(credPath, tokenPath string) (*GoogleDriveUploader, error) {
	ctx := context.Background()

	creds, err := google.ReadCredentialsFile(ctx, credPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials file: %v", err)
	}

	config, err := creds.Config(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get OAuth2 config: %v", err)
	}

	client, err := getClient(ctx, config, tokenPath)
	if err != nil {
		return nil, fmt.Errorf("unable to get client: %v", err)
	}

	srv, err := drive.New(client)
	if err != nil {
		return nil, fmt.Errorf("unable to create Drive service: %v", err)
	}

	return &GoogleDriveUploader{
		service: srv,
	}, nil
}

func (u *GoogleDriveUploader) UploadFile(filePath, folderID string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	mimeType := mime.TypeByExtension(strings.ToLower(filepath.Ext(filePath)))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	f := &drive.File{
		Name:    filepath.Base(filePath),
		Parents: []string{folderID},
	}

	ctx := context.Background()

	uploadedFile, err := u.service.Files.Create(f).Media(file, googleapi.ContentType(mimeType)).Do()
	if err != nil {
		return "", fmt.Errorf("error uploading file: %v", err)
	}

	return uploadedFile.WebViewLink, nil
}

func getClient(ctx context.Context, config *oauth2.Config, tokenPath string) (*http.Client, error) {
	tokenFile, err := os.Open(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open token file: %v", err)
	}
	defer tokenFile.Close()

	token, err := config.TokenSource(ctx, &oauth2.Token{}).Token()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token: %v", err)
	}

	return config.Client(ctx, token), nil
}
