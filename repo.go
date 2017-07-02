package main

// All logic for Git clone and deploy commands

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	git "github.com/daidokoro/go-git"
	"github.com/daidokoro/go-git/plumbing/transport/http"
	"github.com/daidokoro/go-git/plumbing/transport/ssh"
	"github.com/daidokoro/go-git/storage/memory"
	billy "gopkg.in/src-d/go-billy.v2"
	"gopkg.in/src-d/go-billy.v2/memfs"
	yaml "gopkg.in/yaml.v2"
)

// Config - defines gzs3 config file
type Config struct {
	Bucket string `json:"bucket" yaml:"bucket"`
	Key    string `json:"key" yaml:"key"`
}

// Repo used to manage git repo based deployments
type Repo struct {
	URL     string
	fs      *memfs.Memory
	Files   map[string]string
	Config  string
	RSA     string
	User    string
	Secret  string
	zipData *bytes.Buffer
	conf    Config
}

// NewRepo - returns pointer to a new repo struct
func NewRepo(url, user string) (*Repo, error) {
	r := &Repo{
		fs:    memfs.New(),
		Files: make(map[string]string),
		URL:   url,
		RSA:   gitrsa,
	}

	if user != "" {
		r.User = user
	}

	if err := r.clone(); err != nil {
		return r, err
	}

	root, err := r.fs.ReadDir("/")
	if err != nil {
		return r, err
	}

	if err := r.readFiles(root, ""); err != nil {
		return r, err
	}

	// parse config
	if err := r.parseConfig(); err != nil {
		return r, err
	}

	// create zip
	if err := r.createZip(); err != nil {
		return r, err
	}

	return r, nil
}

func (r *Repo) clone() error {
	// memory store for git objects
	store := memory.NewStorage()

	// clone options
	opts := &git.CloneOptions{
		URL:      r.URL,
		Progress: os.Stdout,
	}

	// set authentication
	if err := r.getAuth(opts); err != nil {
		return err
	}

	log.Debug(fmt.Sprintln("calling [git clone] with params:", opts))

	// Clones the repository into the worktree (fs) and storer all the .git
	fmt.Printf("fetching git repo: [%s]\n--\n", r.URL)
	if _, err := git.Clone(store, r.fs, opts); err != nil {
		return err
	}

	fmt.Println("--")

	return nil
}

func (r *Repo) readFiles(root []billy.FileInfo, dirname string) error {
	log.Debug(fmt.Sprintf("writing repo files to memory filesystem [%s]\n", dirname))
	for _, i := range root {
		if i.IsDir() {
			dir, _ := r.fs.ReadDir(i.Name())
			r.readFiles(dir, i.Name())
			continue
		}

		path := filepath.Join(dirname, i.Name())

		out, err := r.fs.Open(path)
		if err != nil {
			return err
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(out)

		// update file map
		r.Files[path] = buf.String()
	}
	return nil
}

func (r *Repo) getAuth(opts *git.CloneOptions) error {
	if strings.HasPrefix(r.URL, "git@") {
		fmt.Println("SSH Source URL detected, attempting to use SSH Keys")

		sshAuth, err := ssh.NewPublicKeysFromFile("git", r.RSA, "")
		if err != nil {
			return err
		}

		opts.Auth = sshAuth
		return nil
	}
	if r.User != "" {
		if r.Secret == "" {
			fmt.Printf(`Password for '%s':`, r.URL)
			p, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			fmt.Printf("\n")

			r.Secret = string(p)
		}
		opts.Auth = http.NewBasicAuth(r.User, r.Secret)
	}

	return nil
}

func (r *Repo) parseConfig() error {
	log.Debug(fmt.Sprintf("checking for %s config file in repo", config))
	// check config is there
	if _, ok := r.Files[config]; !ok {
		return fmt.Errorf("[%s] not found in repo", config)
	}

	c := r.Files[config]
	if err := yaml.Unmarshal([]byte(c), &r.conf); err != nil {
		return err
	}

	return nil
}

func (r *Repo) createZip() error {
	log.Debug("creating ZIP file")
	buf := new(bytes.Buffer)
	zipwriter := zip.NewWriter(buf)

	for f, body := range r.Files {
		zfile, err := zipwriter.Create(f)
		if err != nil {
			return err
		}

		if _, err := zfile.Write([]byte(body)); err != nil {
			return err
		}
	}

	if err := zipwriter.Close(); err != nil {
		return err
	}

	// set zip data
	r.zipData = buf
	return nil
}
