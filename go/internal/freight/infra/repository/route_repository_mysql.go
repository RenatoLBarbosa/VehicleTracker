package repository

import (
	"database/sql"

	"github.com/RenatoLBarbosa/RastreabilidadeVeicular/internal/freight/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	sql := "INSERT INTO routes (id, name, distance, status, freight_price) VALUES (?,?,?,?,?)"
	_, err := r.db.Exec(sql, route.ID, route.Name, route.Distance, route.Status, route.FreightPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMysql) FindByID(id string) (*entity.Route, error) {
	sqlSmt := "SELECT id, name, distance, status, freight_price, started_at, finished_at FROM routes WHERE ID = ?"
	row := r.db.QueryRow(sqlSmt, id)

	var started_at, finished_at sql.NullTime
	var route entity.Route

	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		&started_at,
		&finished_at,
	)
	if err != nil {
		return nil, err
	}

	if started_at.Valid {
		route.StartedAt = started_at.Time
	}
	if finished_at.Valid {
		route.FinishedAt = finished_at.Time
	}

	return &route, nil
}

func (r *RouteRepositoryMysql) Update(route *entity.Route) error {
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	finishedAt := route.FinishedAt.Format("2006-01-02 15:04:05")
	sql := "UPDATE routes SET status = ?, freight_price= ?, started_at= ?, finished_at= ? WHERE id = ?"
	_, err := r.db.Exec(sql, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)
	if err != nil {
		return err
	}
	return nil
}
