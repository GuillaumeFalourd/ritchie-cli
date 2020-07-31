/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"time"

	"github.com/kaduartur/go-cli-spinner/pkg/spinner"

	"github.com/spf13/cobra"

	"github.com/ZupIT/ritchie-cli/pkg/formula"
	"github.com/ZupIT/ritchie-cli/pkg/github"
	"github.com/ZupIT/ritchie-cli/pkg/prompt"
	"github.com/ZupIT/ritchie-cli/pkg/rtutorial"
)

var CommonsRepoURL = "https://github.com/ZupIT/ritchie-formulas"

type initCmd struct {
	repo formula.RepositoryAdder
	git  github.Repositories
	rt   rtutorial.Finder
}

func NewInitCmd(repo formula.RepositoryAdder, git github.Repositories, rtf rtutorial.Finder) *cobra.Command {
	o := initCmd{repo: repo, git: git, rt: rtf}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize rit configuration",
		Long:  "Initialize rit configuration",
		RunE:  o.runPrompt(),
	}

	return cmd
}

func (in initCmd) runPrompt() CommandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		repo := formula.Repo{
			Name:     "commons",
			Url:      CommonsRepoURL,
			Priority: 0,
		}

		s := spinner.StartNew("Adding the commons repository...")
		time.Sleep(time.Second * 2)

		repoInfo := github.NewRepoInfo(repo.Url, repo.Token)

		tag, err := in.git.LatestTag(repoInfo)
		if err != nil {
			return err
		}

		repo.Version = formula.RepoVersion(tag.Name)

		if err := in.repo.Add(repo); err != nil {
			return err
		}

		s.Success(prompt.Green("Initialization successful!"))

		tutorialHolder, err := in.rt.Find()
		if err != nil {
			return err
		}
		tutorialInit(tutorialHolder.Current)
		return nil
	}
}

func tutorialInit(tutorialStatus string) {
	const tagTutorial = "\n[TUTORIAL]"
	const MessageTitle = "How to create new formulas:"
	const MessageBody = ` ∙ Run "rit create formula"
 ∙ Open the project with your favorite text editor.` + "\n"

	if tutorialStatus == tutorialStatusEnabled {
		prompt.Info(tagTutorial)
		prompt.Info(MessageTitle)
		fmt.Println(MessageBody)
	}
}
