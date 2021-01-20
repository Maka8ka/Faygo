package command

// 命令行接口
type Commander interface {
	// 执行命令行并返回结果
	// args: 命令行参数
	// return: 进程的pid, 命令行结果, 错误消息
	Exec(args ...string) (int, string, error)

	// 异步执行命令行并通过channel返回结果
	// stdout: chan结果
	// args: 命令行参数
	// return: 进程的pid
	// exception: 协程内的命令行发生错误时,会panic异常
	ExecAsync(stdout chan string, args ...string) int

	// 执行命令行(忽略返回值)
	// args: 命令行参数
	// return: 错误消息
	ExecIgnoreResult(args ...string) error
}
