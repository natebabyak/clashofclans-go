package coc

// List is one API response page. Items contains only this page's results.
// Paging exposes cursors for explicit calls with After or Before.
type List[T any] struct {
	// Items contains the results in API order.
	Items []T `json:"items"`
	// Paging contains navigation cursors supplied by the API.
	Paging Paging `json:"paging"`
}

// Paging describes navigation between response pages.
type Paging struct {
	// Cursors identifies adjacent pages when more results are available.
	Cursors Cursors `json:"cursors"`
}

// Cursors contains opaque tokens. An empty token means no page in that direction.
// Pass tokens unchanged to After or Before; do not decode or escape them.
type Cursors struct {
	// After identifies the next page.
	After string `json:"after"`
	// Before identifies the previous page.
	Before string `json:"before"`
}

// BuilderBaseLeague identifies a builder base league.
type BuilderBaseLeague struct {
	Name string
	Id   int
}

// CapitalLeague identifies a clan capital league.
type CapitalLeague struct {
	Name string
	Id   int
}

// Clan contains clan details. Search results may omit fields returned by GetClan.
type Clan struct {
	MemberList                  []ClanMember
	WarLeague                   WarLeague
	CapitalLeague               CapitalLeague
	Tag                         string
	ClanLevel                   int
	WarWinStreak                int
	WarWins                     int
	WarTies                     int
	WarLosses                   int
	ClanPoints                  int
	ChatLanguage                Language
	WarFrequency                WarFrequency
	ClanBuilderBasePoints       int
	ClanCapitalPoints           int
	RequiredTrophies            int
	RequiredBuilderBaseTrophies int
	RequiredTownhallLevel       int
	IsFamilyFriendly            bool
	IsWarLogPublic              bool
	Labels                      []Label
	Name                        string
	Location                    Location
	Type                        ClanType
	Members                     int
	Description                 string
	ClanCapital                 ClanCapital
	BadgeUrls                   struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanBuilderBaseRanking is one clan entry in builder base rankings.
type ClanBuilderBaseRanking struct {
	ClanBuilderBasePoints int
	ClanPoints            int
	ClanLevel             int
	Location              Location
	Members               int
	Tag                   string
	Name                  string
	Rank                  int
	PreviousRank          int
	BadgeUrls             struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanCapital describes a clan capital and its districts.
type ClanCapital struct {
	CapitalHallLevel int
	Districts        []ClanDistrictData
}

// ClanCapitalRanking is one clan entry in capital rankings.
type ClanCapitalRanking struct {
	ClanCapitalPoints int
	ClanPoints        int
	ClanLevel         int
	Location          Location
	Members           int
	Tag               string
	Name              string
	Rank              int
	PreviousRank      int
	BadgeUrls         struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanCapitalRaidSeason describes a raid weekend, including its logs and member rewards.
type ClanCapitalRaidSeason struct {
	AttackLog               []ClanCapitalRaidSeasonAttackLogEntry
	DefenseLog              []ClanCapitalRaidSeasonDefenseLogEntry
	State                   string
	StartTime               string
	EndTime                 string
	CapitalTotalLoot        int
	RaidsCompleted          int
	TotalAttacks            int
	EnemyDistrictsDestroyed int
	OffensiveReward         int
	DefensiveReward         int
	Members                 []ClanCapitalRaidSeasonMember
}

// ClanCapitalRaidSeasonAttack describes an attack on a capital district.
type ClanCapitalRaidSeasonAttack struct {
	Attacker           ClanCapitalRaidSeasonAttacker
	DestructionPercent int
	Stars              int
}

// ClanCapitalRaidSeasonAttackLogEntry describes a capital raid attack log entry.
type ClanCapitalRaidSeasonAttackLogEntry struct {
	Defender           ClanCapitalRaidSeasonClanInfo
	AttackCount        int
	DistrictCount      int
	DistrictsDestroyed int
	Districts          []ClanCapitalRaidSeasonDistrict
}

// ClanCapitalRaidSeasonAttacker identifies a capital raid attacker.
type ClanCapitalRaidSeasonAttacker struct {
	Tag  string
	Name string
}

// ClanCapitalRaidSeasonClanInfo identifies a clan in a capital raid log.
type ClanCapitalRaidSeasonClanInfo struct {
	Tag       string
	Name      string
	Level     int
	BadgeUrls struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanCapitalRaidSeasonDefenseLogEntry describes a capital raid defense log entry.
type ClanCapitalRaidSeasonDefenseLogEntry struct {
	Defender           ClanCapitalRaidSeasonClanInfo
	AttackCount        int
	DistrictCount      int
	DistrictsDestroyed int
	Districts          []ClanCapitalRaidSeasonDistrict
}

// ClanCapitalRaidSeasonDistrict describes a district targeted during a capital raid.
type ClanCapitalRaidSeasonDistrict struct {
	Stars              int
	Name               string
	Id                 int
	DestructionPercent int
	AttackCount        int
	TotalLooted        int
	Attacks            []ClanCapitalRaidSeasonAttack
	DistrictHallLevel  int
}

// ClanCapitalRaidSeasonMember contains a member's raid weekend contribution.
type ClanCapitalRaidSeasonMember struct {
	Tag                    string
	Name                   string
	Attacks                int
	AttackLimit            int
	BonusAttackLimit       int
	CapitalResourcesLooted int
}

// ClanDistrictData describes a clan capital district.
type ClanDistrictData struct {
	Name              string
	Id                int
	DistrictHallLevel int
}

// ClanMember contains a clan roster entry; use GetPlayer for a full player profile.
type ClanMember struct {
	League              League
	LeagueTier          LeagueTier
	BuilderBaseLeague   BuilderBaseLeague
	Tag                 string
	Name                string
	Role                Role
	TownHallLevel       int
	ExpLevel            int
	ClanRank            int
	PreviousClanRank    int
	Donations           int
	DonationsReceived   int
	Trophies            int
	BuilderBaseTrophies int
	PlayerHouse         PlayerHouse
}

// ClanRanking is one clan entry in home village rankings.
type ClanRanking struct {
	ClanLevel    int
	ClanPoints   int
	Location     Location
	Members      int
	Tag          string
	Name         string
	Rank         int
	PreviousRank int
	BadgeUrls    struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanWar describes a current or clan war league war.
type ClanWar struct {
	Clan                 WarClan
	TeamSize             int
	AttacksPerMember     int
	BattleModifier       BattleModifier
	Opponent             WarClan
	StartTime            string
	State                State
	EndTime              string
	PreparationStartTime string
}

// ClanWarLeagueWar is the response for an individual CWL war (same schema as ClanWar).
type ClanWarLeagueWar = ClanWar

// ClanWarAttack contains the result of an individual war attack.
type ClanWarAttack struct {
	Order                 int
	AttackerTag           string
	DefenderTag           string
	Stars                 int
	DestructionPercentage int
	Duration              int
}

// ClanWarLeagueClan contains a clan's clan war league roster.
type ClanWarLeagueClan struct {
	Tag       string
	ClanLevel int
	Name      string
	Members   []ClanWarLeagueClanMember
	BadgeUrls struct {
		Small  string
		Large  string
		Medium string
	}
}

// ClanWarLeagueClanMember identifies a player on a clan war league roster.
type ClanWarLeagueClanMember struct {
	Tag           string
	TownHallLevel int
	Name          string
}

// ClanWarLeagueGroup contains the clans and rounds of a clan war league group.
type ClanWarLeagueGroup struct {
	Tag    string
	State  State
	Season string
	Clans  []ClanWarLeagueClan
	Rounds []ClanWarLeagueRound
}

// ClanWarLeagueRound contains war tags for one clan war league round.
type ClanWarLeagueRound struct {
	WarTags []string
}

// ClanWarLogEntry summarizes a completed clan war.
type ClanWarLogEntry struct {
	Clan             WarClan
	TeamSize         int
	Opponent         WarClan
	BattleModifier   BattleModifier
	AttacksPerMember int
	EndTime          string
	Result           Result
}

// ClanWarMember contains a participant's war attacks and defenses.
type ClanWarMember struct {
	Tag                string
	Name               string
	MapPosition        int
	TownhallLevel      int
	OpponentAttacks    int
	BestOpponentAttack ClanWarAttack
	Attacks            []ClanWarAttack
}

// GoldPassSeason contains the current gold pass season's API timestamps.
type GoldPassSeason struct {
	StartTime string
	EndTime   string
}

// Label identifies an API label for clan or player filtering.
type Label struct {
	Name     string
	Id       int
	IconUrls struct {
		Small  string
		Medium string
	}
}

// Language identifies a clan chat language.
type Language struct {
	Name         string
	Id           int
	LanguageCode string
}

// League identifies a home village league.
type League struct {
	Name     string
	Id       int
	IconUrls struct {
		Small string
		Tiny  string
	}
}

// LeagueBattleLogEntry describes an attack or defense in a ranked league group.
type LeagueBattleLogEntry struct {
	OpponentPlayerTag     string
	OpponentName          string
	Stars                 int
	DestructionPercentage int
	Trophies              int
	CreationTime          string
}

// LeagueGroup contains ranked league members and battle logs.
type LeagueGroup struct {
	Members     []LeagueGroupMember
	AttackLogs  []LeagueBattleLogEntry
	DefenseLogs []LeagueBattleLogEntry
}

// LeagueGroupMember contains a player's ranked league group standing.
type LeagueGroupMember struct {
	PlayerTag        string
	PlayerName       string
	ClanTag          string
	ClanName         string
	LeagueTrophies   int
	AttackWinCount   int
	AttackLoseCount  int
	DefenseWinCount  int
	DefenseLoseCount int
}

// LeagueSeason identifies a Legend League season.
type LeagueSeason struct {
	Id string
}

// LeagueSeasonResult contains a player's ranked league season result.
type LeagueSeasonResult struct {
	LeagueSeasonId int
	LeagueTrophies int
	LeagueTierId   int
	Placement      int
	AttackWins     int
	AttackLosses   int
	AttackStars    int
	DefenseWins    int
	DefenseLosses  int
	DefenseStars   int
	MaxBattles     int
}

// LeagueTier identifies a ranked league tier.
type LeagueTier struct {
	Name     string
	Id       int
	IconUrls struct {
		Small string
		Large string
	}
}

// LegendLeagueTournamentSeasonResult contains a player's Legend League season standing.
type LegendLeagueTournamentSeasonResult struct {
	Trophies int
	Id       string
	Rank     int
}

// Location identifies a geographic region; use its Id with LocationID for rankings.
type Location struct {
	LocalizedName string
	Id            int
	Name          string
	IsCountry     bool
	CountryCode   string
}

// BattleLogEntry describes a recent player battle.
type BattleLogEntry struct {
	BattleType            BattleType
	Attack                bool
	ArmyShareCode         string
	OpponentPlayerTag     string
	Stars                 int
	DestructionPercentage int
	LootedResources       []Resource
	ExtraLootedResources  []Resource
	AvailableLoot         []Resource
}

// Player contains a player profile, troops, heroes, and equipment.
type Player struct {
	Clan                PlayerClan
	League              League
	LeagueTier          LeagueTier
	BuilderBaseLeague   BuilderBaseLeague
	Role                Role
	WarPreference       WarPreference
	AttackWins          int
	DefenseWins         int
	TownHallLevel       int
	TownHallWeaponLevel int
	LegendStatistics    PlayerLegendStatistics
	Troops              []PlayerItemLevel
	// Heroes contains hero levels and their currently equipped items.
	Heroes []PlayerItemLevel
	// HeroEquipment contains the player's equipment inventory and upgrade levels.
	HeroEquipment            []PlayerItemLevel
	Spells                   []PlayerItemLevel
	Labels                   []Label
	Tag                      string
	Name                     string
	ExpLevel                 int
	Trophies                 int
	BestTrophies             int
	Donations                int
	DonationsReceived        int
	BuilderHallLevel         int
	BuilderBaseTrophies      int
	BestBuilderBaseTrophies  int
	WarStars                 int
	Achievements             []PlayerAchievementProgress
	ClanCapitalContributions int
	PlayerHouse              PlayerHouse
	CurrentLeagueGroupTag    string
	CurrentLeagueSeasonId    int
	PreviousLeagueGroupTag   string
	PreviousLeagueSeasonId   int
}

// PlayerAchievementProgress contains progress toward an in-game achievement.
type PlayerAchievementProgress struct {
	Stars          int
	Value          int
	Name           string
	Target         int
	Info           string
	CompletionInfo string
	Village        Village
}

// PlayerBuilderBaseRanking is one player entry in builder base rankings.
type PlayerBuilderBaseRanking struct {
	Clan                PlayerRankingClan
	BuilderBaseLeague   BuilderBaseLeague
	Tag                 string
	Name                string
	ExpLevel            int
	Rank                int
	PreviousRank        int
	BuilderBaseTrophies int
}

// PlayerClan identifies the clan in a player profile.
type PlayerClan struct {
	Tag       string
	ClanLevel int
	Name      string
	BadgeUrls struct {
		Small  string
		Large  string
		Medium string
	}
}

// PlayerHouse contains a player's capital house decorations.
type PlayerHouse struct {
	Elements []PlayerHouseElement
}

// PlayerItemLevel describes a troop, spell, hero, or equipment upgrade. Heroes include equipped items.
type PlayerItemLevel struct {
	Level              int
	Name               string
	MaxLevel           int
	Village            Village
	SuperTroopIsActive bool
	// Equipment contains the items equipped by this hero, when present.
	Equipment []Equipment
}

// PlayerLegendStatistics contains current and historical Legend League results.
type PlayerLegendStatistics struct {
	CurrentSeason             LegendLeagueTournamentSeasonResult
	BestSeason                LegendLeagueTournamentSeasonResult
	PreviousSeason            LegendLeagueTournamentSeasonResult
	PreviousBuilderBaseSeason LegendLeagueTournamentSeasonResult
	BestBuilderBaseSeason     LegendLeagueTournamentSeasonResult
	LegendTrophies            int
}

// PlayerRanking is one player ranking entry; use GetPlayer for the full profile.
type PlayerRanking struct {
	Clan         PlayerRankingClan
	League       League
	LeagueTier   LeagueTier
	AttackWins   int
	DefenseWins  int
	Tag          string
	Name         string
	ExpLevel     int
	Rank         int
	PreviousRank int
	Trophies     int
}

// PlayerRankingClan identifies the clan in a player ranking entry.
type PlayerRankingClan struct {
	Tag       string
	Name      string
	BadgeUrls struct {
		Small  string
		Large  string
		Medium string
	}
}

// Resource contains a named resource amount.
type Resource struct {
	Name   string
	Amount int
}

// VerifyTokenResponse contains the API's player-token verification result.
type VerifyTokenResponse struct {
	Tag    string
	Token  string
	Status string
}

// WarClan contains a clan's participants and performance in a war.
type WarClan struct {
	DestructionPercentage float64
	Tag                   string
	Name                  string
	BadgeUrls             struct {
		Small  string
		Large  string
		Medium string
	}
	ClanLevel int
	Attacks   int
	Stars     int
	ExpEarned int
	Members   []ClanWarMember
}

// WarLeague identifies a clan war league.
type WarLeague struct {
	Name string
	Id   int
}

// Equipment describes an equipped hero item and its upgrade level.
type Equipment struct {
	// Name is the in-game equipment name.
	Name string
	// Level is the current upgrade level.
	Level int
	// MaxLevel is the maximum upgrade level.
	MaxLevel int
	// Village identifies the village where this equipment is used.
	Village Village
}

// PlayerHouseElement describes a decoration in a player's capital house.
type PlayerHouseElement struct {
	Type PlayerHouseElementType
	Id   int
}
