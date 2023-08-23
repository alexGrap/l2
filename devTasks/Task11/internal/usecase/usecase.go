package usecase

import (
	models "L2/devTasks/Task11/internal"
	"L2/devTasks/Task11/internal/repository"
	"encoding/json"
	"log"
	"time"
)

type UseCase struct {
	Repository models.Repository
}

func InitUseCase() models.UseCase {
	useCase := UseCase{}
	useCase.Repository = repository.RepInit()
	return &useCase
}

func (useCase *UseCase) Create(model []byte) error {
	var body models.Model
	err := json.Unmarshal(model, &body)
	if err != nil {
		return err
	}
	err = useCase.Repository.Create(body)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *UseCase) Update(model []byte) error {
	var body models.Model
	err := json.Unmarshal(model, &body)
	if err != nil {
		return err
	}
	err = useCase.Repository.Update(body)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *UseCase) Delete(model []byte) error {
	var body models.Model
	err := json.Unmarshal(model, &body)
	if err != nil {
		return err
	}
	err = useCase.Repository.Delete(body)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *UseCase) Get(flag string) ([]byte, error) {
	storage := useCase.Repository.Get()
	events := make([]models.Model, 0)
	for _, tmp := range storage {
		switch flag {
		case "day":
			if ifTime(time.Now().UTC().Truncate(time.Hour*24), time.Now().UTC().Truncate(time.Hour*24).Add(time.Hour*24), tmp.Date) {
				events = append(events, tmp)
			}
		case "week":
			if ifTime(time.Now().UTC().Truncate(time.Hour*24*7), time.Now().UTC().Truncate(time.Hour*24*7).Add(time.Hour*24*7), tmp.Date) {
				events = append(events, tmp)
			}
		case "month":
			if ifTime(time.Now().UTC().Truncate(time.Hour*24*7*30), time.Now().UTC().Add(time.Hour*24*7*30).Truncate(time.Hour*24*7*30), tmp.Date) {
				events = append(events, tmp)
			}
		}
	}
	res, err := json.Marshal(events)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil

}

func ifTime(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
