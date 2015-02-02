package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/dchest/uniuri"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	server      = flag.String("server", "", "server ip")
	username    = flag.String("username", "", "username")
	password    = flag.String("password", "", "password")
	port        = flag.Int("port", 22, "port")
	upload_path = flag.String("upload_path", "", "upload path")
	upload_url  = flag.String("upload_url", "", "upload url")
)

func main() {
	flag.Parse()

	var auths []ssh.AuthMethod
	if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))

	}
	auths = append(auths, ssh.Password(*password))

	config := ssh.ClientConfig{
		User: *username,
		Auth: auths,
	}

	addr := fmt.Sprintf("%s:%d", *server, *port)
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalf("Can't connect to ssh.")
	}

	//create ssh client
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Can't create ssh session.")
	}

	defer conn.Close()

	//create sftp client
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Can't create sftp client.")
	}

	defer client.Close()

	//get date
	currenttime := time.Now().Local()
	year := strconv.Itoa(currenttime.Year())
	month := strconv.Itoa(int(currenttime.Month()))
	day := strconv.Itoa(currenttime.Day())

	//create dir
	path := fmt.Sprintf("%s/%s/%s/%s/", *upload_path, year, month, day)

	if err := session.Run("mkdir -p " + path); err != nil {
		log.Fatalf("Can't create directory.")
	}

	//create file
	filename := uniuri.NewLen(32)
	f, err := client.Create(path + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//write content
	if _, err := io.Copy(f, os.Stdin); err != nil {
		log.Fatalf("Can't upload image")
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s.png", *upload_url, year, month, day, filename)
	clipboard.WriteAll(url)
}
