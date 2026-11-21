package main

import "fmt"

func main() {
	fmt.Println("Slices will be covered in this module.")
	var userNames = []string{"Shubham"}
	fmt.Printf("Type of userNames is %T\n", userNames)

	userNames = append(userNames, "Akshay", "Hitesh")
	fmt.Println("Usernames: ", userNames)
	userNames = append(userNames[1:]) // append firstIndex
	fmt.Println("RemainingUsers:", userNames)

	authTokens := make([]string, 4)
	authTokens[0] = "Bearer-dflsajdfldsurwrn23irlefjsdfasdf"
	authTokens[1] = "Bearer-xxxxffdsldsurwrn23irlefjsdfasdf"
	authTokens[2] = "Bearer-dflsajdfldsfdsfdfsdulefjsdfasdf"
	authTokens[3] = "Bearer-faabaadaadsurwrn23irlefjsdfasdf"
	// authTokens[4] = "Bearer-canNotDoThis"
	authTokens = append(authTokens, "Bearer-flsadfjsigigsgdsigu") // reallocates the memory

	fmt.Println(authTokens)
}
