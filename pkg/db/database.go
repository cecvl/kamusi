package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SearchWord(word string) (string, []string, error) {
   db, err := sql.Open("sqlite3", "C:\\Users\\user\\kamusi\\test4.db")
   if err != nil {
       return "", nil, err
   }
   defer db.Close()

   var meaning string
   err = db.QueryRow("SELECT meaning FROM Words WHERE word = ?", word).Scan(&meaning)
   if err != nil {
       if err == sql.ErrNoRows {
           return "", nil, fmt.Errorf("word not found")
       }
       return "", nil, err
   }

   var synonyms []string
   rows, err := db.Query("SELECT synonym FROM Synonyms WHERE word = ?", word)
   if err != nil {
       return "", nil, err
   }
   defer rows.Close()

   for rows.Next() {
       var synonym string
       if err := rows.Scan(&synonym); err != nil {
           return "", nil, err
       }
       synonyms = append(synonyms, synonym)
   }

   return meaning, synonyms, nil
}
