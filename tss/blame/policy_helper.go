package blame

// getBlamePartyIDsInList returns the nodes partyID.id who are in the partiesID list
func (m *Manager) getBlamePartyIDsInList(partiesID []string) ([]string, error) {
	var partiesInList []string
	for partyID := range m.partyInfo.PartyIDMap {
		for _, el := range partiesID {
			if el == partyID {
				partiesInList = append(partiesInList, partyID)
			}
		}
	}
	return partiesInList, nil
}

// getBlamePartyIDsNotInList returns the nodes partyID.id who are not in the partiesID list
func (m *Manager) getBlamePartyIDsNotInList(partysID []string) ([]string, error) {
	var partiesNotInList []string
	for partyID, _ := range m.partyInfo.PartyIDMap {
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
	return partiesNotInList, nil
}

// GetBlamePartyIDsLists returns the nodes partyID.id who are in and not in the partiesID list
func (m *Manager) GetBlamePartyIDsLists(partiesID []string) ([]string, []string, error) {
	inList, err := m.getBlamePartyIDsInList(partiesID)
	if err != nil {
		return nil, nil, err
	}

	notInlist, err := m.getBlamePartyIDsNotInList(partiesID)
	if err != nil {
		return nil, nil, err
	}

	return inList, notInlist, err
}
