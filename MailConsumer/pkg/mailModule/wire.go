//+build wireinject

package mailModule

import (
	"github.com/google/wire"
	"mailConsumer/pkg/repos/implement"
)

func InitialMailController() *MailController{
	wire.Build(
		NewMailController,
		NewMailService,
		implement.NewMailRepo,
	)
	return &MailController{}
}