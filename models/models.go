package models

type Message struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Symbol         string              `json:"symbol"`
	Rank           int                 `json:"rank"`
	IsNew          bool                `json:"is_new"`
	IsActive       bool                `json:"is_active"`
	Type           string              `json:"type"`
	Logo           string              `json:"logo"`
	Tags           []Tag               `json:"tags"`
	Team           []Team              `json:"team"`
	Description    string              `json:"description"`
	Message        string              `json:"message"`
	OpenSource     bool                `json:"open_source"`
	StartedAt      string              `json:"started_at"`
	DevStatus      string              `json:"development_status"`
	HardwareWallet bool                `json:"hardware_wallet"`
	ProofType      string              `json:"proof_type"`
	OrgStructure   string              `json:"org_structure"`
	HashAlgo       string              `json:"hash_algorithm"`
	Links          map[string][]string `json:"links"`
	ExtendedLinks  []ExtendedLink      `json:"links_extended"`
	Whitepaper     Whitepaper          `json:"whitepaper"`
	FirstDataAt    string              `json:"first_data_at"`
	LastDataAt     string              `json:"last_data_at"`
}

type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CoinCounter int    `json:"coin_counter"`
	IcoCounter  int    `json:"ico_counter"`
}

type Team struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type ExtendedLink struct {
	URL   string `json:"url"`
	Type  string `json:"type"`
	Stats Stats  `json:"stats"`
}

type Stats struct {
	Subscribers  int `json:"subscribers"`
	Contributors int `json:"contributors"`
	Stars        int `json:"stars"`
	Followers    int `json:"followers"`
}

type Whitepaper struct {
	Link      string `json:"link"`
	Thumbnail string `json:"thumbnail"`
}
