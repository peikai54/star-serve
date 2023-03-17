package model

type StatusModel struct {
	StatusName string
}

func (StatusModel) TableName() string {
	return "status"
}

func GetStatusList() ([]StatusModel, error) {
	var list = []StatusModel{}
	err := DbConnect.Find(&list).Error
	return list, err
}
