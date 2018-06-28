package application

import (
	"context"
	"net/http"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"domain/role"
	"log"
)

type roleRequest struct{
	RoleName string
	Desc string
	Permissions []*permissionRequest
}

type permissionRequest struct{
	Name string
	Operation string
	Path string
	Desc string
	RoleId int
}

type roleResponse struct{
	Success bool
	ErrMsg string
}

func DecodeRoleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request role.Role

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeRoleEndpoint(svc RoleAppService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//todo use different biz function by request method , like restful
		log.Println(ctx.Value("rm"))

		req := request.(role.Role)
		svc.AddRole(&req)

		return roleResponse{Success:true},nil
	}
}

func RecordRequestMethodfunc(ctx context.Context, r *http.Request) context.Context{
	ctx=context.WithValue(ctx,"rm",r.Method)

	return ctx
}







