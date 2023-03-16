package model

type StroyHandlerModel struct {
	StoryId int64
	Handler int64
}

func CreateStoryHandlerInBatch(list []StroyHandlerModel) error {
	err := DbConnect.CreateInBatches(list, 100).Error
	return err
}

func (StroyHandlerModel) TableName() string {
	return "story_connect"
}
