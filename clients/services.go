package clients

// Create a new client is the clientId doesn't exists already.
func AddClient (clientId string, clientSecret string) error  {

	clientExists , err := findFromClientId(clientId)
	if err != nil {
		return nil
	}
	if len(clientExists) < 1 {
		aClient := client{ClientId:clientId, Secret:clientSecret}

		err = createClient(aClient)
		if err != nil {
			return err
		}
	}
	return nil

}
