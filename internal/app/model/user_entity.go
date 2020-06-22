package model

type SpUser struct {
	UserId        int    `json:"userId" gorm:"not null pk autoincr comment('自增id') INT(11) 'user_id'"`
	Username      string `json:"username" gorm:"not null default '' comment('登录名') VARCHAR(128) 'username'"`
	QqOpenId      string `json:"qqOpenId" gorm:"default 'NULL' comment('qq官方唯一编号信息') CHAR(32) 'qq_open_id'"`
	Password      string `json:"password" gorm:"not null default '' comment('登录密码') CHAR(64) 'password'"`
	UserEmail     string `json:"userEmail" gorm:"not null default '' comment('邮箱') VARCHAR(64) 'user_email'"`
	UserEmailCode string `json:"userEmailCode" gorm:"default 'NULL' comment('新用户注册邮件激活唯一校验码') CHAR(13) 'user_email_code'"`
	IsActive      string `json:"isActive" gorm:"default '否' comment('新用户是否已经通过邮箱激活帐号') ENUM('否','是') 'is_active'"`
	UserSex       string `json:"userSex" gorm:"not null default '男' comment('性别') ENUM('保密','女','男') 'user_sex'"`
	UserQq        string `json:"userQq" gorm:"not null default '' comment('qq') VARCHAR(32) 'user_qq'"`
	UserTel       string `json:"userTel" gorm:"not null default '' comment('手机') VARCHAR(32) 'user_tel'"`
	UserXueli     string `json:"userXueli" gorm:"not null default '本科' comment('学历') ENUM('专科','初中','博士','小学','本科','硕士','高中') 'user_xueli'"`
	UserHobby     string `json:"userHobby" gorm:"not null default '' comment('爱好') VARCHAR(32) 'user_hobby'"`
	UserIntroduce string `json:"userIntroduce" gorm:"comment('简介') TEXT 'user_introduce'"`
}
