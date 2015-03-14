package eveapi

import (
	"net/url"
	"strconv"
)

const (
	CorpCorporationSheetURL = "/corp/CorporationSheet.xml.aspx"
	CorpContactListURL      = "/corp/ContactList.xml.aspx"
	CorpAccountBalanceURL   = "/corp/AccountBalance.xml.aspx"
	CorpStarbaseListURL     = "/corp/StarbaseList.xml.aspx"
	CorpStarbaseDetailsURL  = "/corp/StarbaseDetail.xml.aspx"
)

type CorporationSheet struct {
	APIResult
	ID           int64  `xml:"result>corporationID"`
	Name         string `xml:"result>corporationName"`
	Ticker       string `xml:"result>ticker"`
	CEOID        int64  `xml:"result>ceoID"`
	CEOName      string `xml:"result>ceoName"`
	StationID    int64  `xml:"result>stationID"`
	StationName  string `xml:"result>stationName"`
	Description  string `xml:"result>description"`
	URL          string `xml:"result>url"`
	AllianceID   int64  `xml:"result>allianceID"`
	AllianceName string `xml:"result>allianceName"`
	FactionID    int64  `xml:"result>factionID"`
	TaxRate      int64  `xml:"result>taxRate"`
	MemberCount  int64  `xml:"result>memberCount"`
	Shares       int64  `xml:"result>shares"`
	// Logo ignored
}

func (api API) CorpCorporationSheet(corporationID int64) (*CorporationSheet, error) {
	output := CorporationSheet{}
	args := url.Values{}
	args.Set("corporationID", strconv.FormatInt(corporationID, 10))
	err := api.Call(CorpCorporationSheetURL, args, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

type Contact struct {
	ID       int64  `xml:"contactID,attr"`
	Name     string `xml:"contactName,attr"`
	Standing int64  `xml:"standing,attr"`
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
		ID      int64   `xml:"accountID,attr"`
		Key     int64   `xml:"accountKey,attr"`
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
	Starbases []*Starbase `xml:"result>rowset>row"`
}

type Starbase struct {
	ID              int64   `xml:"itemID,attr"`
	TypeID          int64   `xml:"typeID,attr"`
	LocationID      int64   `xml:"locationID,attr"`
	MoonID          int64   `xml:"moonID,attr"`
	State           int64   `xml:"state,attr"`
	StateTimestamp  eveTime `xml:"stateTimestamp,attr"`
	OnlineTimestamp eveTime `xml:"onlineTimestamp,attr"`
	StandingOwnerID int64   `xml:"standingOwnerID,attr"`
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
	State           int64   `xml:"result>state"`
	StateTimestamp  eveTime `xml:"result>stateTimestamp"`
	OnlineTimestamp eveTime `xml:"result>onlineTimestamp"`
	GeneralSettings struct {
		UsageFlags              int64 `xml:"usageFlags"`
		DeployFlags             int64 `xml:"deployFlags"`
		AllowCorporationMembers bool  `xml:"allowCorporationMembers"`
		AllowAllianceMembers    bool  `xml:"allowAllianceMembers"`
	} `xml:"result>generalSettings"`
	CombatSettings struct {
		UseStandingsFrom struct {
			OwnerID int64 `xml:"ownerID,attr"`
		} `xml:"useStandingsFrom"`
		OnStandingDrop struct {
			Standing int64 `xml:"standing,attr"`
		} `xml:"onStandingDrop"`
		OnStatusDrop struct {
			Enabled  bool  `xml:"enabled,attr"`
			Standing int64 `xml:"standing,attr"`
		} `xml:"onStatusDrop"`
		OnAgression struct {
			Enabled bool `xml:"enabled,attr"`
		} `xml:"onAgression"`
		OnCorporationWar struct {
			Enabled bool `xml:"enabled, attr"`
		} `xml:"onCorporationWar"`
	} `xml:"result>combatSettings"`
	Fuel []struct {
		TypeID   int64 `xml:"typeID,attr"`
		Quantity int64 `xml:"quantity,attr"`
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
