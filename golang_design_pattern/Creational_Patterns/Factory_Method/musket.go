package main

type musket struct{
	gun
}

func newMusket() iGun{
	return &musket{
		gun: gun{
			name: "musket",
			power: 8,
		},
	}
}