// Package queue defines queue constants.
package queue

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Queue int

func (q *Queue) UnmarshalJSON(b []byte) error {
	var (
		s string
		i int
	)

	// First see if it is stored as native int.
	err := json.Unmarshal(b, &i)
	if err == nil {
		*q = Queue(i)
		return nil
	}

	// Must be a string.
	err = json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch strings.ToUpper(s) {
	case "RANKED_SOLO_5X5":
		*q = RankedSolo5x5
	case "RANKED_FLEX_SR":
		*q = RankedFlexSR
	case "RANKED_FLEX_TT":
		*q = RankedFlexTT
	case "ARAM_games_5x5":
		*q = ARAM_games_5x5
	default:
		return fmt.Errorf("invalid queue %q", s)
	}
	return nil
}

func (q Queue) String() string {
	switch q {
	case RankedSolo5x5:
		return "RANKED_SOLO_5x5"
	case RankedFlexSR:
		return "RANKED_FLEX_SR"
	case RankedFlexTT:
		return "RANKED_FLEX_TT"
	case ARAM_games_5x5:
		return "ARAM_games_5x5 "
	default:
		panic(fmt.Sprintf("invalid Queue %d", q))
	}
}

const (
	RankedSolo5x5                   Queue = 420
	RankedFlexSR                    Queue = 440
	RankedFlexTT                    Queue = 470
	CUSTOM                          Queue = 0   // Custom games
	NORMAL_3x3                      Queue = 8   // Normal 3v3 games
	NORMAL_5x5_BLIND                Queue = 2   // Normal 5v5 Blind Pick games
	NORMAL_5x5_DRAFT                Queue = 14  // Normal 5v5 Draft Pick games
	RANKED_SOLO_5x5                 Queue = 4   // Ranked Solo 5v5 games
	RANKED_PREMADE_5x5              Queue = 6   // Ranked Premade 5v5 games (Deprecated)
	RANKED_PREMADE_3x3              Queue = 9   // Historical Ranked Premade 3v3 games (Deprecated)
	RANKED_FLEX_TT                  Queue = 9   // Ranked Flex Twisted Treeline games
	RANKED_TEAM_3x3                 Queue = 41  // Ranked Team 3v3 games (Deprecated)
	RANKED_TEAM_5x5                 Queue = 42  // Ranked Team 5v5 games
	ODIN_5x5_BLIND                  Queue = 16  // Dominion 5v5 Blind Pick games
	ODIN_5x5_DRAFT                  Queue = 17  // Dominion 5v5 Draft Pick games
	BOT_5x5                         Queue = 7   // Historical Summoner's Rift Coop vs AI games (Deprecated)
	BOT_ODIN_5x5                    Queue = 25  // Dominion Coop vs AI games
	BOT_5x5_INTRO                   Queue = 31  // Summoner's Rift Coop vs AI Intro Bot games
	BOT_5x5_BEGINNER                Queue = 32  // Summoner's Rift Coop vs AI Beginner Bot games
	BOT_5x5_INTERMEDIATE            Queue = 33  // Historical Summoner's Rift Coop vs AI Intermediate Bot games
	BOT_TT_3x3                      Queue = 52  // Twisted Treeline Coop vs AI games
	GROUP_FINDER_5x5                Queue = 61  // Team Builder games
	ARAM_5x5                        Queue = 65  // ARAM games
	ONEFORALL_5x5                   Queue = 70  // One for All games
	FIRSTBLOOD_1x1                  Queue = 72  // Snowdown Showdown 1v1 games
	FIRSTBLOOD_2x2                  Queue = 73  // Snowdown Showdown 2v2 games
	SR_6x6                          Queue = 75  // Summoner's Rift 6x6 Hexakill games
	URF_5x5                         Queue = 76  // Ultra Rapid Fire games
	ONEFORALL_MIRRORMODE_5x5        Queue = 78  // One for All (Mirror mode)
	BOT_URF_5x5                     Queue = 83  // Ultra Rapid Fire games played against AI games
	NIGHTMARE_BOT_5x5_RANK1         Queue = 91  // Doom Bots Rank 1 games
	NIGHTMARE_BOT_5x5_RANK2         Queue = 92  // Doom Bots Rank 2 games
	NIGHTMARE_BOT_5x5_RANK5         Queue = 93  // Doom Bots Rank 5 games
	ASCENSION_5x5                   Queue = 96  // Ascension games
	HEXAKILL                        Queue = 98  // Twisted Treeline 6x6 Hexakill games
	BILGEWATER_ARAM_5x5             Queue = 100 // Butcher's Bridge games
	KING_PORO_5x5                   Queue = 300 // King Poro games
	COUNTER_PICK                    Queue = 310 // Nemesis games
	BILGEWATER_5x5                  Queue = 313 // Black Market Brawlers games
	SIEGE                           Queue = 315 // Nexus Siege games
	DEFINITELY_NOT_DOMINION_5x5     Queue = 317 // Definitely Not Dominion games
	ARURF_5X5                       Queue = 318 // All Random URF games
	ARSR_5x5                        Queue = 325 // All Random Summoner's Rift games
	TEAM_BUILDER_DRAFT_UNRANKED_5x5 Queue = 400 // Normal 5v5 Draft Pick games
	TEAM_BUILDER_DRAFT_RANKED_5x5   Queue = 410 // Ranked 5v5 Draft Pick games (Deprecated)
	TEAM_BUILDER_RANKED_SOLO        Queue = 420 // Ranked Solo games from current season that use Team Builder matchmaking
	RANKED_FLEX_SR                  Queue = 440 // Ranked Flex Summoner's Rift games
	ASSASSINATE_5x5                 Queue = 600 // Blood Hunt Assassin games
	DARKSTAR_3x3                    Queue = 610 // Darkstar games
	ARAM_games_5x5                  Queue = 450 // ARAM games 5x5, 65 deprecated
)
