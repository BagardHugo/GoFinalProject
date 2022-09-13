package account

import (
	"database/sql"
	"errors"
	"finalProject/constants"
	"finalProject/wallet"

	"github.com/jackc/pgconn"
)

type PostgreSqlRepository struct {
	db *sql.DB
}

func NewPostgreSqlRepository(db *sql.DB) *PostgreSqlRepository {
	return &PostgreSqlRepository{
		db: db,
	}
}

func (r *PostgreSqlRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS accounts(
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    	username text COLLATE pg_catalog."default" NOT NULL,
    	password text COLLATE pg_catalog."default" NOT NULL,
    	pincode text COLLATE pg_catalog."default" NOT NULL,
    	CONSTRAINT players_pkey PRIMARY KEY (id),
    	CONSTRAINT username UNIQUE (username)
    );

	CREATE TABLE IF NOT EXISTS wallets(
		id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    	address text COLLATE pg_catalog."default" NOT NULL,
    	id_player integer,
    	CONSTRAINT wallets_pkey PRIMARY KEY (id),
    	CONSTRAINT unique_address UNIQUE (address),
    	CONSTRAINT fk_id_player FOREIGN KEY (id_player)
        REFERENCES public.accounts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
	)
    `
	_, err := r.db.Exec(query)
	return err
}

func (r *PostgreSqlRepository) Create(account Account, wallet wallet.MockWalletReturn) (Account, error) {
	var id int
	var query = `INSERT INTO accounts(userName, password, pinCode) values($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(query, account.UserName, account.Password, account.PinCode).Scan(&id)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return Account{}, constants.ErrDuplicate
			}
		}
		return Account{}, err
	}
	account.Id = id

	// Create wallet with account id
	wallet, _ = r.LinkWallet(account.Id, wallet)
	account.Wallet = wallet

	return account, nil
}

func (r *PostgreSqlRepository) LinkWallet(accountId int, wallet wallet.MockWalletReturn) (wallet.MockWalletReturn, error) {
	var id int

	_ = r.db.QueryRow(
		"insert into wallets(id_player, address) values ($1, $2)",
		accountId,
		wallet.WalletAddress).Scan(&id)
	return wallet, nil
}
