/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

type GameRule interface {
	IsGameOver(info interface{}) (result interface{}) //游戏是否结束
}
