package api

import (
	"github.com/teris-io/shortid"
	"net/url"
)

const baseDomain = "http://t-link.com"

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkHolder struct {
	Links map[string]string
}

func (link LinkHolder) Shorten(urlStr string) string {
	shortPath, _ := shortid.Generate()
	parsedUrl, _ := url.Parse(baseDomain)
	parsedUrl.Path = shortPath

	_, ok := link.Links[parsedUrl.String()]

	if ok {
		return Shorten(link, urlStr)
	}

	link.Links[parsedUrl.String()] = urlStr
	return parsedUrl.String()
}

func (link LinkHolder) Resolve(url string) string {

	val, ok := link.Links[url]

	if !ok {
		return ""
	}

	return val
}

func Shorten(i Shortener, url string) string {
	return i.Shorten(url)
}

func Resolve(i Shortener, url string) string {

	return i.Resolve(url)
}
