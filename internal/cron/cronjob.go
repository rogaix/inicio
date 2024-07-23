package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

var cronScheduler *cron.Cron

type Job struct {
	Name      string
	Callback  func()
	StartTime time.Time
	Interval  string
	EntryID   cron.EntryID
}

var registeredCronJobs []Job

func init() {
	cronScheduler = cron.New(cron.WithSeconds())
	cronScheduler.Start()
}

// RegisterCronJob registers a new cron job with the given parameters
func RegisterCronJob(name string, callback func(), startTime time.Time, interval string) error {
	entryID, err := cronScheduler.AddFunc(interval, func() {
		now := time.Now()
		if now.After(startTime) {
			message := fmt.Sprintf("Cronjob '%s' executed: %s", name, now.String())
			logToFile(message)
			log.Println(message)
			callback()
		}
	})

	if err != nil {
		return fmt.Errorf("error adding the cron job '%s': %v", name, err)
	}

	job := Job{
		Name:      name,
		Callback:  callback,
		StartTime: startTime,
		Interval:  interval,
		EntryID:   entryID,
	}
	registeredCronJobs = append(registeredCronJobs, job)
	return nil
}

func logToFile(message string) {
	f, err := os.OpenFile("/var/log/cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening Log file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		log.Println("error writing to log file:", err)
	}
}

// StartCronJobs starts all the cron jobs and blocks the execution of the program.
func StartCronJobs() {
	startTestCronJob()
	select {}
}

// startTestCronJob starts a test cron job with a specified start time and interval.
// It registers the cron job using the RegisterCronJob function and logs any errors encountered.
// The callback function is executed when the cron job is triggered.
func startTestCronJob() {
	startTime := time.Now().Add(10 * time.Second) // Start in 10 seconds
	interval := "*/5 * * * * *"                   // Every 5 seconds

	err := RegisterCronJob("exampleJob", func() {
		log.Println("Callback function executed")
	}, startTime, interval)

	if err != nil {
		log.Println("error registering cronjob:", err)
	}
}
