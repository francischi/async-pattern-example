//+build wireinject

package memberModule

import (
	"golang/pkg/helpers"
	impl "golang/pkg/repos/implement"
	"golang/pkg/repos/interfaces"
	"golang/pkg/tokenModule"

	"github.com/google/wire"
)

var MemberRepo = wire.NewSet(impl.NewMemberRepo,wire.Bind(new(interfaces.MemberRepo), new(*impl.MemberRepo)))

func InitMemberController() *MemberController{
	wire.Build(
		NewMemberController ,
		tokenModule.NewTokenService ,  
		NewMemberService , 
		MemberRepo , 
		helpers.NewSqlSession ,
	)
	return &MemberController{}
}