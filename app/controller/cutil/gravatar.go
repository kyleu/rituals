package cutil

import (
	"crypto/md5" //nolint:gosec
	"fmt"
	"strconv"
	"strings"
)

type Gravatar struct {
	hash         string
	defaultURL   string
	defaultValue string
	size         int
	forceDefault bool
	rating       string
}

func NewGravatar(email string, size int) *Gravatar {
	hash := md5.Sum([]byte(email)) //nolint:gosec
	return &Gravatar{hash: fmt.Sprintf("%x", hash), size: size, forceDefault: true, rating: "pg"}
}

func (g *Gravatar) URL() string {
	return "https://www.gravatar.com/" + g.hash
}

func (g *Gravatar) AvatarURL() string {
	url := "https://www.gravatar.com/avatar/" + g.hash
	if g.forceDefault {
		url = g.addParameter(url, "f", "y")
	}
	if g.defaultURL != "" {
		url = g.addParameter(url, "d", g.defaultURL)
	} else if g.defaultValue != "" {
		url = g.addParameter(url, "d", g.defaultValue)
	}
	if g.rating != "" {
		url = g.addParameter(url, "r", g.rating)
	}
	if g.size > 0 {
		url = g.addParameter(url, "s", strconv.Itoa(g.size))
	}
	return url
}

func (g *Gravatar) addParameter(url string, key string, value string) string {
	if strings.HasSuffix(url, g.hash) || strings.HasSuffix(url, ".json") {
		url += "?"
	} else {
		url += "&"
	}

	return url + key + "=" + value
}
