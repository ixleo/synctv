package cache

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/synctv-org/synctv/internal/db"
	"github.com/synctv-org/synctv/internal/model"
	"github.com/synctv-org/synctv/internal/vendor"
	"github.com/synctv-org/synctv/utils"
	"github.com/synctv-org/vendors/api/emby"
	"github.com/zijiren233/gencontainer/refreshcache"
)

type EmbyUserCache = MapCache[*EmbyUserCacheData, struct{}]

type EmbyUserCacheData struct {
	Host     string
	ServerID string
	ApiKey   string
	Backend  string
}

func NewEmbyUserCache(userID string) *EmbyUserCache {
	return newMapCache(func(ctx context.Context, key string, args ...struct{}) (*EmbyUserCacheData, error) {
		return EmbyAuthorizationCacheWithUserIDInitFunc(userID, key)
	}, 0)
}

func EmbyAuthorizationCacheWithUserIDInitFunc(userID, serverID string) (*EmbyUserCacheData, error) {
	if serverID == "" {
		return nil, errors.New("serverID is required")
	}
	v, err := db.GetEmbyVendor(userID, serverID)
	if err != nil {
		return nil, err
	}
	if v.ApiKey == "" || v.Host == "" {
		return nil, db.ErrNotFound("vendor")
	}
	return &EmbyUserCacheData{
		Host:     v.Host,
		ServerID: v.ServerID,
		ApiKey:   v.ApiKey,
		Backend:  v.Backend,
	}, nil
}

type EmbySource struct {
	URLs []struct {
		URL  string
		Name string
	}
	Subtitles []struct {
		URL   string
		Type  string
		Name  string
		Cache *refreshcache.RefreshCache[[]byte, struct{}]
	}
}

type EmbyMovieCacheData struct {
	Sources []EmbySource
}

type EmbyMovieCache = refreshcache.RefreshCache[*EmbyMovieCacheData, *EmbyUserCache]

func NewEmbyMovieCache(movie *model.Movie) *EmbyMovieCache {
	return refreshcache.NewRefreshCache(NewEmbyMovieCacheInitFunc(movie), 0)
}

func NewEmbyMovieCacheInitFunc(movie *model.Movie) func(ctx context.Context, args ...*EmbyUserCache) (*EmbyMovieCacheData, error) {
	return func(ctx context.Context, args ...*EmbyUserCache) (*EmbyMovieCacheData, error) {
		if len(args) == 0 {
			return nil, errors.New("need emby user cache")
		}

		var (
			serverID string
			err      error
		)
		serverID, movie.Base.VendorInfo.Emby.Path, err = model.GetEmbyServerIdFromPath(movie.Base.VendorInfo.Emby.Path)
		if err != nil {
			return nil, err
		}

		aucd, err := args[0].LoadOrStore(ctx, serverID)
		if err != nil {
			return nil, err
		}
		if aucd.Host == "" || aucd.ApiKey == "" {
			return nil, errors.New("not bind emby vendor")
		}
		u, err := url.Parse(aucd.Host)
		if err != nil {
			return nil, err
		}
		cli := vendor.LoadEmbyClient(aucd.Backend)
		data, err := cli.GetItem(ctx, &emby.GetItemReq{
			Host:   aucd.Host,
			Token:  aucd.ApiKey,
			ItemId: movie.Base.VendorInfo.Emby.Path,
		})
		if err != nil {
			return nil, err
		}
		if data.IsFolder {
			return nil, errors.New("path is dir")
		}
		var resp EmbyMovieCacheData = EmbyMovieCacheData{
			Sources: make([]EmbySource, len(data.MediaSourceInfo)),
		}
		for i, v := range data.MediaSourceInfo {
			if v.Container == "" {
				continue
			}
			result, err := url.JoinPath("emby", "Videos", data.Id, fmt.Sprintf("stream.%s", v.Container))
			if err != nil {
				return nil, err
			}
			u.Path = result
			query := url.Values{}
			query.Set("api_key", aucd.ApiKey)
			query.Set("Static", "true")
			query.Set("MediaSourceId", v.Id)
			u.RawQuery = query.Encode()
			resp.Sources[i].URLs = append(resp.Sources[i].URLs, struct {
				URL  string
				Name string
			}{
				URL:  u.String(),
				Name: v.Name,
			})
			for _, msi := range v.MediaStreamInfo {
				switch msi.Type {
				case "Subtitle":
					subtutleType := "srt"
					result, err = url.JoinPath("emby", "Videos", data.Id, v.Id, "Subtitles", fmt.Sprintf("%d", msi.Index), fmt.Sprintf("Stream.%s", subtutleType))
					if err != nil {
						return nil, err
					}
					u.Path = result
					u.RawQuery = ""
					url := u.String()
					name := msi.DisplayTitle
					if name == "" {
						if msi.Title != "" {
							name = msi.Title
						} else {
							name = msi.DisplayLanguage
						}
					}
					resp.Sources[i].Subtitles = append(resp.Sources[i].Subtitles, struct {
						URL   string
						Type  string
						Name  string
						Cache *refreshcache.RefreshCache[[]byte, struct{}]
					}{
						URL:   url,
						Type:  subtutleType,
						Name:  name,
						Cache: refreshcache.NewRefreshCache(newEmbySubtitleCacheInitFunc(url), 0),
					})
				}
			}
		}
		return &resp, nil
	}
}

func newEmbySubtitleCacheInitFunc(url string) func(ctx context.Context, args ...struct{}) ([]byte, error) {
	return func(ctx context.Context, args ...struct{}) ([]byte, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", utils.UA)
		req.Header.Set("Referer", req.URL.Host)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("bad status code")
		}
		return io.ReadAll(resp.Body)
	}
}
