package log

// Interface represents the API of both Logger and Entry.
type Interface interface {
	SetEnvProject(env string, project string) *Entry
	WithFields(fields Fielder) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) *Entry
	/********* Interface simple *********/
/* WARNING: This part is autogenerated! Do not change! */

	Debug(msg string)
	Info(msg string)
	Notice(msg string)
	Warn(msg string)
	Error(msg string)
	Critical(msg string)
	Alert(msg string)
	Fatal(msg string)
	Emergency(msg string)

/* END OF WARNING */
	/********* End Interface simple *********/

	/********* Interface formated *********/
/* WARNING: This part is autogenerated! Do not change! */

	Debugf(msg string, v ...interface{})
	Infof(msg string, v ...interface{})
	Noticef(msg string, v ...interface{})
	Warnf(msg string, v ...interface{})
	Errorf(msg string, v ...interface{})
	Criticalf(msg string, v ...interface{})
	Alertf(msg string, v ...interface{})
	Fatalf(msg string, v ...interface{})
	Emergencyf(msg string, v ...interface{})

/* END OF WARNING */
	/********* End Interface formated *********/

	Trace(msg string) *Entry
	Exit(code int)
	AddExitHandler(handler func())
}
