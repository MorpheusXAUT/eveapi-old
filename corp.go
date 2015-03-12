package eveapi

import (
	"net/url"
	"strconv"
)

const (
	CorpContactListURL     = "/corp/ContactList.xml.aspx"
	CorpAccountBalanceURL  = "/corp/AccountBalance.xml.aspx"
	CorpStarbaseListURL    = "/corp/StarbaseList.xml.aspx"
	CorpStarbaseDetailsURL = "/corp/StarbaseDetail.xml.aspx"
)

type Contact struct {
	ID       string `xml:"contactID,attr"`
	Name     string `xml:"contactName,attr"`
	Standing int    `xml:"standing,attr"`
}

type ContactSubList struct {
	Name     string    `xml:"name,attr"`
	Contacts []Contact `xml:"row"`
}

func (api API) CorpContactList() (*ContactList, error) {
	output := ContactList{}
	err := api.Call(CorpContactListURL, nil, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

type ContactList struct {
	APIResult
	ContactList []ContactSubList `xml:"result>rowset"`
}

func (c ContactList) Corporate() []Contact {
	for _, v := range c.ContactList {
		if v.Name == "corporateContactList" {
			return v.Contacts
		}
	}
	return nil
}
func (c ContactList) Alliance() []Contact {
	for _, v := range c.ContactList {
		if v.Name == "allianceContactList" {
			return v.Contacts
		}
	}
	return nil
}

type AccountBalance struct {
	APIResult
	Accounts []struct {
		ID      int     `xml:"accountID,attr"`
		Key     int     `xml:"accountKey,attr"`
		Balance float64 `xml:"balance,attr"`
	} `xml:"result>rowset>row"`
}

func (api API) CorpAccountBalances() (*AccountBalance, error) {
	output := AccountBalance{}
	err := api.Call(CorpAccountBalanceURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

type StarbaseList struct {
	APIResult
	Starbases []Starbase `xml:"result>rowset>row"`
}

type Starbase struct {
	ID              int64   `xml:"itemID,attr"`
	TypeID          int     `xml:"typeID,attr"`
	LocationID      int     `xml:"locationID,attr"`
	MoonID          int     `xml:"moonID,attr"`
	State           int     `xml:"state,attr"`
	StateTimestamp  eveTime `xml:"stateTimestamp,attr"`
	OnlineTimestamp eveTime `xml:"onlineTimestamp,attr"`
	StandingOwnerID int     `xml:"standingOwnerID,attr"`
}

func (api API) CorpStarbaseList() (*StarbaseList, error) {
	output := StarbaseList{}
	err := api.Call(CorpStarbaseListURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

type StarbaseDetails struct {
	APIResult
	State           int     `xml:"result>state"`
	StateTimestamp  eveTime `xml:"result>stateTimestamp"`
	OnlineTimestamp eveTime `xml:"result>onlineTimestamp"`
	GeneralSettings struct {
		UsageFlags              int  `xml:"usageFlags"`
		DeployFlags             int  `xml:"deployFlags"`
		AllowCorporationMembers bool `xml:"allowCorporationMembers"`
		AllowAllianceMembers    bool `xml:"allowAllianceMembers"`
	} `xml:"result>generalSettings"`
	CombatSettings struct {
		UseStandingsFrom struct {
			OwnerID int `xml:"ownerID,attr"`
		} `xml:"useStandingsFrom"`
		OnStandingDrop struct {
			Standing int `xml:"standing,attr"`
		} `xml:"onStandingDrop"`
		OnStatusDrop struct {
			Enabled  bool `xml:"enabled,attr"`
			Standing int  `xml:"standing,attr"`
		} `xml:"onStatusDrop"`
		OnAgression struct {
			Enabled bool `xml:"enabled,attr"`
		} `xml:"onAgression"`
		OnCorporationWar struct {
			Enabled bool `xml:"enabled, attr"`
		} `xml:"onCorporationWar"`
	} `xml:"result>combatSettings"`
	Fuel []struct {
		TypeID   int `xml:"typeID,attr"`
		Quantity int `xml:"quantity,attr"`
	} `xml:"result>rowset>row"`
}

func (api API) CorpStarbaseDetails(starbaseID int64) (*StarbaseDetails, error) {
	output := StarbaseDetails{}
	args := url.Values{}
	args.Set("itemID", strconv.FormatInt(starbaseID, 10))
	err := api.Call(CorpStarbaseDetailsURL, args, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}
