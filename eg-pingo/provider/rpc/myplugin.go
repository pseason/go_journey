package rpc

type MyPlugin struct {
}

func NewMyPlugin() *MyPlugin {
	return new(MyPlugin)
}

func (mp *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "hello," + name
	return nil
}
