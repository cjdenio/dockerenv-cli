package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Raw image data returned from the API
type rawImageData struct {
	Data struct {
		Image struct {
			ImageData
		}
	}
	Errors []struct {
		Message string
	}
}

// Image outline
type ImageData struct {
	URL       string
	Variables []struct {
		Name        string
		Description string
		Default     string
		Required    bool
		Uncommon    bool
	}
}

// Get data for an image from the graphql API
func Image(imageName string) (ImageData, error) {
	// Formulating query
	query, err := json.Marshal(
		map[string]string{
			"query": fmt.Sprintf(`
			query {
				image(name: %q) {
					url
					variables {
						name
						description
						default
						required
						uncommon
					}
				}
			}
		`, imageName),
		},
	)
	if err != nil {
		return ImageData{}, err
	}

	// Make the query
	resp, err := http.Post(
		"https://dockerenv.calebden.io/graphql",
		"application/json",
		bytes.NewBuffer(query),
	)
	if err != nil {
		return ImageData{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ImageData{}, errors.New(
			fmt.Sprintf(
				"Failed to get data for %v. Returned status code of %v",
				imageName,
				resp.Status,
			),
		)
	}

	// Parsing the response
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return ImageData{}, err
	}
	var data rawImageData
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		return ImageData{}, err
	}
	if len(data.Errors) != 0 {
		return ImageData{}, errors.New(
			fmt.Sprintf(
				"Failed to get data for %v as the following error occurred:\n%v",
				imageName,
				data.Errors[0].Message,
			),
		)
	}

	return data.Data.Image.ImageData, nil
}
