package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/exec"
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
	startBackupCronJob()
	select {}
}

// startBackupCronJob starts a backup cron job with a specified interval and callback.
func startBackupCronJob() {
	// Daily at 2 o'clock in the morning
	interval := "0 0 0,6,12,18 * * *"

	err := RegisterCronJob("BackupCron", func() {
		runBackupScript()
	}, time.Now(), interval)

	if err != nil {
		log.Println("error registering backup cronjob:", err)
	}
}

// runBackupScript executes the backup script located at /root/backup_script.sh
func runBackupScript() {
	if os.Getenv("MYSQL_HOST") == "" || os.Getenv("MYSQL_USER") == "" || os.Getenv("MYSQL_DATABASE") == "" {
		log.Println("error: one or more mysql variables are not defined.")
		return
	}

	log.Println("Starting backup process...")
	cmd := exec.Command("/bin/bash", "/root/backup_script.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("error executing backup script: %v", err)
	}

	log.Println("Finished backup process...")
}
