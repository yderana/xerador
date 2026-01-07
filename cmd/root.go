package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
	"github.com/yderana/xerador/generator"
)

// build flags:
// go build -ldflags "-X github.com/yderana/xerador/cmd.Version=1.2.3 -X github.com/yderana/xerador/cmd.AppName=xerador -X github.com/yderana/xerador/cmd.License=MIT"
var (
	AppName = "xerador"
	Version = "dev"
	License = ""
)

func resolvedVersion() string {
	// kalau kamu build pakai -ldflags, tetap dipakai
	if Version != "" && Version != "dev" {
		return Version
	}

	// kalau install via go install @vX.Y.Z, ambil dari build info
	if bi, ok := debug.ReadBuildInfo(); ok {
		if bi.Main.Version != "" && bi.Main.Version != "(devel)" {
			return bi.Main.Version // contoh: "v0.1.0"
		}
	}

	return "dev"
}

func banner() {
	v := resolvedVersion()

	fmt.Printf("\x1b[36m%s\x1b[0m", fmt.Sprintf(`
      ██╗  ██╗███████╗██████╗  █████╗ ██████╗  ██████╗ ██████╗ 
      ╚██╗██╔╝██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔═══██╗██╔══██╗
       ╚███╔╝ █████╗  ██████╔╝███████║██║  ██║██║   ██║██████╔╝
       ██╔██╗ ██╔══╝  ██╔══██╗██╔══██║██║  ██║██║   ██║██╔══██╗
      ██╔╝ ██╗███████╗██║  ██║██║  ██║██████╔╝╚██████╔╝██║  ██║
      ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝  %s

xerador will generate a new service or module in your nestjs monorepo services
`, v))
}

func Execute() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	v := resolvedVersion()

	// pakai generator.Options supaya tidak mapping-mapping lagi
	opts := &generator.Options{}

	cmd := &cobra.Command{
		Use:   AppName,
		Short: "Generate a new service or module in your NestJS monorepo",
		Example: fmt.Sprintf(
			`  %s --add-service -s <project-name> -m <module-name>
  %s --add-module -s <project-name> -m <module-name>
  %s --init-repo --repo <repo-name> -s <project-name> -m <module-name> -i

  %s -r <service-name>
  %s --test <service-name>
  %s --test <service-name> --cov`,
			AppName, AppName, AppName, AppName, AppName, AppName,
		),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			banner()

			// flags template
			initRepo, _ := c.Flags().GetBool("init-repo")
			addService, _ := c.Flags().GetBool("add-service")
			addModule, _ := c.Flags().GetBool("add-module")

			switch {
			case initRepo:
				opts.Template = "repo"
			case addService:
				opts.Template = "project"
			case addModule:
				opts.Template = "module"
			}

			// version
			showVersion, _ := c.Flags().GetBool("version")
			if showVersion {
				fmt.Printf("name   : %s\n", AppName)
				fmt.Printf("version: %s\n", v)
				if License != "" {
					fmt.Printf("license: %s\n", License)
				}
				return nil
			}

			// run service / test service
			if opts.RunService != "" {
				return generator.ExecRunService(*opts)
			}
			if opts.RunTestService != "" {
				return generator.ExecRunTestService(*opts)
			}

			// prompt untuk yang kurang (pilihan template / repo / service / module / install)
			// Kamu perlu update PromptForMissingOptions agar pakai generator.Options (lihat catatan di bawah)
			installChanged := c.Flags().Changed("install")
			opts2, err := PromptForMissingOptions(*opts, installChanged)
			if err != nil {
				return err
			}
			*opts = opts2

			// dispatch generator
			switch opts.Template {
			case "repo":
				_, err := generator.RepoGenerator(*opts)
				return err
			case "project":
				_, err := generator.ServiceGenerator(*opts)
				return err
			default:
				_, err := generator.ModuleGenerator(*opts)
				return err
			}
		},
	}

	// --- Flags: run/test
	cmd.Flags().StringVarP(&opts.RunService, "run", "r", "", "Run Service")
	cmd.Flags().StringVar(&opts.RunTestService, "test", "", "Run Unit Test")
	cmd.Flags().BoolVar(&opts.TestCov, "cov", false, "Run Coverage Test")
	cmd.Flags().BoolVar(&opts.TestCov, "coverage", false, "Alias of --cov")
	cmd.Flags().BoolP("version", "v", false, "Show version number")

	// --- Flags: generator mode
	cmd.Flags().Bool("add-service", false, "Generate service")
	cmd.Flags().Bool("add-module", false, "Generate module")
	cmd.Flags().Bool("init-repo", false, "Init repo")

	// Node: -s itu service/project name
	cmd.Flags().StringVarP(&opts.ProjectName, "service", "s", "", "Service/Project Name")
	cmd.Flags().StringVarP(&opts.ModuleName, "module", "m", "", "Module Name")

	// repo name untuk init-repo
	cmd.Flags().StringVar(&opts.Repo, "repo", "", "Repo Name (for --init-repo)")

	// install deps
	cmd.Flags().BoolVarP(&opts.RunInstall, "install", "i", false, "Install dependencies packages")
	cmd.PreRun = func(c *cobra.Command, args []string) {
		// kalau kamu masih butuh InstallSet, taruh di generator.Options juga.
		// atau handle di PromptForMissingOptions dengan c.Flags().Changed("install").
		// Di sini kita tidak simpan InstallSet karena generator.Options tidak punya field itu.
	}

	// optional: override base dir (pengganti targetDirectory || process.cwd())
	cmd.Flags().StringVar(&opts.BaseDirectory, "base-dir", "", "Base directory (default: current working directory)")

	// optional identity (dipakai license, dll)
	cmd.Flags().StringVar(&opts.Name, "name", "", "Author name")
	cmd.Flags().StringVar(&opts.Email, "email", "", "Author email")

	// template explicit (kalau kamu mau dukung manual)
	cmd.Flags().StringVar(&opts.Template, "template", "", "Template name (repo/project/module). Usually set by --init-repo/--add-service/--add-module")

	return cmd
}
