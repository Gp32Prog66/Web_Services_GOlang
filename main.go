package main

//Data.
type desktop struct {
    ID     string  `json:"id"`
    Manufacturer  string  `json:"manufacturer"`
    OperatingSystem string  `json:"OperatingSystem"`
    Price  float64 `json:"price"`
	RAM  float32 `json:"RAM"`
	GPU string 	`json"GPU"`
	Storage string `json"Storage"`
}

//Insert Data
var desktops []desktop{
	{ID: "1", Manufacturer: "Dell", OperatingSystem: "Windows 11", Price: 1,249.99, RAM: 16, GPU: "NVIDIA GeForce RTX 4060", Storage: "1 TB"},
	{ID: "2", Manufacturer: "ASUS", OperatingSystem: "Windows 10", Price: 1,449.99, RAM: 32, GPU: "NVIDIA GeForce RTX 4060 Ti", Storage: "2 TB"},
	{ID: "3", Manufacturer: "Lenovo", OperatingSystem: "Windows 10", Price: 999.99, RAM: 16, GPU: "AMD Radeon RX 7600", Storage: "1 TB"},
	{ID: "4", Manufacturer: "HP", OperatingSystem: "Windows 11", Price: 1,499.00, RAM: 32, GPU: "AMD Radeon RX 6300", Storage: "1 TB"},
}