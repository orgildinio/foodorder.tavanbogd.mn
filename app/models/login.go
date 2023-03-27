package models

import "time"

type UserInfo struct {
	Affectedrows int `json:"affectedrows"`
	Retdata      []struct {
		Empid                 float32     `json:"empid"`
		Code                  string      `json:"code"`
		Familyname            interface{} `json:"familyname"`
		Familynameenus        interface{} `json:"familynameenus"`
		Lastname              string      `json:"lastname"`
		Lastnameenus          string      `json:"lastnameenus"`
		Firstname             string      `json:"firstname"`
		Firstnameenus         string      `json:"firstnameenus"`
		Gender                string      `json:"gender"`
		Gendername            interface{} `json:"gendername"`
		Gendernameenus        interface{} `json:"gendernameenus"`
		Birthdate             string      `json:"birthdate"`
		Birthcountryid        interface{} `json:"birthcountryid"`
		Birthcountrycode      interface{} `json:"birthcountrycode"`
		Birthcountryname      interface{} `json:"birthcountryname"`
		BirthcountrynameEnus  interface{} `json:"birthcountryname_enus"`
		Birthdivisionid       interface{} `json:"birthdivisionid"`
		Birthdivisioncode     interface{} `json:"birthdivisioncode"`
		Birthdivisionname     interface{} `json:"birthdivisionname"`
		BirthdivisionnameEnus interface{} `json:"birthdivisionname_enus"`
		Birthdistrictid       interface{} `json:"birthdistrictid"`
		Birthdistrictcode     interface{} `json:"birthdistrictcode"`
		Birthdistrictname     interface{} `json:"birthdistrictname"`
		BirthdistrictnameEnus interface{} `json:"birthdistrictname_enus"`
		Birthplace            interface{} `json:"birthplace"`
		Countryid             interface{} `json:"countryid"`
		Countrycode           interface{} `json:"countrycode"`
		Countryname           interface{} `json:"countryname"`
		CountrynameEnus       interface{} `json:"countryname_enus"`
		Divisionid            interface{} `json:"divisionid"`
		Divisioncode          interface{} `json:"divisioncode"`
		Divisionname          interface{} `json:"divisionname"`
		DivisionnameEnus      interface{} `json:"divisionname_enus"`
		Districtid            interface{} `json:"districtid"`
		Districtcode          interface{} `json:"districtcode"`
		Districtname          interface{} `json:"districtname"`
		DistrictnameEnus      interface{} `json:"districtname_enus"`
		Nationalityid         interface{} `json:"nationalityid"`
		Nationalitycode       interface{} `json:"nationalitycode"`
		Nationalityname       interface{} `json:"nationalityname"`
		NationalitynameEnus   interface{} `json:"nationalityname_enus"`
		Socialoriginid        interface{} `json:"socialoriginid"`
		Socialorigincode      interface{} `json:"socialorigincode"`
		Socialoriginname      interface{} `json:"socialoriginname"`
		SocialoriginnameEnus  interface{} `json:"socialoriginname_enus"`
		Maritalstatus         string      `json:"maritalstatus"`
		Maritalstatusname     interface{} `json:"maritalstatusname"`
		MaritalstatusnameEnus interface{} `json:"maritalstatusname_enus"`
		Bloodgroup            string      `json:"bloodgroup"`
		Bloodgroupname        interface{} `json:"bloodgroupname"`
		BloodgroupnameEnus    interface{} `json:"bloodgroupname_enus"`
		Regno                 string      `json:"regno"`
		Passno                interface{} `json:"passno"`
		Forpassno             interface{} `json:"forpassno"`
		Emdno                 interface{} `json:"emdno"`
		Nddno                 interface{} `json:"nddno"`
		Addr1                 interface{} `json:"addr1"`
		Addr2                 interface{} `json:"addr2"`
		Homephone             interface{} `json:"homephone"`
		Mobilephone           string      `json:"mobilephone"`
		Workphone             interface{} `json:"workphone"`
		Workphoneext          interface{} `json:"workphoneext"`
		Fax                   interface{} `json:"fax"`
		Email                 string      `json:"email"`
		Postaddress           interface{} `json:"postaddress"`
		Contactname           interface{} `json:"contactname"`
		Contactphone          interface{} `json:"contactphone"`
		Withpicture           string      `json:"withpicture"`
		Status                string      `json:"status"`
		Statusname            interface{} `json:"statusname"`
		StatusnameEnus        interface{} `json:"statusname_enus"`
		Companyid             string      `json:"companyid"`
		Companyname           string      `json:"companyname"`
		Statusdate            string      `json:"statusdate"`

		Email2          interface{} `json:"email2"`
		Contactname2    interface{} `json:"contactname2"`
		Contactphone2   interface{} `json:"contactphone2"`
		Isblacklisted   interface{} `json:"isblacklisted"`
		Blacklistreason interface{} `json:"blacklistreason"`
		Blacklisteddate interface{} `json:"blacklisteddate"`

		Militaryserved  string      `json:"militaryserved"`
		Trainingsport   interface{} `json:"trainingsport"`
		Interest        interface{} `json:"interest"`
		Appid           interface{} `json:"appid"`
		Militaryno      interface{} `json:"militaryno"`
		Ecompanyid      string      `json:"ecompanyid"`
		Emptype         string      `json:"emptype"`
		Emptypename     interface{} `json:"emptypename"`
		EmptypenameEnus interface{} `json:"emptypename_enus"`
		Depid           string      `json:"depid"`
		Depcode         string      `json:"depcode"`
		Depname         string      `json:"depname"`
		DepnameEnus     interface{} `json:"depname_enus"`

		Poscode       string      `json:"poscode"`
		Posname       string      `json:"posname"`
		PosnameEnus   interface{} `json:"posname_enus"`
		Gradeid       int         `json:"gradeid"`
		Gradecode     string      `json:"gradecode"`
		Gradename     string      `json:"gradename"`
		GradenameEnus interface{} `json:"gradename_enus"`
		Levelcode     string      `json:"levelcode"`
		Levelname     string      `json:"levelname"`
		LevelnameEnus string      `json:"levelname_enus"`
		Hireddate     string      `json:"hireddate"`

		Basewagecurrcode string `json:"basewagecurrcode"`

		Joineddate         interface{} `json:"joineddate"`
		Terminateddate     interface{} `json:"terminateddate"`
		Terminatedreasonid interface{} `json:"terminatedreasonid"`
		Contractid         interface{} `json:"contractid"`

		Picturedata       interface{} `json:"picturedata"`
		Picturetumb       interface{} `json:"picturetumb"`
		Signaturedata     interface{} `json:"signaturedata"`
		Countryworkedyear string      `json:"countryworkedyear"`
		Groupworkedyear   string      `json:"groupworkedyear"`
		Companyworkedyear string      `json:"companyworkedyear"`
		Reason            interface{} `json:"reason"`
	} `json:"retdata"`
	Rettype       int         `json:"rettype"`
	Depfilter     interface{} `json:"depfilter"`
	Retparams     interface{} `json:"retparams"`
	Retmsg        string      `json:"retmsg"`
	Retstacktrace string      `json:"retstacktrace"`
	Traceno       int         `json:"traceno"`
	Totalrow      int         `json:"totalrow"`
	Totalpage     int         `json:"totalpage"`
}

type UserERP struct {
	ID             int64      `gorm:"column:id;primary_key" json:"id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Status         string     `gorm:"column:status;type:ENUM('0','1','2')" json:"status"`
	Role           int64      `gorm:"column:role" json:"role"`
	Login          string     `gorm:"column:login;unique_index;not null;unique" json:"login"`
	Email          string     `gorm:"column:email;unique_index;not null;unique" json:"email"`
	RegisterNumber string     `gorm:"column:register_number;not null;unique" json:"register_number"`
	Avatar         string     `gorm:"column:avatar;type:TEXT" json:"avatar"`
	Bio            string     `gorm:"column:bio;type:TEXT" json:"bio"`
	FirstName      string     `gorm:"column:first_name" json:"first_name"`
	LastName       string     `gorm:"column:last_name" json:"last_name"`
	Birthday       time.Time  `gorm:"column:birthday;type:DATE" json:"birthday"`
	Phone          string     `gorm:"column:phone" json:"phone"`
	//DistrictID     int        `gorm:"column:district_id" json:"district_id"`
	//RegionId       int        `gorm:"column:region_id" json:"region_id"`
	//OrgId          string     `gorm:"column:org_id" json:"org_id"`
	//UnitId         string     `gorm:"column:unit_id" json:"unit_id"`
	//PositionId     string     `gorm:"column:position_id" json:"position_id"`
	Gender      string `gorm:"column:gender;type:ENUM('f','m')" json:"gender"`
	Password    string `gorm:"column:password" json:"password"`
	FcmToken    string `gorm:"column:fcm_token" json:"fcm_token"`
	EmpID       int    `gorm:"column:emp_id" json:"emp_id"`
	CompanyID   string `gorm:"column:company_id" json:"company_id"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	PosName     string `gorm:"column:pos_name" json:"pos_name"`
	DepName     string `gorm:"column:dep_name" json:"dep_name"`
}

// TableName sets the insert table name for this struct type
func (v *UserERP) TableName() string {
	return "users"
}
