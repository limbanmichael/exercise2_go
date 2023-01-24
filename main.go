package main

import (
	"bufio"
	"database/sql"
	"fmt"

	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// fmt.Println("Connecting Database....")

	// db, err := sql.Open("mysql", "root:Password123456@/go_exercise2@tcp(127.0.0.1:3306)")
	// db, err := sql.Open("mysql", "root:Password123456@/user@tcp(127.0.0.1:3306)")
	db, err := sql.Open("mysql", `root:Password123456@/go_exercise2`)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	var (
		iduser   int
		name     string
		username string
		password string
	)
	rows, err := db.Query("select iduser, name, username, password from user where username = ?", "mik_bpi")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&iduser, &name, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(iduser, name, username, password)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	login()

	// defer db.Close()
}

func login() {
	var userExist string = "michael"
	var passwordExist string = "limban"
	var username, password string

	for {
		var isEmpty, notExist bool = false, false

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("INPUT USERNAME: ")
		username, _ = reader.ReadString('\n')

		fmt.Print("INPUT Password: ")
		password, _ = reader.ReadString('\n')

		if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
			fmt.Println("Username and Password should have a value")
			isEmpty = true
		} else {
			if strings.TrimSpace(username) != userExist || strings.TrimSpace(password) != passwordExist {
				fmt.Println("NOT EXIST ", password)
				notExist = true
			}
		}

		if !isEmpty && !notExist {
			break
		}
	}
	fmt.Println("Username Value Accepted", username)
	fmt.Println("Password Value Accepted", password)
	showMenu()
}

func showMenu() {
	for {
		var userSelect string
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("***************Dashboard***************")
		fmt.Println("\n\nPress a for Bank Transfer")
		fmt.Println("Press b for Bank Statement")
		fmt.Println("Press c to Display User Information")
		fmt.Print("Press d to Update User Password")

		fmt.Print("\n\n***************************************\n")

		fmt.Print("Select: ")
		userSelect, _ = reader.ReadString('\n')

		fmt.Print("selected: ", userSelect)
		var selected = strings.TrimSpace(userSelect)

		if selected == "a" ||
			selected == "b" ||
			selected == "c" ||
			selected == "d" {
			fmt.Print("You Selected ", userSelect)
			break
		}

	}
}

func transferMoney() {

}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// )

// type AccountDetails struct {
// 	Amount       float64 `json:"Amount"`
// 	BankCode     string  `json:"BankCode"`
// 	CurrencyCode string  `json:"CurrencyCode"`
// 	AccountNo    string  `json:"AccountNo"`
// }

// func createBAS(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var post AccountDetails
// 	json.Unmarshal(reqBody, &post)

// 	json.NewEncoder(w).Encode(post)

// 	newData, err := json.Marshal(post)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(newData))
// 	}

// 	var finalTxt string = fmt.Sprint("Amount : ", post.Amount, "\nAccount Number : ", post.AccountNo, "\nBank Code : ", post.BankCode, "\nCurrency Code : ", post.CurrencyCode)
// 	fmt.Println(finalTxt)

// 	f, err := os.Create("BAS.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	_, err2 := f.WriteString(string(finalTxt))

// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}

// 	fmt.Println("file BAS.txt has been saved!")
// }

// func handleReqs() {
// 	r := mux.NewRouter().StrictSlash(true)
// 	r.HandleFunc("/create", createBAS).Methods("POST")

// 	log.Fatal(http.ListenAndServe(":10000", r))
// }

// func main() {
// 	handleReqs()
// }
