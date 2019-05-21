package main

var configs SrvcConfigs

func init() {
	configs = SrvcConfigs{
		"echo": SrvcSingleConfig{
			Input: []SrvcServiceIO{
				SrvcServiceIO{
					EventName: "trig.request.get./echo.out.form",
					VarName:   "form",
				},
			},
			Output: []SrvcServiceIO{
				SrvcServiceIO{
					EventName: "trig.response.get./echo.in.code",
					VarName:   "code",
				},
				SrvcServiceIO{
					EventName: "trig.response.get./echo.in.content",
					VarName:   "content",
				},
			},
			Function: WrappedEcho,
		},
	}
}

func main() {
	// consume and publish forever
	ch := make(chan bool)
	SrvcConsumeAndPublish(configs)
	<-ch
}