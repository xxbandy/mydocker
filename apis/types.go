package apis

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfEnv struct {
	Logpath    string `yaml:"logpath"`
	Conconfdir string `yaml:"configdir"`
	Pauseimage string `yaml:"pauseimage"`
	DataPath   string `yaml:"datapath,omitempty"`
}

const config string = "/etc/JFDocker/config.yml"

var LogPath, ConConfDir, Pauseimage, DataPath string

func init() {
	if _, err := os.Stat(config); err == nil {
		content, _ := ioutil.ReadFile(config)
		env := ConfEnv{}
		if err := yaml.Unmarshal(content, &env); err == nil {
			LogPath = env.Logpath
			ConConfDir = env.Conconfdir
			Pauseimage = env.Pauseimage
			DataPath = env.DataPath
		} else {
			fmt.Println("配置文件有误，使用默认配置")
		}
	} else {
		LogPath = "/tmp/Logs/"
		ConConfDir = "containers_config"
		Pauseimage = "xxbandy123/k8s-pause"
		DataPath = "/export/"
	}
}

/*
//注意LogPath必须以/结尾
const LogPath string = "/tmp/Logs/"
const ConConfDir string = "containers_config"
const Pauseimage string = "xxbandy123/k8s-pause"
*/

const Netnamespace string = "/var/run/netns/"

const Usages string = `Usage:
    JFDocker run ywid,sn,image,container,ip,mask,gateway,cpu,mem,label,vnet
    JFDocker rungpu ywid,sn,image,container,ip,mask,gateway,cpu,mem,label,vnet,gpus
    JFDocker update container,images
    JFDocker updategpu container,images
    JFDocker resize container,cpu,4/container,mem,2048
    JFDocker delete container(gpu or cpu)
    JFDocker  rebuilt container(cpu)

`

type ConSpec struct {
	Ywid string `json:"ywid"`
	SN   string `json:"sn"`
	Mem  string `json:"mems"`
	Cpus string `json:"cpus"`
	Gpus string `json:"gpus,omitempty"`
}

type NetSpec struct {
	Ipv4    string `json:"ipaddress"`
	Mask    string `json:"mask"`
	Gateway string `json:gateway`
	Vnet    string `json:"vnet"`
}

type Netns struct {
	Cid   string `json:pausecontainerid,omitempty`
	Pid   string `json:pausecontainerpid,omitempty`
	Netns string `json:netns,omitempty`
	Net   NetSpec
}

type JFDocker struct {
	Appname string `json:"containername"`
	Image   string `json:"Image"`
	Conspec ConSpec
	Netspec NetSpec
}

type RspJFDocker struct {
	Result  uint   `json:"result,omitempty"`
	Appname string `json:"containername,omitempty"`
	ConID   string `json:"containerid,omitempty"`
	Ipv4    string `json:"ipaddress,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
	ErrInfo string `json:"errinfo,omitempty"`
}

type JFDockererr struct {
	JFDockerVersion string `json:"JFDockerversion,omitempty"`
	Code            uint   `json:"result,omitempty`
	Msg             string `json:"msg,omitempty"`
	Details         string `json:"details,omitempty"`
}
