package init

import (
	"todolist/commands"
	"todolist/controllers"
)

func init()  {
	commands.Register("退出",controllers.Logout)
}
