package dataprovider

import (
	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
	"login-micro/entities"
	dbmodel "login-micro/dataprovider/model"
	"context"
)

const createUser = "INSERT INTO users VALUES(nextval('u_id_seq'), $1, $2, $3, $4, $5) RETURNING u_id"
const retrieveUserByID = "SELECT * FROM users WHERE u_id=$1"
const retrieveAllUsers = "SELECT * FROM users"

type Postgresql struct {
	Conn *sqlx.DB
	Context context.Context
}

func NewPostgresConn(conn *dbmodel.ConnectionInfo) *Postgresql {
	var postgres Postgresql
	postgres.CreateConnection(conn)
	return &postgres
}

func (p *Postgresql) CreateConnection(conn *dbmodel.ConnectionInfo) Database {
	p.Conn = sqlx.MustConnect("pgx", conn.ToConnectionChain())
	return p
}

func (p *Postgresql) SaveUser(user entities.User) (int64, error) {
	var id int64
	result := p.Conn.QueryRowx(createUser, user.Name, user.Lastname, user.Password, user.Email, user.Number)
	err := result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Postgresql) DeleteUser(id int) error {
	return nil
}

func (p *Postgresql) EditUser(user entities.User) error {
	return nil
}

func (p *Postgresql) GetUserByID(id int) (*entities.User, error) {
	var user dbmodel.User
	userRow := p.Conn.QueryRowx(retrieveUserByID, id)
	err := userRow.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return user.ToUserEntity(), nil
}

func (p *Postgresql) GetUserByName(name string) *entities.User {
	return &entities.User{}
}

func (p *Postgresql) GetAllUsers() (*[]dbmodel.User, error) {
	var users []dbmodel.User
	rows, err := p.Conn.Queryx(retrieveAllUsers)
	defer rows.Close()
	for rows.Next() {
		var user dbmodel.User
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}