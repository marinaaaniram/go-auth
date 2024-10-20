package main

import (
	"context"
	"flag"
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

	"github.com/marinaaaniram/go-auth/internal/config"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	serverInstance := &server{
		pool: pool,
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
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

// CreateUser - create new user
func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
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

	return &desc.CreateUserResponse{
		Id: int64(userID),
	}, nil
}

// GetUser - get user by id
func (s *server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	builderSelect := sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("auth_user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		OrderBy("id ASC")

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	var (
		userId            int
		userName          string
		userEmail         string
		userRole          string
		createdAt         time.Time
		updatedAtNullable pq.NullTime
	)

	err = s.pool.QueryRow(ctx, query, args...).Scan(&userId, &userName, &userEmail, &userRole, &createdAt, &updatedAtNullable)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, status.Errorf(codes.NotFound, "User with id %d not found", req.GetId())
		}
		return nil, status.Errorf(codes.Internal, "Failed to query user: %v", err)
	}

	var roleMap = map[string]desc.Enum{
		"USER":  desc.Enum_USER,
		"ADMIN": desc.Enum_ADMIN,
	}

	var role desc.Enum
	if val, ok := roleMap[userRole]; ok {
		role = val
	} else {
		role = desc.Enum_UNKNOWN
	}

	var updatedAt *timestamppb.Timestamp
	if updatedAtNullable.Valid {
		updatedAt = timestamppb.New(updatedAtNullable.Time)
	}

	return &desc.GetUserResponse{
		Id:        int64(userId),
		Name:      userName,
		Email:     userEmail,
		Role:      role,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: updatedAt,
	}, nil
}

// UpdateUser - update user by id
func (s *server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
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
func (s *server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
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
