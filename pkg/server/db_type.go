package server

import "time"

type TChAuth struct {
	Fid         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Furl        string    `json:"url" xorm:"not null default '' VARCHAR(128)"`
	Fip         string    `json:"ip" xorm:"not null default '' VARCHAR(1024)"`
	Fuser       string    `json:"user" xorm:"not null default '' VARCHAR(64)"`
	Fsudo       int       `json:"sudo" xorm:"not null default 0 INT(10)"`
	FblackIps   string    `json:"black_ips" xorm:"TEXT"`
	Fctime      int       `json:"ctime" xorm:"created"`
	Fgroup      string    `json:"group" xorm:"not null default '' VARCHAR(512)"`
	Fsalt       string    `json:"salt" xorm:"not null default '' VARCHAR(36)"`
	Ftoken      string    `json:"token" json:"token" xorm:"not null default '' unique VARCHAR(64)"`
	FsudoIps    string    `json:"white_ips" xorm:"TEXT"`
	Fhit        int       `json:"hit" xorm:"not null default 0 INT(10)"`
	FlastUpdate int       `json:"last_update" xorm:"updated"`
	Fdesc       string    `json:"desc" xorm:"not null default '' VARCHAR(1024)"`
	Fenv        string    `json:"env" xorm:"not null default '' VARCHAR(64)"`
	Fenable     int       `json:"enable" xorm:"not null default 1 INT(11)"`
	FmodifyTime time.Time `json:"modify_time" xorm:"updated index DATETIME"`
	Fversion    int       `json:"version" xorm:"not null default 0 INT(11)"`
}

type TChUser struct {
	Fid         int       `xorm:"not null pk autoincr INT(11)"`
	Femail      string    `xorm:"not null default '' VARCHAR(256)"`
	Fuser       string    `xorm:"not null default '' unique VARCHAR(64)"`
	Fpwd        string    `xorm:"not null default '' VARCHAR(256)"`
	Fip         string    `xorm:"not null default '' VARCHAR(32)"`
	Flogincount int       `xorm:"not null default 0 INT(11)"`
	Ffailcount  int       `xorm:"not null default 0 INT(11)"`
	Flasttime   time.Time `xorm:"updated"`
	Fstatus     int       `xorm:"not null default 0 INT(11)"`
	FmodifyTime time.Time `xorm:"updated index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChResults struct {
	Fid         int64     `json:"id" xorm:"not null pk autoincr BIGINT(20)"`
	FtaskId     string    `json:"task_id" xorm:"not null default '' unique VARCHAR(36)"`
	Fip         string    `json:"i" xorm:"not null default '' VARCHAR(16)"`
	Fcmd        string    `json:"cmd"  xorm:"TEXT"`
	Fresult     string    `json:"result" xorm:"TEXT"`
	Fctime      int       `json:"ctime" xorm:"not null default 0  index INT(11)"`
	Futime      int       `json:"utime" xorm:"created"`
	FopUser     string    `json:"user" xorm:"not null default '' VARCHAR(32)"`
	Fuuid       string    `json:"ip" xorm:"not null default '' index VARCHAR(36)"`
	FsysUser    string    `json:"sys_user" xorm:"not null default '' VARCHAR(32)"`
	FmodifyTime time.Time `json:"modifyTime" xorm:"created index DATETIME"`
	Fversion    int       `json:"version" xorm:"not null default 0 INT(11)"`
}

type TChResultsHistory struct {
	Fid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	FtaskId     string    `xorm:"not null default '' unique VARCHAR(36)"`
	Fip         string    `xorm:"not null default '' VARCHAR(16)"`
	Fcmd        string    `xorm:"TEXT"`
	Fresult     string    `xorm:"TEXT"`
	Fctime      int       `xorm:"not null default 0 index INT(11)"`
	Futime      int       `xorm:"not null default 0 INT(11)"`
	FopUser     string    `xorm:"not null default '' VARCHAR(32)"`
	Fuuid       string    `xorm:"not null default '' index VARCHAR(36)"`
	FsysUser    string    `xorm:"not null default '' VARCHAR(32)"`
	FmodifyTime time.Time `xorm:"not null default '1970-01-01 00:00:00' index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChGoogleAuth struct {
	Fid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Fseed       string    `json:"s" xorm:"not null default '' VARCHAR(32)"`
	Fuser       string    `json:"u" xorm:"not null default '' unique(uniq_user_platform) VARCHAR(32)"`
	Fplatform   string    `json:"p" xorm:"not null default '' unique(uniq_user_platform) VARCHAR(32)"`
	Ffail       int       `xorm:"not null default 0 INT(11)"`
	Fhit        int       `xorm:"not null default 0 INT(11)"`
	Fstatus     int       `xorm:"not null default 1 INT(11)"`
	Fctime      time.Time `xorm:"created"`
	Futime      time.Time `xorm:"created"`
	FmodifyTime time.Time `xorm:"created index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChObjs struct {
	Fid         int       `xorm:"not null pk autoincr INT(11)"`
	Fip         string    `xorm:"not null default '' VARCHAR(16)"`
	Fkey        string    `json:"k" xorm:"not null default '' unique(uniq_otype_key) VARCHAR(36)"`
	Fotype      string    `json:"o" xorm:"not null default '' unique(uniq_otype_key) VARCHAR(32)"`
	Fname       string    `json:"n" xorm:"not null default '' VARCHAR(64)"`
	Fbody       string    `xorm:"TEXT"`
	Fuid        int       `xorm:"not null default 0 INT(11)"`
	Fgid        int       `xorm:"not null default 0 INT(11)"`
	Fstatus     int       `xorm:"not null default 0 INT(11)"`
	FmodifyTime time.Time `xorm:"updated index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChLog struct {
	Fid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Furl        string    `json:"url" xorm:"not null default '' VARCHAR(2048)"`
	Fparams     string    `json:"params" xorm:"TEXT"`
	Fmessage    string    `json:"message" xorm:"not null default '' VARCHAR(255)"`
	Fip         string    `json:"ip" xorm:"not null default '' CHAR(15)"`
	Fuser       string    `json:"user" xorm:"not null default '' VARCHAR(64)"`
	Ftime       int       `xorm:"created not null default 0 index INT(11)"`
	FmodifyTime time.Time `xorm:"updated index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChHeartbeat struct {
	Fuuid          string    `json:"uuid" xorm:"not null pk default '' VARCHAR(36)"`
	Fhostname      string    `json:"hostname" xorm:"not null default '' VARCHAR(255)"`
	Fip            string    `json:"ip" xorm:"not null default '' VARCHAR(32)"`
	Fgroup         string    `json:"group" xorm:"not null default '' VARCHAR(32)"`
	FserverUri     string    `json:"server_uri" xorm:"not null default '' VARCHAR(256)"`
	FetcdUri       string    `json:"etcd_uri" xorm:"not null default '' VARCHAR(256)"`
	Fsalt          string    `json:"salt" xorm:"not null default '' VARCHAR(36)"`
	Fplatform      string    `json:"platform" xorm:"not null default '' VARCHAR(36)"`
	Futime         string    `json:"utime"  xorm:"not null default '' VARCHAR(32)"`
	Fnettype       string    `json:"nettype" xorm:"not null default '' VARCHAR(16)"`
	Fstatus        string    `json:"status" xorm:"not null default '' VARCHAR(16)"`
	FsystemStatus  string    `json:"system_status" xorm:"TEXT"`
	FpythonVersion string    `json:"python_version" xorm:"not null default '' VARCHAR(16)"`
	FcliVersion    string    `json:"cli_version" xorm:"not null default '' VARCHAR(32)"`
	FmodifyTime    time.Time `xorm:"updated index DATETIME"`
	Fversion       int       `xorm:"not null default 0 INT(11)"`
}

type TChFiles struct {
	Fid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Fuser       string    `xorm:"not null default '' unique(uniq_user_file) VARCHAR(128)"`
	Fpath       string    `xorm:"not null default '' VARCHAR(128)"`
	Furl        string    `xorm:"VARCHAR(256)"`
	Fmd5        string    `xorm:"VARCHAR(32)"`
	Ffilename   string    `xorm:"not null default '' unique(uniq_user_file) VARCHAR(64)"`
	Fctime      string    `xorm:"not null default '' VARCHAR(32)"`
	Futime      string    `xorm:"not null default '' VARCHAR(32)"`
	Fatime      string    `xorm:"not null default '' VARCHAR(32)"`
	Fhit        int       `xorm:"not null default 0 INT(11)"`
	FmodifyTime time.Time `xorm:"updated index DATETIME"`
	Fversion    int       `xorm:"not null default 0 INT(11)"`
}

type TChDoc struct {
	Fid         int64     `json:"id" xorm:"not null pk autoincr BIGINT(20)"`
	Fcmd        string    `json:"cmd" xorm:"TEXT"`
	Fdoc        string    `json:"doc" xorm:"LONGTEXT"`
	Fremark     string    `json:"remark" xorm:"not null default '' VARCHAR(512)"`
	FmodifyTime time.Time `json:"modifyTime" xorm:"updated index DATETIME"`
	Fversion    int       `json:"version" xorm:"not null default 0 INT(11)"`
}

type TChConfig struct {
	Fid         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Fgroup      string    `json:"group" xorm:"not null default '' VARCHAR(36)"`
	Fip         string    `json:"ip" xorm:"not null default '' VARCHAR(32)"`
	Fuuid       string    `json:"uuid" xorm:"not null default '' VARCHAR(36)"`
	FisGateway  int       `json:"isGrateway" xorm:"not null default 0 TINYINT(1)"`
	Fconfig     string    `json:"config" xorm:"LONGTEXT"`
	FmodifyTime time.Time `json:"modifyTime" xorm:"updated index DATETIME"`
	Fversion    int       `json:"version" xorm:"not null default 0 INT(11)"`
}

type TChBasedata struct {
	Fid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Fname       string    `xorm:"not null default '' VARCHAR(50)"`
	Fpid        int64     `xorm:"not null default 0 BIGINT(20)"`
	Fsort       int       `xorm:"not null default 0 INT(11)"`
	Fvalue      string    `xorm:"TEXT"`
	Fcode       string    `xorm:"not null default '' unique VARCHAR(64)"`
	Fdesc       string    `xorm:"not null default '' VARCHAR(512)"`
	FmodifyTime time.Time `json:"modifyTime" xorm:"updated index DATETIME"`
	Fversion    int       `json:"version" xorm:"not null default 0 INT(11)"`
}
