package session

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/hieven/go-instagram/src/config"
	"gopkg.in/redis.v5"
)

type RedisSession struct {
	prefix    string
	delimiter string
	username  string
	client    *redis.Client
}

func newRedisSession(cnf *config.Config) (*RedisSession, error) {
	redisUrl, _ := url.Parse(cnf.SessionStorage)

	redisHost := redisUrl.Host
	redisPassword, _ := redisUrl.User.Password()

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	redisSession := RedisSession{
		prefix:    "go-instagram",
		delimiter: ":",
		username:  cnf.Username,
		client:    client,
	}

	return &redisSession, nil
}

func (session *RedisSession) GetCookies() []*http.Cookie {
	val, _ := session.client.Get(session.getRedisKey()).Result()

	var cookies []*http.Cookie
	json.Unmarshal([]byte(val), &cookies)

	return cookies
}

func (session *RedisSession) SetCookies(cookies []*http.Cookie) error {
	cookiesByte, err := json.Marshal(cookies)
	if err != nil {
		return err
	}
	cookiesStr := string(cookiesByte)

	return session.client.Set(session.getRedisKey(), cookiesStr, 60*time.Minute).Err()
}

func (session *RedisSession) getRedisKey() string {
	return session.prefix + session.delimiter + "sessions" + session.delimiter + session.username
}
