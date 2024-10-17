package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Data.
type desktop struct {
    ID     string  `json:"id"`
    Manufacturer  string  `json:"manufacturer"`
    OperatingSystem string  `json:"operatingSystem"`
    Price  float64 `json:"price"`
	RAM  int `json:"ram"`
	GPU string 	`json:"gpu"`
	Storage string `json:"storage"`
}

//Insert Data
var desktops = []desktop{
	{ID: "1", Manufacturer: "Dell", OperatingSystem: "Windows11", Price: 1249.99, RAM: 16, GPU: "NVIDIA GeForce RTX 4060", Storage: "1 TB"},
	{ID: "2", Manufacturer: "ASUS", OperatingSystem: "Windows10", Price: 1449.99, RAM: 32, GPU: "NVIDIA GeForce RTX 4060 Ti", Storage: "2 TB"},
	{ID: "3", Manufacturer: "Lenovo", OperatingSystem: "Windows10", Price: 999.99, RAM: 16, GPU: "AMD Radeon RX 7600", Storage: "1 TB"},
	{ID: "4", Manufacturer: "HP", OperatingSystem: "Windows11", Price: 1499.00, RAM: 32, GPU: "AMD Radeon RX 6300", Storage: "1 TB"},
}



//Get Desktops as JSON
func getDesktops(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, desktops)
}

func postDesktops(c *gin.Context) {
	var newDesktop desktop
	
	//BindJSON

	if err := c.BindJSON(&newDesktop); err != nil {
		return
	}

	//Adding New Desktop
	desktops = append(desktops, newDesktop)
	c.IndentedJSON(http.StatusCreated, newDesktop)
}

//Gather Manufacturer of Each Desktop
func getDesktopsByManufacturer(c *gin.Context) {
	manufacturer := c.Param("manufacturer")

	//Display Desktops
	for _, x := range desktops {
		if x.Manufacturer == manufacturer {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Desktop Manufacturer Couldn't be Found"})
}

func main() {

	//Use gin package
	router := gin.Default()

	//GET and POST Methods for Desktops
	router.GET("/desktops", getDesktops)
	router.GET("/desktops/:manufacturer", getDesktopsByManufacturer)
	router.POST("/desktops", postDesktops)

	//Run on Port 7000
	router.Run("localhost:7000")
}