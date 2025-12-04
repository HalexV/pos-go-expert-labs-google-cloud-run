package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/HalexV/pos-go-expert-labs-google-cloud-run/configs"
	"github.com/HalexV/pos-go-expert-labs-google-cloud-run/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{cep}", BuscaCepHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), r)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	cepParam := chi.URLParam(r, "cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	re := regexp.MustCompile(`^\d{8}$`)

	if !re.MatchString(cepParam) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	cep, error := BuscaCep(cepParam)
	if error != nil {
		if error.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("can not find zipcode"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	temp, error := BuscaTemp(cep.Location.Coordinates.Latitude, cep.Location.Coordinates.Longitude)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	output := dto.APIOutput{
		TempC: temp.Current.TempC,
		TempF: temp.Current.TempF,
		TempK: temp.Current.TempC + 273,
	}

	json.NewEncoder(w).Encode(output)
}

func BuscaCep(cep string) (*dto.GetBrasilAPIOutput, error) {
	var getBrasilAPIOutput dto.GetBrasilAPIOutput

	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://brasilapi.com.br/api/cep/v2/"+cep, nil)
	if err != nil {
		log.Println("Error to create request with context")
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode))
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &getBrasilAPIOutput)

	return &getBrasilAPIOutput, nil
}

func BuscaTemp(latitude, longitude string) (*dto.GetWeatherAPIOutput, error) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	var getWeatherAPIOutput dto.GetWeatherAPIOutput

	req, err := http.NewRequestWithContext(context.Background(), "GET", fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s,%s&key=%s", latitude, longitude, configs.WeatherApiKey), nil)
	if err != nil {
		log.Println("Error to create request with context")
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, nil
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &getWeatherAPIOutput)

	return &getWeatherAPIOutput, nil
}
