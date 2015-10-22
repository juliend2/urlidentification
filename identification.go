package urlidentification

import (
	"errors"
	// "fmt"
	"net/url"
	"strings"
)

const (
	Unidentifiable                    = iota
	Absolute                          = iota
	NetworkPathReference              = iota
	RelativeReferenceWithAbsolutePath = iota
	RelativePathReference             = iota
)

// I based my nomenclature on this SO answer: http://stackoverflow.com/a/17407082/242404

func hasNoScheme(url *url.URL) bool {
	return url.Scheme == ""
}

func hasHost(url *url.URL) bool {
	return url.Host != ""
}

func hasNoHost(url *url.URL) bool {
	return !hasHost(url)
}

func startsWithDoubleSlash(url *url.URL) bool {
	return strings.HasPrefix(url.String(), "//")
}

func startsWithSlash(url *url.URL) bool {
	return strings.HasPrefix(url.String(), "/")
}

func pathStartsWithSingleSlash(url *url.URL) bool {
	return startsWithSlash(url) && !startsWithDoubleSlash(url)
}

// Exported functions:

func IsAbsolute(url *url.URL) bool {
	return url.IsAbs()
}

func IsNetworkPathReferenceURL(url *url.URL) bool {
	return hasNoScheme(url) && hasHost(url) && startsWithDoubleSlash(url)
}

// Returns whether the URL looks like /subdir/page.html
func IsRelativeReferenceWithAbsolutePathURL(url *url.URL) bool {
	return hasNoHost(url) && pathStartsWithSingleSlash(url)
}

func IsRelativePathReferenceURL(url *url.URL) bool {
	return hasNoScheme(url) && hasNoHost(url) && !strings.HasPrefix(url.Path, "/")
}

func Identify(url *url.URL) (int, error) {
	if IsAbsolute(url) {
		return Absolute, nil
	} else if IsNetworkPathReferenceURL(url) {
		return NetworkPathReference, nil
	} else if IsRelativeReferenceWithAbsolutePathURL(url) {
		return RelativeReferenceWithAbsolutePath, nil
	} else if IsRelativePathReferenceURL(url) {
		return RelativePathReference, nil
	} else {
		return Unidentifiable, errors.New("urlidentification: URL pattern is Unidentifiable")
	}
}

func IdentifyURLString(urlString string) (int, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return Unidentifiable, err
	}
	return Identify(parsedURL)
}
