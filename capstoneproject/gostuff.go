package main

import (
    "encoding/json"
   "fmt"
    "io/ioutil"
    "log"
    "os"
)

type Card struct{
    ManaCost string `json:"mana_cost"`
    Keywords []interface{} `json:"keywords"`
    Name string `json:"name"`
    Power string `json:"power"`
    Toughness string `json:"toughness"`
    ID string `json:"id"`
    Colors []string `json:"colors"`
    ImageUris ImageUris `json:"image_uris"`
    Prices Prices `json:"prices"`

}

  type Prices struct {
	 Usd       string      `json:"usd"`

}

type ImageUris struct {
//	 Small      string `json:"small"`
	// Normal     string `json:"normal"`
	// Large      string `json:"large"`
	 Png        string `json:"png"`
//	 ArtCrop    string `json:"art_crop"`
//	 BorderCrop string `json:"border_crop"`
}


  func main() {


      filename, err := os.Open("default-cards.json")
      if err != nil {
          log.Fatal(err)
      }

      defer filename.Close()

      data, err := ioutil.ReadAll(filename)

      if err != nil {
          log.Fatal(err)
      }
      var result []Card
  //    var result2 []Prices
  //    var result3 []ImageUris
      jsonErr := json.Unmarshal(data, &result)
    //  jsonErr2 := json.Unmarshal(data, &result2)
    //  jsonErr3 := json.Unmarshal(data, &result3)
      if jsonErr != nil {
          log.Fatal(jsonErr)
      }

      data, err = json.MarshalIndent(result, "", "  ")
        if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("%+v\n", result)
}
