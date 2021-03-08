package main

import (
	//"strings"
	"github.com/tamama9527/notion_bot/util"
    "os"
    "os/exec"
    "fmt"

)

// You more than likely want your "Bot User OAuth Access Token" which starts with "xoxb-"
var client = util.NotionClient(os.Getenv("NOTION_AUTH_token"))

func main() {
    util.NotionExport(client, "03970d328a2f4278859915d0a88ae830")
    util.NotionPage(client, "03970d328a2f4278859915d0a88ae830")
    output, err := exec.Command("./gitpush.sh","test3").CombinedOutput()
    if err != nil {
        fmt.Println("Error when running command.  Output:")
        fmt.Println(string(output))
        fmt.Printf("Got command status: %s\n", err.Error())
        return
    }

}
