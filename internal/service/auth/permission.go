package auth

import (
	"github.com/Confialink/wallet-settings/internal/srvdiscovery"
	"context"
	"net/http"

	"github.com/Confialink/wallet-permissions/rpc/permissions"
	"github.com/Confialink/wallet-pkg-discovery/v2"
)

const (
	PermissionModifySettings = Permission("modify_settings")
)

type PermissionService struct {
	sdResolver discovery.Resolver
}

func NewPermissionService(resolver discovery.Resolver) *PermissionService {
	return &PermissionService{
		sdResolver: resolver,
	}
}

//Check checks if specified user is granted permission to perform some action
func (p *PermissionService) Check(userId, actionKey string) (bool, error) {
	request := &permissions.PermissionReq{UserId: userId, ActionKey: actionKey}

	checker, err := p.checker()
	if nil != err {
		return false, err
	}

	response, err := checker.Check(context.Background(), request)
	if nil != err {
		return false, err
	}
	return response.IsAllowed, nil
}

func (p *PermissionService) checker() (permissions.PermissionChecker, error) {
	url, err := p.sdResolver.Resolve(srvdiscovery.PortNameRpc, srvdiscovery.ServiceNamePermissions)
	if nil != err {
		return nil, err
	}
	checker := permissions.NewPermissionCheckerProtobufClient(url.String(), http.DefaultClient)
	return checker, nil
}
