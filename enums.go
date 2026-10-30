package coc

type ClanType string
type Role string
type Village string
type WarFrequency string
type WarPreference string
type PlayerHouseElementType string
type State string
type Result string
type BattleModifier string
type BattleType string

const (
	OPEN        ClanType = "OPEN"
	INVITE_ONLY ClanType = "INVITE_ONLY"
	CLOSED      ClanType = "CLOSED"
)

const (
	GROUND PlayerHouseElementType = "GROUND"
	ROOF   PlayerHouseElementType = "ROOF"
	FOOT   PlayerHouseElementType = "FOOT"
	DECO   PlayerHouseElementType = "DECO"
)

const (
	NOT_MEMBER Role = "NOT_MEMBER"
	MEMBER     Role = "MEMBER"
	LEADER     Role = "LEADER"
	ADMIN      Role = "ADMIN"
	COLEADER   Role = "COLEADER"
)

const (
	HOME_VILLAGE Village = "HOME_VILLAGE"
	BUILDER_BASE Village = "BUILDER_BASE"
	CLAN_CAPITAL Village = "CLAN_CAPITAL"
)

const (
	UNKNOWN                 WarFrequency = "UNKNOWN"
	ALWAYS                  WarFrequency = "ALWAYS"
	MORE_THAN_ONCE_PER_WEEK WarFrequency = "MORE_THAN_ONCE_PER_WEEK"
	ONCE_PER_WEEK           WarFrequency = "ONCE_PER_WEEK"
	LESS_THAN_ONCE_PER_WEEK WarFrequency = "LESS_THAN_ONCE_PER_WEEK"
	NEVER                   WarFrequency = "NEVER"
	ANY                     WarFrequency = "ANY"
)

const (
	OUT WarPreference = "OUT"
	IN  WarPreference = "IN"
)

const (
	GROUP_NOT_FOUND State = "GROUP_NOT_FOUND"
	CLAN_NOT_FOUND  State = "CLAN_NOT_FOUND"
	ACCESS_DENIED   State = "ACCESS_DENIED"
	NOT_IN_WAR      State = "NOT_IN_WAR"
	IN_MATCHMAKING  State = "IN_MATCHMAKING"
	ENTER_WAR       State = "ENTER_WAR"
	MATCHED         State = "MATCHED"
	PREPARATION     State = "PREPARATION"
	WAR             State = "WAR"
	IN_WAR          State = "IN_WAR"
	ENDED           State = "ENDED"
)

const (
	LOSE Result = "LOSE"
	WIN  Result = "WIN"
	TIE  Result = "TIE"
)

const (
	NONE        BattleModifier = "NONE"
	HARD_MODE   BattleModifier = "HARD_MODE"
	MINUS_ONE   BattleModifier = "MINUS_ONE"
	MINUS_TWO   BattleModifier = "MINUS_TWO"
	MINUS_THREE BattleModifier = "MINUS_THREE"
)

const (
	BattleTypeHomeVillage BattleType = "HOME_VILLAGE"
	BattleTypeRanked      BattleType = "RANKED"
	BattleTypeLegend      BattleType = "LEGEND"
)
