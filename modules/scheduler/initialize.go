package scheduler

import (
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/version"
	"github.com/erda-project/erda/modules/scheduler/conf"
	"github.com/erda-project/erda/modules/scheduler/i18n"
	"github.com/erda-project/erda/modules/scheduler/server"
	"github.com/erda-project/erda/pkg/dumpstack"
)

// Initialize 应用相关的初始化操作
func Initialize() error {
	logrus.Infof(version.String())
	conf.Load()
	// control log's level.
	if conf.Debug() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
	// open the function of dump stack
	dumpstack.Open()
	logrus.Infof("start the service and listen on address: \"%s\"", conf.ListenAddr())
	logrus.Errorf("[alert] starting scheduler instance")
	i18n.InitI18N()

	//if err := server.NewServer(conf.ListenAddr()).ListenAndServe(); err != nil {
	//	logrus.Error(err)
	//	os.Exit(2)
	//}
	return server.NewServer(conf.ListenAddr()).ListenAndServe()
}
