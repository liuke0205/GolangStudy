package main

type Gun struct {
	name  string
	power int
}

func (gun *Gun) setName(name string) {
	gun.name = name
}

func (gun *Gun) getName() string {
	return gun.name
}

func (gun *Gun) setPower(power int) {
	gun.power = power
}

func (gun *Gun) getPower() int {
	return gun.power
}
