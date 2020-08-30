package services

import (
	"go-crash-course/entities"
	"go-crash-course/repositories"
)

type TagService struct {
	TagRepository repositories.ITagRepository
}

func (a *TagService) FetchTag() []*entities.Tag {
	tags := a.TagRepository.FindAll()
	return tags
}

func (a *TagService) SaveTag(data *entities.Tag) *entities.Tag {
	result := a.TagRepository.Store(data)
	return result
}

func (a *TagService) FetchTagById(id int) *entities.Tag {
	tag := a.TagRepository.FindById(id)
	return tag
}

func (a *TagService) Destroy(id int) error {
	err := a.TagRepository.Destroy(id)

	if err != nil {
		return err
	}
	return nil
}

func (a *TagService) Update(tag *entities.Tag) (*entities.Tag, error) {
	result, err := a.TagRepository.Update(tag)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewTagService() *TagService {
	return &TagService{
		TagRepository: repositories.NewTagRepository(),
	}
}
