package models

type kategori []kategoriElement

type kategoriElement struct {
	MainKategori string        `json:"main_kategori"`
	SubKategori  []subKategori `json:"sub_kategori"`
}

type subKategori struct {
	Nama string `json:"nama"`
	ID   int64  `json:"id"`
}

type student struct {
	ID    string
	Name  string
	Grade int
}

var kat = []kategoriElement{
	kategoriElement{"Telekomunikasi", []subKategori{
		subKategori{"pulsa", 11},
		subKategori{"Pascabayar", 12},
	}},
	kategoriElement{"Tagihan", []subKategori{
		subKategori{"Listrik", 21},
		subKategori{"PDAM", 22},
		subKategori{"TV Kabel", 23},
		subKategori{"Internet", 24},
		subKategori{"Telepon", 25},
		subKategori{"Gas", 26},
		subKategori{"Rekening Virtual", 27},
		subKategori{"E-Commerce", 28},
	}},
	kategoriElement{"Transportasi", []subKategori{
		subKategori{"Kereta", 31},
		subKategori{"taksi", 32},
		subKategori{"Bis", 33},
	}},
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

func GetKat() kategori {
	return kat
}
