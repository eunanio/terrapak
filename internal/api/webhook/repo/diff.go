package repo

import (
	"fmt"
	"terrapak/internal/api/webhook/helpers"
	"terrapak/internal/config/mid"
	"terrapak/internal/db/services"
)

func DiffModule(localpath, modulePath, base_branch string) bool {
	cmd := helpers.Cmd{Dir: localpath}
	stdout, _ := cmd.Execute("git", "diff", "--compact-summary", "HEAD", fmt.Sprintf("origin/%s", base_branch), "--", modulePath)

	return stdout != ""
}

// returns has_changes, createOrReplace (0 for create, 1 for replace), error
func DiffFiles(mid mid.MID, modulePath string) (bool, int, error) {
	ms := services.ModulesService{}
	module := ms.Find(mid)
	hash, err := helpers.HashFiles(modulePath); if err != nil {
		return false,-1, err
	}

	if module.SHAChecksum == "" {
		return true,0, nil
	}

	if module.SHAChecksum != hash {
		return true,1, nil
	}

	return false,-1,nil
}