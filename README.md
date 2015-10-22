URL Identification
==================

urlidentification is a library to help you identify what kind of URL your have.

It is based on [RFC 3986][rfc]'s definitions for:

* An [absolute URL](http://tools.ietf.org/html/rfc3986#page-27)
* A relative reference to an [absolute path](http://tools.ietf.org/html/rfc3986#section-4.2)
* A relative path reference
* A network-path reference

## Example usage

```go
package main

import "fmt"
import "net/url"
import "github.com/statusmachine/urlidentification"

func main() {
	absoluteUrl, _ := url.Parse("https://www.statusmachine.com/subdir/page.html")
	networkPathReference, _ := url.Parse("//www.statusmachine.com/subdir/page.html")
	relativeRefWithAbsPath, _ := url.Parse("/subdir/page.html")
	relativePathRef, _ := url.Parse("subdir/page.html")

	fmt.Println("is absolute url:", urlidentification.IsAbsoluteURL(absoluteUrl))
	fmt.Println("is network path reference:", urlidentification.IsNetworkPathReferenceURL(networkPathReference))
	fmt.Println("is relative reference with absolute path:", urlidentification.IsRelativeReferenceWithAbsolutePathURL(relativeRefWithAbsPath))
	fmt.Println("is relative path reference:", urlidentification.IsRelativePathReferenceURL(relativePathRef))

	identifiedURL, _ := urlidentification.IdentifyURLString("http://google.com")
	fmt.Println(identifiedURL == urlidentification.Absolute)
}
```

## License

The [BSD 3-Clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[rfc]: http://tools.ietf.org/html/rfc3986
