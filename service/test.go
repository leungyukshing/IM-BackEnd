package service

func (this *Handler) Test() {
	this.Ctx.Output.Body([]byte("connection success"))
}
