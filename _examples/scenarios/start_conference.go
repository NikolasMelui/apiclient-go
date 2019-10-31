package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.StartConferenceParams{ConferenceName:"boss", RuleId:1, UserId:1, ScriptCustomData:"mystr"}
	res, verr, err := client.Scenarios.StartConference(params)
	fmt.Println(res, verr, err)
}