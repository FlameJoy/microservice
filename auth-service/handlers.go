package main

import (
	"context"
	"fmt"
	"math/rand"
	"microsvc/api-gateway/data"
	"microsvc/auth-service/proto"
	"microsvc/common/utils"
	"microsvc/storage"
	"net/smtp"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	logger  *utils.CustomLogger
	storage storage.Storage
}

func (s *AuthServer) UserExist(username string) (bool, error) {
	query := "SELECT 1 FROM users WHERE name = $1 LIMIT 1"

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
	// Send email
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
	err := s.SendEmail(subject, HTMLbody, email)
	if err != nil {
		fmt.Println("Can't send verification email")
		return err
	}
	return nil
}

func (s *AuthServer) SendEmail(subject, HTMLbody, email string) error {
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	var fromEmail = os.Getenv("from_Email_Addr")
	var smtpPswd = os.Getenv("SMTP_pswd")
	var entityName = os.Getenv("entity_name")
	auth := smtp.PlainAuth("", fromEmail, smtpPswd, host)
	msg := []byte(
		"From: " + entityName + ": <" + fromEmail + ">\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-type: text/html; charset=\"utf8\";\r\n" +
			"\r\n" +
			HTMLbody)
	err := smtp.SendMail(address, auth, fromEmail, to, msg)
	if err != nil {
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
	u.Timeout = time.Now().Local().AddDate(0, 0, 1)

	// Save user
	query := "INSERT INTO TABLE users (username, pswdHash, email, verHash, timeout, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if err = s.storage.ExecuteQuery(query, req.Username, u.PswdHash, req.Email, u.VerHash, u.Timeout, time.Now().Local(), time.Now().Local()); err != nil {
		s.logger.Error("Storage error: %v", err)
		return err
	}

	// Send ver email
	if err = s.SendVerLink(emailVerPswd, req.Username, req.Email); err != nil {
		s.logger.Error("Error: %v", err)
		return err
	}

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
		return &proto.RegResponse{Message: "User already exist"}, nil
	}

	// Check if email already exist
	exists, err = s.EmailExist(req.Email)
	if err != nil {
		return &proto.RegResponse{Message: "EmailExist error"}, err
	}
	if exists {
		return &proto.RegResponse{Message: "Email already using"}, nil
	}

	// Create new user
	if err = s.NewUser(req); err != nil {
		return &proto.RegResponse{Message: "NewUser error"}, err
	}

	s.logger.Info("User %s registered successfully", req.Username)
	return &proto.RegResponse{Message: "User registered successfully"}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	if req.Username == "Ilya" && req.Password == "qwerty123" {
		return &proto.LoginResponse{Token: "some_token"}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}
