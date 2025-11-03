package character

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/unknown321/fuse/tppmessage"
)

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	// Store entire MGOCharacterData as JSON blob per player
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS mgo_character_data (
			player_id INTEGER PRIMARY KEY,
			data BLOB NOT NULL,
			FOREIGN KEY (player_id) REFERENCES player(id)
		)
	`)
	return err
}

func LoadDefaultCharacterData() (tppmessage.MGOCharacterData, error) {
	var chars tppmessage.MGOCharacterData
	data, err := os.ReadFile("default_character.json")
	if err != nil {
		return tppmessage.MGOCharacterData{}, err
	}
	if err := json.Unmarshal(data, &chars); err != nil {
		return tppmessage.MGOCharacterData{}, err
	}
	return chars, nil
}

func (r *Repo) FindAllByPlayer(ctx context.Context, playerID int) (tppmessage.MGOCharacterData, error) {
	row := r.db.QueryRowContext(ctx, "SELECT data FROM mgo_character_data WHERE player_id = ?", playerID)

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			// No character data found, load and return defaults
			return LoadDefaultCharacterData()
		}
		return tppmessage.MGOCharacterData{}, err
	}

	var charData tppmessage.MGOCharacterData
	if err := json.Unmarshal(data, &charData); err != nil {
		return tppmessage.MGOCharacterData{}, err
	}

	return charData, nil
}

// upsert
func (r *Repo) Upsert(ctx context.Context, playerID int, char tppmessage.MGOCharacterData) error {
	data, err := json.Marshal(char)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, `
		INSERT OR REPLACE INTO mgo_character_data (player_id, data)
		VALUES (?, ?)
	`, playerID, data)

	return err
}

// delete a character by player ID and character index
func (r *Repo) Delete(ctx context.Context, playerID int, characterID int) error {
	// Load current character data
	charData, err := r.FindAllByPlayer(ctx, playerID)
	if err != nil {
		return err
	}

	// Filter out the character to delete
	var updatedCharList []tppmessage.MGOCharacter
	for _, char := range charData.CharacterList {
		if char.CharacterID != characterID {
			updatedCharList = append(updatedCharList, char)
		}
	}

	// Update the character list
	charData.CharacterList = updatedCharList

	// Save the updated data back
	return r.Upsert(ctx, playerID, charData)
}
