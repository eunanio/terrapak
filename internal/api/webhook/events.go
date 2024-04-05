package webhook

import (
	"fmt"
	"log/slog"
	"os"
	"terrapak/internal/api/modules"
	"terrapak/internal/api/webhook/hclconfig"
	"terrapak/internal/api/webhook/helpers"
	"terrapak/internal/api/webhook/repo"
	"terrapak/internal/config/mid"
)

func onSyncPullRequest(event PullRequestEvent) {
	dir, err := os.MkdirTemp("","terrapak"); if err != nil {
		slog.Error(err.Error())
	}
	defer os.RemoveAll(dir)
	opts := repo.CloneOptions{
		Url: event.PullRequest.Head.Repo.CloneURL,
		OutputDir: dir,
		Branch: event.PullRequest.Head.Ref,
	}

	err = repo.Clone(event.Installation.ID, opts); if err != nil {
		slog.Error(err.Error())
	}

	err = syncRepo(dir, event); if err != nil {
		slog.Error(err.Error())
	}
}

func onClosePullRequest(event PullRequestEvent) {
	dir, err := os.MkdirTemp("","terrapak"); if err != nil {
		slog.Error(err.Error())
	}
	configpath := fmt.Sprintf("%s/terrapak.hcl", dir)
	defer os.RemoveAll(dir)
	fmt.Printf("Dir and configpath %s %s\n", dir, configpath)
	opts := repo.CloneOptions{
		Url: event.PullRequest.Head.Repo.CloneURL,
		OutputDir: dir,
		Branch: event.PullRequest.Head.Ref,
	}

	err = repo.Clone(event.Installation.ID, opts); if err != nil {
		slog.Error(err.Error())
	}

	repo_config, err  := hclconfig.Load(&configpath); if err != nil {
		slog.Error(err.Error())
	}

	for _, module := range repo_config.Modules {
		mid, err := mid.NewMID(module.GetNamespace(module.Namespace),module.Name,module.Provider,module.Version); if err != nil {
			slog.Error(err.Error())
		}
		if event.PullRequest.Merged {
			modules.PublishDraft(mid); if err != nil {
				slog.Error("Error removing module")
			}
		} else {
			modules.RemoveDraft(mid); if err != nil {
				slog.Error("Error removing module")
			}
		}
	}
}

func syncRepo(localpath string, event PullRequestEvent) (error) {
	configpath := fmt.Sprintf("%s/terrapak.hcl", localpath)
	syncReport := SyncReport{}
	fmt.Println("Syncing repo", localpath)
	var readme string
	repo_config, err  := hclconfig.Load(&configpath); if err != nil {
		return err
	}

	for _, module := range repo_config.Modules {
		mid, err := mid.NewMID(module.GetNamespace(module.Namespace),module.Name,module.Provider,module.Version); if err != nil {
			slog.Error(err.Error())
		}
		extModule := modules.Read(mid); if extModule == nil {
			slog.Info("module not found")
		}
		has_changes := repo.DiffModule(localpath, module.Path,event.PullRequest.Base.Ref);
		zippath, hash, err := helpers.Pack(localpath,module.Path,module.Name); if err != nil {
			slog.Error(err.Error())
		}

		modulePath := fmt.Sprintf("%s/%s",localpath,module.Path)
		fmt.Printf("Changes detected in %s/%s\n", localpath, module.Path)

		file, err := os.ReadFile(zippath); if err != nil {
			slog.Error(err.Error())
		}

		readmePath := fmt.Sprintf("%s/README.md", modulePath)
		if helpers.FileExists(readmePath){
			readme_data, err := os.ReadFile(readmePath); if err != nil {
				slog.Error(err.Error())
			}
			readme = string(readme_data)
		}

		opts := modules.UploadOptions{
			File: file,
			Hash: hash,
			Readme: readme,
		}

		if extModule != nil {
			if has_changes {

				if extModule.PublishedAt.IsZero() {
					res := modules.Upload(mid,opts)

					if res.Code != 201 {
						slog.Error(fmt.Sprintf("Error uploading module %s", module.Name))
						return fmt.Errorf("error uploading module %s", res.Message)
					}
					syncReportItem := SyncReportItem{ Name: module.Name, Version: module.Version}
					syncReport.Items = append(syncReport.Items, syncReportItem)
				} else {
					syncReportItem := SyncReportItem{ Name: module.Name, Version: "new version reqired :warning:" }
					syncReport.Items = append(syncReport.Items, syncReportItem)
				}
			}
		} else {
			res := modules.Upload(mid,opts)
					if res.Code != 201 {
						slog.Error(fmt.Sprintf("Error uploading module %s", module.Name))
						return fmt.Errorf("error uploading module %s", res.Message)
					}
					syncReportItem := SyncReportItem{ Name: module.Name, Version: module.Version}
					syncReport.Items = append(syncReport.Items, syncReportItem)
		}
	}

	err = CreateSyncReport(event, syncReport); if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}