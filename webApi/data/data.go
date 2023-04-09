package data

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/sijms/go-ora/v2"
)

type Dane struct {
	val1 string
}

type BodyData struct {
	AllTable    float64
	SingleTable float64
	TableChoice string
}

func DbGenerate(bodyData BodyData) {

	connString := "oracle://s101598:o8N2C5Q259@217.173.198.135:1521/tpdb"

	db, err := sql.Open("oracle", connString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	if bodyData.AllTable != 0 {
		//file create
		hour, min, sec := time.Now().Clock()
		fileName := fmt.Sprintf("inserted_AllTables_%d_%d_%d.txt", hour, min, sec)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		warsztatID := CheckLastID("SELECT id_warsztatu FROM Warsztat WHERE id_warsztatu = (SELECT MAX(id_warsztatu) FROM Warsztat)", db)
		autoZastepczeID := CheckLastID("SELECT id_auta_zastepczego FROM Warsztat WHERE id_auta_zastepczego = (SELECT MAX(id_auta_zastepczego) FROM auto_zastepcze)", db)
		histAutaZastepczegoID := CheckLastID("SELECT ID_HISTORI_AUTA_ZASTEPCZEGO FROM Warsztat WHERE ID_HISTORI_AUTA_ZASTEPCZEGO = (SELECT MAX(ID_HISTORI_AUTA_ZASTEPCZEGO) FROM HISTORIA_AUTA_ZASTEPCZEGO)", db)
		klientID := CheckLastID("SELECT ID_KLIENTA FROM Warsztat WHERE ID_KLIENTA = (SELECT MAX(ID_KLIENTA) FROM KLIENT)", db)
		zamowieniaID := CheckLastID("SELECT ID_ZAMOWIENIA FROM Warsztat WHERE ID_ZAMOWIENIA = (SELECT MAX(ID_ZAMOWIENIA) FROM ZAMOWIENIA)", db)
		magazyID := CheckLastID("SELECT ID_MAGAZYNU FROM Warsztat WHERE ID_MAGAZYNU = (SELECT MAX(ID_MAGAZYNU) FROM MAGAZYN)", db)
		pracownikID := CheckLastID("SELECT ID_PRACOWNIKA FROM Warsztat WHERE ID_PRACOWNIKA = (SELECT MAX(ID_PRACOWNIKA) FROM PRACOWNIK_PIELEGNACJI)", db)
		samochodKlientaID := CheckLastID("SELECT ID_SAMOCHODU_KLIENTA FROM Warsztat WHERE ID_SAMOCHODU_KLIENTA = (SELECT MAX(ID_SAMOCHODU_KLIENTA) FROM SAMOCHOD_KLIENTA)", db)
		uslugaID := CheckLastID("SELECT ID_USLUGI FROM Warsztat WHERE ID_USLUGI = (SELECT MAX(ID_USLUGI) FROM USLUGI)", db)

		for i := 1; i <= int(bodyData.AllTable); i++ {

			idA := autoZastepczeID + i

			data := time.Now().AddDate(0, 1, autoZastepczeID+idA)
			var d string
			if data.Month() < 10 && data.Day() < 10 {
				d = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
			} else if data.Month() < 10 {
				d = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
			} else if data.Day() < 10 {
				d = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
			} else {

				d = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
			}

			queryStirng := fmt.Sprintf("INSERT INTO AUTO_ZASTEPCZE (ID_AUTA_ZASTEPCZEGO, NUMER_REJESTRACYJNY, TERMIN_PRZEGLADU, TERMIN_UBEZPIECZENIA) VALUES (%d, 'ok%d', '%s', '%s')", idA, idA, d, d)
			_, err = db.Query(queryStirng)
			if err != nil {
				fmt.Printf("error in  1 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idH := histAutaZastepczegoID + i
			var dH string
			data = time.Now().AddDate(0, -2, histAutaZastepczegoID+idH)
			if data.Month() < 10 && data.Day() < 10 {
				dH = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
			} else if data.Month() < 10 {
				dH = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
			} else if data.Day() < 10 {
				dH = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
			} else {

				dH = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
			}

			queryStirng = fmt.Sprintf("INSERT INTO HISTORIA_AUTA_ZASTEPCZEGO(ID_HISTORI_AUTA_ZASTEPCZEGO, TERMIN_WYDANIA, TERMIN_ODEBRANIA, ZDARZENIE, AUTO_ZASTEPCZE_ID_AUTA_ZASTEPCZEGO) VALUES(%d, '%s', '%s', NULL , %d)", idH, dH, dH, idH)
			_, err = db.Query(queryStirng)
			if err != nil {
				fmt.Printf("error in  2 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idK := klientID + i
			numerLokalu := strconv.Itoa(rand.Intn(5))
			if numerLokalu == "0" {
				numerLokalu = "Null"
			}
			queryStirng = fmt.Sprintf("INSERT INTO KLIENT VALUES('%d', 'Imie_%d', 'NAzwisko_%d', 'Ulica_%d', %d, %s)", idK, idK, idK, idK, idK, numerLokalu)
			_, err = db.Query(queryStirng)
			if err != nil {
				fmt.Printf("error in  3 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idW := warsztatID + i
			queryStirng = fmt.Sprintf("INSERT INTO WARSZTAT VALUES (%d, 'miasto_%d', 'Ulica_%d', %d)", idW, idW, idW, idW)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 4 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idZ := zamowieniaID + i
			czyZrealizowane := strconv.Itoa(rand.Intn(2))
			if czyZrealizowane == "1" {
				czyZrealizowane = "N"
			} else {
				czyZrealizowane = "T"
			}
			queryStirng = fmt.Sprintf("INSERT INTO ZAMOWIENIA VALUES(%d, '%s', '%s', %d)", idZ, czyZrealizowane, d, idZ)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 5 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idM := magazyID + i
			var dM string
			data = time.Now().AddDate(0, -2, magazyID+idM)
			if data.Month() < 10 && data.Day() < 10 {
				dM = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
			} else if data.Month() < 10 {
				dM = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
			} else if data.Day() < 10 {
				dM = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
			} else {

				dM = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
			}
			queryStirng = fmt.Sprintf("INSERT INTO MAGAZYN VALUES(%d, 'produkt_%d', 'dostawca_%d', %d, '%s', %d, %d)", idM, idM, idM, idM, dM, idM, idM)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 6 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idP := pracownikID + i
			queryStirng = fmt.Sprintf("INSERT INTO PRACOWNIK_PIELEGNACJI VALUES (%d, '20310230%d', 'imie_%d', 'nazwisko_%d', '1%d', 'dzial_%d', %d)", idP, idP, idP, idP, idP, idP, idP)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 7 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			relID := CheckLastID("SELECT ID_KLIENTA FROM Warsztat WHERE ID_KLIENTA = (SELECT MAX(ID_KLIENTA) FROM KLIENT)", db)
			idR := rand.Intn(relID)
			queryStirng = fmt.Sprintf("INSERT INTO RELATION_17 VALUES (%d, %d)", idR, idR)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 8 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idSK := samochodKlientaID + i
			queryStirng = fmt.Sprintf("INSERT INTO SAMOCHOD_KLIENTA VALUES(%d, 'DW03%d', 'model_%d', %d)", idSK, idSK, idSK, idSK)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 9 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

			idU := uslugaID + i
			var dU string
			data = time.Now().AddDate(0, 3, uslugaID+idU)
			if data.Month() < 10 && data.Day() < 10 {
				dU = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
			} else if data.Month() < 10 {
				dU = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
			} else if data.Day() < 10 {
				dU = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
			} else {

				dU = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
			}
			queryStirng = fmt.Sprintf("INSERT INTO USLUGA VALUES(%d, 'usluga_%d', '%s', %d, %d, %d)", idU, idU, dU, idU, idU, idU)
			_, err = db.Exec(queryStirng)
			if err != nil {
				fmt.Printf("error in 10 qaury: %v", err)
			}
			f.WriteString(queryStirng + "\n")

		}
	}

	if bodyData.SingleTable != 0 {

		hour, min, sec := time.Now().Clock()
		fileName := fmt.Sprintf("inserted_SingleTable_%d_%d_%d.txt", hour, min, sec)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		switch bodyData.TableChoice {
		case "AutoZas":
			autoZastepczeID := CheckLastID("SELECT id_auta_zastepczego FROM Warsztat WHERE id_auta_zastepczego = (SELECT MAX(id_auta_zastepczego) FROM auto_zastepcze)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idA := autoZastepczeID + i

				data := time.Now().AddDate(0, 1, autoZastepczeID+idA)
				var d string
				if data.Month() < 10 && data.Day() < 10 {
					d = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
				} else if data.Month() < 10 {
					d = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
				} else if data.Day() < 10 {
					d = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
				} else {

					d = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
				}

				queryStirng := fmt.Sprintf("INSERT INTO AUTO_ZASTEPCZE (ID_AUTA_ZASTEPCZEGO, NUMER_REJESTRACYJNY, TERMIN_PRZEGLADU, TERMIN_UBEZPIECZENIA) VALUES (%d, 'ok%d', '%s', '%s')", idA, idA, d, d)
				_, err = db.Query(queryStirng)
				if err != nil {
					fmt.Printf("error in  1 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "HisAutaZAs":
			histAutaZastepczegoID := CheckLastID("SELECT ID_HISTORI_AUTA_ZASTEPCZEGO FROM Warsztat WHERE ID_HISTORI_AUTA_ZASTEPCZEGO = (SELECT MAX(ID_HISTORI_AUTA_ZASTEPCZEGO) FROM HISTORIA_AUTA_ZASTEPCZEGO)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {

				idH := histAutaZastepczegoID + i
				var dH string
				data := time.Now().AddDate(0, -2, histAutaZastepczegoID+idH)
				if data.Month() < 10 && data.Day() < 10 {
					dH = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
				} else if data.Month() < 10 {
					dH = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
				} else if data.Day() < 10 {
					dH = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
				} else {

					dH = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
				}

				queryStirng := fmt.Sprintf("INSERT INTO HISTORIA_AUTA_ZASTEPCZEGO(ID_HISTORI_AUTA_ZASTEPCZEGO, TERMIN_WYDANIA, TERMIN_ODEBRANIA, ZDARZENIE, AUTO_ZASTEPCZE_ID_AUTA_ZASTEPCZEGO) VALUES(%d, '%s', '%s', NULL , %d)", idH, dH, dH, idH)
				_, err = db.Query(queryStirng)
				if err != nil {
					fmt.Printf("error in  2 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "Klient":
			klientID := CheckLastID("SELECT ID_KLIENTA FROM Warsztat WHERE ID_KLIENTA = (SELECT MAX(ID_KLIENTA) FROM KLIENT)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {

				idK := klientID + i
				numerLokalu := strconv.Itoa(rand.Intn(5))
				if numerLokalu == "0" {
					numerLokalu = "Null"
				}
				queryStirng := fmt.Sprintf("INSERT INTO KLIENT VALUES('%d', 'Imie_%d', 'NAzwisko_%d', 'Ulica_%d', %d, %s)", idK, idK, idK, idK, idK, numerLokalu)
				_, err = db.Query(queryStirng)
				if err != nil {
					fmt.Printf("error in  3 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "Warsztat":
			warsztatID := CheckLastID("SELECT id_warsztatu FROM Warsztat WHERE id_warsztatu = (SELECT MAX(id_warsztatu) FROM Warsztat)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idW := warsztatID + i
				queryStirng := fmt.Sprintf("INSERT INTO WARSZTAT VALUES (%d, 'miasto_%d', 'Ulica_%d', %d)", idW, idW, idW, idW)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 4 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")

			}
		case "ZAmowienia":
			zamowieniaID := CheckLastID("SELECT ID_ZAMOWIENIA FROM Warsztat WHERE ID_ZAMOWIENIA = (SELECT MAX(ID_ZAMOWIENIA) FROM ZAMOWIENIA)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idZ := zamowieniaID + i
				czyZrealizowane := strconv.Itoa(rand.Intn(2))
				if czyZrealizowane == "1" {
					czyZrealizowane = "N"
				} else {
					czyZrealizowane = "T"
				}
				var dZ string
				data := time.Now().AddDate(0, 3, zamowieniaID+idZ)
				if data.Month() < 10 && data.Day() < 10 {
					dZ = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
				} else if data.Month() < 10 {
					dZ = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
				} else if data.Day() < 10 {
					dZ = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
				} else {

					dZ = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
				}
				queryStirng := fmt.Sprintf("INSERT INTO ZAMOWIENIA VALUES(%d, '%s', '%s', %d)", idZ, czyZrealizowane, dZ, idZ)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 5 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "Magazyn":
			magazyID := CheckLastID("SELECT ID_MAGAZYNU FROM Warsztat WHERE ID_MAGAZYNU = (SELECT MAX(ID_MAGAZYNU) FROM MAGAZYN)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idM := magazyID + i
				var dM string
				data := time.Now().AddDate(0, -2, magazyID+idM)
				if data.Month() < 10 && data.Day() < 10 {
					dM = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
				} else if data.Month() < 10 {
					dM = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
				} else if data.Day() < 10 {
					dM = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
				} else {

					dM = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
				}
				queryStirng := fmt.Sprintf("INSERT INTO MAGAZYN VALUES(%d, 'produkt_%d', 'dostawca_%d', %d, '%s', %d, %d)", idM, idM, idM, idM, dM, idM, idM)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 6 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "Pracownik":
			pracownikID := CheckLastID("SELECT ID_PRACOWNIKA FROM Warsztat WHERE ID_PRACOWNIKA = (SELECT MAX(ID_PRACOWNIKA) FROM PRACOWNIK_PIELEGNACJI)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idP := pracownikID + i
				queryStirng := fmt.Sprintf("INSERT INTO PRACOWNIK_PIELEGNACJI VALUES (%d, '20310230%d', 'imie_%d', 'nazwisko_%d', '1%d', 'dzial_%d', %d)", idP, idP, idP, idP, idP, idP, idP)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 7 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "SamochodKlienta":
			samochodKlientaID := CheckLastID("SELECT ID_SAMOCHODU_KLIENTA FROM Warsztat WHERE ID_SAMOCHODU_KLIENTA = (SELECT MAX(ID_SAMOCHODU_KLIENTA) FROM SAMOCHOD_KLIENTA)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idSK := samochodKlientaID + i
				queryStirng := fmt.Sprintf("INSERT INTO SAMOCHOD_KLIENTA VALUES(%d, 'DW03%d', 'model_%d', %d)", idSK, idSK, idSK, idSK)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 9 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")
			}
		case "USLUGA ":
			uslugaID := CheckLastID("SELECT ID_USLUGI FROM Warsztat WHERE ID_USLUGI = (SELECT MAX(ID_USLUGI) FROM USLUGI)", db)

			for i := 1; i < int(bodyData.SingleTable); i++ {
				idU := uslugaID + i
				var dU string
				data := time.Now().AddDate(0, 3, uslugaID+idU)
				if data.Month() < 10 && data.Day() < 10 {
					dU = fmt.Sprintf("%d-0%d-0%d", data.Year(), data.Month(), data.Day())
				} else if data.Month() < 10 {
					dU = fmt.Sprintf("%d-0%d-%d", data.Year(), data.Month(), data.Day())
				} else if data.Day() < 10 {
					dU = fmt.Sprintf("%d-%d-0%d", data.Year(), data.Month(), data.Day())
				} else {

					dU = fmt.Sprintf("%d-%d-%d", data.Year(), data.Month(), data.Day())
				}
				queryStirng := fmt.Sprintf("INSERT INTO USLUGA VALUES(%d, 'usluga_%d', '%s', %d, %d, %d)", idU, idU, dU, idU, idU, idU)
				_, err = db.Exec(queryStirng)
				if err != nil {
					fmt.Printf("error in 10 qaury: %v", err)
				}
				f.WriteString(queryStirng + "\n")

			}
		default:
		}
	}

}

func CheckLastID(query string, db *sql.DB) int {
	lastID := 0
	rows, err := db.Query("SELECT id_warsztatu FROM Warsztat WHERE id_warsztatu = (SELECT MAX(id_warsztatu) FROM Warsztat)")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&lastID)
		if err != nil {
			log.Fatal(err)
		}
	}
	return lastID
}
