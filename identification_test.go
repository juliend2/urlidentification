package urlidentification

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var absoluteURL, _ = url.Parse("https://www.statusmachine.com")
var networkPathReferenceURL, _ = url.Parse("//statusmachine.com/subdir/page.html")
var relativeRefWithAbsPathURL, _ = url.Parse("/subdir/page.html")
var relativePathRefURL, _ = url.Parse("subdir/page.html")

func TestAbsolute(t *testing.T) {
	assert.True(t, IsAbsolute(absoluteURL), "URL should be Absolute")
	assert.False(t, IsAbsolute(networkPathReferenceURL), "URL should NOT be absolute")
	assert.False(t, IsAbsolute(relativeRefWithAbsPathURL), "URL should NOT be absolute")
	assert.False(t, IsAbsolute(relativePathRefURL), "URL should NOT be absolute")
}

func TestNetworkPathReference(t *testing.T) {
	assert.True(t, IsNetworkPathReferenceURL(networkPathReferenceURL), "URL should look like //statusmachine.com/somepath")
	assert.False(t, IsNetworkPathReferenceURL(absoluteURL), "URL doesn't look like //statusmachine.com/somepath")
	assert.False(t, IsNetworkPathReferenceURL(relativeRefWithAbsPathURL), "URL doesn't look like //statusmachine.com/somepath")
	assert.False(t, IsNetworkPathReferenceURL(relativePathRefURL), "URL doesn't look like //statusmachine.com/somepath")
}

func TestRelativeReferenceWithAbsolutePath(t *testing.T) {
	assert.True(t, IsRelativeReferenceWithAbsolutePathURL(relativeRefWithAbsPathURL), "URL looks like /somepath")
	assert.False(t, IsRelativeReferenceWithAbsolutePathURL(absoluteURL), "URL doesn't look like /somepath")
	assert.False(t, IsRelativeReferenceWithAbsolutePathURL(networkPathReferenceURL), "URL doesn't look like /somepath")
	assert.False(t, IsRelativeReferenceWithAbsolutePathURL(relativePathRefURL), "URL doesn't look like /somepath")
}

func TestRelativePathReference(t *testing.T) {
	assert.True(t, IsRelativePathReferenceURL(relativePathRefURL), "URL looks like subdir/page.html")
	assert.False(t, IsRelativePathReferenceURL(absoluteURL), "URL DOESN'T looks like subdir/page.html")
	assert.False(t, IsRelativePathReferenceURL(networkPathReferenceURL), "URL DOESN'T looks like subdir/page.html")
	assert.False(t, IsRelativePathReferenceURL(relativeRefWithAbsPathURL), "URL DOESN'T looks like subdir/page.html")
}

func TestIdentify(t *testing.T) {
	identifiedURL, err := Identify(absoluteURL)
	assert.Equal(t, identifiedURL, Absolute, "Should be Absolute")
	assert.Nil(t, err, "Error should be nil")

	identifiedURL2, err := Identify(networkPathReferenceURL)
	assert.Equal(t, identifiedURL2, NetworkPathReference, "Should be NetworkPathReference")
	assert.Nil(t, err, "Error should be nil")

	identifiedURL3, err := Identify(relativeRefWithAbsPathURL)
	assert.Equal(t, identifiedURL3, RelativeReferenceWithAbsolutePath, "Should be RelativeReferenceWithAbsolutePath")
	assert.Nil(t, err, "Error should be nil")

	identifiedURL4, err := Identify(relativePathRefURL)
	assert.Equal(t, identifiedURL4, RelativePathReference, "Should be RelativePathReference")
	assert.Nil(t, err, "Error should be nil")
}

func TestIdentifyURLString(t *testing.T) {
	identifiedURL, err := IdentifyURLString("https://www.statusmachine.com")
	assert.Equal(t, identifiedURL, Absolute, "Should be absolute")
	assert.Nil(t, err, "Error should be nil")

	unparseableURL, err := IdentifyURLString("%")
	assert.Equal(t, unparseableURL, Unidentifiable, "Should NOT be identifiable")
	assert.NotNil(t, err, "Error should NOT be nil")
}
