package client

import (
	"io/ioutil"
	"net"
	"strconv"

	"github.com/z1son/leap/config"
	"golang.org/x/crypto/ssh"
)

// 用户名密码
func DialWithPasswd(auth *config.AllConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: auth.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(auth.Password),
			ssh.KeyboardInteractive(keyboardInteractivePassword(auth.GoogleAuthToken)),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return ssh.Dial("tcp", auth.Address+":"+strconv.Itoa(auth.Port), config)
}

// 通过key文件
func DialWithKey(auth *config.AllConfig) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(auth.KeyFile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: auth.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.KeyboardInteractive(keyboardInteractivePassword(auth.GoogleAuthToken)),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return ssh.Dial("tcp", auth.Address+":"+strconv.Itoa(auth.Port), config)
}

// key+password
func DialWithKeyAndPassword(auth *config.AllConfig) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(auth.KeyFile)
	passwordBytes := []byte(auth.Password)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, passwordBytes)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: auth.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.KeyboardInteractive(keyboardInteractivePassword(auth.GoogleAuthToken)),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	return ssh.Dial("tcp", auth.Address+":"+strconv.Itoa(auth.Port), config)
}
func GetClient(host config.AllConfig) (*ssh.Client, error) {
	var (
		cli *ssh.Client
		err error
	)
	switch type_ := host.Type; type_ {
	case "host":
		cli, err = DialWithPasswd(&host)
	case "key":
		cli, err = DialWithKey(&host)
	case "keyWithPass":
		cli, err = DialWithKeyAndPassword(&host)
	default:
		cli = nil
		err = nil
	}
	return cli, err
}
