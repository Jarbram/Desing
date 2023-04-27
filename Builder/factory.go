package builder

func Factory(t int) MessageBuilder {
	switch t {
	case 1:
		return &JSONMessageBuilder{}
	case 2:
		return &XMLMessageBuilder{}
	default:
		return nil
	}
}
