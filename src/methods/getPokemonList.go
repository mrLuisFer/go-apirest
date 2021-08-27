package methods

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getPokemonList() {
	var baseEndpoint string = "https://pokeapi.co/api/v2/pokemon?limit=100&offset=200"
	response, error := http.Get(baseEndpoint)

	if error != nil {
		fmt.Print(error)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
