package eveapi

import "fmt"

const (
	// AccountCharactersURL includes all characters and API
	AccountCharactersURL = "/account/Characters.xml.aspx"
)

// Character contains information about a specific character
type Character struct {
	ID              string `xml:"characterID,attr"`
	Name            string `xml:"name,attr"`
	CorporationID   string `xml:"corporationID,attr"`
	CorporationName string `xml:"corporationName,attr"`
	AllianceID      string `xml:"allianceID,attr"`
	AllianceName    string `xml:"allianceName,attr"`
}

func (c Character) String() string {
	return fmt.Sprintf("%s (%s) | %s (%s)", c.Name, c.ID, c.CorporationName, c.CorporationID)
}

// CharacterList is the list of characters returned by the API
type CharacterList struct {
	APIResult
	Characters []Character `xml:"result>rowset>row"`
}

// AccountCharacters calls the API
// Returns a CharacterList struct
func (api API) AccountCharacters() ([]Character, error) {
	output := CharacterList{}
	err := api.Call(AccountCharactersURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return output.Characters, nil
}
