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
	RAM  float32 `json:"ram"`
	GPU string 	`json"gpu"`
	Storage string `json"storage"`
}

//Insert Data
var desktops []desktop {
	{ID: "1", Manufacturer: "Dell", OperatingSystem: "Windows 11", Price: 1,249.99, RAM: 16, GPU: "NVIDIA GeForce RTX 4060", Storage: "1 TB"},
	{ID: "2", Manufacturer: "ASUS", OperatingSystem: "Windows 10", Price: 1,449.99, RAM: 32, GPU: "NVIDIA GeForce RTX 4060 Ti", Storage: "2 TB"},
	{ID: "3", Manufacturer: "Lenovo", OperatingSystem: "Windows 10", Price: 999.99, RAM: 16, GPU: "AMD Radeon RX 7600", Storage: "1 TB"},
	{ID: "4", Manufacturer: "HP", OperatingSystem: "Windows 11", Price: 1,499.00, RAM: 32, GPU: "AMD Radeon RX 6300", Storage: "1 TB"},
}

//Get Desktops as JSON
func getDesktops(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, desktops)
}

func postDesktops(c *gin.Context)
{
	var newDesktop desktop
	
	//BindJSON

	if err := c.BindJSON(&newDesktop); err != nil {
		return
	}

	//Adding New Desktop
	desktops = append(desktops, newDesktop)
	c.IndentedJSON(http.StatusCreated, newDesktop)
}

func getDesktopsByOS(c *gin.Context) {
	operatingSystem := c.Param("id")

	/for _, x := range desktops {
		if x.OperatingSystem == operatingSystem {
			c.IndentedJSON(http.StatusOk, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Operating System Couldn't be Found"})
}

func main() {
	router := gin.Default()
	router.GET("/desktops", getDesktops)
	router.GET("/desktops", getDesktopsByOS)
	router.POST("/desktops", postDesktops)

	router.Run("localhost:7000")
}