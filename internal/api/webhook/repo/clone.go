package repo

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"terrapak/internal/api/webhook/helpers"
)

type CloneOptions struct {
	Url string
	Sha string
	Branch string
	OutputDir string
}

func Clone(installationId int, opts CloneOptions) (error) {
	cmd := helpers.Cmd{Dir: opts.OutputDir}
	token, err := GetAccessToken(installationId); if err != nil {
		slog.Error(err.Error())
		return err
	}
	cloneUrl := createCloneUrl(opts.Url, token)
	stdout, _ := cmd.Execute("git", "clone", cloneUrl, "./"); if stdout != "" {
		slog.Error(stdout)
		return errors.New(stdout)
	}

	if opts.Branch != "" && opts.Sha != "" {
		return errors.New("cannot clone with both branch and sha")
	
	}

	if opts.Sha != "" {
		stdout, err = cmd.Execute("git", "checkout", opts.Sha); if stdout != "" || err != nil{
			slog.Error(err.Error())
			return errors.New(stdout)
		}
		fmt.Println("Checked out to ", opts.Sha)
	}

	if opts.Branch != "" {
		stdout, err = cmd.Execute("git", "checkout", opts.Branch); if stdout != "" || err != nil{
			slog.Error(err.Error())
			return errors.New(stdout)
		}
		fmt.Println("Checked out to ", opts.Branch)
	}
			

	return nil
}

func createCloneUrl(url string, token string) string {
	root_url := strings.Replace(url, "https://", "", 1)
	return "https://oauth2:" + token + "@" + root_url
}