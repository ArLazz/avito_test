package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func (slugService) UsersHistory(ctx context.Context, year string, month string) (string, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "DB opening error", err
	}
	defer db.Close()
	rows, err := db.Query("SELECT user_id, segment_name, operation, time FROM usershistory WHERE EXTRACT(YEAR FROM time) = $1 AND EXTRACT(MONTH FROM time) = $2", year, month)
	if err != nil {
		return "Error", err
	}
	defer rows.Close()
	var userId, segmentName, operation, time string
	b := bytes.Buffer{}
	writer := csv.NewWriter(&b)
	for rows.Next() {
		err := rows.Scan(&userId, &segmentName, &operation, &time)
		if err != nil {
			return "Error", err
		}
		data := []string{userId, segmentName, operation, time}
		writer.Write(data)
	}
	writer.Flush()


	client := ServiceAccount("client_secret.json")

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return "Connection error to GDrive", err
	}

	folderId := "1_tmocipkxUoJl9rSzXhtmorpfwMADa08"
	f := &drive.File{
		MimeType: "text/plain",
		Name:     year + month + ".csv",
		Parents:  []string{folderId},
	}
	file, err := srv.Files.Create(f).Media(bytes.NewReader(b.Bytes())).Do()
	str := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", file.Id)
	if err != nil {
		return "Error creating a file in", err
	}
	return str, nil
}


func ServiceAccount(secretFile string) *http.Client {
	b, err := os.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}
