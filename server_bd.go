package whois

import (
	"net/url"
	"strings"
)

var bd = &Server{
	Resolve: func(req *Request) error {
		labels := strings.SplitN(req.Query, ".", 2)
		values := url.Values{}
		values.Set("dom", labels[0])
		values.Set("ext", labels[1])
		// "http://www.whois.com.bd/?dom=google&ext=com.bd"
		req.URL = "http://www.whois.com.bd/?" + values.Encode()
		req.Body = ""
		return nil
	},
}

func init() {
	register(
		bd,
		"www.whois.com.bd",
	)
}
