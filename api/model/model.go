package model

type User struct {
	Id int `gorm:"type:int NOT NULL auto_increment;primary_key;" json:"id"`
	Username string `gorm:"type:varchar(50) NOT NULL;unique" json:"username"`
	Password string `json:"-"`
	Email string `gorm:"type:varchar(50) NOT NULL;unique" json:"email"`
	Bio string `gorm:"type:text" json:"bio"`
	Location string `gorm:"type:varchar(50)" json:"location"`
	ProfileImage string `gorm:"type:varchar(255) NOT NULL" json:"profile_image"`
	TotalCollections int `gorm:"type:int DEFAULT:0" json:"total_collections"`
	TotalPhotos int	`gorm:"type:int DEFAULT:0" json:"total_photos"`
	Link string `gorm:"type:varchar(255)" json:"link"`
	LockStatus bool `gorm:"type:bool DEFAULT:False" json:"lock_status"`
	Confirmed bool `gorm:"type:bool DEFAULT:False" json:"confirmed"`
	RoleId int `gorm:"foreignKey:Id"`
	MethodId int `gorm:"foreignKey:Id" json:"-"`
	Method Method `gorm:"foreignKey:MethodId"`
	//Photos Photo `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	//Collections []Collection `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	//Following []Follow `gorm:"foreignKey:Follow.followerId, constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	//Followers []Follow `gorm:"foreignKey:Follow.followedId, constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

type Photo struct {
	Id          uint   `gorm:"type:int NOT NULL auto_increment;primary_key;" json:"id"`
	Imageurl    string `gorm:"type:varchar(255)" json:"imageUrl"`
	Description string `gorm:"type:text" json:"description"`
	Download    string `gorm:"type:varchar(255)" json:"download"`
	Timestamp   string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"timestamp"`
	Label       string `gorm:"type:varchar(100) DEFAULT ''" json:"label"`
	User        User   `gorm:"foreignKey:AuthorId"`
	AuthorId    int    `gorm:"foreignKey:Id" json:"-"`
}

type Video struct {
	Id uint `gorm:"type:int NOT NULL auto_increment; primary_key;" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
	Videourl string `gorm:"type:varchar(255)" json:"videourl"`
	Link string `gorm:"type:varchar(255)" json:"link"`
	CategoryId int `gorm:"foreignKey:Category.id, type:int NOT NULL" json:"-"`

}

type Role struct {
	Id int
}

type Permission struct {
	id uint `gorm:"type:int NOT NULL auto_increment; primary_key;"`
	name string `gorm:"type:varchar(20) unique"`
}

type Method struct {
	Id int `json:"-"`
	Name string `gorm:"type:varchar(20) NOT NULL" json:"user_method"`
	users []User `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

type Follow struct {
	followerId int `gorm:"foreignKey:User.id, type:int NOT NULL auto_increment; primary_key;"`
	followedId int `gorm:"foreignKey:User.id, type:int NOT NULL auto_increment; primary_key;"`
	timestamp string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"timestamp"`
	follower []User `gorm:"foreignKey:Follow.followerId"`
	followed []User `gorm:"foreignKey:Follow.followedId"`
}

type Collection struct {
	CollectorId int `gorm:"foreignKey:User.id, type:int NOT NULL auto_increment; primary_key;" json:"collector_id"`
	CollectedId int `form:"foreignKet:Photo.id, type:int NOT NULL auto_increment; primary_key" json:"collected_id"`
	Timestamp string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"timestamp"`

}

type Category struct {
	Id uint `gorm:"type:int NOT NULL auto_increment;primary_key;" json:"id"`
	Name string `gorm:"type:varchar(20) NOT NULL" json:"name"`
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	Videos []Video `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

