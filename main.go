package main

import (
    "fmt"
    "io"
    "os"
    "net/http"
    // // "context"
    "main/Util"
    "github.com/go-rod/rod"
    // "github.com/go-rod/rod/lib/launcher"
    // // "path/filepath"
    "github.com/joho/godotenv"
    // "github.com/JustinBeckwith/go-yelp"
	"github.com/gorilla/mux"
)

type Station struct {
    Name string
    Date string
    Review string
}

func ScrapeStationNames() {
    browser := *rod.New().MustConnect()
    page := *browser.MustPage("https://www.itsmarta.com/train-stations-and-schedules.aspx").MustWaitStable(). MustWaitLoad()
    stationElementsDiv := page.MustElements("stations__items isotope")
    fmt.Println("stationElementsDiv: ", stationElementsDiv)
    stationElements := stationElementsDiv.MustElements("stations__item route-gold route-red")
    fmt.Println("stationElements: ", stationElements)
    

}

const staticDir string = "/static/"

func main(){
    godotenv.Load()
    fmt.Println("\n---------------------------------------------------------\n")
    // port := os.Getenv("PORT")

    r := mux.NewRouter()
    // router := router.NewRouter(service.NewService(database.NewDatabase(database.GetDB())))

    staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("static/")))
    r.PathPrefix(staticDir).Handler(staticHandler)

    r.HandleFunc("/getstationrowelement", ScrapeStationNames
}
