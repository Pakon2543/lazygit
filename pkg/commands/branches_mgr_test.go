package commands_test

import (
	"github.com/go-errors/errors"
	. "github.com/jesseduffield/lazygit/pkg/commands"
	. "github.com/jesseduffield/lazygit/pkg/commands/commandsfakes"
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BranchesMgr", func() {
	var (
		commander   *FakeICommander
		gitconfig   *FakeIGitConfig
		branchesMgr *BranchesMgr
		userConfig  *config.UserConfig
	)

	BeforeEach(func() {
		commander = NewFakeCommander()
		gitconfig = &FakeIGitConfig{}
		userConfig = &config.UserConfig{}
		gitconfig.GetUserConfigCalls(func() *config.UserConfig { return userConfig })
		gitconfig.ColorArgCalls(func() string { return "always" })
		branchesMgr = NewBranchesMgr(commander, gitconfig)
	})

	Describe("NewBranch", func() {
		It("runs expected command", func() {
			SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
				{"git checkout -b newName master", "", nil},
			})

			branchesMgr.NewBranch("newName", "master")
		})
	})

	Describe("CurrentBranchName", func() {
		Context("On master branch", func() {
			It("returns 'master'", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git symbolic-ref --short HEAD", "master\n", nil},
				})

				name, displayname, err := branchesMgr.CurrentBranchName()
				Expect(name).To(Equal("master"))
				Expect(displayname).To(Equal("master"))
				Expect(err).To(BeNil())
			})
		})

		Context("symbolic-ref fails", func() {
			Context("when git branch command says we're on master", func() {
				It("falls back to 'git branch --contains'", func() {
					SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
						{"git symbolic-ref --short HEAD", "", errors.New("my error")},
						{"git branch --contains", "* master\n  otherbranch\n", nil},
					})

					name, displayname, err := branchesMgr.CurrentBranchName()
					Expect(name).To(Equal("master"))
					Expect(displayname).To(Equal("master"))
					Expect(err).To(BeNil())
				})
			})

			Context("when git branch command says we're on a detached head", func() {
				It("falls back to 'git branch --contains'", func() {
					SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
						{"git symbolic-ref --short HEAD", "", errors.New("my error")},
						{"git branch --contains", "* (HEAD detached at 264fc6f5)\n  otherbranch\n", nil},
					})

					name, displayname, err := branchesMgr.CurrentBranchName()
					Expect(name).To(Equal("264fc6f5"))
					Expect(displayname).To(Equal("(HEAD detached at 264fc6f5)"))
					Expect(err).To(BeNil())
				})
			})

			Context("when both commands return an error", func() {
				It("bubbles up error", func() {
					SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
						{"git symbolic-ref --short HEAD", "", errors.New("my error")},
						{"git branch --contains", "", errors.New("my other error")},
					})

					name, displayname, err := branchesMgr.CurrentBranchName()
					Expect(name).To(Equal(""))
					Expect(displayname).To(Equal(""))
					Expect(err).To(MatchError("my other error"))
				})
			})
		})
	})

	Describe("AllBranchesCmdObj", func() {
		It("runs expected command", func() {
			userConfig.Git.AllBranchesLogCmd = "git log --graph --all"

			cmdObj := branchesMgr.AllBranchesCmdObj()
			Expect(cmdObj.ToString()).To(Equal("git log --graph --all"))
		})
	})

	Describe("GetBranchGraphCmdObj", func() {
		It("runs expected command", func() {
			userConfig.Git.BranchLogCmd = "git log --graph {{branchName}} --"

			cmdObj := branchesMgr.GetBranchGraphCmdObj("mybranch")
			Expect(cmdObj.ToString()).To(Equal("git log --graph mybranch --"))
		})
	})

	Describe("Delete", func() {
		Context("when force flag is true", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git branch -D mybranch", "", nil},
				})

				err := branchesMgr.Delete("mybranch", true)
				Expect(err).To(BeNil())
			})
		})

		Context("when force flag is false", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git branch -d mybranch", "", nil},
				})

				err := branchesMgr.Delete("mybranch", false)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Merge", func() {
		Context("default case", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git merge --no-edit mybranch", "", nil},
				})

				err := branchesMgr.Merge("mybranch", MergeOpts{})
				Expect(err).To(BeNil())
			})
		})

		Context("when fast-forward only arg is passed", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git merge --no-edit --ff-only mybranch", "", nil},
				})

				err := branchesMgr.Merge("mybranch", MergeOpts{FastForwardOnly: true})
				Expect(err).To(BeNil())
			})

			Context("when user has configured additional args", func() {
				It("runs expected command", func() {
					userConfig.Git.Merging.Args = "--extra-arg"

					SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
						{"git merge --no-edit --ff-only --extra-arg mybranch", "", nil},
					})

					err := branchesMgr.Merge("mybranch", MergeOpts{FastForwardOnly: true})
					Expect(err).To(BeNil())
				})
			})
		})
	})

	Describe("Checkout", func() {
		Context("non-forced", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git checkout mybranch", "", nil},
				})

				err := branchesMgr.Checkout("mybranch", CheckoutOpts{})
				Expect(err).To(BeNil())
			})
		})

		Context("forced", func() {
			It("runs expected command", func() {
				SetExpectedRunWithOutputCalls(commander, []ExpectedRunWithOutputCall{
					{"git checkout --force mybranch", "", nil},
				})

				err := branchesMgr.Checkout("mybranch", CheckoutOpts{Force: true})
				Expect(err).To(BeNil())
			})
		})
	})
})