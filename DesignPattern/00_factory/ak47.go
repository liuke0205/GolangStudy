package main

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "Ak47",
			power: 10,
		}}
}
