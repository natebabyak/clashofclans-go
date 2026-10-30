package coc

type List[T any] struct {
	Items  []T
	Paging Paging
}

type Paging struct {
	Cursors struct {
		After  string
		Before string
	}
}

type BuilderBaseLeague struct {
	Name string
	Id   int
}

type CapitalLeague struct {
	Name string
	Id   int
}

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

type ClanCapital struct {
	CapitalHallLevel int
	Districts        []ClanDistrictData
}

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

type ClanCapitalRaidSeasonAttack struct {
	Attacker           ClanCapitalRaidSeasonAttacker
	DestructionPercent int
	Stars              int
}

type ClanCapitalRaidSeasonAttackLogEntry struct {
	Defender           ClanCapitalRaidSeasonClanInfo
	AttackCount        int
	DistrictCount      int
	DistrictsDestroyed int
	Districts          []ClanCapitalRaidSeasonDistrict
}

type ClanCapitalRaidSeasonAttacker struct {
	Tag  string
	Name string
}

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

type ClanCapitalRaidSeasonDefenseLogEntry struct {
	Defender           ClanCapitalRaidSeasonClanInfo
	AttackCount        int
	DistrictCount      int
	DistrictsDestroyed int
	Districts          []ClanCapitalRaidSeasonDistrict
}

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

type ClanCapitalRaidSeasonMember struct {
	Tag                    string
	Name                   string
	Attacks                int
	AttackLimit            int
	BonusAttackLimit       int
	CapitalResourcesLooted int
}

type ClanDistrictData struct {
	Name              string
	Id                int
	DistrictHallLevel int
}

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

type ClanWarAttack struct {
	Order                 int
	AttackerTag           string
	DefenderTag           string
	Stars                 int
	DestructionPercentage int
	Duration              int
}

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

type ClanWarLeagueClanMember struct {
	Tag           string
	TownHallLevel int
	Name          string
}

type ClanWarLeagueGroup struct {
	Tag    string
	State  State
	Season string
	Clans  []ClanWarLeagueClan
	Rounds []ClanWarLeagueRound
}

type ClanWarLeagueRound struct {
	WarTags []string
}

type ClanWarLogEntry struct {
	Clan             WarClan
	TeamSize         int
	Opponent         WarClan
	BattleModifier   BattleModifier
	AttacksPerMember int
	EndTime          string
	Result           Result
}

type ClanWarMember struct {
	Tag                string
	Name               string
	MapPosition        int
	TownhallLevel      int
	OpponentAttacks    int
	BestOpponentAttack ClanWarAttack
	Attacks            []ClanWarAttack
}

type GoldPassSeason struct {
	StartTime string
	EndTime   string
}

type Label struct {
	Name     string
	Id       int
	IconUrls struct {
		Small  string
		Medium string
	}
}

type Language struct {
	Name         string
	Id           int
	LanguageCode string
}

type League struct {
	Name     string
	Id       int
	IconUrls struct {
		Small string
		Tiny  string
	}
}

type LeagueBattleLogEntry struct {
	OpponentPlayerTag     string
	OpponentName          string
	Stars                 int
	DestructionPercentage int
	Trophies              int
	CreationTime          string
}

type LeagueGroup struct {
	Members     []LeagueGroupMember
	AttackLogs  []LeagueBattleLogEntry
	DefenseLogs []LeagueBattleLogEntry
}

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

type LeagueSeason struct {
	Id string
}

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

type LeagueTier struct {
	Name     string
	Id       int
	IconUrls struct {
		Small string
		Large string
	}
}

type LegendLeagueTournamentSeasonResult struct {
	Trophies int
	Id       string
	Rank     int
}

type Location struct {
	LocalizedName string
	Id            int
	Name          string
	IsCountry     bool
	CountryCode   string
}

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

type Player struct {
	Clan                     PlayerClan
	League                   League
	LeagueTier               LeagueTier
	BuilderBaseLeague        BuilderBaseLeague
	Role                     Role
	WarPreference            WarPreference
	AttackWins               int
	DefenseWins              int
	TownHallLevel            int
	TownHallWeaponLevel      int
	LegendStatistics         PlayerLegendStatistics
	Troops                   []PlayerItemLevel
	Heroes                   []PlayerItemLevel
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

type PlayerAchievementProgress struct {
	Stars          int
	Value          int
	Name           string
	Target         int
	Info           string
	CompletionInfo string
	Village        Village
}

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

type PlayerHouse struct {
	Elements []struct {
		Type PlayerHouseElementType
		Id   int
	}
}

type PlayerItemLevel struct {
	Level              int
	Name               string
	MaxLevel           int
	Village            Village
	SuperTroopIsActive bool
	Equipment          []struct {
		Name     string
		Level    int
		MaxLevel int
		Village  Village
	}
}

type PlayerLegendStatistics struct {
	CurrentSeason             LegendLeagueTournamentSeasonResult
	BestSeason                LegendLeagueTournamentSeasonResult
	PreviousSeason            LegendLeagueTournamentSeasonResult
	PreviousBuilderBaseSeason LegendLeagueTournamentSeasonResult
	BestBuilderBaseSeason     LegendLeagueTournamentSeasonResult
	LegendTrophies            int
}

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

type PlayerRankingClan struct {
	Tag       string
	Name      string
	BadgeUrls struct {
		Small  string
		Large  string
		Medium string
	}
}

type Resource struct {
	Name   string
	Amount int
}

type VerifyTokenResponse struct {
	Tag    string
	Token  string
	Status string
}

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

type WarLeague struct {
	Name string
	Id   int
}
