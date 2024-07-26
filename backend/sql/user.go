package sql

type UserInfoDataStruct struct {
	UserID        string `db:"user_id" json:"user_id"` // UNIQUE
	UserName      string `db:"user_name" json:"user_name"`
	UserPassword  string `db:"user_password" json:"user_password"` // Password was store in hasher.
	UserEmail     string `db:"user_email" json:"user_email"`
	UserAuthority int64  `db:"user_authority" json:"user_authority"` // 0 => root | 1 => admin | 2 => normal User.
	UserSex       int64  `db:"user_sex" json:"user_sex"`             // 0 ==> male | 1 ==> female
}

type UserInfoDataNoPasswordStruct struct {
	UserID        string `db:"user_id" json:"user_id"`
	UserName      string `db:"user_name" json:"user_name"`
	UserEmail     string `db:"user_email" json:"user_email"`
	UserAuthority int64  `db:"user_authority" json:"user_authority"` // 0 => root | 1 => admin | 2 => normal User.
	UserSex       int64  `db:"user_sex" json:"user_sex"`             // 0 ==> male | 1 ==> female
}

func RemoveReferUserInfo(userID string) {
	SqlDataBase.Del(ReturnDatabaseTableName(UserInfoDataStruct{}), "Where user_id is '"+userID+"'")
}

func QueryDataBaseByID(userID string) UserInfoDataStruct {
	var DataBaseReturner UserInfoDataStruct
	SqlDataBase.Find(ReturnDatabaseTableName(UserInfoDataStruct{}), &DataBaseReturner, "Where user_id is '"+userID+"'")
	return DataBaseReturner
}

func ModifyReferData(dataStruct UserInfoDataStruct) error {
	err := SqlDataBase.Insert(ReturnDatabaseTableName(UserInfoDataStruct{}), &dataStruct)
	if err != nil {
		return err
	}
	return nil
}

func QueryAllUserInfoData() []UserInfoDataStruct {
	var VariesList UserInfoDataStruct
	var QueryListAll []UserInfoDataStruct

	SqlDataBase.FindFor(ReturnDatabaseTableName(UserInfoDataStruct{}), &VariesList, "", func() error {
		QueryListAll = append(QueryListAll, VariesList)
		return nil
	})
	return QueryListAll
}

func QueryAllUserInfoDataWithNoPassword() []UserInfoDataNoPasswordStruct {
	var VariesList UserInfoDataStruct
	var QueryListAll []UserInfoDataNoPasswordStruct

	SqlDataBase.FindFor(ReturnDatabaseTableName(UserInfoDataStruct{}), &VariesList, "", func() error {
		QueryListAll = append(QueryListAll, UserInfoDataNoPasswordStruct{
			UserID:        VariesList.UserID,
			UserName:      VariesList.UserName,
			UserEmail:     VariesList.UserEmail,
			UserAuthority: VariesList.UserAuthority,
			UserSex:       VariesList.UserSex,
		})
		return nil
	})
	return QueryListAll
}
