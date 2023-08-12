package wallet

import (
	"database/sql"
	"dbgo/db"
	"log"
)

func InsertWallet(owner string, balance float64) (int, error) {
	query := `
			INSERT INTO wallet (owner,balance)
			VALUES ($1, $2) RETURNING id
	`
	row := db.Conn.QueryRow(query, owner, balance)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	log.Println("last insert id: ", id)
	return id, nil
}

func updateWalletBalance(walletID int, newBalance float64) error {
	query := `
			UPDATE wallet SET balance = $1 WHERE id = $2
	`
	_, err := db.Conn.Exec(query, newBalance, walletID)
	if err != nil {
		return err
	}

	return nil
}

func deleteWallet(db *sql.DB, walletID int) error {
	query := `
			DELETE FROM wallet WHERE id = $1
	`
	_, err := db.Exec(query, walletID)
	if err != nil {
		return err
	}

	return nil
}

func GetWalletById(walletID int) (Wallet, error) {
	query := `
			SELECT *
			FROM wallet w WHERE id = $1
	`
	row := db.Conn.QueryRow(query, walletID)
	var id int
	var owner string
	var balance float64
	err := row.Scan(&id, &owner, &balance)
	var wal Wallet
	if err != nil {
		log.Println("can not parse wallet to struct")
		return wal, err
	}
	wal.ID = id
	wal.Balance = balance
	wal.Owner = owner
	return wal, nil
}