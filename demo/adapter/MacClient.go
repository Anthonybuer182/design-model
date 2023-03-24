package adapter

type MacClient struct {
}

func (m MacClient) InsertCharge(c Typec) {
	c.insert()
}
