package handler

import (
	"net/http"
	"strconv"
	"xixi-movie/model"
	"xixi-movie/service"

	"github.com/gin-gonic/gin"
)

// 1.

type MoiveServer struct {
	yts *service.YTS
	qb  *service.QbClient
}

func SetupMoiveServer() (*MoiveServer, error) {
	yts := service.Setup()
	qb, err := service.SetupQbClient()
	if err != nil {
		return nil, err
	}
	ms := &MoiveServer{yts: yts, qb: qb}

	return ms, nil
}

// add torrent with magnet url
func (ms *MoiveServer) HandlerAddQb(c *gin.Context) {
	magnet := c.Query("magnet")
	err := ms.qb.AddByMagnet(magnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "torrent added"})
}

// login qbittorrent
func (ms *MoiveServer) HandlerLoginQb(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	host := c.Query("host")
	port := c.Query("port")
	ms.qb.AuthLogin(username, password, host, port)
}

func (ms *MoiveServer) HandlerSearch(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	quality := c.Query("quality")
	minimumRating := c.Query("minimum_rating")
	queryTerm := c.Query("query_term")
	genre := c.Query("genre")
	sortBy := c.Query("sort_by")
	orderBy := c.Query("order_by")
	withRtRatings := c.Query("with_rt_ratings")

	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)
	minimumRatingInt, _ := strconv.Atoi(minimumRating)
	withRtRatingsBool, _ := strconv.ParseBool(withRtRatings)

	req := &model.YTSListMoviesRequest{
		Limit:         limitInt,
		Page:          pageInt,
		Quality:       quality,
		MinimumRating: minimumRatingInt,
		QueryTerm:     queryTerm,
		Genre:         genre,
		SortBy:        sortBy,
		OrderBy:       orderBy,
		WithRtRatings: withRtRatingsBool,
	}
	m, err := ms.yts.ListMovies(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
