package src

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

type FirebaseServer struct {
	CredentialPath string
	Ctx            context.Context
	client         *firestore.Client
}

func (f *FirebaseServer) Init() {
	sa := option.WithCredentialsFile(f.CredentialPath)
	var app *firebase.App
	app, err = firebase.NewApp(f.Ctx, nil, sa)
	checkErr()

	f.client, err = app.Firestore(f.Ctx)
	checkErr()
	//defer f.client.Close()
}

func (f *FirebaseServer) UpdateData(data interface{}) {
	//var tmpData []byte
	//tmpData, err = json.Marshal(data)
	//log.Println(string(tmpData))

	var res *firestore.WriteResult
	res, err = f.client.Collection("youtube-dl").Doc("server").Set(f.Ctx, map[string]interface{}{"hello": 3333})
	checkErr()
	log.Println(res)
}
