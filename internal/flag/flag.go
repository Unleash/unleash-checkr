package flag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

var (
	all      = "all"
	active   = "actives"
	inactive = "inactive"
)

// Get all flags
func Get(ctx *context.Context) ([]Flag, error) {
	log.Info("Getting flags")

	url := fmt.Sprintf("%s/api/client/features", ctx.Config.URL)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var flags Flags
	err = json.NewDecoder(res.Body).Decode(&flags)
	if err != nil {
		return nil, err
	}

	ff := filter(ctx, flags.Features)

	return ff, nil
}

func filter(ctx *context.Context, flags []Flag) []Flag {
	const day = 24 * time.Hour
	dueDate := time.Now().Add(-time.Duration(ctx.Config.Expires) * day)
	dueDateUnix := dueDate.Unix()

	log.WithField("due-date", dueDate).
		Info("Filtering flags")

	var ff []Flag
	for _, f := range flags {
		if f.CreatedAt.Unix() < dueDateUnix {
			ff = append(ff, f)
		}
	}

	return ff
}
