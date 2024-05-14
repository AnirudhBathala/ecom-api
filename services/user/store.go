package user

import (
	"context"
	"fmt"

	"github.com/AnirudhBathala/ecom-api/db"
	"github.com/AnirudhBathala/ecom-api/models"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *db.Postgres
}

func NewStore(db *db.Postgres) *Store{
	return &Store{
		db: db, 
	}
}

func ScanRowIntoUser(rows pgx.Rows) (*models.User,error){
	user:=new(models.User)

	err:=rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err!=nil {
		return nil,err
	}

	return user,nil
}

func (s *Store) GetUserByEmail(email string) (*models.User,error){
	
	rows,err:=s.db.Pool.Query(context.Background(),"SELECT * from users WHERE email= $1 ",email)
	if err!=nil {
		return nil,err
	}
	defer rows.Close()

	u:=new(models.User)
	for rows.Next() {
		u,err=ScanRowIntoUser(rows)
		if err!=nil {
			return nil,err
		}
	}

	if u.ID==0 {
		return nil,fmt.Errorf("user not found")
	}
	return u,nil
}

func (s *Store) GetUserByID(id int) (*models.User,error) {
	rows,err:=s.db.Pool.Query(context.Background(),"SELECT * from users WHERE id= $1 ",id)
	if err!=nil {
		return nil,err
	}
	defer rows.Close()

	u:=new(models.User)
	for rows.Next() {
		u,err=ScanRowIntoUser(rows)
		if err!=nil {
			return nil,err
		}
	}

	if u.ID==0 {
		return nil,fmt.Errorf("user not found")
	}

	return u,nil
}

func (s *Store) CreateUser(user models.User) error{
	query:=`INSERT INTO USERS (firstname,lastname,email,password) VALUES (@firstName,@lastName,@email,@password)`
	args:=pgx.NamedArgs{
		"firstName":user.FirstName,
		"lastName":user.LastName,
		"email":user.Email,
		"password":user.Password,
	}

	_,err:=s.db.Pool.Exec(context.Background(),query,args)
	if err!=nil {
		return fmt.Errorf("unable to create user:%w",err)
	}

	return nil
}