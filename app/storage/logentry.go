package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/tibeahx/growers-dairy/app/types"
)

func (d *DB) CrateLogEntry(ctx context.Context, req types.CreateLogEntryReq) (types.CreateLogEntryResp, error) {
	const query = `INSERT INTO log_entries
					(entry_date, comment, growlog_id)
					values ($1, $2, (SELECT id FROM grow_logs WHERE id = $3))
					returning id, entry_date, comment, growlog_id;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var resp types.CreateLogEntryResp
	err := d.db.QueryRowContext(ctx, query,
		time.Now().UTC(),
		req.Comment,
		req.GrowLogID,
	).Scan(
		&resp.ID,
		&resp.EntryDate,
		&resp.Comment,
		&resp.GrowLogID,
	)
	return resp, err
}

func (d *DB) LogEntries(ctx context.Context) ([]types.LogEntry, error) {
	const query = `SELECT id, entry_date, comment, photos, growlog_id FROM log_entries order by id asc;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var entries []types.LogEntry
	for rows.Next() {
		var entry types.LogEntry
		err = rows.Scan(
			&entry.ID,
			&entry.EntryDate,
			&entry.Comment,
			pq.Array(&entry.Photos),
			&entry.GrowLogID,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (d *DB) UpdateEntry(ctx context.Context, req types.UpdateLogEntryReq, bucketName string, objName string) (types.UpdateLogEntryResp, error) {
	const query = `UPDATE log_entries
					SET comment = $1, growlog_id = (SELECT id FROM grow_logs WHERE id = $2)
					WHERE id = $3
					RETURNING id, comment, growlog_id;
					`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var resp types.UpdateLogEntryResp
	err := d.db.QueryRowContext(ctx, query,
		req.Comment,
		req.GrowLogID,
		req.ID,
	).Scan(
		&resp.ID,
		&resp.Comment,
		&resp.GrowLogID,
	)
	if err == sql.ErrNoRows {
		return types.UpdateLogEntryResp{}, errEntryNotFound
	}
	return resp, err
}

func (d *DB) DeleteEntryByID(ctx context.Context, id int) error {
	const query = `DELETE FROM log_entries WHERE id = $1;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if exists, _ := d.AssertLogEntryIDExists(id); !exists {
		return errEntryNotFound
	}
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errEntryNotFound
		}
		return errScanningRows
	}
	return nil
}

func (d *DB) EntryByID(ctx context.Context, id int) (types.LogEntry, error) {
	const query = `SELECT id, entry_date, comment, photos, growlog_id FROM log_entries WHERE id = $1 order by id asc;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if exists, _ := d.AssertLogEntryIDExists(id); !exists {
		return types.LogEntry{}, errEntryNotFound
	}
	rows, err := d.db.QueryContext(ctx, query, id)
	if err != nil {
		return types.LogEntry{}, err
	}
	var entry types.LogEntry
	for rows.Next() {
		err := rows.Scan(
			&entry.ID,
			&entry.EntryDate,
			&entry.Comment,
			pq.Array(&entry.Photos),
			&entry.GrowLogID,
		)
		if err != nil {
			return types.LogEntry{}, err
		}
	}
	return entry, nil
}

func (d *DB) InsertS3Link(ctx context.Context, id int, url string) error {
	const query = `UPDATE log_entries
					SET photos = $1
					WHERE id = $2;
					`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	urls := d.getCurrentURLs(ctx, id)
	urls = append(urls, url)
	_, err := d.db.ExecContext(ctx, query, pq.Array(urls), id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) getCurrentURLs(ctx context.Context, id int) []string {
	const query = `SELECT photos FROM log_entries WHERE id = $1;`
	var urls []string
	err := d.db.QueryRowContext(ctx, query, id).Scan(pq.Array(&urls))
	if err != nil {
		return nil
	}
	return urls
}

func (d *DB) AssertLogEntryIDExists(id int) (bool, error) {
	var exists bool
	const query = `SELECT id FROM log_entries WHERE id = $1;`
	err := d.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	return true, err
}
