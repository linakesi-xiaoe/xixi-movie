package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"xixi-movie/model"
)

type YTS struct {
	apiurl string
	c      *http.Client
}

func Setup() *YTS {
	yts := &YTS{apiurl: "https://yts.mx/api/v2/", c: http.DefaultClient}
	return yts
}
func (y *YTS) route(path string) (string, error) {
	return url.JoinPath(y.apiurl, path)
}

func (y *YTS) ListMovies(req *model.YTSListMoviesRequest) ([]model.YTSMovies, error) {
	// exmaple: https://yts.mx/api/v2/list_movies.json?quality=3D
	query := url.Values{}
	query.Add("quality", req.Quality)
	query.Add("limit", strconv.Itoa(req.Limit))
	query.Add("page", strconv.Itoa(req.Page))
	query.Add("minimum_rating", strconv.Itoa(req.MinimumRating))
	query.Add("query_term", req.QueryTerm)
	query.Add("genre", req.Genre)
	query.Add("sort_by", req.SortBy)
	query.Add("order_by", req.OrderBy)
	query.Add("with_rt_ratings", strconv.FormatBool(req.WithRtRatings))

	url, err := y.route("list_movies.json")
	if err != nil {
		return nil, err
	}
	log.Println("url:", url)

	resp, err := y.c.Get(url + "?" + query.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mr model.YTSListMoviesResponse
	err = json.NewDecoder(resp.Body).Decode(&mr)
	if err != nil {
		return nil, err
	}
	return mr.Data.Movies, nil
}
