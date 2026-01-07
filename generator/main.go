package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yderana/xerador/helpers"
)

type Options struct {
	// base: pengganti (options.targetDirectory || process.cwd()) di Node
	BaseDirectory string

	// mode/template: "repo" | "project" | "module"
	Template string

	// nama repo (untuk init-repo)
	Repo string

	// Node: -s/--service itu project/service name
	ProjectName string

	// Node: -m/--module itu module name
	ModuleName string

	// runtime runner
	RunService     string
	RunTestService string
	TestCov        bool

	// install dependencies
	RunInstall bool

	// identity (untuk LICENSE, dsb)
	Name  string
	Email string

	// path yang dihitung saat generator berjalan
	TargetDirectory string
	TargetService   string
}

func ModuleGenerator(opts Options) (bool, error) {
	// targetService: <cwd>/packages/<project>
	targetService := filepath.Join(baseDir(opts), "packages", helpers.KebabCase(opts.ProjectName))
	if !helpers.PathExists(targetService) {
		return false, fmt.Errorf("ERROR Service not found: %s", targetService)
	}

	// set options
	opts.TargetService = targetService
	opts.TargetDirectory = filepath.Join(
		baseDir(opts),
		"packages",
		helpers.KebabCase(opts.ProjectName),
		"src",
		helpers.KebabCase(opts.ModuleName),
	)

	// template untuk moduleGenerator selalu "module"
	opts.Template = "module"

	// default identity
	opts.Name = defaultName(opts.Name)
	opts.Email = defaultEmail(opts.Email)

	// module must not exist
	if helpers.PathExists(opts.TargetDirectory) {
		return false, fmt.Errorf("ERROR Module already exist: %s", opts.TargetDirectory)
	}

	tasks := []helpers.Task{
		{Title: "Create directories", Run: func() error { return CreateDirectory(opts) }},
		{Title: "Create interfaces", Run: func() error { return CreateInterfaces(opts) }},
		{Title: "Create repository", Run: func() error { return CreateRepository(opts) }},
		{Title: "Create providers", Run: func() error { return CreateProviders(opts) }},
		{Title: "Create schema", Run: func() error { return CreateSchema(opts) }},
		{Title: "Create controllers", Run: func() error { return CreateControllers(opts) }},
		{Title: "Create module", Run: func() error { return CreateModuleFiles(opts) }},
	}

	if err := helpers.RunTasks(tasks); err != nil {
		return false, err
	}

	fmt.Println(helpers.Green("DONE"), "Module ready")
	return true, nil
}

func RepoGenerator(opts Options) (bool, error) {
	repoKebab := helpers.KebabCase(opts.Repo)
	projectKebab := helpers.KebabCase(opts.ProjectName)
	moduleKebab := helpers.KebabCase(opts.ModuleName)

	base := baseDir(opts)

	// repo root options
	repoOpts := opts
	repoOpts.Template = "repo"
	repoOpts.TargetDirectory = filepath.Join(base, repoKebab)
	repoOpts.Name = defaultName(opts.Name)
	repoOpts.Email = defaultEmail(opts.Email)

	// repo must not exist
	if helpers.PathExists(repoOpts.TargetDirectory) {
		return false, fmt.Errorf("ERROR Service already exist: %s", repoOpts.TargetDirectory)
	}

	// optionsProject (createService untuk project template)
	optionsProject := opts
	optionsProject.Template = "project"
	optionsProject.TargetDirectory = filepath.Join(base, repoKebab, "packages", projectKebab)
	optionsProject.Name = repoOpts.Name
	optionsProject.Email = repoOpts.Email

	// optionsModule (module default di dalam project)
	optionsModule := opts
	optionsModule.Template = "module"
	optionsModule.TargetService = filepath.Join(base, repoKebab, "packages", projectKebab)
	optionsModule.TargetDirectory = filepath.Join(base, repoKebab, "packages", projectKebab, "src", moduleKebab)
	optionsModule.Name = repoOpts.Name
	optionsModule.Email = repoOpts.Email

	tasks := []helpers.Task{
		{Title: "Create repo", Run: func() error { return CreateRepo(repoOpts) }},
		{Title: "Create project", Run: func() error { return CreateService(optionsProject) }},
		{Title: "Create directories", Run: func() error { return CreateDirectory(optionsModule) }},
		{Title: "Create interfaces", Run: func() error { return CreateInterfaces(optionsModule) }},
		{Title: "Create repository", Run: func() error { return CreateRepository(optionsModule) }},
		{Title: "Create providers", Run: func() error { return CreateProviders(optionsModule) }},
		{Title: "Create schema", Run: func() error { return CreateSchema(optionsModule) }},
		{Title: "Create controllers", Run: func() error { return CreateControllers(optionsModule) }},
		{Title: "Create module", Run: func() error { return CreateModuleFiles(optionsModule) }},
		{Title: "Implement module", Run: func() error { return ImplementModule(optionsProject) }},
		{
			Title: "Install dependencies",
			Run: func() error {
				if !opts.RunInstall {
					return helpers.Skip("Use --install to install dependencies automatically")
				}
				return YarnInstall(optionsProject.TargetDirectory)
			},
		},
	}

	if err := helpers.RunTasks(tasks); err != nil {
		return false, err
	}

	fmt.Println(helpers.Green("DONE"), "Project ready")
	return true, nil
}

func ServiceGenerator(opts Options) (bool, error) {
	projectKebab := helpers.KebabCase(opts.ProjectName)
	moduleKebab := helpers.KebabCase(opts.ModuleName)
	base := baseDir(opts)

	// project root
	projectOpts := opts
	projectOpts.Template = "project"
	projectOpts.TargetDirectory = filepath.Join(base, "packages", projectKebab)
	projectOpts.Name = defaultName(opts.Name)
	projectOpts.Email = defaultEmail(opts.Email)

	// service must not exist
	if helpers.PathExists(projectOpts.TargetDirectory) {
		return false, fmt.Errorf("ERROR Service already exist: %s", projectOpts.TargetDirectory)
	}

	// optionsModule (module default di dalam project)
	optionsModule := opts
	optionsModule.Template = "module"
	optionsModule.TargetService = filepath.Join(base, "packages", projectKebab)
	optionsModule.TargetDirectory = filepath.Join(base, "packages", projectKebab, "src", moduleKebab)
	optionsModule.Name = projectOpts.Name
	optionsModule.Email = projectOpts.Email

	tasks := []helpers.Task{
		{Title: "Create project", Run: func() error { return CreateService(projectOpts) }},
		{Title: "Create directories", Run: func() error { return CreateDirectory(optionsModule) }},
		{Title: "Create interfaces", Run: func() error { return CreateInterfaces(optionsModule) }},
		{Title: "Create repository", Run: func() error { return CreateRepository(optionsModule) }},
		{Title: "Create providers", Run: func() error { return CreateProviders(optionsModule) }},
		{Title: "Create schema", Run: func() error { return CreateSchema(optionsModule) }},
		{Title: "Create controllers", Run: func() error { return CreateControllers(optionsModule) }},
		{Title: "Create module", Run: func() error { return CreateModuleFiles(optionsModule) }},
		{Title: "Implement module", Run: func() error { return ImplementModule(projectOpts) }},
		{
			Title: "Install dependencies",
			Run: func() error {
				if !opts.RunInstall {
					return helpers.Skip("Pass --install to automatically install dependencies")
				}
				return YarnInstall(projectOpts.TargetDirectory)
			},
		},
	}

	if err := helpers.RunTasks(tasks); err != nil {
		return false, err
	}

	fmt.Println(helpers.Green("DONE"), "Project ready")
	return true, nil
}

func ExecRunService(opts Options) error {
	targetDirectory := filepath.Join(baseDir(opts), "packages", helpers.KebabCase(opts.RunService))
	if !helpers.PathExists(targetDirectory) {
		return fmt.Errorf("ERROR Service not found: %s", targetDirectory)
	}
	return RunCmdStream(targetDirectory, "npm", "run", "start:dev")
}

func ExecRunTestService(opts Options) error {
	targetDirectory := filepath.Join(baseDir(opts), "packages", helpers.KebabCase(opts.RunTestService))
	if !helpers.PathExists(targetDirectory) {
		return fmt.Errorf("ERROR Service not found: %s", targetDirectory)
	}

	testScript := "test"
	if opts.TestCov {
		testScript = "test:cov"
	}
	return RunCmdStream(targetDirectory, "npm", "run", testScript)
}

func baseDir(opts Options) string {
	if opts.BaseDirectory != "" {
		return opts.BaseDirectory
	}
	cwd, _ := os.Getwd()
	return cwd
}

func defaultName(v string) string {
	if v != "" {
		return v
	}
	return "Yoga Derana"
}

func defaultEmail(v string) string {
	if v != "" {
		return v
	}
	return "yderana@gmail.com"
}
