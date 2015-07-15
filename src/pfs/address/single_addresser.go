package address

type singleAddresser struct {
	address string
}

func newSingleAddresser(address string) *singleAddresser {
	return &singleAddresser{address}
}

func (l *singleAddresser) GetMasterAddress(shard int) (string, error) {
	return l.address, nil
}

func (l *singleAddresser) GetSlaveAddresses(shard int) ([]string, error) {
	return []string{l.address}, nil
}

func (l *singleAddresser) GetServerAddress() (string, error) {
	return l.address, nil
}
