package tppmessage

// MGOUserData is for CMD_GET_MGO_USER_DATA
type MGOUserData struct {
	Name     string `json:"name"`
	PlayTime int    `json:"play_time"`
	Version  int    `json:"version"`
}

// MGOCharacterData represents the entire character data structure from default_character.json
type MGOCharacterData struct {
	Match               MatchData        `json:"match"`
	Version             int64            `json:"version"`
	LastActive          int              `json:"last_active"`
	SelectedBgm         int              `json:"selected_bgm"`
	CharacterList       []MGOCharacter   `json:"character_list"`
	PresetRadioRuleList []PresetRadioSet `json:"preset_radio_rule_list"`
}

type MatchData struct {
	AutoLeave        int           `json:"auto_leave"`
	PlayerNum        int           `json:"player_num"`
	HostComment      int           `json:"host_comment"`
	MaxCapacity      int           `json:"max_capacity"`
	BriefingTime     int           `json:"briefing_time"`
	MissionSlotList  []MissionSlot `json:"mission_slot_list"`
	MissionSlotCount int           `json:"mission_slot_count"`
}

type MissionSlot struct {
	Map             int `json:"map"`
	Rule            int `json:"rule"`
	Rush            int `json:"rush"`
	Time            int `json:"time"`
	Flags           int `json:"flags"`
	Night           int `json:"night"`
	Ticket          int `json:"ticket"`
	Weather         int `json:"weather"`
	UniqueCharacter int `json:"unique_character"`
}

// MGOCharacter represents a single character from the character_list
type MGOCharacter struct {
	CharacterID int    `json:"-"` // Not in original JSON, for DB use
	Name        string `json:"name"`
	Avatar      Avatar `json:"avatar"`
	PlayerType  int    `json:"player_type"`
	LastLoadout int    `json:"last_loadout"`
	PlayerClass int    `json:"player_class"`
}

type Avatar struct {
	Voice                 int   `json:"voice"`
	FaceRace              int   `json:"face_race"`
	FaceType              int   `json:"face_type"`
	FaceColor             int   `json:"face_color"`
	HairColor             int   `json:"hair_color"`
	HairStyle             int   `json:"hair_style"`
	BeardStyle            int   `json:"beard_style"`
	BeardLength           int   `json:"beard_length"`
	TattooColor           int   `json:"tattoo_color"`
	EyebrowStyle          int   `json:"eyebrow_style"`
	EyebrowWidth          int   `json:"eyebrow_width"`
	FaceVariation         int   `json:"face_variation"`
	LeftEyeColor          int   `json:"left_eye_color"`
	AccessoryFlags        int   `json:"accessory_flags"`
	RightEyeColor         int   `json:"right_eye_color"`
	MotionFrameList       []int `json:"motion_frame_list"`
	LeftEyeBrightness     int   `json:"left_eye_brightness"`
	RightEyeBrightness    int   `json:"right_eye_brightness"`
	GashOrTattooVariation int   `json:"gash_or_tattoo_variation"`
}

type PresetRadioSet struct {
	PresetRadioIDList []int `json:"preset_radio_id_list"`
}

// MGOLoadoutData represents the entire loadout data structure from default_loadout.json
type MGOLoadoutData struct {
	Version       int64               `json:"version"`
	CharacterList []CharacterLoadouts `json:"character_list"`
}

type CharacterLoadouts struct {
	LoadoutList []MGOLoadout `json:"loadout_list"`
}

// MGOLoadout represents a single loadout.
type MGOLoadout struct {
	LoadoutIndex      int             `json:"-"` // Not in original JSON, for DB use
	Name              string          `json:"name"`
	GearList          []Gear          `json:"gear_list"`
	ItemList          []Item          `json:"item_list"`
	SkillList         []Skill         `json:"skill_list"`
	WeaponList        []Weapon        `json:"weapon_list"`
	SupportWeaponList []SupportWeapon `json:"support_weapon_list"`
}

type Gear struct {
	ID        uint   `json:"id"`
	Model     uint32 `json:"model"`
	ColorList []uint `json:"color_list"`
}

type Item struct {
	ID   uint32 `json:"id"`
	Slot int    `json:"slot"`
}

type Skill struct {
	ID   uint32 `json:"id"`
	Slot int    `json:"slot"`
}

type Weapon struct {
	ID        uint32   `json:"id"`
	Slot      int      `json:"slot"`
	PartList  []uint32 `json:"part_list,omitempty"`
	ColorList []uint   `json:"color_list,omitempty"`
}

type SupportWeapon struct {
	ID   uint32 `json:"id"`
	Slot int    `json:"slot"`
}