package vulndb

import (
	"log/slog"
	"time"
)

type vulnDBService struct {
	leaderElector leaderElector

	mitreService mitreService
	epssService  epssService
	nvdService   NVDService
}

func newVulnDBService(leaderElector leaderElector, mitreService mitreService, epssService epssService, nvdService NVDService) *vulnDBService {
	return &vulnDBService{
		leaderElector: leaderElector,

		mitreService: mitreService,
		epssService:  epssService,
		nvdService:   nvdService,
	}
}

func (v *vulnDBService) mirror() {
	for {
		// first mirror mitre
		// then mirror nvd
		// then mirror epss
		// then sleep for 2 hours
		if v.leaderElector.IsLeader() {
			if err := v.mitreService.mirror(); err != nil {
				slog.Error("could not mirror mitre cwes", "err", err)
			} else {
				slog.Info("successfully mirrored mitre cwes")
			}
			if err := v.nvdService.mirror(); err != nil {
				slog.Error("could not mirror nvd", "err", err)
				panic(err)
			} else {
				slog.Info("successfully mirrored nvd")
			}
			if err := v.epssService.mirror(); err != nil {
				slog.Error("could not mirror epss", "err", err)
			} else {
				slog.Info("successfully mirrored epss")
			}
			time.Sleep(2 * time.Hour)
		} else {
			// if we are not the leader, sleep for 5 minutes
			slog.Info("not the leader. Waiting for 5 minutes to try again")
			time.Sleep(5 * time.Minute)
		}
	}
}

func (v *vulnDBService) startMirrorDaemon() {
	v.mirror()
}