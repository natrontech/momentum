package main

import (
	"fmt"
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
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

const ENABLE_CLEANUP = true

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

	app := pocketbase.New()

	var publicDirFlag string

	// add "--publicDir" option flag
	app.RootCmd.PersistentFlags().StringVar(
		&publicDirFlag,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)
	migrationsDir := "" // default to "pb_migrations" (for js) and "migrations" (for go)

	// load js files to allow loading external JavaScript migrations
	jsvm.MustRegisterMigrations(app, &jsvm.MigrationsOptions{
		Dir: migrationsDir,
	})

	// register the `migrate` command
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		TemplateLang: migratecmd.TemplateLangJS, // or migratecmd.TemplateLangGo (default)
		Dir:          migrationsDir,
		Automigrate:  true,
	})

	// call this only if you want to use the configurable "hooks" functionality
	hooks.PocketBaseInit(app)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDirFlag), true))

		return nil
	})

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {

		if ENABLE_CLEANUP {
			testCleanUp(app.Dao().Clone())
		}

		dispatcher := momentumcore.NewDispatcher(momentumConfig, app)

		// momentum core features must run before executing DB statements.
		// like this invalid/inconsistent state is prevented.
		dispatcher.Pocketbase.OnRecordBeforeCreateRequest().Add(dispatcher.DispatchCreate)
		dispatcher.Pocketbase.OnRecordBeforeUpdateRequest().Add(dispatcher.DispatchUpdate)
		dispatcher.Pocketbase.OnRecordBeforeDeleteRequest().Add(dispatcher.DispatchDelete)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func testCleanUp(dao *daos.Dao) {

	keyvalues, err := dao.FindRecordsByExpr("keyValues")
	depls, err := dao.FindRecordsByExpr("deployments")
	stages, err := dao.FindRecordsByExpr("stages")
	apps, err := dao.FindRecordsByExpr("applications")
	reps, err := dao.FindRecordsByExpr("repositories")

	if err != nil {
		fmt.Println("TEST CLEAN FAILED:", err.Error())
	}

	for _, e := range keyvalues {
		dao.DeleteRecord(e)
	}
	for _, e := range depls {
		dao.DeleteRecord(e)
	}
	for _, e := range stages {
		dao.DeleteRecord(e)
	}
	for _, e := range apps {
		dao.DeleteRecord(e)
	}
	for _, e := range reps {
		dao.DeleteRecord(e)
	}
}
