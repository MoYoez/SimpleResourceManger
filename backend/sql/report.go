package sql

type ReportBaseStruct struct {
	ReportID              string `db:"report_id" json:"report_id"`
	ReportUser            string `db:"report_user" json:"report_user"`
	ReportReferRes        string `db:"report_refer_res" json:"report_refer_res"` // can be null
	ReportContent         string `db:"report_content" json:"report_content"`
	ReportTime            int64  `db:"report_time" json:"report_time"`
	ReportIsEmergency     bool   `db:"report_emergency" json:"report_emergency"`
	ReportProcessStatus   int64  `db:"report_process" json:"report_process"` // 0 => no start 1 => started 2 => done.
	ReportFeedbackTime    int64  `db:"report_feedback_time" json:"report_feedback_time"`
	ReportFeedbackRate    int64  `db:"report_feedback_rate" json:"report_feedback_rate"` // feedback Rate 0-5
	ReportFeedbackContent string `db:"report_feedback_content" json:"report_feedback_content"`
	ReportFeedbackLogs    string `db:"report_feedback_logs" json:"report_feedback_logs"`
	ReportFeedbackUser    string `db:"report_feedback_user" json:"report_feedback_user"`
}

func RemoveReferReportInfo(ReportID string) {
	SqlDataBase.Del(ReturnDatabaseTableName(ReportBaseStruct{}), "Where report_id is '"+ReportID+"'")
}

func QueryReportDataBaseByID(ReportID string) ReportBaseStruct {
	var DataBaseReturner ReportBaseStruct
	SqlDataBase.Find(ReturnDatabaseTableName(ReportBaseStruct{}), &DataBaseReturner, "Where report_id is '"+ReportID+"'")
	return DataBaseReturner
}

func ModifyReportReferData(dataStruct ReportBaseStruct) error {
	err := SqlDataBase.Insert(ReturnDatabaseTableName(ReportBaseStruct{}), &dataStruct)
	if err != nil {
		return err
	}
	return nil
}

func QueryAllReportInfoData() []ReportBaseStruct {
	var VariesList ReportBaseStruct
	var QueryListAll []ReportBaseStruct

	SqlDataBase.FindFor(ReturnDatabaseTableName(ReportBaseStruct{}), &VariesList, "", func() error {
		QueryListAll = append(QueryListAll, VariesList)
		return nil
	})
	return QueryListAll
}
