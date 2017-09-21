/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

const (
	cmd_query_version    = 0x10001
	cmd_player_net_error = 0x10002
)

type Process struct {
}

var GProcess = &Process{}

// 处理各种请求
func (this *Process) ProcessCmd(cmd int, info interface{}, p *PlayerEx) error {
	switch cmd {
	case cmd_query_version:
		// do same thing
	case cmd_player_net_error:
		return this.prcess_error(p)
	default:
		return this.prcess_error(p)
	}
	return nil
}

func (this *Process) prcess_error(p *PlayerEx) error {
	p.CreateTask(task_player_cmd_kick, "", p, nil)
	p.desk.CreateTask(task_desk_cmd_leave, "", p, nil)
	GHall.CreateTask(task_hall_cmd_player_leave, "", p, nil)
	return nil
}
