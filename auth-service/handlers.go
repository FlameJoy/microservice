package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"microsvc/auth-service/proto"
	"microsvc/common/utils"
	"microsvc/data"
	"microsvc/storage"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	logger  *utils.CustomLogger
	storage storage.Storage
}

func (s *AuthServer) UserExist(username string) (bool, error) {
	query := "SELECT 1 FROM users WHERE username = $1 LIMIT 1"

	data, err := s.storage.GetData(query, username)
	if err != nil {
		s.logger.Error("Can't get data from storage: %v", err)
		return false, err
	}

	if len(data) > 0 {
		return true, nil
	}

	return false, nil
}

func (s *AuthServer) EmailExist(email string) (bool, error) {
	selectQuery := "SELECT 1 FROM users WHERE email = $1 LIMIT 1"

	data, err := s.storage.GetData(selectQuery, email)
	if err != nil {
		s.logger.Error("Can't get data from storage: %v", err)
		return false, err
	}

	if len(data) > 0 {
		return true, nil
	}

	return false, nil
}

func (s *AuthServer) SendVerLink(emailVerPswd, username, email string) error {
	var domName string
	var serveLoc = os.Getenv("SERVER_LOC")
	var useHTTPS = os.Getenv("USE_HTTPS")
	var domainName = os.Getenv("DOMAIN_NAME")

	if serveLoc == "local" {
		domName = "http://localhost:8080"
	} else {
		if useHTTPS == "true" {
			domName = "https://" + domainName
		} else {
			domName = "http://" + domainName
		}
	}
	subject := "Email Verification"
	HTMLbody := `<html><h1>CLick link to verify email</h1><a href="` + domName + `/verify-email/` + username + `/` + emailVerPswd + `">Click here to verify email</a></html>`

	if err := s.SendEmail(subject, HTMLbody, email); err != nil {
		s.logger.Error("Can't send verification email")
		return err
	}

	return nil
}

func (s *AuthServer) SendEmail(subject, HTMLbody, email string) error {
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	fromEmail := os.Getenv("from_Email_Addr")
	smtpPswd := os.Getenv("SMTP_pswd")
	entityName := os.Getenv("entity_name")
	auth := smtp.PlainAuth("", fromEmail, smtpPswd, host)
	msg := []byte(
		"From: " + entityName + ": <" + fromEmail + ">\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-type: text/html; charset=\"utf8\";\r\n" +
			"\r\n" +
			HTMLbody)

	if err := smtp.SendMail(address, auth, fromEmail, to, msg); err != nil {
		s.logger.Error("smtp.SendMail error")
		return err
	}

	return nil
}

func (s *AuthServer) NewUser(req *proto.RegRequest) error {
	var u data.User

	// Create hash from pswd
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	u.PswdHash = string(hash)
	if err != nil {
		s.logger.Error("Hash generation error: %v", err)
		return err
	}

	// Init rand source
	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	randBytes := make([]byte, 64)

	// Create a random slice of chars to create emailVerPswd
	for i := 0; i < 64; i++ {
		randBytes[i] = chars[rng.Intn(len(chars)-1)]
	}
	emailVerPswd := string(randBytes)
	b, err := bcrypt.GenerateFromPassword([]byte(emailVerPswd), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("GenerateFromPassword error: %v", err)
		return err
	}
	u.VerHash = string(b)

	// Create user timeout after 24 hours
	u.TimeoutAt = time.Now().Local().AddDate(0, 0, 1)

	// Save user
	query := "INSERT INTO users (username, email, pswd_hash, ver_hash, timeout_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err = s.storage.ExecuteQuery(query, req.Username, req.Email, u.PswdHash, u.VerHash, u.TimeoutAt, time.Now().Local(), time.Now().Local()); err != nil {
		s.logger.Error("Storage error: %v", err)
		return err
	}

	// Send ver email
	// if err = s.SendVerLink(emailVerPswd, req.Username, req.Email); err != nil {
	// 	s.logger.Error("Error: %v", err)
	// 	return err
	// }

	return nil
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegRequest) (*proto.RegResponse, error) {
	s.logger.Info("auth svc: starts gRPC server Register func")

	// Check if user already exist
	exists, err := s.UserExist(req.Username)
	if err != nil {
		return &proto.RegResponse{Message: "UserExist error"}, err
	}
	if exists {
		return &proto.RegResponse{Message: "User already exist"}, errors.New("user already exist")
	}

	// Check if email already exist
	exists, err = s.EmailExist(req.Email)
	if err != nil {
		return &proto.RegResponse{Message: "EmailExist error"}, err
	}
	if exists {
		return &proto.RegResponse{Message: "Email already using"}, errors.New("email already using")
	}

	// Create new user
	if err = s.NewUser(req); err != nil {
		return &proto.RegResponse{Message: "NewUser error"}, err
	}

	s.logger.Info("User %s registered successfully", req.Username)
	return &proto.RegResponse{Message: "User registered successfully"}, nil
}

func (s *AuthServer) NewToken(userID int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expTimeHours, err := strconv.Atoi(os.Getenv("TOKEN_EXP_TIME"))
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Local().Add(time.Hour * time.Duration(expTimeHours)).Unix()
	claims["user_id"] = strconv.FormatInt(userID, 10)

	secretKey := os.Getenv("SECRET")
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	s.logger.Info("auth svc: starts gRPC server Login func")

	query := "SELECT id, username, pswd_hash FROM users WHERE username = $1 AND deleted_at IS NULL LIMIT 1"
	data, err := s.storage.GetData(query, req.Username)
	if err != nil {
		return &proto.LoginResponse{Message: "storage.GetData error"}, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}

	pswdHash, ok := data[0]["pswd_hash"]
	if !ok {
		return &proto.LoginResponse{Message: "storage: missing user pswdHash data"}, errors.New("storage: missing user pswdHash data")
	}

	err = bcrypt.CompareHashAndPassword([]byte(pswdHash.(string)), []byte(req.Password))
	if err != nil {
		return &proto.LoginResponse{Message: "bcrypt.CompareHashAndPassword error"}, err
	}

	userID, ok := data[0]["id"]
	if !ok {
		return &proto.LoginResponse{Message: "storage: missing user id data"}, errors.New("storage: missing user id data")
	}

	id, ok := userID.(int64)
	if !ok {
		return &proto.LoginResponse{Message: "userID interface conversion error"}, errors.New("interface {} is not int64")
	}

	token, err := s.NewToken(id)
	if err != nil {
		return &proto.LoginResponse{Message: "NewToken error"}, err
	}

	return &proto.LoginResponse{Message: "Access granted", Token: token}, nil
}
