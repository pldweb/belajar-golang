package main
import "fmt"

type Jamaah struct {
  nama string
  umur int
  kota string
  jenis_kelamin string
}

func main() {
//   var orang1 Jamaah
//   var orang2 Jamaah

//   orang1.nama = "Medi"
//   orang1.umur = 45
//   orang1.kota = "IRT"
//   orang1.jenis_kelamin = "Perempuan"

//   orang2.nama = "Somad"
//   orang2.umur = 24
//   orang2.kota = "Marketing"
//   orang2.jenis_kelamin = "Laki-laki"

//   fmt.Println("Nama Jamaah: ", orang1.nama)
//   fmt.Println("Umur: ", orang1.umur)
//   fmt.Println("Kota: ", orang1.kota)
//   fmt.Println("Gaji: ", orang1.jenis_kelamin)

//   fmt.Println()

//   fmt.Println("Nama: ", orang2.nama)
//   fmt.Println("Umur: ", orang2.umur)
//   fmt.Println("Kota: ", orang2.kota)
//   fmt.Println("Gaji: ", orang2.jenis_kelamin)4


//   var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
//   b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

//   fmt.Printf("a\t%v\n", a)
//   fmt.Printf("b\t%v\n", b)



  var dataJamaah = map[string]Jamaah{
    "1": {nama: "Medi", umur: 45, kota: "IRT", jenis_kelamin: "Perempuan"},
    "2": {nama: "Somad", umur: 24, kota: "Jakarta", jenis_kelamin: "Laki-laki"},
    "3": {nama: "Gme", umur: 24, kota: "Jakarta", jenis_kelamin: "Laki-laki"},

  }

  fmt.Println("Data Jamaah")
  fmt.Println("============")

  for key, value := range dataJamaah {
    fmt.Println("Data Jamaah:", key)
    jamahBayar(value)
    fmt.Println()
  }
   
}

func jamahBayar(orang Jamaah) {
  fmt.Println("Nama Jamaah: ", orang.nama)
  fmt.Println("Umur: ", orang.umur)
  fmt.Println("Kota: ", orang.kota)
  fmt.Println("Jenis Kelamin: ", orang.jenis_kelamin)
}