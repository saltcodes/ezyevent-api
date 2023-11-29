package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

// CreateOrganization takes a model struct and saves it
func CreateOrganization(organization *model.Organization) error {
	return (*db).Create(organization).Error
}

// ListOrganization Queries the organizations tab;e and stored them in organization list array
func ListOrganization(organizationList *[]model.Organization) error {
	return (*db).Preload("User").Find(&organizationList).Error
}

// GetOrganization Queries and return a single organization
func GetOrganization(id string, organization *model.Organization) error {
	return (*db).Preload("User").First(&organization, "id=?", id).Error
}

// UpdateOrganization update organizations
func UpdateOrganization(id string, organization *model.Organization) error {

	if strings.Compare(id, organization.Id) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(organization).Error
}

// DeleteOrganization delete organization
func DeleteOrganization(id string) error {
	var organization model.Organization

	err := (*db).First(&organization, "id=?", id).Error
	if organization.Id == "" {
		return err
	}
	return (*db).Delete(&organization).Error
}
