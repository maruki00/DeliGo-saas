package app

import (
	"context"
	"deligo/cmd/user/configs"
	grpc_services "deligo/internal/user/app/grpc/services"
	"deligo/internal/user/domain/contracts"
	"deligo/internal/user/infra/repositories"
	pkgPostgres "deligo/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db          *pkgPostgres.PGHandler
	UserRepo    contracts.IUserRepository
	ProfileRepo contracts.IProfileRepository
	UserSVC     *grpc_services.UserService
	ProfileSVC  *grpc_services.ProfileService
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *configs.Config) (*App, func(), error) {

	fmt.Println("dsn : ", cfg.Postgres.Dsn)
	db, err := pkgPostgres.NewPG(cfg.Postgres.Dsn)
	if err != nil {
		return nil, func() {}, err
	}

	userRepo := repositories.NewUserRepository(*db)
	userSVC := grpc_services.NewUserService(userRepo)

	profileRepo := repositories.NewProfileRepository(*db)
	profileSVC := grpc_services.NewProfileService(profileRepo)

	app := &App{
		db:          db,
		UserRepo:    userRepo,
		ProfileRepo: profileRepo,
		UserSVC:     userSVC,
		ProfileSVC:  profileSVC,
	}

	return app, func() { _ = db.DB.Close() }, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.deligo) {

	for {
		select {
		case <-ctx.Done():
			slog.Info("Shuting Down the client.")
			break
		default:
			slog.Info("default interception ....")
		}
	}
	// forever := struct{}{}
	// <-forever
}
