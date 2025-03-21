package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type ApiResponse struct {
	Aciklama          string
	Alis              string
	Satis             string
	GuncellenmeZamani string
}
type ApiResponsePokemon struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var wg sync.WaitGroup
var chanelAltin chan []ApiResponse
var chanelPokemon chan ApiResponsePokemon

func init() {

	chanelAltin = make(chan []ApiResponse)
	chanelPokemon = make(chan ApiResponsePokemon)
}
func AltinKaynak() {
	response, err := http.Get("https://rest.altinkaynak.com/Currency.json")
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200 {
		data, _ := io.ReadAll(response.Body)
		var apiResponse []ApiResponse
		json.Unmarshal(data, &apiResponse)
		chanelAltin <- apiResponse

	} else {
		fmt.Println("Http call failed with status: ", response.Status)
		chanelAltin <- nil

	}
}
func Pokemon() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0")
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200 {
		data, _ := io.ReadAll(response.Body)
		var apiResponsePokemon ApiResponsePokemon
		json.Unmarshal(data, &apiResponsePokemon)
		chanelPokemon <- apiResponsePokemon
	} else {
		fmt.Println("Http call failed with status: ", response.Status)
		chanelPokemon <- ApiResponsePokemon{}
	}
}

func main() {
	ctxAltin, cancelAltin := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelAltin()
	ctxPokemon, cancelPokemon := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelPokemon()
	defer cancelAltin()
	go AltinKaynak()
	go Pokemon()
	select {
	case AltinBilgi := <-chanelAltin:
		for a := 0; a < 20; a++ {
			fmt.Println(AltinBilgi[a].Aciklama, "                                     Güncelleme Zamanı: ", AltinBilgi[a].GuncellenmeZamani)
			fmt.Println("Alış Fiyatı: ", AltinBilgi[a].Alis)
			fmt.Println("Satış Fiyatı: ", AltinBilgi[a].Satis)
			fmt.Println("------------------------------------------------------------------------------------------------")
		}
	case <-ctxAltin.Done():
		fmt.Println("TimeOut AltınKaynak")
	}
	select {
	case PokemonBilgi := <-chanelPokemon:
		for i := 0; i < 20; i++ {
			fmt.Println(PokemonBilgi.Results[i].Name)
			fmt.Println(PokemonBilgi.Results[i].Url)
			fmt.Println("-------------------------------------------------------------------------------------------------")
		}
	case <-ctxPokemon.Done():
		fmt.Println("TimeOut Pokemn")
	}

}
