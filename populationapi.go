package main

import (

    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type APIAIRequest struct {
    ID        string    `json:"id"`
    Timestamp time.Time `json:"timestamp"`
    Result    struct {
        Parameters map[string]string `json:"parameters"`
        Contexts   []interface{}     `json:"contexts"`
        Metadata   struct {
            IntentID                  string `json:"intentId"`
            WebhookUsed               string `json:"webhookUsed"`
            WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
            IntentName                string `json:"intentName"`
        } `json:"metadata"`
        Score float32 `json:"score"`
    } `json:"result"`
    Status struct {
        Code      int    `json:"code"`
        ErrorType string `json:"errorType"`
    } `json:"status"`
    SessionID       string      `json:"sessionId"`
    OriginalRequest interface{} `json:"originalRequest"`
}

type APIAIMessage struct {
    Speech      string `json:"speech"`
    DisplayText string `json:"displayText"`
    Source      string `json:"source"`
}

type YearCountryAgeResponse struct {
    Females int
    Males   int
    Country string
    Year    int
    Total   int
    Age     int
}

type PopulationAPIResp []YearCountryAgeResponse

func APIAIPopulationEndpoint(w http.ResponseWriter, req *http.Request) {

    if req.Method == "POST" {
        decoder := json.NewDecoder(req.Body)

        var t APIAIRequest
        err := decoder.Decode(&t)
        if err != nil {
            fmt.Println(err)
            http.Error(w, "Error in decoding the Request data", http.StatusInternalServerError)
        }

        log.Println(t.Result.Parameters["geo-country"], t.Result.Parameters["year"], t.Result.Parameters["age"])
        country := t.Result.Parameters["geo-country"]
        year := t.Result.Parameters["year"]
        age := t.Result.Parameters["age"]

        apiResponse, err := http.Get("http://api.population.io/1.0/population/" + year + "/" + country + "/" + age + "/?format=json")
        if err != nil {
            fmt.Println(err)
            http.Error(w, "Error in decoding the Request data", http.StatusInternalServerError)
        }
        defer apiResponse.Body.Close()

        body, _ := ioutil.ReadAll(apiResponse.Body)
        var populationAPIResponse PopulationAPIResp
        err = json.Unmarshal(body, &populationAPIResponse)
        if err != nil {
            fmt.Println(err)
            http.Error(w, "Error in decoding the Request data", http.StatusInternalServerError)
        }

        speechResponse := fmt.Sprintf("In %s, the Population Statistics shows a total of %d males and %d females in %s in the age group of %s years", year, populationAPIResponse[0].Males, populationAPIResponse[0].Females, country, age)
        textResponse := fmt.Sprintf("Total males: %d, Total females: %d", populationAPIResponse[0].Males, populationAPIResponse[0].Females)
        msg := APIAIMessage{Source: "Population.io API", Speech: speechResponse, DisplayText: textResponse}

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(msg)
    } else {
        http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/apiai", APIAIPopulationEndpoint)
    log.Fatal(http.ListenAndServe(":9000", mux))
}