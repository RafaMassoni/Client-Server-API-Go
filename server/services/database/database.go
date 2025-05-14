package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"server/model/tableModel"
	"time"

	_ "modernc.org/sqlite"
)

func getDatabase() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "./quotes.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func InitDataBase() {

	println("O banco de dados está inicializando")

	db, err := getDatabase()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS quote_dollar (id integer not null primary key, quote text);
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	quotes, err := getAllDollarQuotes()

	if err != nil {
		log.Printf("%q", err)
		return
	}

	for _, quote := range quotes {
		log.Printf("ID: %d, Quote: %s\n", quote.ID, quote.DollarValue)
	}

	println("O banco de dados foi inicializado com sucesso")
	println("")
}

func InsertDollarQuote(dollarQuote tableModel.DollarQuote) error {
	db, err := getDatabase()
	if err != nil {
		log.Println("Erro ao obter conexão com o banco de dados:", err)
		return fmt.Errorf("falha ao conectar no banco de dados: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO quote_dollar(quote) VALUES(?)")
	if err != nil {
		log.Println("Erro ao preparar a query de inserção:", err)
		return fmt.Errorf("erro ao preparar query SQL: %w", err)
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err = stmt.ExecContext(ctx, dollarQuote.DollarValue)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("tempo limite excedido para inserir cotação no banco")
			return fmt.Errorf("tempo limite excedido para inserir cotação no banco: %w", err)
		}
		log.Println("Erro ao executar a query de inserção:", err)
		return fmt.Errorf("erro ao inserir cotação no banco: %w", err)
	}

	log.Println("Registro inserido com sucesso: DollarValue ->", dollarQuote.DollarValue)
	return nil
}

func getAllDollarQuotes() ([]tableModel.DollarQuote, error) {
	db, err := getDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, quote FROM quote_dollar")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []tableModel.DollarQuote

	for rows.Next() {
		var dq tableModel.DollarQuote
		err := rows.Scan(&dq.ID, &dq.DollarValue)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, dq)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}
