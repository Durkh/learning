package Domain

import "strings"

type RolePermissions struct {
	RolePermissions map[string][]string
}

func (r RolePermissions) IsAuthorizedFor(role string, routeName string) bool {

	perms := r.RolePermissions[role]

	for _, i := range perms {
		if i == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}
