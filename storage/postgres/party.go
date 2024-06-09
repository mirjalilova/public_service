package postgres

import (
	"database/sql"
	"fmt"
	pb "public/genproto"
	"strings"
	"time"
)

type PartyRepo struct {
	db *sql.DB
}

func NewPartyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{}
}

func (pr *PartyRepo) Create(party *pb.PartyCreate) (*pb.Void, error) {
	res := &pb.Void{}

	query := `INSERT INTO party(name, slogan, opened_date, description) VALUES ($1, $2, $3, $4)`

	_, err := pr.db.Exec(query, party.Name, party.Slogan, party.OpenedDate, party.Description)

	if err != nil {
		return nil, err
	}
	return res, nil
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

func (pr *PartyRepo) GetAll(party *pb.GetAllPartysRequest) (*pb.GetAllPartysResponse, error) {
	res := &pb.GetAllPartysResponse{
		Party: []*pb.PartyRes{},
	}

	query := `SELECT id, name, slogan, opened_date, description FROM party`

	var args []interface{}
	var conditions []string

	if party.Name != "" {
		args = append(args, party.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}
	if party.OpenedDate != "" {
		args = append(args, party.OpenedDate)
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
	if party.Filter.Limit == 0 {
		party.Filter.Limit = defaultLimit
	}

	args = append(args, party.Filter.Limit, party.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	fmt.Println(query, args)
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
		res.Party = append(res.Party, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (pr *PartyRepo) Update(party *pb.PartyUpdate) (*pb.Void, error) {
	res := &pb.Void{}

	query := `UPDATE party SET name = $1, slogan $ 2, opened_date = $3, description = $4, updated_at=now() WHERE id = $5`

	_, err := pr.db.Exec(query,
		party.UpdateParty.Name,
		party.UpdateParty.Slogan,
		party.UpdateParty.OpenedDate,
		party.UpdateParty.Description,
		party.Id,
	)

	return res, err
}

func (pr *PartyRepo) Delete(id *pb.GetByIdReq) (*pb.Void, error) {
	res := &pb.Void{}

	query := `UPDATE party SET deleted_at=$1 WHERE id = $2`

	_, err := pr.db.Exec(query, time.Now(), id)

	return res, err

}
