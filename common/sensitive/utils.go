package sensitive

import (
	"errors"
	"github.com/opentreehole/backend/common"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"mvdan.cc/xurls/v2"
	"net/url"
	"regexp"
	"strings"
)

var imageRegex = regexp.MustCompile(
	`!\[(.*?)]\(([^" )]*?)\s*(".*?")?\)`,
)

var (
	ErrUrlParsing        = errors.New("error parsing url")
	ErrInvalidImageHost  = errors.New("不允许使用外部图片链接")
	ErrImageLinkTextOnly = errors.New("image link only contains text")
)

// findImagesInMarkdown 从Markdown文本中查找所有图片链接，检查图片链接是否合法，并且返回清除链接之后的文本
func findImagesInMarkdownContent(content string) (imageUrls []string, clearContent string, err error) {
	err = nil
	clearContent = imageRegex.ReplaceAllStringFunc(content, func(s string) string {
		if err != nil {
			return ""
		}
		submatch := imageRegex.FindStringSubmatch(s)
		altText := submatch[1]

		var imageUrl string
		if len(submatch) > 2 && submatch[2] != "" {
			imageUrl = submatch[2]
			innerErr := checkValidUrl(imageUrl)
			if innerErr != nil {
				if errors.Is(innerErr, ErrInvalidImageHost) {
					err = innerErr
					return ""
				}
				// if the url is not valid, treat as text only
			} else {
				// append only valid image url
				imageUrls = append(imageUrls, imageUrl)
				imageUrl = ""
			}
		}

		var title string
		if len(submatch) > 3 && submatch[3] != "" {
			title = strings.Trim(submatch[3], "\"")
		}

		var ret strings.Builder
		if altText != "" {
			ret.WriteString(altText)
		}
		if imageUrl != "" {
			if ret.String() != "" {
				ret.WriteString(" ")
			}
			ret.WriteString(imageUrl)
		}
		if title != "" {
			if ret.String() != "" {
				ret.WriteString(" ")
			}
			ret.WriteString(title)
		}
		return ret.String()
	})
	return
}

func checkType(params ParamsForCheck) bool {
	return slices.Contains(checkTypes, params.TypeName)
}

func containsUnsafeURL(content string) (bool, string) {
	xurlsRelaxed := xurls.Relaxed()
	matchedURLs := xurlsRelaxed.FindAllString(content, -1)
	if len(matchedURLs) == 0 {
		return false, ""
	}

	for _, matchedURL := range matchedURLs {
		if !strings.Contains(matchedURL, "://") {
			matchedURL = "http://" + matchedURL
		}
		parsedURL, err := url.Parse(matchedURL)
		if err != nil || parsedURL == nil {
			return true, matchedURL
		}
		checked := slices.ContainsFunc(viper.GetStringSlice(common.EnvUrlHostnameWhitelist), func(s string) bool {
			return strings.HasSuffix(parsedURL.Host, s)
		})
		if !checked {
			return true, parsedURL.Host
		}
	}
	return false, ""
}

func checkValidUrl(input string) error {
	imageUrl, err := url.Parse(input)
	if err != nil {
		return ErrUrlParsing
	}
	// if the url is text only, skip check
	if imageUrl.Scheme == "" && imageUrl.Host == "" {
		return ErrImageLinkTextOnly
	}
	if !slices.Contains(viper.GetStringSlice(common.EnvValidImageUrl), imageUrl.Hostname()) {
		return ErrInvalidImageHost
	}
	return nil
}

var reHole = regexp.MustCompile(`[^#]#(\d+)`)
var reFloor = regexp.MustCompile(`##(\d+)`)

func removeIDReprInContent(content string) string {
	content = " " + content
	content = reHole.ReplaceAllString(content, "")
	content = reFloor.ReplaceAllString(content, "")
	return strings.TrimSpace(content)
}
