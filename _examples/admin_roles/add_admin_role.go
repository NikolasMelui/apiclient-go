package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.AddAdminRoleParams{AdminRoleName:"read_only", AllowedEntries:"GetAccountInfo;GetCallHistory"}
	res, verr, err := client.AdminRoles.AddAdminRole(params)
	fmt.Println(res, verr, err)
}