package inject

import "engineeringexercise/api/handlers/healthcheck"

func Reset() {
	SetupHealthcheckRoutes = defaultSetupHealthcheckRoutes
}

var SetupHealthcheckRoutes = defaultSetupHealthcheckRoutes
var defaultSetupHealthcheckRoutes = healthcheck.SetupRoutes
