package postgres

import (
	"database/sql"
	"fmt"
	pb "public_service/genproto"
	"strings"
	"time"
)

type PartyRepo struct {
	db *sql.DB
}

func NewPartyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{}
}

func (pr *PartyRepo) Create(party *pb.PartyCreate) error {
	query := `INSERT INTO party(name, slogan, opened_date, description) VALUES ($1, $2, $3, $4)`

	_, err := pr.db.Exec(query, party.Name, party.Slogan, party.OpenedDate, party.Description)

	if err != nil {
		return err
	}
	return nil
}

func (pr *PartyRepo) GetByID(id *pb.GetByIdReq) (*pb.PartyRes, error) {
	res := &pb.PartyRes{}
	query := `SELECT id, name, slogan, opened_date, description FROM party WHERE id = $1`
	row := pr.db.QueryRow(query, id)

	err := row.Scan(
		&res.Id,
		&res.Name,
		&res.Slogan,
		&res.OpenedDate,
		&res.Description,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *PartyRepo) GetAll(filter *pb.Filter, name, opened_date string) (*pb.GetAllPartysResponse, error) {
	res := &[]pb.PartyRes{}

	query := `SELECT id, name, slogan, opened_date, description FROM party`

	var args []interface{}
	var conditions []string

	if name != "" {
		args = append(args, name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}
	if opened_date != "" {
		args = append(args, opened_date)
		conditions = append(conditions, fmt.Sprintf("opened_date = $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int32
	err := pr.db.QueryRow("SELECT COUNT(1) FROM provision WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := pr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pr := &pb.PartyRes{}
		err := rows.Scan(
			&pr.Id,
			&pr.Name,
			&pr.Slogan,
			&pr.OpenedDate,
			&pr.Description,
		)
		if err != nil {
			return nil, err
		}
		*res = append(*res, *pr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := &pb.GetAllPartysResponse{
		Party: res,
	}
	return response, nil
}

func (pr *PartyRepo) Update(party *pb.PartyUpdate) error {
	query := `UPDATE party SET name = $1, slogan $ 2, opened_date = $3, description = $4, updated_at=now() WHERE id = $5`

	_, err := pr.db.Exec(query, party.Name, party.Slogan, party.OpenedDate, party.Description)

	return err
}

func (pr *PartyRepo) Delete(id *pb.GetByIdReq) error {
	query := `UPDATE party SET deleted_at=$1 WHERE id = $2`

	_, err := pr.db.Exec(query, time.Now(), id)

	return err

}
