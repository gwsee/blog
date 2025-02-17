// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountColumns holds the columns for the "account" table.
	AccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "账户ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "nickname", Type: field.TypeString, Comment: "昵称", Default: "好运连连"},
		{Name: "account", Type: field.TypeString, Unique: true, Comment: "账户", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "password", Type: field.TypeString, Comment: "密码", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "email", Type: field.TypeString, Comment: "邮箱", Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "description", Type: field.TypeString, Comment: "想说啥?", SchemaType: map[string]string{"mysql": "tinytext"}},
		{Name: "avatar", Type: field.TypeString, Comment: "头像", Default: ""},
		{Name: "blog_num", Type: field.TypeInt, Comment: "博客数量", Default: 0},
		{Name: "status", Type: field.TypeInt8, Comment: "状态:0失效,1正常", Default: 1, SchemaType: map[string]string{"mysql": "tinyint"}},
	}
	// AccountTable holds the schema information for the "account" table.
	AccountTable = &schema.Table{
		Name:       "account",
		Comment:    "站点",
		Columns:    AccountColumns,
		PrimaryKey: []*schema.Column{AccountColumns[0]},
	}
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "博客ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "account_id", Type: field.TypeInt, Comment: "账户ID"},
		{Name: "title", Type: field.TypeString, Comment: "标题", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "description", Type: field.TypeString, Comment: "描述", SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "is_hidden", Type: field.TypeInt8, Comment: "是否隐藏:0否,1是", Default: 0, SchemaType: map[string]string{"mysql": "tinyint"}},
		{Name: "tags", Type: field.TypeJSON, Comment: "标签"},
		{Name: "cover", Type: field.TypeString, Comment: "封面", SchemaType: map[string]string{"mysql": "varchar(200)"}},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Comment:    "博客",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
	}
	// BlogsCommentsColumns holds the columns for the "blogs_comments" table.
	BlogsCommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "评论ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "account_id", Type: field.TypeInt, Comment: "账户ID"},
		{Name: "blog_id", Type: field.TypeInt, Comment: "博客ID"},
		{Name: "top_id", Type: field.TypeInt, Comment: "顶级ID", Default: 0},
		{Name: "parent_id", Type: field.TypeInt, Comment: "父评论", Default: 0},
		{Name: "level", Type: field.TypeInt, Comment: "第几楼", Default: 0},
		{Name: "total", Type: field.TypeInt, Comment: "当前级有多少数据;状态未0的", Default: 0},
		{Name: "status", Type: field.TypeInt8, Comment: "0显示,1隐藏", Default: 0, SchemaType: map[string]string{"mysql": "tinyint"}},
		{Name: "content", Type: field.TypeString, Comment: "评论内容"},
	}
	// BlogsCommentsTable holds the schema information for the "blogs_comments" table.
	BlogsCommentsTable = &schema.Table{
		Name:       "blogs_comments",
		Comment:    "博客评论",
		Columns:    BlogsCommentsColumns,
		PrimaryKey: []*schema.Column{BlogsCommentsColumns[0]},
	}
	// BlogsContentColumns holds the columns for the "blogs_content" table.
	BlogsContentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "博客ID"},
		{Name: "content", Type: field.TypeString, Comment: "内容", SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "files", Type: field.TypeJSON, Comment: "博客内容可以关联的文件"},
	}
	// BlogsContentTable holds the schema information for the "blogs_content" table.
	BlogsContentTable = &schema.Table{
		Name:       "blogs_content",
		Comment:    "博客内容",
		Columns:    BlogsContentColumns,
		PrimaryKey: []*schema.Column{BlogsContentColumns[0]},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Comment: "文件的ID(内容MD5)"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "type", Type: field.TypeString, Comment: "文件类型", Default: "", SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "size", Type: field.TypeInt64, Comment: "文件大小", Default: 0},
		{Name: "name", Type: field.TypeString, Comment: "文件名称", Default: ""},
		{Name: "path", Type: field.TypeString, Comment: "文件真实存放地址", Default: "", SchemaType: map[string]string{"mysql": "varchar(250)"}},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Comment:    "文件",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// FilesExtendColumns holds the columns for the "files_extend" table.
	FilesExtendColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "当前账户关联的文件的ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "user_id", Type: field.TypeInt, Comment: "所属的用户", Default: 0},
		{Name: "from", Type: field.TypeString, Comment: "文件来源的表"},
		{Name: "from_id", Type: field.TypeInt, Comment: "文件来源表的ID"},
		{Name: "is_hidden", Type: field.TypeBool, Comment: "是否隐藏"},
		{Name: "file_id", Type: field.TypeString, Nullable: true, Comment: "文件的ID"},
	}
	// FilesExtendTable holds the schema information for the "files_extend" table.
	FilesExtendTable = &schema.Table{
		Name:       "files_extend",
		Comment:    "文件关系表",
		Columns:    FilesExtendColumns,
		PrimaryKey: []*schema.Column{FilesExtendColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_extend_files_extends",
				Columns:    []*schema.Column{FilesExtendColumns[11]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "filesextend_file_id_user_id_from",
				Unique:  true,
				Columns: []*schema.Column{FilesExtendColumns[11], FilesExtendColumns[7], FilesExtendColumns[8]},
			},
		},
	}
	// TravelExtendsColumns holds the columns for the "travel_extends" table.
	TravelExtendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "account_id", Type: field.TypeInt, Comment: "账户ID"},
		{Name: "is_thumb", Type: field.TypeBool, Comment: "是否点赞", Default: false},
		{Name: "is_collect", Type: field.TypeBool, Comment: "收藏量", Default: false},
		{Name: "travel_id", Type: field.TypeInt, Nullable: true, Comment: "旅行的ID"},
	}
	// TravelExtendsTable holds the schema information for the "travel_extends" table.
	TravelExtendsTable = &schema.Table{
		Name:       "travel_extends",
		Comment:    "旅行关联关系",
		Columns:    TravelExtendsColumns,
		PrimaryKey: []*schema.Column{TravelExtendsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "travel_extends_travels_travel_extends",
				Columns:    []*schema.Column{TravelExtendsColumns[10]},
				RefColumns: []*schema.Column{TravelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "travelextends_account_id_travel_id",
				Unique:  true,
				Columns: []*schema.Column{TravelExtendsColumns[7], TravelExtendsColumns[10]},
			},
		},
	}
	// TravelsColumns holds the columns for the "travels" table.
	TravelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "旅行记录的ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "title", Type: field.TypeString, Comment: "标题", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "description", Type: field.TypeString, Comment: "旅行简介", SchemaType: map[string]string{"mysql": "text"}},
		{Name: "video", Type: field.TypeString, Comment: "旅行视频", Default: "", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "is_hidden", Type: field.TypeBool, Comment: "是否隐藏:0否,1是", Default: false},
		{Name: "photos", Type: field.TypeJSON, Comment: "旅行的照片"},
		{Name: "browse_num", Type: field.TypeInt, Comment: "浏览量", Default: 0},
		{Name: "thumb_num", Type: field.TypeInt, Comment: "点赞量", Default: 0},
		{Name: "collect_num", Type: field.TypeInt, Comment: "收藏量", Default: 0},
		{Name: "account_id", Type: field.TypeInt, Nullable: true, Comment: "账户ID"},
	}
	// TravelsTable holds the schema information for the "travels" table.
	TravelsTable = &schema.Table{
		Name:       "travels",
		Comment:    "旅行",
		Columns:    TravelsColumns,
		PrimaryKey: []*schema.Column{TravelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "travels_account_travels",
				Columns:    []*schema.Column{TravelsColumns[15]},
				RefColumns: []*schema.Column{AccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "用户ID=账户ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "name", Type: field.TypeString, Comment: "姓名", Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "avatar", Type: field.TypeString, Comment: "头像", Default: "", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "email", Type: field.TypeString, Comment: "工作邮箱", Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "professional", Type: field.TypeString, Comment: "职称", Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "address", Type: field.TypeString, Comment: "地址", Default: "", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "skills", Type: field.TypeJSON, Comment: "技能"},
		{Name: "description", Type: field.TypeString, Comment: "个人简介", SchemaType: map[string]string{"mysql": "text"}},
		{Name: "experience", Type: field.TypeInt, Comment: "经历数", Default: 0},
		{Name: "project", Type: field.TypeInt, Comment: "项目数", Default: 0},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Comment:    "用户",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// UserExperienceColumns holds the columns for the "user_experience" table.
	UserExperienceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "user_id", Type: field.TypeInt, Comment: "用户ID"},
		{Name: "company", Type: field.TypeString, Comment: "公司名称", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "role", Type: field.TypeString, Comment: "职位名称", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "location", Type: field.TypeString, Comment: "公司地址", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "start", Type: field.TypeInt64, Comment: "开始时间"},
		{Name: "end", Type: field.TypeInt64, Comment: "结束时间"},
		{Name: "description", Type: field.TypeString, Comment: "职位描述", SchemaType: map[string]string{"mysql": "tinytext"}},
		{Name: "responsibilities", Type: field.TypeString, Comment: "主要职责", SchemaType: map[string]string{"mysql": "tinytext"}},
		{Name: "achievements", Type: field.TypeString, Comment: "工作成就", SchemaType: map[string]string{"mysql": "tinytext"}},
		{Name: "skills", Type: field.TypeJSON, Comment: "使用技能"},
		{Name: "project", Type: field.TypeInt, Comment: "项目数", Default: 0},
		{Name: "image", Type: field.TypeString, Comment: "公司名称", Default: "863f7821fa42eb9d61091b5c6df1c4b0", SchemaType: map[string]string{"mysql": "varchar(255)"}},
	}
	// UserExperienceTable holds the schema information for the "user_experience" table.
	UserExperienceTable = &schema.Table{
		Name:       "user_experience",
		Comment:    "经历表",
		Columns:    UserExperienceColumns,
		PrimaryKey: []*schema.Column{UserExperienceColumns[0]},
	}
	// UserProjectColumns holds the columns for the "user_project" table.
	UserProjectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, Comment: "ID"},
		{Name: "created_at", Type: field.TypeInt64, Comment: "创建时间", Default: 1736627068},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建人", Default: 0},
		{Name: "updated_at", Type: field.TypeInt64, Comment: "更新时间", Default: 1736627068},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新人", Default: 0},
		{Name: "deleted_at", Type: field.TypeInt64, Comment: "软删除时间", Default: 0},
		{Name: "deleted_by", Type: field.TypeInt64, Comment: "删除人", Default: 0},
		{Name: "user_id", Type: field.TypeInt, Comment: "用户ID"},
		{Name: "experience_id", Type: field.TypeInt, Comment: "经历ID"},
		{Name: "title", Type: field.TypeString, Comment: "项目名称", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "description", Type: field.TypeString, Comment: "项目描述", SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "skills", Type: field.TypeJSON, Comment: "使用技能"},
		{Name: "start", Type: field.TypeInt64, Comment: "开始时间"},
		{Name: "end", Type: field.TypeInt64, Comment: "结束时间"},
		{Name: "link", Type: field.TypeString, Comment: "项目地址", SchemaType: map[string]string{"mysql": "varchar(200)"}},
		{Name: "photos", Type: field.TypeJSON, Comment: "项目照片"},
	}
	// UserProjectTable holds the schema information for the "user_project" table.
	UserProjectTable = &schema.Table{
		Name:       "user_project",
		Comment:    "项目表",
		Columns:    UserProjectColumns,
		PrimaryKey: []*schema.Column{UserProjectColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountTable,
		BlogsTable,
		BlogsCommentsTable,
		BlogsContentTable,
		FilesTable,
		FilesExtendTable,
		TravelExtendsTable,
		TravelsTable,
		UserTable,
		UserExperienceTable,
		UserProjectTable,
	}
)

func init() {
	AccountTable.Annotation = &entsql.Annotation{
		Table: "account",
	}
	BlogsTable.Annotation = &entsql.Annotation{
		Table: "blogs",
	}
	BlogsCommentsTable.Annotation = &entsql.Annotation{
		Table: "blogs_comments",
	}
	BlogsContentTable.Annotation = &entsql.Annotation{
		Table: "blogs_content",
	}
	FilesTable.Annotation = &entsql.Annotation{
		Table: "files",
	}
	FilesExtendTable.ForeignKeys[0].RefTable = FilesTable
	FilesExtendTable.Annotation = &entsql.Annotation{
		Table: "files_extend",
	}
	TravelExtendsTable.ForeignKeys[0].RefTable = TravelsTable
	TravelExtendsTable.Annotation = &entsql.Annotation{
		Table: "travel_extends",
	}
	TravelsTable.ForeignKeys[0].RefTable = AccountTable
	TravelsTable.Annotation = &entsql.Annotation{
		Table: "travels",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	UserExperienceTable.Annotation = &entsql.Annotation{
		Table: "user_experience",
	}
	UserProjectTable.Annotation = &entsql.Annotation{
		Table: "user_project",
	}
}
