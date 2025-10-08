package character

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/unknown321/fuse/player"
	"github.com/unknown321/fuse/tppmessage"
)

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS mgo_character (
			player_id INTEGER NOT NULL,
			character_id INTEGER NOT NULL,
			data BLOB,
			PRIMARY KEY (player_id, character_id),
			FOREIGN KEY (player_id) REFERENCES player(id)
		)
	`)
	return err
}

func (r *Repo) Find(ctx context.Context, playerID player.ID, characterID int) (*tppmessage.MGOCharacter, error) {
	row := r.db.QueryRowContext(ctx, "SELECT data FROM mgo_character WHERE player_id = ? AND character_id = ?", playerID, characterID)

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var char tppmessage.MGOCharacter
	if err := json.Unmarshal(data, &char); err != nil {
		return nil, err
	}

	return &char, nil
}

func (r *Repo) Save(ctx context.Context, playerID player.ID, char *tppmessage.MGOCharacter) error {
	data, err := json.Marshal(char)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, `
		INSERT OR REPLACE INTO mgo_character (player_id, character_id, data)
		VALUES (?, ?, ?)
	`, playerID, char.CharacterID, data)

	return err
}

func (r *Repo) FindAllByPlayer(ctx context.Context, playerID player.ID) ([]*tppmessage.MGOCharacter, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT data FROM mgo_character WHERE player_id = ?", playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chars []*tppmessage.MGOCharacter
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, err
		}

		var char tppmessage.MGOCharacter
		if err := json.Unmarshal(data, &char); err != nil {
			return nil, err
		}
		chars = append(chars, &char)
	}

	return chars, nil
}