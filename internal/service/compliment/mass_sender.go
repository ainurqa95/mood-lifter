package compliment

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/service"
	"golang.org/x/sync/errgroup"
	"log"
)

const (
	DEFAULT_LIMIT = 1000
	WORKERS_COUNT = 100
)

type MassSender struct {
	bot         bots.BotManager
	userService service.UserService
}

func NewMassSender(
	bot bots.BotManager,
	userService service.UserService,
) MassSender {
	return MassSender{bot: bot, userService: userService}
}

func (m *MassSender) SendMassCompliments(ctx context.Context, periodTypes []int) error {
	toSendUsers := make(chan model.UserInfo, 10000)
	offset := 0
	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		for {
			users, err := m.userService.GetUsersByPeriodWithOffset(ctx, periodTypes, DEFAULT_LIMIT, offset)
			if err != nil {
				return err
			}
			if len(users) == 0 {
				close(toSendUsers)
				break
			}
			for _, usr := range users {
				toSendUsers <- usr
			}
			offset += DEFAULT_LIMIT
		}
		return nil
	})

	for i := 0; i < WORKERS_COUNT; i++ {
		eg.Go(func() error {
			for user := range toSendUsers {
				err := m.bot.SendCompliment(gCtx, user.Name, user.ChatID)
				if err != nil {
					log.Println("error sending compliment:", err)
				}
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
