package mailModule

type SendDto struct{
	Context string
	To      string
}

func(dto *SendDto) Check() error {
	return nil
}
