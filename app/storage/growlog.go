package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/tibeahx/growers-dairy/app/types"
)

func (d *DB) CreateGrowLog(ctx context.Context, req types.CreateGrowLogReq) (types.CreateGrowLogResp, error) {
	const query = `INSERT INTO grow_logs 
	(name, start_date, end_date, description, strain_id)
	values($1, $2, $3, $4, (SELECT id FROM strains WHERE id = $5))
	returning id, name, start_date, end_date, description, strain_id;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var resp types.CreateGrowLogResp
	err := d.db.QueryRowContext(ctx, query,
		req.Name,
		time.Now().UTC(),
		req.EndDate,
		req.Description,
		req.StrainID,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.StartDate,
		&resp.EndDate,
		&resp.Description,
		&resp.StrainID,
	)
	return resp, err
}

func (d *DB) UpdateGrowLog(ctx context.Context, req types.UpdateGrowLogReq) (types.UpdateGrowLogResp, error) {
	const query = `
	UPDATE grow_logs
	SET name = $1, end_date = $2, description = $3
	where id = $4
	returning id, name, start_date, end_date, description, strain_id; 
	`
	exists, _ := d.AssertGrowLogIDExists(req.ID)
	if !exists {
		return types.UpdateGrowLogResp{}, errGrowLogNotFound
	}
	var resp types.UpdateGrowLogResp
	err := d.db.QueryRowContext(ctx, query,
		req.Name,
		req.EndDate,
		req.Description,
		req.ID,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.StartDate,
		&resp.EndDate,
		&resp.Description,
		&resp.StrainID,
	)
	if err == sql.ErrNoRows {
		return types.UpdateGrowLogResp{}, errGrowLogNotFound
	}
	return resp, err
}
func (d *DB) GrowLogByID(ctx context.Context, id int) (types.GrowLog, error) {
	const query = `SELECT * FROM grow_logs WHERE id = $1;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	exists, _ := d.AssertGrowLogIDExists(id)
	if !exists {
		return types.GrowLog{}, errGrowLogNotFound
	}
	rows, err := d.db.QueryContext(ctx, query, id)
	if err != nil {
		return types.GrowLog{}, err
	}
	var growLog types.GrowLog
	for rows.Next() {
		err := rows.Scan(
			&growLog.ID,
			&growLog.Name,
			&growLog.StartDate,
			&growLog.EndDate,
			&growLog.Description,
			&growLog.StrainID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return types.GrowLog{}, errGrowLogNotFound
			}
			return types.GrowLog{}, err
		}
	}
	return growLog, nil
}
func (d *DB) GrowLogs(ctx context.Context) ([]types.GrowLog, error) {
	const query = `SELECT * FROM grow_logs order by id asc;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var growLogs []types.GrowLog
	for rows.Next() {
		var growLog types.GrowLog
		err := rows.Scan(
			&growLog.ID,
			&growLog.Name,
			&growLog.StartDate,
			&growLog.EndDate,
			&growLog.Description,
			&growLog.StrainID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errGrowLogNotFound
			}
		}
		growLogs = append(growLogs, growLog)
	}
	return growLogs, err
}
func (d *DB) DeleteGrowLogByID(ctx context.Context, id int) error {
	const query = `DELETE FROM grow_logs WHERE id = $1;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if exists, _ := d.AssertGrowLogIDExists(id); !exists {
		return errGrowLogNotFound
	}
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errGrowLogNotFound
		}
		return err
	}
	return nil
}

func (d *DB) AssertGrowLogIDExists(id int) (bool, error) {
	var exists bool
	const query = `SELECT id FROM grow_logs WHERE id = $1;`
	err := d.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	return true, err
}
