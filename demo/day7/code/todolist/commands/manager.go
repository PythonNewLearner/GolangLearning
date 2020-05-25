package commands

import (
	"fmt"
	"strconv"
	"todolist/commands/command"
	"todolist/utils/ioutils"
)

type manager struct {
	cmds map[int]*command.Command
}

func newManager() *manager  {
	return &manager{
		cmds: make(map[int]*command.Command),
	}
}

func (mgr *manager)register(name string, callback command.Callback)  {
	mgr.cmds[len(mgr.cmds) + 1] = command.New(name,callback)
}

func (mgr *manager)prompt()  {
	 for i:= 1;i<=len(mgr.cmds);i++{
	 	fmt.Printf("%d %s\n",i,mgr.cmds[i].Name)
	 }
}
func (mgr *manager)get(key int) (command.Callback , error) {
	if cmd,ok :=mgr.cmds[key];ok{
		return cmd.Callback,nil
	}
	return nil,fmt.Errorf("指令不存在")
}

func (mgr *manager)run()  {
	for {
		mgr.prompt()
		key,err := strconv.Atoi(ioutils.Input("请输入指令："))
		if err!=nil{
			ioutils.Error("输入指令错误！")
			continue
		}
		if callback,err := mgr.get(key); err!=nil{
			ioutils.Error(err.Error())
		}else{
			callback()
		}
	}
}

var mgr = newManager()

func Register(name string,callback command.Callback)  {
	mgr.register(name,callback)
}

func Run()  {
	mgr.run()
}