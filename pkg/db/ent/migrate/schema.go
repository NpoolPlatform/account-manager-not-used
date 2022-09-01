// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultAccountUsedFor"},
		{Name: "platform_hold_private_key", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "active", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "locked", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "locked_by", Type: field.TypeString, Nullable: true, Default: "DefaultLockedBy"},
		{Name: "blocked", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// DepositsColumns holds the columns for the "deposits" table.
	DepositsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "incoming", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "outcoming", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "collecting_tid", Type: field.TypeUUID, Nullable: true},
		{Name: "scannable_at", Type: field.TypeUint32, Nullable: true},
	}
	// DepositsTable holds the schema information for the "deposits" table.
	DepositsTable = &schema.Table{
		Name:       "deposits",
		Columns:    DepositsColumns,
		PrimaryKey: []*schema.Column{DepositsColumns[0]},
	}
	// GoodBenefitsColumns holds the columns for the "good_benefits" table.
	GoodBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "backup", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// GoodBenefitsTable holds the schema information for the "good_benefits" table.
	GoodBenefitsTable = &schema.Table{
		Name:       "good_benefits",
		Columns:    GoodBenefitsColumns,
		PrimaryKey: []*schema.Column{GoodBenefitsColumns[0]},
	}
	// LimitationsColumns holds the columns for the "limitations" table.
	LimitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "limitation", Type: field.TypeString, Nullable: true, Default: "DefaultLimitationType"},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// LimitationsTable holds the schema information for the "limitations" table.
	LimitationsTable = &schema.Table{
		Name:       "limitations",
		Columns:    LimitationsColumns,
		PrimaryKey: []*schema.Column{LimitationsColumns[0]},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "collecting_tid", Type: field.TypeUUID, Nullable: true},
		{Name: "available_at", Type: field.TypeUint32, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
	}
	// PlatformsColumns holds the columns for the "platforms" table.
	PlatformsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultAccountUsedFor"},
		{Name: "backup", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// PlatformsTable holds the schema information for the "platforms" table.
	PlatformsTable = &schema.Table{
		Name:       "platforms",
		Columns:    PlatformsColumns,
		PrimaryKey: []*schema.Column{PlatformsColumns[0]},
	}
	// TransfersColumns holds the columns for the "transfers" table.
	TransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "target_user_id", Type: field.TypeUUID},
	}
	// TransfersTable holds the schema information for the "transfers" table.
	TransfersTable = &schema.Table{
		Name:       "transfers",
		Columns:    TransfersColumns,
		PrimaryKey: []*schema.Column{TransfersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultAccountUsedFor"},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		DepositsTable,
		GoodBenefitsTable,
		LimitationsTable,
		PaymentsTable,
		PlatformsTable,
		TransfersTable,
		UsersTable,
	}
)

func init() {
}
