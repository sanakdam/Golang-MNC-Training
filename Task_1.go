package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name       string
	address    string
	profession string
	reason     string
}

func getStudent(students []*Student, absence int) (Student, error) {
	maxStudent := len(students)

	if absence > maxStudent {
		return Student{}, fmt.Errorf("Jumlah Peserta MNC Golang hanya terdapat %d orang", maxStudent)
	}

	return *students[absence-1], nil
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Silahkan masukkan Nomor Absensi Peserta")
		return
	}

	absence, errAbsence := strconv.Atoi(args[1])
	if errAbsence != nil {
		fmt.Println("Nomor Absensi Peserta tidak valid")
		return
	}

	students := []*Student{
		{
			name:       "Teguh Afdilla",
			address:    "Alamat 1",
			profession: "Profesi 1",
			reason:     "Alasan 1",
		},
		{
			name:       "Fajar Agus Maulana",
			address:    "Alamat 1",
			profession: "Profesi 1",
			reason:     "Alasan 1",
		},
		{
			name:       "Kadek Bintang A",
			address:    "Alamat 1",
			profession: "Profesi 1",
			reason:     "Alasan 1",
		},
	}

	student, errorStudent := getStudent(students, absence)
	if errorStudent != nil {
		fmt.Println(errorStudent)
		return
	}

	fmt.Printf("Nama : %s\nAlamat : %s\nProfesi : %s\nAlasan : %s\n", student.name, student.address, student.profession, student.reason)
}
