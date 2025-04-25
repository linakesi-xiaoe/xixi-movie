package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// cookiejar 保存http.client请求的

type QbClient struct {
	hc *http.Client

	qbs *QbServer
}
type QbServer struct {
	host     string
	port     string
	username string
	password string
}

func SetupQbClient() (*QbClient, error) {
	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	hc := &http.Client{
		Jar: cj,
	}
	return &QbClient{hc: hc, qbs: &QbServer{}}, nil
}

func (q *QbClient) AuthLogin(username string, password string, host string, port string) error {
	// curl -i -c cookies.txt -d 'username=admin&password=adminadmin' http://localhost:8080/api/v2/auth/login
	u := fmt.Sprintf("http://%s:%s/api/v2/auth/login", host, port)
	query := url.Values{}
	query.Add("username", username)
	query.Add("password", password)
	resp, err := http.Get(u + "?" + query.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("login failed")
	}
	q.qbs.host = host
	q.qbs.port = port
	q.qbs.username = username
	q.qbs.password = password

	// save response cookie

	//save cookie to q.hc.jar
	url, err := url.Parse(u)
	if err != nil {
		return err
	}
	cookies := q.hc.Jar.Cookies(url)
	for _, cookie := range cookies {
		log.Println(cookie.Name, cookie.Value)
	}
	return nil
}

func (q *QbClient) AddByMagnet(magnet string) error {
	// 	curl 'https://vuetorrent.blink.heiyu.space/api/v2/torrents/add' \
	// --data-raw $'------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="urls"\r\n\r\nmagnet:?xt=urn:btih:35E38EB6225EA407A6DA679A59FEC79ECCC8B941&dn=Life+Is+Beautiful+%281997%29+%5B1080p%5D+%5BYTS.MX%5D&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fopen.tracker.cl%3A1337%2Fannounce&tr=udp%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce&tr=https%3A%2F%2Fopentracker.i2p.rocks%3A443%2Fannounce\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="paused"\r\n\r\nfalse\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="skip_checking"\r\n\r\nfalse\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="autoTMM"\r\n\r\ntrue\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="sequentialDownload"\r\n\r\ntrue\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="firstLastPiecePrio"\r\n\r\ntrue\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="contentLayout"\r\n\r\nOriginal\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="stopCondition"\r\n\r\nNone\r\n------WebKitFormBoundarypun6f6nZyscf1jNp\r\nContent-Disposition: form-data; name="tags"\r\n\r\n\r\n------WebKitFormBoundarypun6f6nZyscf1jNp--\r\n'

	data := url.Values{}
	data.Add("urls", magnet)
	data.Add("paused", "false")
	data.Add("skip_checking", "false")
	data.Add("autoTMM", "true")
	data.Add("sequentialDownload", "true")
	data.Add("firstLastPiecePrio", "true")
	data.Add("contentLayout", "Original")
	// post form with q.hc.cookie jar
	resp, err := q.hc.PostForm(fmt.Sprintf("http://%s:%s/api/v2/torrents/add", q.qbs.host, q.qbs.port), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyStr := string(body)
	log.Println(bodyStr)
	if resp.StatusCode != 200 {
		return fmt.Errorf("add torrent failed")
	}
	return nil
}
