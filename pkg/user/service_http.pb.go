// Code generated by protoc-gen-go-http. DO NOT EDIT.

package user

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.user.UserService/QueryAccount",
				FunctionName:     "QueryAccount",
				Path:             "/users",
				Method:           "GET",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/DescribeAccount",
				FunctionName:     "DescribeAccount",
				Path:             "/users/:account",
				Method:           "GET",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/CreateAccount",
				FunctionName:     "CreateAccount",
				Path:             "/users",
				Method:           "POST",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/BlockAccount",
				FunctionName:     "BlockAccount",
				Path:             "/users/:account/block",
				Method:           "POST",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/DeleteAccount",
				FunctionName:     "DeleteAccount",
				Path:             "/users/:account",
				Method:           "DELETE",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/UpdateAccountProfile",
				FunctionName:     "UpdateAccountProfile",
				Path:             "/users/:account",
				Method:           "PUT",
				Resource:         "user",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/UpdateAccountPassword",
				FunctionName:     "UpdateAccountPassword",
				Path:             "/users/:account/password",
				Method:           "PUT",
				Resource:         "password",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/GeneratePassword",
				FunctionName:     "GeneratePassword",
				Path:             "/passwords",
				Method:           "POST",
				Resource:         "password",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "*"},
			},
		},
	}
	return set
}
