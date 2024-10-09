package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
)


type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`  // 'user' or 'admin'
	CreatedAt time.Time `json:"created_at"`
}

// Hash the password before saving to the database
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}


func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}


func (u *User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"userId": u.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("SECRET_KEY environment variable is not set")
	}
	return token.SignedString([]byte(secretKey))
}


func (u *User) Create(db *sql.DB) error {
	u.CreatedAt = time.Now()

	_, err := db.Exec("INSERT INTO users (username, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5)",
		u.Username, u.Email, u.Password, u.Role, u.CreatedAt)
	return err
}


func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, role, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, password, role, created_at FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func GetUserById(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, role, created_at FROM users WHERE id = $1", id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (u *User) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET username = $1, email = $2, role = $3 WHERE id = $4",
		u.Username, u.Email, u.Role, u.ID)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
