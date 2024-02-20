package helper

import "gorm.io/gorm"

type PaginateOptions struct {
	PageSize int
	Page     int
}

// Paginate returns a function that applies pagination to a GORM database query.
//
// It takes PaginateOptions as a parameter and returns a function that takes a GORM database pointer and returns a GORM database pointer.
//
// db.Scopes(Paginate(dbfunc.opts)).Find(&users)
func Paginate(opts PaginateOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if opts.Page <= 0 {
			opts.Page = 1
		}

		switch {
		case opts.PageSize > 100:
			opts.PageSize = 100
		case opts.PageSize <= 0:
			opts.PageSize = 10
		}

		offset := (opts.Page - 1) * opts.PageSize
		return db.Offset(offset).Limit(opts.PageSize)
	}
}
