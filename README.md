# lazyporker
基于go语言的 扑克服务器框架   a easy poker service base on golang

结构组成

登陆服务器 ---登陆凭证管理---游戏服务器---玩家管理
    |                             |			 |
注册/登陆					 根据凭证登陆	 |
										   大厅
											 |
										    桌子
											
全部通过channel 减少线程间同步  

应用案例：余杭大板同服务器 www.yhdbt.pw