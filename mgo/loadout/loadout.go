package loadout

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
	// Store entire MGOLoadoutData as JSON blob per player
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS mgo_loadout_data (
			player_id INTEGER PRIMARY KEY,
			data BLOB NOT NULL,
			FOREIGN KEY (player_id) REFERENCES player(id)
		)
	`)
	return err
}

func LoadDefaultLoadouts() (tppmessage.MGOLoadoutData, error) {
	var loadouts tppmessage.MGOLoadoutData
	data, err := os.ReadFile("default_loadout.json")
	if err != nil {
		return tppmessage.MGOLoadoutData{}, err
	}
	if err := json.Unmarshal(data, &loadouts); err != nil {
		return tppmessage.MGOLoadoutData{}, err
	}
	return loadouts, nil
}

func (r *Repo) FindAllByPlayer(ctx context.Context, playerID int) (tppmessage.MGOLoadoutData, error) {
	row := r.db.QueryRowContext(ctx, "SELECT data FROM mgo_loadout_data WHERE player_id = ?", playerID)

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			// No loadout data found, load and return defaults
			return LoadDefaultLoadouts()
		}
		return tppmessage.MGOLoadoutData{}, err
	}

	var loadoutData tppmessage.MGOLoadoutData
	if err := json.Unmarshal(data, &loadoutData); err != nil {
		return tppmessage.MGOLoadoutData{}, err
	}

	return loadoutData, nil
}

// upsert
func (r *Repo) Upsert(ctx context.Context, playerID int, loadoutData tppmessage.MGOLoadoutData) error {
	data, err := json.Marshal(loadoutData)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, `
		INSERT OR REPLACE INTO mgo_loadout_data (player_id, data)
		VALUES (?, ?)
	`, playerID, data)

	return err
}
