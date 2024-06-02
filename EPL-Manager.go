// TUBES KELOMPOK 5
// EPL Manager
// Satria Febry Andanu   103012300372
// Muhammad Nabiel Salim 103012300022

package main

import "fmt"

const NMAX = 20
const PekanMAX = 38
const jumTandMAX = 10

type club struct {
	nama               string
	jumlahPertandingan int
	jumlahMenang       int
	jumlahKalah        int
	jumlahDraw         int
	jumlahGol          int
	jumlahKebobolan    int
	jumlahSelisihGol   int
	jumlahPoin         int
}
type pekanPertandingan struct {
	sudahTanding                 bool
	home, away, golHome, golAway int
}
type Pekan [PekanMAX][jumTandMAX]pekanPertandingan
type klub [NMAX]club

func main() {
	var tempInput, nKlub int
	var Klub klub
	var pekan Pekan

	//Data Klub EPL
	namaKlub := [18]string{"ARS", "AVL", "BOU", "BRE", "BHA", "BUR", "CHE", "CRY",
		"EVE", "FUL", "LIV", "MCI", "MUN", "NEW", "NFO", "SHU",
		"TOT", "WHU", /*"WOL", "LEE"*/}
	//Memasukkan Data Klub EPL ke array Klub
	for i, nama := range namaKlub {
		Klub[i].nama = nama
		nKlub++
	}
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4/5/6)? ")
		fmt.Scan(&tempInput)
		switch tempInput {
		case 1:
			fmt.Println("1. Klub baru")
			fmt.Println("2. Ubah klub")
			fmt.Println("3. Hapus klub")
			fmt.Println("4. Back")
			fmt.Print("Pilih (1/2/3/4)? ")
			fmt.Scan(&tempInput)
			switch tempInput {
			case 1:
				tambahkanKlub(&Klub, &nKlub, pekan)
			case 2:
				ubahDataKlub(&Klub, nKlub)
			case 3:
				hapusKlub(&Klub, &nKlub)
			case 4:
				break
			default:
				fmt.Println("Angka yang anda masukkan salah")
			}
		case 2:
			inputHasilPertandingan(&Klub, nKlub, &pekan)
		case 3:
			tampilkanJadwal(Klub, nKlub, &pekan)
		case 4:
			riwayatPertandingan(Klub, nKlub, pekan)
		case 5:
			tampilkanKlasemen(Klub, nKlub)
		case 6:
			return
		default:
			fmt.Println("Angka yang anda masukkan salah")
		}
	}

}
func menu() {
	/*I.S -
	  F.S Mencetak menu*/
	fmt.Println("---------------------------------")
	fmt.Println("           EPL Manager           ")
	fmt.Println("---------------------------------")
	fmt.Println("1. Klub")
	fmt.Println("2. Input data pertandingan")
	fmt.Println("3. Tampilkan Jadwal")
	fmt.Println("4. Riwayat pertandingan")
	fmt.Println("5. Tampilkan Klasemen")
	fmt.Println("6. Exit")
	fmt.Println("---------------------------------")

}
func tambahkanKlub(K *klub, n *int, P Pekan) {
	/*I.S Array P dan K terdefinisi, nilai n terdifinisi. prosedur akan dijalankan
		  dengan syarat tidak ada klub yang sudah bertanding
	  F.S Mengembalikan string inputan user ke array K yang baru dan menambah jumlah klub (n)*/
	var tempNama string
	var buat bool = true
	adaJadwal := false
	for i := 0; i < (*n-1)*2; i++ {
		for j := 0; j < *n/2; j++ {
			if P[i][j].sudahTanding {
				adaJadwal = true
			}
		}
	}
	if adaJadwal {
		fmt.Println("Sudah ada klub yang tanding")
		fmt.Println("Jika ingin menambahkan klub baru")
		fmt.Println("Hapus semua data pertandingan")
	} else {
		fmt.Print("Masukkan Nama Club (3 Huruf besar): ")
		fmt.Scan(&tempNama)
		if tempNama >= "AAA" && tempNama <= "ZZZ" {
			if *n < NMAX {
				/*Mengecek apakah ada nama klub yang sama dengan inputan user
				dengan menggunakan metode sequential search*/
				for i := 0; i < *n; i++ {
					if tempNama == K[i].nama {
						buat = false
					}
				}

				if buat {
					K[*n].nama = tempNama
					*n++
					fmt.Println("Club berhasil dibuat")
				} else {
					fmt.Println("Nama klub sudah ada")
				}
			} else {
				/*Ketika jumlah klub sudah melebihi kapasitas yaitu NMAX, user tidak
				bisa menambahkan klub baru*/
				fmt.Println("Club sudah mencapai batas maksimum")
			}

		} else if tempNama < "AAA" || tempNama > "ZZZ" {
			/*Nama klub yang dimasukan harus 3 huruf besar
			jika user memasukan angka atau huruf kecil atau simbol
			maka akan kembali ke menu awal*/
			fmt.Println("Masukkan harus 3 huruf besar")
		}
	}
}
func ubahDataKlub(K *klub, n int) {
	/*I.S Array K terdefinisi nilai n terdefinisi. user harus memasuki nama klub yang ada
	  F.S Merubah data array klub yang dipilih oleh user*/
	var tempNama string
	var idx int = -1
	if n < 1 {
		fmt.Println("Belum ada klub")
	} else {
		fmt.Print("Masukkan nama Club yang ingin diubah (3 Huruf besar): ")
		fmt.Scan(&tempNama)
		if tempNama < "AAA" || tempNama > "ZZZ" {
			/*Nama klub yang dimasukan harus 3 huruf besar
			jika user memasukan angka atau huruf kecil atau
			simbol maka akan kembali ke menu awal*/
			fmt.Println("Masukkan harus 3 huruf besar")
		} else {
			//Menggunakan algoritma Sequential Search untuk mencari index klub
			for i := 0; i < n; i++ {
				if K[i].nama == tempNama {
					idx = i
				}
			}

			if idx != -1 {
				fmt.Print("Jumlah Pertandingan: ")
				fmt.Scan(&K[idx].jumlahPertandingan)
				fmt.Print("Jumlah Menang: ")
				fmt.Scan(&K[idx].jumlahMenang)
				fmt.Print("Jumlah Kalah: ")
				fmt.Scan(&K[idx].jumlahKalah)
				fmt.Print("Jumlah Draw: ")
				fmt.Scan(&K[idx].jumlahDraw)
				fmt.Print("Jumlah Gol: ")
				fmt.Scan(&K[idx].jumlahGol)
				fmt.Print("Jumlah Kebobolan: ")
				fmt.Scan(&K[idx].jumlahKebobolan)
				fmt.Print("Jumlah Selisih Gol: ")
				fmt.Scan(&K[idx].jumlahSelisihGol)
				fmt.Print("Jumlah Point: ")
				fmt.Scan(&K[idx].jumlahPoin)
				fmt.Println("Data Klub sudah diubah")
			} else {
				fmt.Println("Klub tidak ditemukan")
			}
		}
	}
}
func hapusKlub(K *klub, n *int) {
	/*I.S Array K, dan nilai n terdefinisi
	  F.S Menghapus data klub sesuai dengan nama inputan klub yang
	  	  dimasukkan oleh user*/
	var idx int = -1
	var x string
	if *n < 1 {
		fmt.Println("Belum ada klub")
	} else {
		fmt.Print("Masukkan nama club yang ingin dihapus (3 Huruf besar): ")
		fmt.Scan(&x)
		if x < "AAA" || x > "ZZZ" {
			fmt.Println("Masukkan harus 3 huruf besar")
		} else {
			for i := 0; i < *n; i++ {
				if K[i].nama == x {
					idx = i
				}
			}

			if idx != -1 {
				*n--
				for i := idx; i < *n; i++ {
					K[i] = K[i+1]
				}
				fmt.Println("Klub sudah dihapus")
			} else {
				fmt.Println("Klub tidak ditemukan")
			}
		}
	}
}
func tampilkanJadwal(K klub, n int, P *Pekan) {
	/*I.S Array K, dan jumlah klub (n) haruslah genap dan minimal terdapat 2 klub
	  F.S Menampilkan jadwal dengan tabel berger menggunakan sistem double round robin
		  menggunakan prosedur buatJadwal dengan syarat jumlah klub yang ada harus genap*/

	if n == 0 {
		fmt.Println("Belum ada klub yang dibuat")
	} else if n%2 == 1 {
		fmt.Println("Jumlah klub harus genap, jumlah klub sekarang adalah", n)
	} else {
		buatJadwal(&*P, n)
		totalPekan := n - 1
		jumlahPertandingan := n / 2

		for i := 0; i < totalPekan*2; i++ {
			fmt.Println("Pekan", i+1)
			for j := 0; j < jumlahPertandingan; j++ {
				fmt.Printf("%d. %s vs %s\n", j+1, K[P[i][j].home].nama, K[P[i][j].away].nama)
			}
			fmt.Println()
		}
	}
}
func riwayatPertandingan(K klub, n int, P Pekan) {
	/*I.S Array K, dan nilai n terdefinisi. Harus ada klub yang sudah tanding
		  jika tidak maka prosedur akan mengeluarkan output "Belum ada pertandingan"
	  F.S Menampilkan jumlah gol sebuah klub di setiap pertandingan dan nama klub yang sudah tanding*/
	var cetak bool
	for i := 0; i < n*2; i++ {
		for j := 0; j < n/2; j++ {
			if P[i][j].sudahTanding {
				cetak = true
			}
		}
	}
	if cetak {
		for i := 0; i < n*2; i++ {
			var j int
			if P[i][j].sudahTanding {
				fmt.Printf("Pekan %d\n", i+1)
			}
			for j = 0; j < n/2; j++ {
				if P[i][j].sudahTanding {
					fmt.Printf("%s vs %s (%d-%d)\n", K[P[i][j].home].nama, K[P[i][j].away].nama, P[i][j].golHome, P[i][j].golAway)
				}
			}
		}
	} else {
		fmt.Println("Belum ada pertandingan")
	}
}
func hapusDataPertandingan(K *klub, i, inputMatch int, P *Pekan) {
	/*I.S Array K, dan nilai n terdefinisi. Nama klub yang di inputkan oleh user
		  terdapat dalam data array K
	  F.S Menghapus data pertandingan sesuai dengan pekan dan urutan pertandingan yang
	  	  dipilih oleh user*/
	K[P[i][inputMatch].home].jumlahGol -= P[i][inputMatch].golHome
	K[P[i][inputMatch].away].jumlahGol -= P[i][inputMatch].golAway
	K[P[i][inputMatch].home].jumlahKebobolan -= P[i][inputMatch].golAway
	K[P[i][inputMatch].away].jumlahKebobolan -= P[i][inputMatch].golHome
	if P[i][inputMatch].golHome > P[i][inputMatch].golAway {
		K[P[i][inputMatch].home].jumlahMenang -= 1
		K[P[i][inputMatch].home].jumlahPoin -= 3
		K[P[i][inputMatch].away].jumlahKalah -= 1
	} else if P[i][inputMatch].golHome < P[i][inputMatch].golAway {
		K[P[i][inputMatch].away].jumlahMenang -= 1
		K[P[i][inputMatch].away].jumlahPoin -= 3
		K[P[i][inputMatch].home].jumlahKalah -= 1
	} else {
		K[P[i][inputMatch].home].jumlahDraw -= 1
		K[P[i][inputMatch].away].jumlahDraw -= 1
		K[P[i][inputMatch].home].jumlahPoin -= 1
		K[P[i][inputMatch].away].jumlahPoin -= 1
	}
	K[P[i][inputMatch].home].jumlahSelisihGol = K[P[i][inputMatch].home].jumlahGol - K[P[i][inputMatch].home].jumlahKebobolan
	K[P[i][inputMatch].away].jumlahSelisihGol = K[P[i][inputMatch].away].jumlahGol - K[P[i][inputMatch].away].jumlahKebobolan
	K[P[i][inputMatch].home].jumlahPertandingan--
	K[P[i][inputMatch].away].jumlahPertandingan--
	P[i][inputMatch].sudahTanding = false
	P[i][inputMatch].golHome = 0
	P[i][inputMatch].golAway = 0
}
func ubahDataPertandingan(K *klub, i, inputMatch int, P *Pekan) {
	/*I.S Array K, dan nilai n terdefinisi. Nama klub yang di inputkan oleh user
		  terdapat dalam data array K
	  F.S Mengubah data pertandingan sesuai dengan pekan dan urutan pertandingan yang
	  	  dipilih oleh user*/
	hapusDataPertandingan(&*K, i, inputMatch, &*P)
	inputDataPertandingan(&*K, i, inputMatch, &*P)
}
func inputDataPertandingan(K *klub, i, inputMatch int, P *Pekan) {
	/*I.S Array K, dan nilai n terdefinisi. Nama klub yang di inputkan oleh user
		  terdapat dalam data array K
	  F.S Menginput data pertandingan sesuai dengan pekan dan urutan pertandingan yang
	  	  dipilih oleh user*/
	var golA, golH int
	fmt.Printf("Masukkan jumlah gol tuan rumah (%s): ", K[P[i][inputMatch].home].nama)
	fmt.Scan(&golH)
	fmt.Printf("Masukkan jumlah gol tamu (%s): ", K[P[i][inputMatch].away].nama)
	fmt.Scan(&golA)
	P[i][inputMatch].golHome = golH
	P[i][inputMatch].golAway = golA
	K[P[i][inputMatch].home].jumlahGol += golH
	K[P[i][inputMatch].away].jumlahGol += golA
	K[P[i][inputMatch].home].jumlahKebobolan += golA
	K[P[i][inputMatch].away].jumlahKebobolan += golH
	if golH > golA {
		K[P[i][inputMatch].home].jumlahMenang += 1
		K[P[i][inputMatch].home].jumlahPoin += 3
		K[P[i][inputMatch].away].jumlahKalah += 1
	} else if golH < golA {
		K[P[i][inputMatch].away].jumlahMenang += 1
		K[P[i][inputMatch].away].jumlahPoin += 3
		K[P[i][inputMatch].home].jumlahKalah += 1
	} else {
		K[P[i][inputMatch].home].jumlahDraw += 1
		K[P[i][inputMatch].away].jumlahDraw += 1
		K[P[i][inputMatch].home].jumlahPoin += 1
		K[P[i][inputMatch].away].jumlahPoin += 1
	}
	K[P[i][inputMatch].home].jumlahSelisihGol = K[P[i][inputMatch].home].jumlahGol - K[P[i][inputMatch].home].jumlahKebobolan
	K[P[i][inputMatch].away].jumlahSelisihGol = K[P[i][inputMatch].away].jumlahGol - K[P[i][inputMatch].away].jumlahKebobolan
	K[P[i][inputMatch].home].jumlahPertandingan++
	K[P[i][inputMatch].away].jumlahPertandingan++
	P[i][inputMatch].sudahTanding = true
}
func inputHasilPertandingan(K *klub, n int, P *Pekan) {
	/*I.S Array K, dan nilai n terdefinisi. Jumlah klub (n) haruslah genap
	  F.S Data array K dan P terisi. sesuai dengan pertandingan yang dipilih*/
	var x int
	if n%2 == 1 && n > 0 {
		fmt.Println("Jumlah klub harus genap, jumlah klub sekarang adalah", n)
	} else {
		buatJadwal(&*P, n)
		fmt.Print("Pekan ke: ")
		fmt.Scan(&x)
		x--
		if x+1 < 1 || x+1 > (n-1)*2 {
			/*Jika user memasukkan nilai lebih besar dari pekan yang ada (Total pekan adalah (n-1)*2)
			algoritma akan mencetak "Pekan tidak valid"*/
			fmt.Printf("Pekan tidak ada, jumlah pekan saat ini %d\n", (n-1)*2)
		} else {
			for i := 0; i < n*2; i++ {
				var inputMatch, inputMatch2, golH, golA int
				if i == x {
					for j := 0; j < n/2; j++ {
						fmt.Printf("%d. %s vs %s\n", j+1, K[P[i][j].home].nama, K[P[i][j].away].nama)
					}
					//User bisa memilih pertandingan yang ingin dimasukkan datanya
					fmt.Print("Pilih pertandingan: ")
					fmt.Scan(&inputMatch)
					inputMatch--
					if inputMatch+1 < 1 || inputMatch+1 > n/2 {
						/*Jika input yang dimasukkan oleh user adalah 0 atau lebih dari jumlah
						pertandingan yang ada di setiap pekan, maka akan mengeluarkan outpu "Pertandingan tidak ada"*/
						fmt.Println("Pertandingan tidak ada")
					} else if P[i][inputMatch].sudahTanding {
						/*Jika data pertandingan yang dipilih sudah diinput maka user akan diberi
						pilihan untuk ubah data atau hapus data atau membiarkan data tersebut*/
						fmt.Println("Data sudah diinput")
						fmt.Println("1. Ubah data")
						fmt.Println("2. Hapus data")
						fmt.Println("3. Back")
						fmt.Print("Pilih (1/2/3)? ")
						fmt.Scan(&inputMatch2)
						switch inputMatch2 {
						case 1:
							ubahDataPertandingan(&*K, i, inputMatch, &*P)
						case 2:
							hapusDataPertandingan(&*K, i, inputMatch, &*P)
						case 3:
							return
						}

					} else if !P[i][inputMatch].sudahTanding {
						fmt.Printf("Masukkan jumlah gol tuan rumah (%s): ", K[P[i][inputMatch].home].nama)
						fmt.Scan(&golH)

						fmt.Printf("Masukkan jumlah gol tamu (%s): ", K[P[i][inputMatch].away].nama)
						fmt.Scan(&golA)
						P[i][inputMatch].golHome = golH
						P[i][inputMatch].golAway = golA
						K[P[i][inputMatch].home].jumlahGol += golH
						K[P[i][inputMatch].away].jumlahGol += golA
						K[P[i][inputMatch].home].jumlahKebobolan += golA
						K[P[i][inputMatch].away].jumlahKebobolan += golH
						if golH > golA {
							K[P[i][inputMatch].home].jumlahMenang += 1
							K[P[i][inputMatch].home].jumlahPoin += 3
							K[P[i][inputMatch].away].jumlahKalah += 1
						} else if golH < golA {
							K[P[i][inputMatch].away].jumlahMenang += 1
							K[P[i][inputMatch].away].jumlahPoin += 3
							K[P[i][inputMatch].home].jumlahKalah += 1
						} else {
							K[P[i][inputMatch].home].jumlahDraw += 1
							K[P[i][inputMatch].away].jumlahDraw += 1
							K[P[i][inputMatch].home].jumlahPoin += 1
							K[P[i][inputMatch].away].jumlahPoin += 1
						}
						K[P[i][inputMatch].home].jumlahSelisihGol = K[P[i][inputMatch].home].jumlahGol - K[P[i][inputMatch].home].jumlahKebobolan
						K[P[i][inputMatch].away].jumlahSelisihGol = K[P[i][inputMatch].away].jumlahGol - K[P[i][inputMatch].away].jumlahKebobolan
						K[P[i][inputMatch].home].jumlahPertandingan++
						K[P[i][inputMatch].away].jumlahPertandingan++
						P[i][inputMatch].sudahTanding = true
					}
				}
			}
		}
	}

}
func sortClub(K *klub, n int) {
	/*I.S Array K, dan nilai n terdefinisi.
	  F.S Mencetak data klub dengan urutan menurun (Descending). Dengan klub di urutan
	  	  pertama adalah klub yang memiliki jumlah poin dan jumlah selisih gol paling tinggi*/
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if K[i].jumlahPoin < K[j].jumlahPoin {
				K[i], K[j] = K[j], K[i]
			} else if K[i].jumlahPoin == K[j].jumlahPoin {
				if K[i].jumlahSelisihGol < K[j].jumlahSelisihGol {
					K[i], K[j] = K[j], K[i]
				} else if K[i].jumlahSelisihGol == K[j].jumlahSelisihGol {
					if K[i].jumlahGol < K[j].jumlahGol {
						K[i], K[j] = K[j], K[i]
					}
				}
			}
		}
	}
}
func tampilkanKlasemen(K klub, n int) {
	/*I.S Array K terdefinisi nilai n terdifinisi
	  F.s Mencetak klasemen menggunakan array K yang sudah diurut menggunakan prosedur sortClub*/
	sortClub(&K, n)
	fmt.Println("                        Klasemen Premier League                       ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-4s | %-4s | %-4s | %-4s | %-4s | %-4s | %-4s | %-4s | %-4s |\n", "No.", "Nama", "P", "W", "D", "L", "GF", "GA", "GD", "Pts")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-3d | %-4s | %-4d | %-4d | %-4d | %-4d | %-4d | %-4d | %-4d | %-4d |\n", i+1, K[i].nama, K[i].jumlahPertandingan, K[i].jumlahMenang, K[i].jumlahDraw, K[i].jumlahKalah, K[i].jumlahGol, K[i].jumlahKebobolan, K[i].jumlahSelisihGol, K[i].jumlahPoin)
	}
	fmt.Println("----------------------------------------------------------------------")
}
func buatJadwal(P *Pekan, n int) {
	/*I.S Array P terdefinisi dan n terdefinisi
	  F.S Array P terisi dengan index klub home dan away*/
	var temp int
	for i := 0; i < (n-1)*2; i++ {
		for j := 0; j < n/2; j++ {
			P[i][j].home = (j + i) % (n - 1)
			if j == 0 {
				P[i][j].away = n - 1
			} else {
				P[i][j].away = (n - 1 - j + i) % (n - 1)
			}

			if i%2 != 0 {
				temp = P[i][j].home         //
				P[i][j].home = P[i][j].away //Untuk menukar home dan away agar seluruh tim bisa main di kandang atau tandang
				P[i][j].away = temp         //
			}
		}
	}
}
