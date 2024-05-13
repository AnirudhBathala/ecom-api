package user

import (
	"context"
	"fmt"

	"github.com/AnirudhBathala/ecom-api/models"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store{
	return &Store{
		db: db, 
	}
}

func (s *Store) GetUserByEmail(email string) (*models.User,error){
	
	rows,err:=s.db.Query(context.Background(),"SELECT * from users WHERE email= $1 ",email)
	if err!=nil {
		return nil,err
	}
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

func ScanRowIntoUser(rows pgx.Rows) (*models.User,error){
	user:=new(models.User)

	err:=rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err!=nil {
		return nil,err
	}

	return user,nil
}

func (s *Store) GetUserByID(id int) (*models.User,error) {
	return &models.User{},nil
}

func (s *Store) CreateUser(user models.User) error{
	return nil
}