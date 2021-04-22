package decorators

import (
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-settings/internal/db/models"
)

var regexSchema = regexp.MustCompile("^http(s)?://") // "http://" or "https://"

type AddSchemeToUrl struct {
	client        *http.Client
	cacheTimeMin  int // use negative value to disable cache
	logger        log15.Logger
	prevDecorator Decorator
	cache         *CachedSetting
	mutex         sync.Mutex
}

type CachedSetting struct {
	setting   *models.Config
	expiredAt time.Time
}

func NewAddSchemeToUrl(client *http.Client, cacheTimeMin int, logger log15.Logger, prevDecorator Decorator) Decorator {
	return &AddSchemeToUrl{
		client:        client,
		cacheTimeMin:  cacheTimeMin,
		logger:        logger.New("decorator", "AddSchemaToUrl"),
		prevDecorator: prevDecorator,
		mutex:         sync.Mutex{},
	}
}

// add http or https scheme to the value of the setting
func (s *AddSchemeToUrl) GetSetting(setting models.Config) (*models.Config, error) {
	if s.prevDecorator != nil {
		setting, err := s.prevDecorator.GetSetting(setting)
		if err != nil {
			return setting, err
		}
	}

	if s.cacheTimeMin >= 0 {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	currentTime := time.Now()
	if s.cache != nil && currentTime.Before(s.cache.expiredAt) {
		return s.cache.setting, nil
	}

	domainName := regexSchema.ReplaceAllString(setting.Value, "") // cut scheme from url
	setting.Value = "https://" + domainName                       // set secure url

	_, err := s.client.Get(setting.Value)
	// we get an error if ssl certificate is not valid
	if err != nil {
		s.logger.Info("failed loading site url", "err", err)

		// we do not return error. we set not secure scheme
		setting.Value = "http://" + domainName
	}

	var cacheTimeMin int
	if s.cacheTimeMin > 0 {
		cacheTimeMin = s.cacheTimeMin
	} else if s.cacheTimeMin == 0 {
		cacheTimeMin = 5256000 // 10 years
	}

	// save result in cache
	if cacheTimeMin > 0 {
		cacheExpiredAt := time.Now().Local().Add(time.Minute * time.Duration(cacheTimeMin))
		s.cache = &CachedSetting{expiredAt: cacheExpiredAt, setting: &setting}
	}

	return &setting, nil
}
