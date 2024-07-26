package sql

type MaintainTableStruct struct {
	MaintainID               string `db:"m_id" json:"m_id"`
	MaintainTime             int64  `db:"m_time" json:"m_time"`
	MaintainContent          string `db:"m_content" json:"m_content"`
	MaintainProcess          int64  `db:"m_process" json:"m_process"`
	MaintainUserName         string `db:"m_name" json:"m_name"`
	MaintainVerifiedStatus   bool   `db:"v_verify" json:"v_verify"`
	MaintainVerifiedTime     int64  `db:"v_time" json:"v_time"`
	MaintainVerifiedUserName string `db:"v_user" json:"v_user"`
}

func RemoveReferMaintainInfo(maintainID string) {
	SqlDataBase.Del(ReturnDatabaseTableName(MaintainTableStruct{}), "Where m_id is '"+maintainID+"'")
}

func QueryMaintainDataBaseByID(maintainID string) MaintainTableStruct {
	var DataBaseReturner MaintainTableStruct
	SqlDataBase.Find(ReturnDatabaseTableName(MaintainTableStruct{}), &DataBaseReturner, "Where m_id is '"+maintainID+"'")
	return DataBaseReturner
}

func ModifyMaintainReferData(dataStruct MaintainTableStruct) error {
	err := SqlDataBase.Insert(ReturnDatabaseTableName(MaintainTableStruct{}), &dataStruct)
	if err != nil {
		return err
	}
	return nil
}

func QueryAllMaintainInfoData() []MaintainTableStruct {
	var VariesList MaintainTableStruct
	var QueryListAll []MaintainTableStruct

	SqlDataBase.FindFor(ReturnDatabaseTableName(MaintainTableStruct{}), &VariesList, "", func() error {
		QueryListAll = append(QueryListAll, VariesList)
		return nil
	})
	return QueryListAll
}
