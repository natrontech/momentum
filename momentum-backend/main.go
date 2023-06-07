package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	hooks "momentum/hooks"
	momentumcore "momentum/momentum-core"
	config "momentum/momentum-core/momentum-config"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}

	return filepath.Join(os.Args[0], "../pb_public")
}

func main() {

	momentumConfig, configErr := config.InitializeMomentumCore()
	if configErr != nil {
		panic("failed initializing momentum. problem: " + configErr.Error())
	}

	dispatcher := momentumcore.NewDispatcher(momentumConfig, pocketbase.New())

	// momentum core features must run before executing DB statements.
	// like this invalid/inconsistent state is prevented.
	dispatcher.App.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		dispatcher.App.OnRecordBeforeCreateRequest().Add(dispatcher.DispatchCreate)
		dispatcher.App.OnRecordBeforeUpdateRequest().Add(dispatcher.DispatchUpdate)
		dispatcher.App.OnRecordBeforeDeleteRequest().Add(dispatcher.DispatchDelete)

		return nil
	})

	var publicDirFlag string

	// add "--publicDir" option flag
	dispatcher.App.RootCmd.PersistentFlags().StringVar(
		&publicDirFlag,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)
	migrationsDir := "" // default to "pb_migrations" (for js) and "migrations" (for go)

	// load js files to allow loading external JavaScript migrations
	jsvm.MustRegisterMigrations(dispatcher.App, &jsvm.MigrationsOptions{
		Dir: migrationsDir,
	})

	// register the `migrate` command
	migratecmd.MustRegister(dispatcher.App, dispatcher.App.RootCmd, &migratecmd.Options{
		TemplateLang: migratecmd.TemplateLangJS, // or migratecmd.TemplateLangGo (default)
		Dir:          migrationsDir,
		Automigrate:  true,
	})

	// call this only if you want to use the configurable "hooks" functionality
	hooks.PocketBaseInit(dispatcher.App)

	dispatcher.App.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDirFlag), true))

		return nil
	})

	if err := dispatcher.App.Start(); err != nil {
		log.Fatal(err)
	}
}
