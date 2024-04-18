package storage

import (
	"context"
	"database/sql"

	"github.com/tibeahx/growers-dairy/app/types"
)

func (d *DB) CreateStrain(ctx context.Context, req types.CreateStrainReq) (resp types.CreateStrainResp, err error) {
	const query = `INSERT INTO strains (name, description, feminized, regular, auto, photo)
	values($1, $2, $3, $4, $5, $6)
	returning id, name, description, feminized, regular, auto, photo;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	err = d.db.QueryRowContext(ctx, query,
		req.Name,
		req.Description,
		req.StrainAttrs.FeminizedType.Feminized,
		req.StrainAttrs.FeminizedType.Regular,
		req.StrainAttrs.StrainType.Auto,
		req.StrainAttrs.StrainType.Photo,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.StrainAttrs.FeminizedType.Feminized,
		&resp.StrainAttrs.FeminizedType.Regular,
		&resp.StrainAttrs.StrainType.Auto,
		&resp.StrainAttrs.StrainType.Photo,
	)
	return resp, err
}

func (d *DB) StrainByID(ctx context.Context, id int) (types.Strain, error) {
	const query = `SELECT * FROM strains WHERE id = $1;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if exists, _ := d.AssertStrainIDExists(id); exists {
		return types.Strain{}, errStrainNotFound
	}
	rows, err := d.db.QueryContext(ctx, query, id)
	if err != nil {
		return types.Strain{}, err
	}
	var strain types.Strain
	for rows.Next() {
		err := rows.Scan(
			&strain.ID,
			&strain.Name,
			&strain.StrainAttrs.StrainType.Auto,
			&strain.StrainAttrs.StrainType.Photo,
			&strain.Description,
			&strain.StrainAttrs.FeminizedType.Feminized,
			&strain.StrainAttrs.FeminizedType.Regular,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return types.Strain{}, errStrainNotFound
			}
			return types.Strain{}, err
		}
	}
	return strain, nil
}

func (d *DB) DeleteStrainByID(ctx context.Context, id int) error {
	const query = `DELETE FROM strains WHERE id = $1;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if exists, _ := d.AssertStrainIDExists(id); !exists {
		return errStrainNotFound
	}
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errStrainNotFound
		}
		return errScanningRows
	}
	return nil
}

func (d *DB) Strains(ctx context.Context) ([]types.Strain, error) {
	const query = `SELECT * FROM strains order by id asc;`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var strains []types.Strain
	for rows.Next() {
		var strain types.Strain
		err := rows.Scan(
			&strain.ID,
			&strain.Name,
			&strain.StrainAttrs.StrainType.Auto,
			&strain.StrainAttrs.StrainType.Photo,
			&strain.Description,
			&strain.StrainAttrs.FeminizedType.Feminized,
			&strain.StrainAttrs.FeminizedType.Regular,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errStrainNotFound
			}
			return nil, err
		}
		strains = append(strains, strain)
	}
	return strains, nil
}

func (d *DB) UpdateStrain(ctx context.Context, req types.UpdateStrainReq) (types.UpdateStrainResp, error) {
	const query = `
	UPDATE strains
	SET name = $1, description = $2, feminized = $3, regular = $4, auto = $5, photo = $6
	WHERE id = $7
	returning id, name, description, feminized, regular, auto, photo; 
	`
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	exists, _ := d.AssertStrainIDExists(req.ID)
	if !exists {
		return types.UpdateStrainResp{}, errStrainNotFound
	}
	var resp types.UpdateStrainResp
	err := d.db.QueryRowContext(ctx, query,
		req.Name,
		req.Description,
		req.StrainAttrs.FeminizedType.Feminized,
		req.StrainAttrs.FeminizedType.Regular,
		req.StrainAttrs.StrainType.Auto,
		req.StrainAttrs.StrainType.Photo,
		req.ID,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.StrainAttrs.FeminizedType.Feminized,
		&resp.StrainAttrs.FeminizedType.Regular,
		&resp.StrainAttrs.StrainType.Auto,
		&resp.StrainAttrs.StrainType.Photo,
	)
	if err == sql.ErrNoRows {
		return types.UpdateStrainResp{}, errStrainNotFound
	}
	return resp, err
}

func (d *DB) AssertStrainIDExists(id int) (bool, error) {
	var exists bool
	const query = `SELECT id FROM strains WHERE id = $1;`
	err := d.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	return true, err
}
