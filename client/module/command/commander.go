package command


type Commander interface {
	
	
	
	Exec(args ...string) (int, string, error)

	
	
	
	
	
	ExecAsync(stdout chan string, args ...string) int

	
	
	
	ExecIgnoreResult(args ...string) error
}
