package blame

// getBlamePartyIDsInList returns the nodes partyID.id who are in the partiesID list
func (m *Manager) getBlamePartyIDsInList(partiesID []string) []string {
	var partiesInList []string
	for partyID := range m.partyInfo.PartyIDMap {
		for _, el := range partiesID {
			if el == partyID {
				partiesInList = append(partiesInList, partyID)
			}
		}
	}
	return partiesInList
}

// getBlamePartyIDsNotInList returns the nodes partyID.id who are not in the partiesID list
func (m *Manager) getBlamePartyIDsNotInList(partysID []string) []string {
	var partiesNotInList []string
	for partyID := range m.partyInfo.PartyIDMap {
		if m.localPartyID == partyID {
			continue
		}
		found := false
		for _, el := range partysID {
			if el == partyID {
				found = true
				break
			}
		}
		if !found {
			partiesNotInList = append(partiesNotInList, partyID)
		}
	}
	return partiesNotInList
}

// GetBlamePartyIDsLists returns the nodes partyID.id who are in and not in the partiesID list
func (m *Manager) GetBlamePartyIDsLists(partiesID []string) ([]string, []string, error) {
	inList := m.getBlamePartyIDsInList(partiesID)

	notInlist := m.getBlamePartyIDsNotInList(partiesID)

	return inList, notInlist, nil
}
