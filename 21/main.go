package main

import "fmt"

type Connector interface {
	getConnection() string
}
type Config struct {
	Scheme string
	Host   string
	Port   int
}

func (config *Config) getConnection() string {
	return fmt.Sprintf("%v://%v:%v", config.Scheme, config.Host, config.Port)
}

type OldConfig struct {
	Scheme   string
	Host     string
	Port     int
	clientID string
}

func (oldConfig *OldConfig) getString() string {
	return fmt.Sprintf("%v://%v:%v/%v", oldConfig.Scheme, oldConfig.Host, oldConfig.Port, oldConfig.clientID)
}

type Adapter struct {
	oldConfig *OldConfig
}

func (adapter *Adapter) getConnection() string {
	return adapter.oldConfig.getString()
}

func showConnection(connector Connector) {
	fmt.Println(connector.getConnection())
}
func main() {
	config := Config{Scheme: "example", Host: "localhost", Port: 22}
	oldConfig := OldConfig{Scheme: "example", Host: "myhost", Port: 80, clientID: "clientExample"}
	adapter := Adapter{&oldConfig}
	/*Теперь мы может получить старую и новую строку подлючения в любом месте без
	необходимости вызывать отдельные методы у каждого объекта,
	вместо этого пользуясь единым методом и интерфейсом*/
	showConnection(&config)
	showConnection(&adapter)
}
