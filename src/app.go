package main

import (
	"log"
	"net/http"
	. "application"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main(){


	/*
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})
*/

	log.Println("Server begin")
	var roleApp RoleAppService
	roleApp=BuildRoleAppService()

	addRoleHandler:=httptransport.NewServer(
		MakeRoleEndpoint(roleApp),
		DecodeRoleRequest,
		EncodeResponse,
		httptransport.ServerBefore(RecordRequestMethodfunc),
	)


	http.Handle("/role",addRoleHandler)

	log.Fatal(http.ListenAndServe(":8921",nil))
}
