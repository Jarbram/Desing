package client_two

import "singleton/singleton"

func SetAge() {
	p := singleton.GetInstance()
	p.SetAge()
}
