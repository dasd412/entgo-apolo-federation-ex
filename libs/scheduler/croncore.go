package croncore

import "github.com/robfig/cron/v3"

type CronCore struct {
	cron *cron.Cron
}

func NewCronCore() *CronCore {
	return &CronCore{
		cron: cron.New(cron.WithSeconds()),
	}
}

func (c *CronCore) AddJob(spec string, job func()) (cron.EntryID, error) {
	entryId, err := c.cron.AddFunc(spec, job)
	return entryId, err
}

func (c *CronCore) Start() {
	c.cron.Start()
}

func (c *CronCore) Stop() {
	c.cron.Stop()
}

func (c *CronCore) RemoveJob(entryID cron.EntryID) {
	c.cron.Remove(entryID)
}
