package Application

import (
	"context"
	"fmt"

	"github.com/Domain/region"
)

type Summoner struct {
	ProfileIconID int    `json:"profileIconID"` // ID of the summoner icon associated with the summoner.
	Name          string `json:"name"`          // Summoner name.
	PUUID         string `json:"puuid"`         // PUUID is the player universally unique identifier.
	SummonerLevel int64  `json:"summonerLevel"` // Summoner level associated with the summoner.
	AccountID     string `json:"accountID"`     // Encrypted account ID.
	ID            string `json:"id"`            // Encrypted summoner ID.
	RevisionDate  int64  `json:"revisionDate"`  // Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
}

func (c *client) GetBySummonerName(ctx context.Context, r region.Region, name string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v4/summoners/by-name", fmt.Sprintf("/%s", name), nil, &res)
	return &res, err
}
