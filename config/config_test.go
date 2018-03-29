package config

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"testing"
)

func TestNewConf(t *testing.T) {
	conf := AppConfInfo{
		AppName:         "bjwt",
		HttpPort:        4009,
		GrpcListen:      ":5009",
		RunMode:         "dev",
		AutoRender:      false,
		CopyRequestBody: true,
		EnableDocs:      true,
		LogLevel:        logs.LevelDebug,
		JwtSalt:         "testsalt",
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", string(data))
	if err := iot.WriteFile("../conf/app_template.yml", data, 0644); err != nil {
		t.Fatal(err)
	}

}

func TestLoadConf(t *testing.T) {
	data, err := LoadConf("../conf/app.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}

func TestLoadConfFromData(t *testing.T) {
	data, err := LoadConfFromData(TestData, "yml")
	if err != nil {
		t.Logf("data:%v", data)
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}
