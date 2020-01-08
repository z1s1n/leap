package client

import (
	"io/ioutil"
	"net"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// 用户名密码
func DialWithPasswd(addr, user, passwd string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return ssh.Dial("tcp", addr, config)
}

// 通过key文件
func DialWithKey(addr, user, keyfile string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return ssh.Dial("tcp", addr, config)
}

// key+password
func DialWithKeyAndPassword(addr, user, keyfile, password string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(keyfile)
	passwordBytes := []byte(password)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, passwordBytes)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	return ssh.Dial("tcp", addr, config)
}
func GetClient(host map[string]interface{}) (*ssh.Client, error) {
	var (
		cli *ssh.Client
		err error
	)

	switch type_ := host["type"].(string); type_ {
	case "host":
		address := host["address"].(string)
		port := strconv.Itoa(host["port"].(int))
		username := host["username"].(string)
		password := host["password"].(string)
		cli, err = DialWithPasswd(address+":"+port, username, password)
	case "key":
		address := host["address"].(string)
		port := strconv.Itoa(host["port"].(int))
		username := host["username"].(string)
		keyFile := host["keyFile"].(string)
		cli, err = DialWithKey(address+":"+port, username, keyFile)
	case "keyWithPass":
		address := host["address"].(string)
		port := strconv.Itoa(host["port"].(int))
		username := host["username"].(string)
		password := host["password"].(string)
		keyFile := host["keyFile"].(string)
		cli, err = DialWithKeyAndPassword(address+":"+port, username, keyFile, password)
	default:
		cli = nil
		err = nil
	}
	return cli, err
}
