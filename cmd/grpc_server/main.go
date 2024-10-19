package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=postgres user=postgres password=postgres sslmode=disable"
)

type server struct {
	desc.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

type AuthUser struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	serverInstance := &server{
		pool: pool,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, serverInstance)

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Create - create new user
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	builderInsert := sq.Insert("auth_user").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(req.GetName(), req.GetPassword(), req.GetEmail(), req.GetRole()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var userID int
	err = s.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert user: %v", err)
	}

	log.Printf("Inserted user with id: %d", userID)

	return &desc.CreateResponse{
		Id: int64(userID),
	}, nil
}

// Get - get user by id
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	builderSelect := sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("auth_user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		OrderBy("id ASC")

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	user := AuthUser{}
	var updatedAt pq.NullTime
	err = s.pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User with id %d not found", req.GetId())
		}
		return nil, status.Errorf(codes.Internal, "Failed to query user: %v", err)
	}

	var roleMap = map[string]desc.Enum{
		"USER":  desc.Enum_USER,
		"ADMIN": desc.Enum_ADMIN,
	}

	var role desc.Enum
	if val, ok := roleMap[user.Role]; ok {
		role = val
	} else {
		role = desc.Enum_UNKNOWN
	}
	var updatedAtPB *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtPB = timestamppb.New(updatedAt.Time)
	} else {
		updatedAtPB = nil
	}

	return &desc.GetResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Role:      role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAtPB,
	}, nil
}

// Update - update user by id
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	builderSelect := sq.Select("COUNT(*)").
		From("auth_user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	var count int
	err = s.pool.QueryRow(ctx, selectQuery, args...).Scan(&count)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to select user: %v", err)
	}

	if count == 0 {
		return nil, status.Errorf(codes.NotFound, "User with id %d not found", req.GetId())
	}

	builderUpdate := sq.Update("auth_user").
		PlaceholderFormat(sq.Dollar).
		Set("name", req.Name.GetValue()).
		Set("role", req.GetRole()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("Failed to build query: %v", err)
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	log.Printf("Updated %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

// Delete - delete user by id
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	builderSelect := sq.Select("COUNT(*)").
		From("auth_user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	var count int
	err = s.pool.QueryRow(ctx, selectQuery, args...).Scan(&count)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to select user: %v", err)
	}

	if count == 0 {
		return nil, status.Errorf(codes.NotFound, "User with id %d not found", req.GetId())
	}

	builderDelete := sq.Delete("auth_user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build delete query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete user: %v", err)
	}

	log.Printf("User with id %d deleted", req.GetId())

	return &emptypb.Empty{}, nil
}
