package bigcommerce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Post is a BC blog post
type Post struct {
	ID                   int64       `json:"id"`
	Title                string      `json:"title"`
	URL                  string      `json:"url"`
	PreviewURL           string      `json:"preview_url"`
	Body                 string      `json:"body"`
	Tags                 []string    `json:"tags"`
	Summary              string      `json:"summary"`
	IsPublished          bool        `json:"is_published"`
	PublishedDate        interface{} `json:"publisheddate"`
	PublishedDateISO8601 string      `json:"publisheddate_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author"`
	ThumbnailPath        string      `json:"thumbnail_path"`
}

type CreatePostPayload struct {
	Title           string   `json:"title" validate:"required"`
	URL             string   `json:"url,omitempty"`
	Body            string   `json:"body" validate:"required"`
	Tags            []string `json:"tags,omitempty"`
	IsPublished     bool     `json:"is_published,omitempty"`
	MetaDescription string   `json:"meta_description,omitempty"`
	MetaKeywords    string   `json:"meta_keywords,omitempty"`
	Author          string   `json:"author,omitempty"`
	ThumbnailPath   string   `json:"thumbnail_path,omitempty"`
	PublishedDate   string   `json:"published_date,omitempty"`
}

// GetAllPosts downloads all posts from BigCommerce, handling pagination
func (bc *Client) GetAllPosts() ([]Post, error) {
	cs := []Post{}
	var csp []Post
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetPosts(page)
		if err != nil {
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		cs = append(cs, csp...)
		page++
	}
	return cs, err
}

// GetPosts downloads all posts from BigCommerce, handling pagination
// page: the page number to download
func (bc *Client) GetPosts(page int) ([]Post, bool, error) {
	url := "/v2/blog/posts?limit=250&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, false, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, false, err
	}

	var pp []Post
	err = json.Unmarshal(body, &pp)
	if err != nil {
		log.Printf("Error unmarshalling posts: %s %s", err, string(body))
		return nil, false, err
	}
	return pp, len(pp) == 250, nil
}

// CreatePost creates a new blog post in BigCommerce and returns the post or error
func (bc *Client) CreatePost(payload *CreatePostPayload) (*Post, error) {
	var b []byte
	b, _ = json.Marshal([]CreatePostPayload{*payload})
	req := bc.getAPIRequest(http.MethodPost, "/v2/blog/posts", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusUnprocessableEntity {
			var errResp ErrorResult
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				log.Printf("Error: %s\nResult: %s", err, string(body))
				return nil, err
			}
			if len(errResp.Errors) > 0 {
				errors := []string{}
				for _, e := range errResp.Errors {
					errors = append(errors, e)
				}
				return nil, fmt.Errorf("%s", strings.Join(errors, ", "))
			}
			return nil, errors.New("unknown error")
		}
		log.Printf("Error: %s\nResult: %s", err, string(body))
		return nil, err
	}

	var ret Post
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
