package main

import (
    "fmt"
    "net/http"
    // "io"
    "github.com/go-rod/rod"
    // "github.com/go-rod/rod/lib/launcher"
    "github.com/joho/godotenv"
    // "github.com/JustinBeckwith/go-yelp"
	"github.com/gorilla/mux"
    // "encoding/json"
)

type Station struct {
    Name string `json:"name"`
    Date string `json:"date"`
    Review string `json:review"`
}

type ScrapeStationResponse struct {
    Name string `json:"name"`
    Date string `json:"date"`
    Review string `json:review"`
}


// func scrapeStationNames (w http.ResponseWriter, r *http.Request) {
func scrapeStationNames () ([]Station, error) {
    stations := []Station{}
    browser := rod.New().MustConnect()
    page := *browser.MustPage("https://www.itsmarta.com/train-stations-and-schedules.aspx").MustWaitLoad()
    stationElementsDiv := page.MustElement("stations__items isotope")
    fmt.Println("stationElementsDiv: ", stationElementsDiv)
    stationElements := stationElementsDiv.MustElements("stations__item")
    fmt.Println("stationElements: ", stationElements)

    for _,stationElement := range stationElements {
        fmt.Println("stationElement: ", stationElement)
        station := Station{}
        fmt.Println("station: ", station)
        stationName := stationElement.MustElement("stations__item-name").MustText()
        station.Name = stationName
        stations = append(stations, station)
    } 
    return stations, nil
}

func scrapeStationNames2 () ([]Station, error) {
    stations := []Station{}
    browser := *rod.New().MustConnect()
    page := browser.MustPage("https://www.itsmarta.com/train-stations-and-schedules.aspx").MustWaitLoad()
    stationElements := page.MustElements("a[class='stations__item-name'")

    for _,stationElement := range stationElements {
        fmt.Println("stationElement: ", stationElement)
        station := Station{}
        fmt.Println("station: ", station)
        stationName := stationElement.MustText()
        station.Name = stationName
        stations = append(stations, station)
    } 
    return stations, nil
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

    // r.HandleFunc("/getstationrowelement", scrapeStationNames)
    //
    // res,err := http.Get("/getstationrowelement")
    // if err != nil {panic(err)}
    //
    // data,err := io.ReadAll(res.Body)
    // if err != nil {panic(err)}
    //
    //
    //
    // testResponse := ScrapeStationResponse{}
    // err = json.Unmarshal(data, &testResponse)
    // fmt.Println("testResponse: ", testResponse)

    stationNames,err := scrapeStationNames2()
        if err != nil {
            panic(fmt.Errorf("Error scaping staion names in main.go: %s", err))
        }
        for i,station := range stationNames {
            fmt.Println(i, ": " , station)
        }
}

