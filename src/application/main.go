package application

import (
	"log"
	"github.com/natefinch/lumberjack"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main(){
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})

	var roleApp RoleAppService
	roleApp=BuildRoleAppService()

	addRoleHandler:=httptransport.NewServer(
		MakeRoleEndpoint(roleApp),
		DecodeRoleRequest,
		EncodeResponse,
	)

	http.Handle("/role",addRoleHandler)

	log.Fatal(http.ListenAndServe(":8921",nil))
}
