package loadout

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
		CREATE TABLE IF NOT EXISTS mgo_loadout (
			player_id INTEGER NOT NULL,
			character_id INTEGER NOT NULL,
			loadout_index INTEGER NOT NULL,
			data BLOB,
			PRIMARY KEY (player_id, character_id, loadout_index),
			FOREIGN KEY (player_id) REFERENCES player(id)
		)
	`)
	return err
}

func (r *Repo) FindByCharacter(ctx context.Context, playerID player.ID, characterID int) ([]*tppmessage.MGOLoadout, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT data FROM mgo_loadout WHERE player_id = ? AND character_id = ?", playerID, characterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loadouts []*tppmessage.MGOLoadout
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, err
		}

		var loadout tppmessage.MGOLoadout
		if err := json.Unmarshal(data, &loadout); err != nil {
			return nil, err
		}
		loadouts = append(loadouts, &loadout)
	}

	return loadouts, nil
}

func (r *Repo) Save(ctx context.Context, playerID player.ID, characterID int, loadout *tppmessage.MGOLoadout) error {
	data, err := json.Marshal(loadout)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, `
		INSERT OR REPLACE INTO mgo_loadout (player_id, character_id, loadout_index, data)
		VALUES (?, ?, ?, ?)
	`, playerID, characterID, loadout.LoadoutIndex, data)

	return err
}

func (r *Repo) Find(ctx context.Context, playerID player.ID, characterID int, loadoutIndex int) (*tppmessage.MGOLoadout, error) {
	row := r.db.QueryRowContext(ctx, "SELECT data FROM mgo_loadout WHERE player_id = ? AND character_id = ? AND loadout_index = ?", playerID, characterID, loadoutIndex)

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var loadout tppmessage.MGOLoadout
	if err := json.Unmarshal(data, &loadout); err != nil {
		return nil, err
	}

	return &loadout, nil
}