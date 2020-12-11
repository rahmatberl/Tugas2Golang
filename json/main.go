package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var db *sql.DB
var err error

type yamlconfig struct {
	Connection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		User     string `yaml:"user"`
		Database string `yaml:"database"`
	}
}
type university struct {
	Idmahasiswa string `json:"id_mahasiswa"`
	Nama        string `json:"nama"`
	Alamat      struct {
		Jalan     string `json:"jalan"`
		Kelurahan string `json:"kelurahan"`
		Kecamatan string `json:"kecamatan"`
		Kabupaten string `json:"kabupaten"`
		Provinsi  string `json:"provinsi"`
	} `json:"alamat"`
	Fakultas string  `json:"fakultas"`
	Jurusan  string  `json:"jurusan"`
	Nilai    []nilai `json:"Nilai"`
}

type nilai struct {
	Idmahasiswa string  `json:"id_mahasiswa"`
	Idmatkul    string  `json:"id_matkul"`
	Mkuliah     string  `json:"m_kuliah"`
	Nilai       float32 `json:"nilai"`
	Semester    int8    `json:"semester"`
}

func getUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var university university
	var nilaidet nilai
	params := mux.Vars(r)

	sql := `select
				university.id_mahasiswa,
				university.nama,
				fakultas.nama as fakultas,
				jurusan.nama as jurusan,
				university.jalan,
				university.kelurahan,
				university.kecamatan,
				university.kabupaten,
				university.provinsi 
				FROM
				university.university
				INNER JOIN university.fakultas
				ON (university.Id_Fakultas = fakultas.id_fakultas)
				INNER JOIN university.jurusan
				ON (university.Id_Jurusan = jurusan.id_jurusan) where university.id_mahasiswa=?`
	result, err := db.Query(sql, params["id"])

	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&university.Idmahasiswa, &university.Nama, &university.Fakultas, &university.Jurusan,
			&university.Alamat.Jalan, &university.Alamat.Kelurahan, &university.Alamat.Kecamatan, &university.Alamat.Kabupaten, &university.Alamat.Provinsi)

		Idmahasiswa := &university.Idmahasiswa

		if err != nil {
			panic(err.Error())
		}

		sqlnilai := `SELECT
						mata_kuliah.nama,nilai.nilai,nilai.semester
						FROM
							mahasiswa.nilai
							INNER JOIN mahasiswa.mata_kuliah
								ON (nilai.Id_matkul = matkul.id_matkul) where nilai.id_mahasiswa=?;`

		resultnilai, errnilai := db.Query(sqlnilai, *Idmahasiswa)

		defer resultnilai.Close()

		if errnilai != nil {
			panic(err.Error())
		}

		for resultnilai.Next() {
			err := resultnilai.Scan(&nilaidet.Mkuliah, &nilaidet.Nilai, &nilaidet.Semester)

			if err != nil {
				panic(err.Error())
			}

			university.Nilai = append(university.Nilai, nilaidet)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(university)
}
func getNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var univP []university

	params := mux.Vars(r)

	sql := `SELECT
				id_mahasiswa,
				IFNULL(nama,'') nama,
				IFNULL(jalan,'') jalan,
				IFNULL(kelurahan,'') kelurahan,
				IFNULL(kecamatan,'') kecamatan,
				IFNULL(kabupaten,'') kabupaten,
				IFNULL(provinsi,'') provinsi,
				IFNULL(fakultas,'') fakultas,
				IFNULL(jurusan,'') jurusan				
			FROM mahasiswa WHERE id_mahasiswa IN (?)`

	result, err := db.Query(sql, params["id"])

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var univ university
		err := result.Scan(&univ.Idmahasiswa, &univ.Nama, &univ.Alamat.Jalan, &univ.Alamat.Kelurahan, &univ.Alamat.Kecamatan, &univ.Alamat.Kabupaten, &univ.Alamat.Provinsi, &univ.Fakultas, &univ.Jurusan)

		if err != nil {
			panic(err.Error())
		}

		sqlNilai := `SELECT
						id_mahasiswa		
						, mata_kuliah.id_matkul
						, mata_kuliah.m_kuliah
						, nilai
						, semester
					FROM
						nilai INNER JOIN mata_kuliah
							ON (nilai.id_matkul = mata_kuliah.id_matkul)
					WHERE nilai.id_mahasiswa = ?`

		resultNilai, errNilai := db.Query(sqlNilai, univ.Idmahasiswa)

		defer resultNilai.Close()

		if errNilai != nil {
			panic(err.Error())
		}

		for resultNilai.Next() {
			var nilaiP nilai
			err := resultNilai.Scan(&nilaiP.Idmahasiswa, &nilaiP.Idmatkul, &nilaiP.Mkuliah, &nilaiP.Nilai, &nilaiP.Semester)
			if err != nil {
				panic(err.Error())
			}
			univ.Nilai = append(univ.Nilai, nilaiP)
		}
		univP = append(univP, univ)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(univP)
}

func getNilaiAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var univG []university

	sql := `SELECT
				id_mahasiswa,
				IFNULL(nama,'') nama,
				IFNULL(jalan,'') jalan,
				IFNULL(kelurahan,'') kelurahan,
				IFNULL(kecamatan,'') kecamatan,
				IFNULL(kabupaten,'') kabupaten,
				IFNULL(provinsi,'') provinsi,
				IFNULL(fakultas,'') fakultas,
				IFNULL(jurusan,'') jurusan				
			FROM mahasiswa`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var univ2 university
		err := result.Scan(&univ2.Idmahasiswa, &univ2.Nama, &univ2.Alamat.Jalan, &univ2.Alamat.Kelurahan, &univ2.Alamat.Kecamatan, &univ2.Alamat.Kabupaten, &univ2.Alamat.Provinsi, &univ2.Fakultas, &univ2.Jurusan)

		if err != nil {
			panic(err.Error())
		}

		sqlNilai := `SELECT
						id_mahasiswa		
						, mata_kuliah.id_matkul
						, mata_kuliah.m_kuliah
						, nilai
						, semester
					FROM
						nilai INNER JOIN mata_kuliah
							ON (nilai.id_matkul = mata_kuliah.id_matkul)
					WHERE nilai.id_mahasiswa = ?`

		resultNilai, errNilai := db.Query(sqlNilai, univ2.Idmahasiswa)

		defer resultNilai.Close()

		if errNilai != nil {
			panic(err.Error())
		}

		for resultNilai.Next() {
			var nilaiG nilai
			err := resultNilai.Scan(&nilaiG.Idmahasiswa, &nilaiG.Idmatkul, &nilaiG.Mkuliah, &nilaiG.Nilai, &nilaiG.Semester)
			if err != nil {
				panic(err.Error())
			}
			univ2.Nilai = append(univ2.Nilai, nilaiG)
		}
		univG = append(univG, univ2)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(univG)
}
func updateUniversity(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newNama := r.FormValue("nama")
		newJalan := r.FormValue("jalan")
		newKelurahan := r.FormValue("kelurahan")
		newKecamatan := r.FormValue("kecamatan")
		newKabupaten := r.FormValue("kabupaten")
		newProvinsi := r.FormValue("provinsi")
		newFakultas := r.FormValue("fakultas")
		newJurusan := r.FormValue("jurusan")

		stmt, err := db.Prepare("UPDATE mahasiswa SET nama = ?, jalan = ?, kelurahan = ?, kecamatan = ?, kabupaten = ?, provinsi = ?, fakultas = ?, jurusan = ? WHERE id_mahasiswa = ?")

		_, err = stmt.Exec(newNama, newJalan, newKelurahan, newKecamatan, newKabupaten, newProvinsi, newFakultas, newJurusan, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Mahasiswa with id_mahasiswa = %s was updated", params["id"])
	}
}
func createUniversity(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		Idmahasiswa := r.FormValue("id_mahasiswa")
		Nama := r.FormValue("nama")
		Jalan := r.FormValue("jalan")
		Kelurahan := r.FormValue("kelurahan")
		Kecamatan := r.FormValue("kecamatan")
		Kabupaten := r.FormValue("kabupaten")
		Provinsi := r.FormValue("provinsi")
		Fakultas := r.FormValue("fakultas")
		Jurusan := r.FormValue("jurusan")

		stmt, err := db.Prepare("INSERT INTO mahasiswa (id_mahasiswa, nama, jalan, kelurahan, kecamatan, kabupaten, provinsi, fakultas, jurusan) VALUES (?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(Idmahasiswa, Nama, Jalan, Kelurahan, Kecamatan, Kabupaten, Provinsi, Fakultas, Jurusan)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}
func main() {
	yamlFile, err := ioutil.ReadFile("../Yaml/config.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	var yamlConfig yamlconfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	host := yamlConfig.Connection.Host
	port := yamlConfig.Connection.Port
	user := yamlConfig.Connection.User
	pass := yamlConfig.Connection.Password
	data := yamlConfig.Connection.Database

	var (
		mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)
	)

	db, err = sql.Open("mysql", mySQL)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/mahasiswa/{id}", getNilai).Methods("GET")
	r.HandleFunc("/mahasiswaG", getNilaiAll).Methods("GET")
	r.HandleFunc("/mahasiswa/{id}", updateUniversity).Methods("PUT")
	r.HandleFunc("/mahasiswaT", createUniversity).Methods("POST")
	r.HandleFunc("/mahasiswa", getUniversity).Methods("GET")

	fmt.Println("Server on :8181")
	log.Fatal(http.ListenAndServe(":8181", r))
}
