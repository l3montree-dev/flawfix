package vulndb

import (
	"errors"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

type flawService interface {
	RecalculateAllRawRiskAssessments() error
}

type vulnDBService struct {
	leaderElector leaderElector

	mitreService           mitreService
	epssService            epssService
	nvdService             NVDService
	osvService             osvService
	exploitDBService       exploitDBService
	githubExploitDBService githubExploitDBService
	dsa                    debianSecurityTracker
	cveList                cvelistService

	configService configService

	flawService flawService
}

func newVulnDBService(leaderElector leaderElector, mitreService mitreService, epssService epssService, nvdService NVDService, configService configService, osvService osvService, exploitDBService exploitDBService, githubExploitDBService githubExploitDBService, flawService flawService, dsa debianSecurityTracker, cveList cvelistService) *vulnDBService {
	return &vulnDBService{
		leaderElector: leaderElector,
		// Add a comma after leaderElector
		osvService:             osvService,
		mitreService:           mitreService,
		epssService:            epssService,
		nvdService:             nvdService,
		exploitDBService:       exploitDBService,
		githubExploitDBService: githubExploitDBService,
		dsa:                    dsa,
		cveList:                cveList,

		configService: configService,

		flawService: flawService,
	}
}

func (v *vulnDBService) mirror() {
	for {
		// first mirror mitre
		// then mirror nvd
		// then mirror epss
		// then sleep for 2 hours
		if v.leaderElector.IsLeader() {
			// check the last time we mirrored
			var lastMirror struct {
				Time time.Time `json:"time"`
			}

			err := v.configService.GetJSONConfig("vulndb.lastMirror", &lastMirror)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Error("could not get last mirror time", "err", err)
				continue
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Info("no last mirror time found. Setting to 0")
				lastMirror.Time = time.Time{}
			}

			if time.Since(lastMirror.Time) > 2*time.Hour {

				slog.Info("last mirror was more than 2 hours ago. Starting mirror process")

				if err := v.mitreService.Mirror(); err != nil {
					slog.Error("could not mirror mitre cwes", "err", err)
				} else {
					slog.Info("successfully mirrored mitre cwes")
				}

				if err := v.nvdService.mirror(); err != nil {
					slog.Error("could not mirror nvd", "err", err)
				} else {
					slog.Info("successfully mirrored nvd")
				}

				if err := v.cveList.Mirror(); err != nil {
					slog.Error("could not mirror cve list", "err", err)
				} else {
					slog.Info("successfully mirrored cve list")
				}

				if err := v.exploitDBService.Mirror(); err != nil {
					slog.Error("could not mirror exploitdb", "err", err)
				} else {
					slog.Info("successfully mirrored exploitdb")
				}
				if err := v.githubExploitDBService.Mirror(); err != nil {
					slog.Error("could not mirror github exploitdb", "err", err)
				} else {
					slog.Info("successfully mirrored github exploitdb")
				}
				if err := v.epssService.Mirror(); err != nil {
					slog.Error("could not mirror epss", "err", err)
				} else {
					slog.Info("successfully mirrored epss")
				}
				if err := v.osvService.Mirror(); err != nil {
					slog.Error("could not mirror osv", "err", err)
				} else {
					slog.Info("successfully mirrored osv")
				}

				if err := v.dsa.Mirror(); err != nil {
					slog.Error("could not mirror dsa", "err", err)
				} else {
					slog.Info("successfully mirrored dsa")
				}

				if err := v.configService.SetJSONConfig("vulndb.lastMirror", struct {
					Time time.Time `json:"time"`
				}{
					Time: time.Now(),
				}); err != nil {
					slog.Error("could not set last mirror time", "err", err)
				}

			} else {
				slog.Info("last mirror was less than 2 hours ago. Not mirroring", "lastMirror", lastMirror.Time, "now", time.Now())
			}
			err = v.flawService.RecalculateAllRawRiskAssessments()
			if err != nil {
				slog.Error("could not recalculate raw risk assessment", "err", err)
			}
			slog.Info("done. Waiting for 2 hours to check again")
			time.Sleep(2 * time.Hour)
		} else {
			// if we are not the leader, sleep for 5 minutes
			slog.Info("not the leader. Waiting for 5 minutes to check again")
			time.Sleep(5 * time.Minute)
		}
	}
}

func (v *vulnDBService) startMirrorDaemon() {
	go v.mirror()
}
