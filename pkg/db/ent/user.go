// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels []string `json:"labels,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldLabels:
			values[i] = new([]byte)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case user.FieldUsedFor:
			values[i] = new(sql.NullString)
		case user.FieldID, user.FieldAppID, user.FieldUserID, user.FieldCoinTypeID, user.FieldAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = uint32(value.Int64)
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = uint32(value.Int64)
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = uint32(value.Int64)
			}
		case user.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				u.AppID = *value
			}
		case user.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				u.UserID = *value
			}
		case user.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				u.CoinTypeID = *value
			}
		case user.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				u.AccountID = *value
			}
		case user.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				u.UsedFor = value.String
			}
		case user.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", u.DeletedAt))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", u.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", u.UserID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", u.CoinTypeID))
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", u.AccountID))
	builder.WriteString(", used_for=")
	builder.WriteString(u.UsedFor)
	builder.WriteString(", labels=")
	builder.WriteString(fmt.Sprintf("%v", u.Labels))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}