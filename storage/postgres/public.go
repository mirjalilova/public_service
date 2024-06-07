package postgres

import (
	"database/sql"
	"fmt"
	pb "public/genproto"
	"strings"
	"time"
)

type PublicRepo struct {
	db *sql.DB
}

func NewPublicRepo(db *sql.DB) *PublicRepo {

	return &PublicRepo{db: db}
}

func (p *PublicRepo) Create(public *pb.PublicCreate) (*pb.Void, error) {
	res := &pb.Void{}

	query := `INSERT INTO public (first_name, last_name, birthday, gender, nation, party_id) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := p.db.Exec(
		query,
		public.FirstName,
		public.LastName,
		public.Gender,
		public.Nation,
		public.PartyId,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *PublicRepo) GetById(id *pb.GetByIdReq) (*pb.PublicRes, error) {
	res := &pb.PublicRes{}

	query := `SELECT 
				pb.id, 
				pb.first_name, 
				pb.last_name, 
				pb.birthday, 
				pb.gender, 
				pb.nation, 
				pr.id, 
				pr.name, 
				pr.slogan, 
				pr.opened_date, 
				pr.description 
			FROM public pb
			join
			party pr
			on pb.party_id=pr.id
			WHERE id = $1`

	err := p.db.QueryRow(query, id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Birthday,
		&res.Gender,
		&res.Nation,
		&res.Party.Id,
		&res.Party.Name,
		&res.Party.Slogan,
		&res.Party.OpenedDate,
		&res.Party.Description,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *PublicRepo) GetAll(public *pb.GetAllPublicsRequest) (*pb.GetAllPublicsResponse, error) {
	res := &pb.GetAllPublicsResponse{
		Public: []*pb.PublicRes{},
	}

	query := `SELECT 
				pb.id, 
				pb.first_name, 
				pb.last_name, 
				pb.birthday, 
				pb.gender, 
				pb.nation, 
				pr.id, 
				pr.name, 
				pr.slogan, 
				pr.opened_date, 
				pr.description 
			FROM public pb
			JOIN party pr
			ON pb.party_id = pr.id`

	var args []interface{}
	var conditions []string

	if public.Gender != "" {
		args = append(args, public.Gender)
		conditions = append(conditions, fmt.Sprintf("gender = $%d", len(args)))
	}
	if public.Nation != "" {
		args = append(args, public.Nation)
		conditions = append(conditions, fmt.Sprintf("nation = $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int32
	err := p.db.QueryRow("SELECT COUNT(1) FROM provision WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		return nil, err
	}
	if public.Filter.Limit == 0 {
		public.Filter.Limit = defaultLimit
	}

	fmt.Println(query, args)

	args = append(args, public.Filter.Limit, public.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		public := &pb.PublicRes{}
		err = rows.Scan(
			&public.Id,
			&public.FirstName,
			&public.LastName,
			&public.Birthday,
			&public.Gender,
			&public.Nation,
			&public.Party.Id,
			&public.Party.Name,
			&public.Party.Slogan,
			&public.Party.OpenedDate,
			&public.Party.Description,
		)
		if err != nil {
			return nil, err
		}
		res.Public = append(res.Public, public)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (p *PublicRepo) Update(public *pb.PublicUpdate) (*pb.Void, error) {
	res := &pb.Void{}
	query := `UPDATE public SET first_name=$1, last_name=$2, birtday=$3, nation=$4, party_id=$5, updated_at=now() WHERE id = $6`
	_, err := p.db.Exec(query,
		public.FirstName,
		public.LastName,
		public.Birthday,
		public.Nation,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PublicRepo) Delete(id *pb.GetByIdReq) (*pb.Void, error) {
	res := &pb.Void{}
	query := `UPDATE public SET deleted_at=$1 WHERE id = $2`

	_, err := p.db.Exec(query, time.Now().Unix(), id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
