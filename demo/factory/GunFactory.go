package factory

func CreateGun(kind string) IGun {
	if kind == "ak" {
		return AK{Gun{"ak", "30"}}
	} else if kind == "gatlin" {
		return Gatlin{Gun{"gatlin", "100"}}
	}
	return nil
}
