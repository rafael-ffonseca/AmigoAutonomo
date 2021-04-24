package models

type Config struct {
	ServicePort				string
	SchedulerInMinutes		int
	DatabaseHost			string
	DatabasePort			string
	DatabaseUser			string
	DatabasePassword		string
	DatabaseName			string
	LogDatabasePath			string
	Homologation 			bool
}
