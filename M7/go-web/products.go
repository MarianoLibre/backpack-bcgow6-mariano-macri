package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
    Id          int         //`json:"id"`
    Name        string      //`json:"name"`
    Colour      string      //`json:"colour"`
    Price       float64     //`json:"price"`
    Stock       int         //`json:"stock"`
    Code        string      //`json:"code"`
    Published   bool        //`json:"published"`
    CreatedAt   string      //`json:"created-at"`
}

type All struct {
    Products    []Product   `json:"products"`
}


func GetAll(c *gin.Context) {
    var data []byte
    var err error

    if data, err = ioutil.ReadFile("./products.json"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops! Something went wrong..."})
    }

    var payload All
    err = json.Unmarshal(data, &payload)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    }

    c.JSON(http.StatusOK, payload)
}


func Filter(c *gin.Context) {
    var data []byte
    var err error

    if data, err = ioutil.ReadFile("./products.json"); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Oops! Something went wrong..."})
    }

    var products All
    err = json.Unmarshal(data, &products)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    }
   
    var product Product
    if c.Bind(&product) != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
    }
    
    var filteredById []Product
    for _, p := range products.Products {
        if product.Id == 0 || product.Id == p.Id {
            filteredById = append(filteredById, p)
        }
    }

    var filteredByName []Product
    for _, p := range filteredById {
        if product.Name == "" || product.Name == p.Name {
            filteredByName = append(filteredByName, p)
        }
    }

    var filteredByPrice []Product
    for _, p := range filteredByName {
        if product.Price == 0.0 || product.Price == p.Price {
            filteredByPrice = append(filteredByPrice, p)
        }
    }

    var filteredByStock []Product
    for _, p := range filteredByPrice {
        if product.Stock == 0 || product.Stock == p.Stock {
            filteredByStock = append(filteredByStock, p)
        }
    }

    var filteredByColour []Product
    for _, p := range filteredByStock {
        if product.Colour == "" || product.Colour == p.Colour {
            filteredByColour = append(filteredByColour, p)
        }
    }

    c.JSON(http.StatusOK, filteredByColour)

}


func FindById(c *gin.Context) {
    var data []byte
    var err error

    if data, err = ioutil.ReadFile("./products.json"); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Oops! Something went wrong..."})
        return
    }

    var products All
    err = json.Unmarshal(data, &products)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }
   
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    for _, p := range products.Products {
        if int(id) == p.Id {
            c.JSON(http.StatusOK, p)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})

}


func ShowQueries(c *gin.Context) {
    var data []byte
    var err error

    if data, err = ioutil.ReadFile("./products.json"); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Oops! Something went wrong..."})
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    }

    var payload All
    err = json.Unmarshal(data, &payload)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    }
   
    var p Product
    if c.Bind(&p) == nil {
        c.String(http.StatusOK, "Query:\n\tId: %d\n\tName: '%s'\n\tColour: '%s'\n\tPrice: %.2f\n\tStock: %d\n\tCode: '%s'\n\tPublished: %v\n", p.Id, p.Name, p.Colour, p.Price, p.Stock, p.Code, p.Published)
    }
}

func FillJson() {
    products := All{
        Products: []Product{
          {
              Id: 1,
              Name: "Guitar", 
              Colour: "red", 
              Price: 500.00,
              Stock: 13,
              Code: "G1123",
              Published: true,
              CreatedAt: "05/10/2022",
          },
          {
              Id: 2,
              Name: "Bass", 
              Colour: "black", 
              Price: 450.00,
              Stock: 10,
              Code: "B1124",
              Published: true,
              CreatedAt: "05/10/2022",
          },
          {
              Id: 3,
              Name: "Keyboard", 
              Colour: "silver", 
              Price: 700.00,
              Stock: 8,
              Code: "K1125",
              Published: true,
              CreatedAt: "05/10/2022",
          },
          {
              Id: 4,
              Name: "Synth", 
              Colour: "blue", 
              Price: 1500.00,
              Stock: 17,
              Code: "S1126",
              Published: false,
              CreatedAt: "05/10/2022",
          },
          {
              Id: 5,
              Name: "Guitar", 
              Colour: "black", 
              Price: 3000.00,
              Stock: 3,
              Code: "G1127",
              Published: true,
              CreatedAt: "05/10/2022",
          },
      }, 
    }
    data, err := json.MarshalIndent(&products, "", "  ")
    if err != nil {
        log.Fatal("Oops! Can't convert array to JSON...")
    }
    ioutil.WriteFile("./products.json", data, 0644)
}

