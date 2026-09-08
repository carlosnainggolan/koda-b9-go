package model

type Biodata struct {
	Nama string
	Foto string
	Email string
	Umur uint8
	NomorTelepon string
	StatusPernikahan bool
	RiwayatPendidikan []Pendidikan
}

type Pendidikan struct {
	Nama string
	Jurusan string
}