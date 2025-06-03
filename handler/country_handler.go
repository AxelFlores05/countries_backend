package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Mapeo para API externa
type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital   []string `json:"capital"`
	Region    string   `json:"region"`
	Population int     `json:"population"`
}

// Estructura de respuesta para frontend
type CountryResponse struct {
	Name      string `json:"name"`
	Capital   string `json:"capital"`
	Region    string `json:"region"`
	Population int   `json:"population"`
}

// Consulta de API externa
func GetCountries(c echo.Context) error {
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo consultar la API externa"})
	}
	defer resp.Body.Close()

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al decodificar respuesta de la API"})
	}

	var result []CountryResponse
	for _, country := range countries {
		capital := ""
		if len(country.Capital) > 0 {
			capital = country.Capital[0]
		}
		result = append(result, CountryResponse{
			Name:      country.Name.Common,
			Capital:   capital,
			Region:    country.Region,
			Population: country.Population,
		})
	}

	return c.JSON(http.StatusOK, result)
}
