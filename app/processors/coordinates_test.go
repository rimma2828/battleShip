package processors

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	gen "battleship/app/generated/models"
	"battleship/app/models"
	"battleship/app/repository"
	"battleship/app/system/db"
	"testing"
)

func Test_coordinatesProcessor_Add(t *testing.T) {
	type fields struct {
		client          db.Storage
		repoCoordinates repository.CoordinatesRepository
		repoUser        repository.UserRepository
		repoShip        repository.ShipRepository
		log             *zap.Logger
	}
	type args struct {
		params *gen.CoordinatesAddParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Existed User",
			fields: fields{
				client: db.NewStorageMock(),
				repoCoordinates: &repository.CoordinatesRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return &models.User{
						ID:   9,
						Name: "User99",
					}, nil
				}},
				repoShip: &repository.ShipRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.Ship, error) {
					return &models.Ship{
						ID:     9,
						Length: 4,
						Count:  4,
					}, nil
				}},
				log: zap.NewNop(),
			},
			args: args{params: &gen.CoordinatesAddParams{
				UserID: 9,
				ShipID: 1,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    true,
			wantErr: false,
		},
		{name: "Unknown User",
			fields: fields{
				client: db.NewStorageMock(),
				repoCoordinates: &repository.CoordinatesRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return nil, fmt.Errorf("пользователь не найден")
				}},
				repoShip: &repository.ShipRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.Ship, error) {
					return &models.Ship{
						ID:     9,
						Length: 4,
						Count:  4,
					}, nil
				}},
				log: zap.NewNop(),
			},
			args: args{params: &gen.CoordinatesAddParams{
				UserID: 9,
				ShipID: 1,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    false,
			wantErr: true,
		},
		{name: "Existed Ship",
			fields: fields{
				client: db.NewStorageMock(),
				repoCoordinates: &repository.CoordinatesRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoShip: &repository.ShipRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.Ship, error) {
					return &models.Ship{
						ID:     9,
						Length: 4,
						Count:  4,
					}, nil
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return &models.User{
						ID:   9,
						Name: "User99",
					}, nil
				}},
				log: zap.NewNop(),
			},
			args: args{params: &gen.CoordinatesAddParams{
				UserID: 9,
				ShipID: 1,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    true,
			wantErr: false,
		},
		{name: "Unknown Ship",
			fields: fields{
				client: db.NewStorageMock(),
				repoCoordinates: &repository.CoordinatesRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoShip: &repository.ShipRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.Ship, error) {
					return nil, fmt.Errorf("корабль не найден")
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return &models.User{
						ID:   9,
						Name: "User99",
					}, nil
				}},
				log: zap.NewNop(),
			},
			args: args{params: &gen.CoordinatesAddParams{
				UserID: 9,
				ShipID: 1,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coordinatesProcessor{
				client:          tt.fields.client,
				repoCoordinates: tt.fields.repoCoordinates,
				repoUser:        tt.fields.repoUser,
				repoShip:        tt.fields.repoShip,
				log:             tt.fields.log,
			}
			got, err := c.Add(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
