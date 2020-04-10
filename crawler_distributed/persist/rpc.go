package persist

import (
	"crawler/engine"
	"crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}

	return err
}
