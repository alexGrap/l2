package repository

import (
	models "L2/devTasks/Task11/internal"
	"errors"
	"sync"
)

type Rep struct {
	storage map[int]models.Model
	mut     sync.RWMutex
}

func RepInit() models.Repository {
	rep := Rep{}
	rep.storage = make(map[int]models.Model)
	return &rep
}

func (rep *Rep) Create(model models.Model) error {
	rep.mut.Lock()
	defer rep.mut.Unlock()
	if _, exist := rep.storage[model.Id]; exist {
		return errors.New("this event already exist")
	}
	rep.storage[model.Id] = model
	return nil
}

func (rep *Rep) Update(model models.Model) error {
	rep.mut.Lock()
	defer rep.mut.Unlock()
	if _, exist := rep.storage[model.Id]; !exist {
		return errors.New("this event not exist")
	}
	rep.storage[model.Id] = model
	return nil
}

func (rep *Rep) Delete(model models.Model) error {
	rep.mut.Lock()
	defer rep.mut.Unlock()
	if _, exist := rep.storage[model.Id]; !exist {
		return errors.New("this event not exist")
	}
	delete(rep.storage, model.Id)
	return nil
}

func (rep *Rep) Get() map[int]models.Model {
	rep.mut.RLock()
	defer rep.mut.RUnlock()
	return rep.storage
}
