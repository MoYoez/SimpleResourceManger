package sql

type ResBaseStruct struct {
	ResID     string `db:"res_id" json:"res_id"`
	ResNumber int64  `db:"res_number" json:"res_number"`
	ResType   string `db:"res_type" json:"res_type"`
	ResName   string `db:"res_name" json:"res_name"`
	ResModel  string `db:"res_model" json:"res_model"`
	ResDate   int64  `db:"res_date" json:"res_date"`
	ResPlace  string `db:"res_place" json:"res_place"`
}

func RemoveReferResInfo(ResID string) {
	SqlDataBase.Del(ReturnDatabaseTableName(ResBaseStruct{}), "Where res_id is '"+ResID+"'")
}

func QueryResDataBaseByID(ResID string) ResBaseStruct {
	var DataBaseReturner ResBaseStruct
	SqlDataBase.Find(ReturnDatabaseTableName(ResBaseStruct{}), &DataBaseReturner, "Where res_id is '"+ResID+"'")
	return DataBaseReturner
}

func ModifyResReferData(dataStruct ResBaseStruct) error {
	err := SqlDataBase.Insert(ReturnDatabaseTableName(ResBaseStruct{}), &dataStruct)
	if err != nil {
		return err
	}
	return nil
}

func QueryAllResInfoData() []ResBaseStruct {
	var VariesList ResBaseStruct
	var QueryListAll []ResBaseStruct

	SqlDataBase.FindFor(ReturnDatabaseTableName(ResBaseStruct{}), &VariesList, "", func() error {
		QueryListAll = append(QueryListAll, VariesList)
		return nil
	})
	return QueryListAll
}
