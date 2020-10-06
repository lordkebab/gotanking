package gotanking

/*

// ListMaps queries the encyclopedia/arenas endpoint
func (c *WOTClient) ListMaps() (model.Arena, error) {
	endpoint := "/encyclopedia/arenas"
	arenas := model.Arena{}

	resp, err := http.Get(c.baseURL + endpoint)
	if err != nil {
		return arenas, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	b := body.Bytes()

	// unmarshall into the data model
	err = json.Unmarshal(b, &arenas)
	if err != nil {
		return model.Arena{}, err
	}

	return arenas, nil
}
*/
