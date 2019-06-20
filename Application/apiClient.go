// Package Application accesses the official Riot API.
//
// Construct a client with the New() function, and call the various client
// methods to retrieve data from the API.
package Application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/AdrianOrtuno/appSumamoos/Domain/region"
	"github.com/AdrianOrtuno/appSumamoos/Persistance/external"
	"github.com/AdrianOrtuno/appSumamoos/Persistance/ratelimit"
)

type client struct {
	key string
	c   external.Doer
	r   ratelimit.Limiter
}

type Client interface {
	// ----- Summoner API -----

	// GetBySummonerName returns a summoner by summoner name.
	GetBySummonerName(ctx context.Context, r region.Region, name string) (*Summoner, error)

	// ------ Match API -----
	// GetRecentMatchlist returns the last 20 matches played on the given account ID.
	GetMatchlist(ctx context.Context, r region.Region, accountID string, opts *GetMatchlistOptions) (*Matchlist, error)

	// GetMatch returns a match by match ID.
	GetMatch(ctx context.Context, r region.Region, matchID int64) (*Match, error)
}

// New returns a Client configured for the given API client and underlying HTTP
// client. The returned Client is threadsafe.
func New(key string, httpClient external.Doer, limiter ratelimit.Limiter) Client {
	return &client{
		key: key,
		c:   httpClient,
		r:   limiter,
	}
}

// dispatchAndUnmarshalWithUniquifier is the same as dispatchAndUnmarshal,
// except with an additional uniquifier parameter that allows special case
// handling of certain methods that have different quota buckets depending on
// the relative path.
func (c *client) dispatchAndUnmarshalWithUniquifier(ctx context.Context, r region.Region, m string, relativePath string, v url.Values, u string, dest interface{}) (*http.Response, error) {
	res, err := c.dispatchMethod(ctx, r, m, relativePath, v, u)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		err, ok := httpErrors[res.StatusCode]
		if !ok {
			err = ErrBadHTTPStatus
		}
		return res, err
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewReader(b))

	// The body is in good state, so now we can return if there was an IO problem.
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(b, dest)

	return res, err
}

// dispatchAndUnmarshal dispatches the method (see dispatchMethod). If the
// method returns HTTP okay, then read the body into a buffer and attempt to
// unmarshal it into the supplied destination. Otherwise, the method returns
// one of the documented errors. In any case, the body is set to read from the
// beginning of the stream and is left open, as if the response were returned
// directly from an HTTP request.
func (c *client) dispatchAndUnmarshal(ctx context.Context, r region.Region, m string, relativePath string, v url.Values, dest interface{}) (*http.Response, error) {
	return c.dispatchAndUnmarshalWithUniquifier(ctx, r, m, relativePath, v, "", dest)
}

// dispatchMethod calls the given API method for the given region. The
// relativePath is appended to the method to form the REST endpoint. The given
// URL values are encoded and passed as URL parameters following the REST
// endpoint.
func (c *client) dispatchMethod(ctx context.Context, r region.Region, m string, relativePath string, v url.Values, uniquifier string) (*http.Response, error) {
	var suffix, separator string

	if len(v) > 0 {
		suffix = fmt.Sprintf("?%s", v.Encode())
	}
	if !strings.HasPrefix(relativePath, "/") {
		separator = "/"
	}
	path := r.Host() + m + separator + relativePath + suffix
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("X-Riot-Token", c.key)

	done, _, err := c.r.Acquire(ctx, ratelimit.Invocation{
		ApplicationKey: c.key,
		Region:         strings.ToUpper(string(r)),
		Method:         strings.ToLower(m),
		Uniquifier:     uniquifier,
	})

	if err != nil {
		return nil, err
	}

	// If either the done() or the HTTP request is an error, then return error.
	res, err := c.c.Do(req)
	derr := done(res)
	if err == nil {
		err = derr
	}
	return res, err
}
