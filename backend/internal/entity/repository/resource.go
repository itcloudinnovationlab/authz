package repository

import (
	"github.com/eko/authz/backend/internal/entity/model"
	"gorm.io/gorm"
)

type ResourceQueryOption func(*resourceQueryOptions)

type resourceQueryOptions struct {
	resourceIDs []string
}

func WithResourceIDs(resourceIDs []string) ResourceQueryOption {
	return func(o *resourceQueryOptions) {
		o.resourceIDs = resourceIDs
	}
}

type Resource interface {
	Base[model.Resource]
	FindMatchingAttributesWithPrincipals(resourceAttribute string, principalAttribute string, options ...ResourceQueryOption) ([]*MatchingAttributeResourcePrincipal, error)
}

// besource struct that allows contacting the database using Gorm.
type resource struct {
	Base[model.Resource]
}

// NewResource initializes a new resource repository.
func NewResource(repository Base[model.Resource]) Resource {
	return &resource{
		repository,
	}
}

type MatchingAttributeResourcePrincipal struct {
	PrincipalID   string
	ResourceKind  string
	ResourceValue string
}

func (r *resource) FindMatchingAttributesWithPrincipals(
	resourceAttribute string,
	principalAttribute string,
	options ...ResourceQueryOption,
) ([]*MatchingAttributeResourcePrincipal, error) {
	matches := []*MatchingAttributeResourcePrincipal{}

	tx := applyResourceOptions(r.DB(), options)

	err := tx.
		Select("authz_principals_attributes.principal_id AS principal_id, authz_resources.kind AS resource_kind, authz_resources.value AS resource_value").
		Model(&model.Resource{}).
		Joins("INNER JOIN authz_resources_attributes ON authz_resources.id = authz_resources_attributes.resource_id").
		Joins("INNER JOIN authz_attributes ON authz_resources_attributes.attribute_id = authz_attributes.id").
		Joins("INNER JOIN authz_principals_attributes ON authz_attributes.id = authz_principals_attributes.attribute_id").
		Where("authz_attributes.key = ? OR authz_attributes.key = ?", resourceAttribute, principalAttribute).
		Where("authz_principals_attributes.principal_id IS NOT NULL").
		Where("authz_resources.value <> ?", "*").
		Scan(&matches).Error
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func applyResourceOptions(tx *gorm.DB, options []ResourceQueryOption) *gorm.DB {
	opts := &resourceQueryOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if len(opts.resourceIDs) > 0 {
		tx = tx.Where("authz_resources.id IN ?", opts.resourceIDs)
	}

	return tx
}