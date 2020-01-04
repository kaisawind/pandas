//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package factory

import (
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/jinzhu/gorm"
)

type viewFactory struct {
	modelDB *gorm.DB
}

func (pf *viewFactory) initialize(factoryServingOptions *FactoryServingOptions) error {
	modelDB, err := gorm.Open(factoryServingOptions.StorePath, "pandas-views.db")
	if err != nil {
		return err
	}
	modelDB.AutoMigrate(&models.Project{})
	pf.modelDB = modelDB
	return nil
}

func (pf *viewFactory) Save(owner Owner, obj models.Model) (models.Model, error) {
	view := obj.(*models.View)
	view.CreatedAt = time.Now()
	pf.modelDB.Save(view)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return view, nil
}

func (pf *viewFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	views := []*models.Project{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(views)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, view := range views {
		results = append(results, view)
	}
	return results, nil
}

func (pf *viewFactory) Get(Owner, string) (models.Model, error) {
	return nil, nil
}

func (pf *viewFactory) Delete(Owner, string) error {
	return nil
}

func (pf *viewFactory) Update(Owner, models.Model) error {
	return nil
}