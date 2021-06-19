package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

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
		GraphQLQuery{
			Query: `
			query($image: String!) {
				image(name: $image) {
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
		`,
			Variables: map[string]interface{}{
				"image": imageName,
			},
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
		return ImageData{}, fmt.Errorf(
			"Failed to get data for %v. Returned status code of %v",
			imageName,
			resp.Status,
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
		return ImageData{}, fmt.Errorf(
			"Failed to get data for %v as the following error occurred:\n%v",
			imageName,
			data.Errors[0].Message,
		)

	}

	return data.Data.Image.ImageData, nil
}
