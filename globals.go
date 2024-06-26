package main

import (
	"time"
)

const FILENAME_LENGTH int = 100
const MAGIC_NUMBER int64 = 68500
const HIDDEN_FILE_DATA_OFFSET int64 = 1500
const MB_IN_BYTES int64 = 1000000
const CONFIG_FILE string = "StegstreamServerConfig.txt"

const DEFAULT_PORT int = 8080
const DEFAULT_STREAM_ONLY bool = false
const DEFAULT_HIDE_ONLY bool = false
const DEFAULT_WIPE_AUDIO bool = false
const DEFAULT_WIPE_HIDDEN bool = false
const DEFAULT_WIPE_AFTER_HIDE = false

var DEFAULT_AUTO_SHUTDOWN time.Time = time.Time{}

const DEBUG = false
const UPDATE_UI = true
