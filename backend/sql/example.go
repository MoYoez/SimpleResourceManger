// Package sql EXAMPLE: 加入样例, 可演示数据
package sql

import (
	utils "github.com/moyoez/starry-backend/utils"
	"time"
)

// ExampleLoggerinsert Example UserLoggin
func ExampleLoggerinsert() {
	UserAddRoot := UserInfoDataStruct{
		UserID:        "admin",
		UserName:      "admin",
		UserPassword:  utils.Hasher("123456"),
		UserEmail:     "123456@qq.com",
		UserAuthority: 0,
		UserSex:       0,
	}
	UserAddAdmin := UserInfoDataStruct{
		UserID:        "MoeMagicMango",
		UserName:      "KoiParadise",
		UserPassword:  utils.Hasher("123456"),
		UserEmail:     "i@himoyo.cn",
		UserAuthority: 1,
		UserSex:       1,
	}
	UserAddNormalUser := UserInfoDataStruct{
		UserID:        "GUEST",
		UserName:      "GUEST",
		UserPassword:  utils.Hasher("123456"),
		UserEmail:     "guest@guest.com",
		UserAuthority: 2,
		UserSex:       0,
	}
	ModifyReferData(UserAddRoot)
	ModifyReferData(UserAddAdmin)
	ModifyReferData(UserAddNormalUser)
	// add maintain data.
	addMaintainInfoRoot := MaintainTableStruct{
		MaintainID:               "114514",
		MaintainTime:             time.Now().UnixMilli(),
		MaintainContent:          "12839128398",
		MaintainProcess:          1,
		MaintainUserName:         "MoeMagicMango",
		MaintainVerifiedStatus:   true,
		MaintainVerifiedTime:     time.Now().UnixMilli() - 1000,
		MaintainVerifiedUserName: "MoeMagicMango",
	}
	addMaintainInfo := MaintainTableStruct{
		MaintainID:               "1145142",
		MaintainTime:             time.Now().UnixMilli(),
		MaintainContent:          "12839128398",
		MaintainProcess:          0,
		MaintainUserName:         "admin",
		MaintainVerifiedStatus:   true,
		MaintainVerifiedTime:     time.Now().UnixMilli() - 1000,
		MaintainVerifiedUserName: "admin",
	}
	ModifyMaintainReferData(addMaintainInfo)
	ModifyMaintainReferData(addMaintainInfoRoot)
	// add res
	resExample := ResBaseStruct{
		ResID:     "11451423",
		ResNumber: 1,
		ResType:   "ji",
		ResName:   "we",
		ResModel:  "ew",
		ResDate:   time.Now().UnixMilli(),
		ResPlace:  "123",
	}
	ModifyResReferData(resExample)
	// add report
	reportExample := ReportBaseStruct{
		ReportID:              "3123123",
		ReportUser:            "MoeMagicMango",
		ReportContent:         "3123123123123",
		ReportTime:            time.Now().UnixMilli(),
		ReportIsEmergency:     true,
		ReportProcessStatus:   1,
		ReportFeedbackTime:    time.Now().UnixMilli(),
		ReportFeedbackRate:    3,
		ReportFeedbackContent: "2313",
		ReportFeedbackLogs:    "5345345345",
		ReportFeedbackUser:    "admin",
	}
	reportExampleTwo := ReportBaseStruct{
		ReportID:              "312312223",
		ReportUser:            "GUEST",
		ReportContent:         "3123123123123",
		ReportTime:            time.Now().UnixMilli(),
		ReportIsEmergency:     true,
		ReportProcessStatus:   2,
		ReportFeedbackTime:    time.Now().UnixMilli(),
		ReportFeedbackRate:    3,
		ReportFeedbackContent: "2313",
		ReportFeedbackLogs:    "5345345345",
		ReportFeedbackUser:    "admin",
	}
	reportExampleThree := ReportBaseStruct{
		ReportID:              "312312223",
		ReportUser:            "GUEST",
		ReportContent:         "3123123123123",
		ReportTime:            time.Now().UnixMilli(),
		ReportIsEmergency:     false,
		ReportProcessStatus:   0,
		ReportFeedbackTime:    time.Now().UnixMilli(),
		ReportFeedbackRate:    3,
		ReportFeedbackContent: "2313",
		ReportFeedbackLogs:    "5345345345",
		ReportFeedbackUser:    "admin",
	}
	ModifyReportReferData(reportExample)
	ModifyReportReferData(reportExampleTwo)
	ModifyReportReferData(reportExampleThree)
}
