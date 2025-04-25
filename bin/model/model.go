package model

type YTSListMoviesResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		MovieCount int         `json:"movie_count"`
		Limit      int         `json:"limit"`
		PageNumber int         `json:"page_number"`
		Movies     []YTSMovies `json:"movies"`
	} `json:"data"`
}

// Parameter	Required	Type	Default	Description
// limit		Integer between 1 - 50 (inclusive)	20	The limit of results per page that has been set
// page		Integer (Unsigned)	1	Used to see the next page of movies, eg limit=15 and page=2 will show you movies 15-30
// quality		String (480p, 720p, 1080p, 1080p.x265, 2160p, 3D)	All	Used to filter by a given quality
// minimum_rating		Integer between 0 - 9 (inclusive)	0	Used to filter movie by a given minimum IMDb rating
// query_term		String	0	Used for movie search, matching on: Movie Title/IMDb Code, Actor Name/IMDb Code, Director Name/IMDb Code
// genre		String	All	Used to filter by a given genre (See http://www.imdb.com/genre/ for full list)
// sort_by		String (title, year, rating, peers, seeds, download_count, like_count, date_added)	date_added	Sorts the results by choosen value
// order_by		String (desc, asc)	desc	Orders the results by either Ascending or Descending order
// with_rt_ratings		Boolean	false	Returns the list with the Rotten Tomatoes rating included
type YTSListMoviesRequest struct {
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	Quality       string `json:"quality"`
	MinimumRating int    `json:"minimum_rating"`
	QueryTerm     string `json:"query_term"`
	Genre         string `json:"genre"`
	SortBy        string `json:"sort_by"`
	OrderBy       string `json:"order_by"`
	WithRtRatings bool   `json:"with_rt_ratings"`
}

type YTSMovies struct {
	Id                      int           `json:"id"`
	TitleLong               string        `json:"title_long"`
	Title                   string        `json:"title"`
	Year                    int           `json:"year"`
	Rating                  float64       `json:"rating"`
	Runtime                 int           `json:"runtime"`
	Genres                  []string      `json:"genres"`
	Summary                 string        `json:"summary"`
	DescriptionFull         string        `json:"description_full"`
	Synopsis                string        `json:"synopsis"`
	YtTrailerCode           string        `json:"yt_trailer_code"`
	Language                string        `json:"language"`
	MpaRating               string        `json:"mpa_rating"`
	BackgroundImage         string        `json:"background_image"`
	BackgroundImageOriginal string        `json:"background_image_original"`
	SmallCoverImage         string        `json:"small_cover_image"`
	MediumCoverImage        string        `json:"medium_cover_image"`
	LargeCoverImage         string        `json:"large_cover_image"`
	State                   string        `json:"state"`
	Torrents                []YTSTorrents `json:"torrents"`
	DateUploaded            string        `json:"date_uploaded"`
	DateUploadedUnix        int           `json:"date_uploaded_unix"`
}
type YTSTorrents struct {
	Url              string `json:"url"`
	Hash             string `json:"hash"`
	Quality          string `json:"quality"`
	Type             string `json:"type"`
	IsRepack         string `json:"is_repack"`
	VideoCodec       string `json:"video_codec"`
	BitDepth         string `json:"bit_depth"`
	AudioChannels    string `json:"audio_channels"`
	Seeds            int    `json:"seeds"`
	Peers            int    `json:"peers"`
	Size             string `json:"size"`
	SizeBytes        int    `json:"size_bytes"`
	DateUploaded     string `json:"date_uploaded"`
	DateUploadedUnix int    `json:"date_uploaded_unix"`
}

// {
//   "status": "ok",
//   "status_message": "Query was successful",
//   "data": {
//     "movie_count": 1,
//     "limit": 20,
//     "page_number": 1,
//     "movies": [
//       {
//         "id": 36,
//         "url": "https://yts.mx/movies/28-weeks-later-2007",
//         "imdb_code": "tt0463854",
//         "title": "28 Weeks Later",
//         "title_english": "28 Weeks Later",
//         "title_long": "28 Weeks Later (2007)",
//         "slug": "28-weeks-later-2007",
//         "year": 2007,
//         "rating": 6.9,
//         "runtime": 100,
//         "genres": [
//           "Action",
//           "Adventure",
//           "Drama",
//           "Horror",
//           "Sci-Fi"
//         ],
//         "summary": "Almost six months after London was decimated by the unstoppable Rage Virus in 28 Days Later (2002), the U.S. Army has restored peace and repopulated the quarantined city. However, the deadly epidemic reawakens when an unsuspecting carrier of the highly transmittable pathogen enters the dead capital with the first wave of returning refugees. This time, the horrible virus is more dangerous than ever. Has the next nightmare begun?—Nick Riganas",
//         "description_full": "Almost six months after London was decimated by the unstoppable Rage Virus in 28 Days Later (2002), the U.S. Army has restored peace and repopulated the quarantined city. However, the deadly epidemic reawakens when an unsuspecting carrier of the highly transmittable pathogen enters the dead capital with the first wave of returning refugees. This time, the horrible virus is more dangerous than ever. Has the next nightmare begun?—Nick Riganas",
//         "synopsis": "Almost six months after London was decimated by the unstoppable Rage Virus in 28 Days Later (2002), the U.S. Army has restored peace and repopulated the quarantined city. However, the deadly epidemic reawakens when an unsuspecting carrier of the highly transmittable pathogen enters the dead capital with the first wave of returning refugees. This time, the horrible virus is more dangerous than ever. Has the next nightmare begun?—Nick Riganas",
//         "yt_trailer_code": "cH-9OTWwjxM",
//         "language": "en",
//         "mpa_rating": "",
//         "background_image": "https://yts.mx/assets/images/movies/28_Weeks_Later_2007/background.jpg",
//         "background_image_original": "https://yts.mx/assets/images/movies/28_Weeks_Later_2007/background.jpg",
//         "small_cover_image": "https://yts.mx/assets/images/movies/28_Weeks_Later_2007/small-cover.jpg",
//         "medium_cover_image": "https://yts.mx/assets/images/movies/28_Weeks_Later_2007/medium-cover.jpg",
//         "large_cover_image": "https://yts.mx/assets/images/movies/28_Weeks_Later_2007/large-cover.jpg",
//         "state": "ok",
//         "torrents": [
//           {
//             "url": "https://yts.mx/torrent/download/84A74935AA5CE794D7159460E66BC58127181F2C",
//             "hash": "84A74935AA5CE794D7159460E66BC58127181F2C",
//             "quality": "720p",
//             "type": "bluray",
//             "is_repack": "0",
//             "video_codec": "x264",
//             "bit_depth": "8",
//             "audio_channels": "2.0",
//             "seeds": 56,
//             "peers": 9,
//             "size": "697.66 MB",
//             "size_bytes": 731549532,
//             "date_uploaded": "2015-10-31 20:51:29",
//             "date_uploaded_unix": 1446321089
//           },
//           {
//             "url": "https://yts.mx/torrent/download/44D0F59C95D555D0EF5BFBB4EBC13426E74B5A1A",
//             "hash": "44D0F59C95D555D0EF5BFBB4EBC13426E74B5A1A",
//             "quality": "1080p",
//             "type": "bluray",
//             "is_repack": "0",
//             "video_codec": "x264",
//             "bit_depth": "8",
//             "audio_channels": "5.1",
//             "seeds": 100,
//             "peers": 27,
//             "size": "1.85 GB",
//             "size_bytes": 1986422374,
//             "date_uploaded": "2022-01-28 01:11:22",
//             "date_uploaded_unix": 1643328682
//           }
//         ],
//         "date_uploaded": "2015-10-31 20:51:29",
//         "date_uploaded_unix": 1446321089
//       }
//     ]
//   },
//   "@meta": {
//     "server_time": 1745566024,
//     "server_timezone": "CET",
//     "api_version": 2,
//     "execution_time": "0 ms"
//   }
// }
