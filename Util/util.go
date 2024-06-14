package Util

import (
    "fmt"
    "os"
    "strings"
)

func GetHostname() string {
    if os.Getenv("ENV_TYPE") == "dev" {
        return os.Getenv("DEV_HOSTNAME")
    } else {
        if os.Getenv("DEPLOYED") == "true" {
            return  os.Getenv("PROD_HOSTNAME")
        } else {
            // return strings.ReplaceAll(os.Getenv("DEV_HOSTNAME"),"http://","")
            return os.Getenv("DEV_HOSTNAME")
        }
    }
}

func ConvertSrcLink(srcLink string) (string, error) {
    newLink := strings.ReplaceAll(srcLink, "https://i.4cdn.org", "http://is2.4chan.org")
    fmt.Println("newLink: ", newLink)
    return newLink, nil
}
