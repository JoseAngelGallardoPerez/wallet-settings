package auth

import (
	"log"

	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const ResourceSettings = "settings"

const ActionCreate = "create"
const ActionUpdate = "update"
const ActionRead = "read"
const ActionDelete = "delete"

const RoleRoot = "root"
const RoleAdmin = "admin"
const RoleClient = "client"

type Service struct {
	dynamicPermissions PermissionMap
	permissionService  *PermissionService
}

type Permission string

type Policy func(*userpb.User) bool
type PermissionMap map[string]map[string]map[string]Policy

func NewService(permissionService *PermissionService) *Service {
	auth := Service{permissionService: permissionService}
	auth.dynamicPermissions = PermissionMap{
		RoleAdmin: {
			ResourceSettings: {ActionUpdate: auth.ProvideCheckSpecificPermission(PermissionModifySettings)},
		},
	}

	return &auth
}

// CanDynamic checks action is allowed by calling associated function
func (auth *Service) CanDynamic(user *userpb.User, action string, resourceName string) bool {
	if user.RoleName == RoleRoot {
		return true
	}

	function := auth.getPermissionFunc(user.RoleName, action, resourceName)
	return function(user)
}

// blockFunc always block access
func blockFunc(_ *userpb.User) bool {
	return false
}

// allowFunc always allows access
func allowFunc(_ *userpb.User) bool {
	return true
}

// getPermissionFunc returns function by role, action and resourceName.
// Returns blockFunc if proposed func not found
func (auth *Service) getPermissionFunc(role string, action string, resourceName string) Policy {
	if rolePermission, ok := auth.dynamicPermissions[role]; ok {
		if resourcePermission, ok := rolePermission[resourceName]; ok {
			if actionPermission, ok := resourcePermission[action]; ok {
				return actionPermission
			}
		}
	}
	return blockFunc
}

// CheckPermission calls permission service in order to check if user granted permission
func (auth *Service) CheckPermission(perm Permission, user *userpb.User) bool {
	result, err := auth.permissionService.Check(user.UID, string(perm))
	if err != nil {
		log.Printf("permission policy failed to check permission: %s", err.Error())
		return false
	}
	return result
}

func (auth *Service) ProvideCheckSpecificPermission(perm Permission) Policy {
	return func(user *userpb.User) bool {
		return auth.CheckPermission(perm, user)
	}
}
