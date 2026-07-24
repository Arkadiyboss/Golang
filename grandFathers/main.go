package main

import (
	"fmt"
	"grandFather/auth"
	"grandFather/database"
	findInfo "grandFather/findInfo/findExactInfo"
	newEntity "grandFather/newEntity"
	"net/http"
	"time"
)

type FindRequestByName struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

type FindRequestSlots struct {
	DateIn time.Time `json:"dateIn"`
}

type NewGrandFather struct {
	NameSecondName NewNameSecondName `json:"NameSecondName"`
	Age            int               `json:"age"`
}

type NewGuardian struct {
	NameSecondName NewNameSecondName `json:"NameSecondName"`
	PhoneNumber    int               `json:"phoneNumber"`
}

type NewEmployee struct {
	NameSecondName NewNameSecondName `json:"NameSecondName"`
	PhoneNumber    int               `json:"phoneNumber"`
	Post           string            `json:"post"`
}

type NewSlots struct {
	Room           int       `json:"room"`
	Cost           int       `json:"cost"`
	DateIn         time.Time `json:"dateIn"`
	DateOut        time.Time `json:"dateOut"`
	GrandFatherId  int       `json:"grandFatherId"`
	NursingHouseId int       `json:"nursingHouseId"`
}

type NewNameSecondName struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

type NewUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func DeleteInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'")
}

func main() {

	pool, err := database.ConnectDB("postgres://arkadiy:123@localhost:5432/grandFatherBase?sslmode=disable")

	if err != nil {
		fmt.Println("Не получается подключиться к БД", err)
		return
	}

	defer pool.Close()

	go database.PingDB(pool)

	http.HandleFunc("/newGrandFather", func(w http.ResponseWriter, r *http.Request) {
		newEntity.NewEntity(w, r, pool, 1)
	})
	http.HandleFunc("/newGuardian", func(w http.ResponseWriter, r *http.Request) {
		newEntity.NewEntity(w, r, pool, 2)
	})
	http.HandleFunc("/newEmployee", func(w http.ResponseWriter, r *http.Request) {
		newEntity.NewEntity(w, r, pool, 3)
	})
	http.HandleFunc("/newSlots", func(w http.ResponseWriter, r *http.Request) {
		newEntity.NewEntity(w, r, pool, 4)
	})
	http.HandleFunc("/newUser", func(w http.ResponseWriter, r *http.Request) {
		auth.NewUser(w, r, pool)
	})
	http.HandleFunc("/newUpdateToken", func(w http.ResponseWriter, r *http.Request) {
		auth.NewUpdatedToken(w, r, pool)
	})
	

	http.HandleFunc("/findGrandFather", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindInfo(w, r, pool, 1)
	})
	http.HandleFunc("/findGuardianInfo", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindInfo(w, r, pool, 2)
	})
	http.HandleFunc("/findEmployee", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindInfo(w, r, pool, 3)
	})
	http.HandleFunc("/findSlots", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindInfo(w, r, pool, 4)
	})

	http.HandleFunc("/allGrandFathers", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindAllInfo(w, r, pool, 1)
	})
	http.HandleFunc("/allGuardianInfo", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindAllInfo(w, r, pool, 2)
	})
	http.HandleFunc("/allEmployee", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindAllInfo(w, r, pool, 3)
	})
	http.HandleFunc("/allSlots", func(w http.ResponseWriter, r *http.Request) {
		findInfo.FindAllInfo(w, r, pool, 4)
	})
	http.HandleFunc("/deleteInfo", DeleteInfo)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
