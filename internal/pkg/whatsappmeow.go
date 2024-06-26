package pkg

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	qrcodelib "github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsmeowClient interface {
	Auth(client *whatsmeow.Client, send func(context string, data any))
	GetClient(deviceStore *store.Device) *whatsmeow.Client
	GetContainerSql() *sqlstore.Container
}

type whatsmeowClient struct {
	containerSql *sqlstore.Container
	clientLog    waLog.Logger
}

func NewWhatsappmeow(database string) WhatsmeowClient {

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	var (
		container *sqlstore.Container
		err error
	)
	
	if database == "posgresql" {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("HOST"), 
			os.Getenv("PORT"), 
			os.Getenv("USERNAME"), 
			os.Getenv("PASSWORD"), 
			os.Getenv("DATABASE_NAME"))
		container, err = sqlstore.New("postgres", psqlInfo, dbLog)
	}else {
		container, err = sqlstore.New("sqlite3", "./database/wa-app.db?_foreign_keys=on", dbLog)
	}

	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)

	return &whatsmeowClient{container, clientLog}
}

func (w *whatsmeowClient) GetContainerSql() *sqlstore.Container {
	return w.containerSql
}

func (w *whatsmeowClient) Auth(client *whatsmeow.Client, send func(context string, data any)) {

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, err := client.GetQRChannel(context.Background())
		if err != nil {
			panic(err)
		}
		if err := client.Connect(); err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				img, err := qrcodelib.Encode(evt.Code, qrcodelib.Low, 200) // evt.Code
				if err != nil {
					panic(err)
				}
				qrCode := base64.StdEncoding.EncodeToString(img)
				send("qrcode", fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="qrcode"/>`, qrCode))
			} else if evt.Event == "success" {
				break
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
		if client.WaitForConnection(time.Duration(10 * time.Second)) {
			send("success", client.Store.ID)
			client.Disconnect()
		}
	} else {
		send("success", client.Store.ID)
	}
}

func (w *whatsmeowClient) GetClient(deviceStore *store.Device) *whatsmeow.Client {
	client := whatsmeow.NewClient(deviceStore, w.clientLog)
	return client
}
