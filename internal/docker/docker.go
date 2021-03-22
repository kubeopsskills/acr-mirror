package docker

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ryanuber/go-glob"
)

type TagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

const defaultSleepDuration time.Duration = 60 * time.Second

func getSleepTime(rateLimitReset string, now time.Time) time.Duration {
	rateLimitResetInt, err := strconv.ParseInt(rateLimitReset, 10, 64)

	if err != nil {
		return defaultSleepDuration
	}

	sleepTime := time.Unix(rateLimitResetInt, 0)
	calculatedSleepTime := sleepTime.Sub(now)

	if calculatedSleepTime < (0 * time.Second) {
		return 0 * time.Second
	}

	return calculatedSleepTime
}

func FilterTags(tags []string, matchTags []string) []string {
	res := make([]string, 0)

	for _, tag := range tags {
		// match tags, with glob
		if len(matchTags) > 0 {
			for _, matchTag := range matchTags {
				if !glob.Glob(matchTag, tag) {
					// m.log.Debugf("Dropping tag '%s', it doesn't match glob pattern '%s'", remoteTag.Name, tag)
					continue
				}
				res = append(res, tag)
			}
		}
	}

	return res
}

// get the remote tags from the remote (v2) compatible registry.
func GetRemoteTags(registryName string, repositoryName string, username string, password string) ([]string, error) {
	url := fmt.Sprintf("https://%s.azurecr.io/v2/%s/tags/list", registryName, repositoryName)

	var allTags []string
	var (
		err     error
		res     *http.Response
		req     *http.Request
		retries int = 5
	)

	for retries > 0 {
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(username+":"+password))))

		httpClient := &http.Client{Timeout: 10 * time.Second}
		res, err = httpClient.Do(req)

		if err != nil {
			// log.Warningf("Failed to get %s, retrying", url)
			retries--
		} else if res.StatusCode == 429 {
			sleepTime := getSleepTime(res.Header.Get("X-RateLimit-Reset"), time.Now())
			// m.log.Infof("Rate limited on %s, sleeping for %s", url, sleepTime)
			time.Sleep(sleepTime)
			retries--
		} else if res.StatusCode < 200 || res.StatusCode >= 300 {
			// m.log.Warningf("Get %s failed with %d, retrying", url, res.StatusCode)
			retries--
		} else {
			break
		}

	}

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var tagResponse TagsResponse
	if err := json.NewDecoder(res.Body).Decode(&tagResponse); err != nil {
		return nil, err
	}

	allTags = append(allTags, tagResponse.Tags...)
	return allTags, nil
}
