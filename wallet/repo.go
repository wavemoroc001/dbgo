package wallet

import (
	"database/sql"
	"log"
)

type Repo struct {
	conn *sql.DB
}

func ProvideRepo(conn *sql.DB) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) InsertWallet(owner string, balance float64) (int, error) {
	query := `
			INSERT INTO wallet (owner,balance)
			VALUES ($1, $2) RETURNING id
	`
	row := r.conn.QueryRow(query, owner, balance)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	log.Println("last insert id: ", id)
	return id, nil
}

func (r *Repo) UpdateWalletBalance(walletID int, newBalance float64) error {
	query := `
			UPDATE wallet SET balance = $1 WHERE id = $2
	`
	_, err := r.conn.Exec(query, newBalance, walletID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteWallet(walletID int) error {
	query := `
			DELETE FROM wallet WHERE id = $1
	`
	_, err := r.conn.Exec(query, walletID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetWalletById(walletID int) (Wallet, error) {
	query := `
			SELECT *
			FROM wallet w WHERE id = $1
	`
	row := r.conn.QueryRow(query, walletID)

	var wal Wallet

	err := row.Scan(&wal.ID, &wal.Owner, &wal.Balance)
	if err != nil {
		log.Println("can not parse wallet to struct")
		return wal, err
	}

	return wal, nil
}
