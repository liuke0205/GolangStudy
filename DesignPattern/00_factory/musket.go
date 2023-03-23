package main

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{Gun{
		name:  "musket",
		power: 3,
	}}
}
