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

func Test_fightProcessor_Act(t *testing.T) {
	type fields struct {
		client    db.Storage
		repoFight repository.FightRepository
		repoUser  repository.UserRepository
		repoCoord repository.CoordinatesRepository
		log       *zap.Logger
	}
	type args struct {
		fight *gen.FightActParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   string
		wantErr bool
	}{
		{
			name: "Existed User",
			fields: fields{
				client: db.NewStorageMock(),
				repoFight: &repository.FightRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return &models.User{
						ID:   9,
						Name: "User99",
					}, nil
				}},
				repoCoord: &repository.CoordinatesRepositoryMock{GetByCoordinatesResponseFn: func(ctx context.Context) (*models.Coordinates, error) {
					return &models.Coordinates{
						UserId: 9,
						ShipId: 4,
						XCoord: 2,
						YCoord: 2,
					}, nil
				}},
				log: zap.NewNop(),
			},
			args: args{fight: &gen.FightActParams{
				UserID: 9,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    true,
			want1:   "ранен",
			wantErr: false,
		},
		{name: "Unknown User",
			fields: fields{
				client: db.NewStorageMock(),
				repoFight: &repository.FightRepositoryMock{SaveResponseFn: func(ctx context.Context) (bool, error) {
					return true, nil
				}},
				repoUser: &repository.UserRepositoryMock{GetResponseFn: func(ctx context.Context) (*models.User, error) {
					return nil, fmt.Errorf("пользователь не найден")
				}},
				log: zap.NewNop(),
			},
			args: args{fight: &gen.FightActParams{
				UserID: 9,
				XCoord: 2,
				YCoord: 2,
			}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &fightProcessor{
				client:    tt.fields.client,
				repoFight: tt.fields.repoFight,
				repoUser:  tt.fields.repoUser,
				repoCoord: tt.fields.repoCoord,
				log:       tt.fields.log,
			}
			got, got1, err := c.Act(tt.args.fight)
			if (err != nil) != tt.wantErr {
				t.Errorf("Act() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Act() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Act() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
