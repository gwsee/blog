// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"blog/internal/ent/schema\",\"Package\":\"blog/internal/ent\",\"Schemas\":[{\"name\":\"Account\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建时间\"},{\"name\":\"created_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建人\"},{\"name\":\"updated_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"update_default\":true,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新时间\"},{\"name\":\"updated_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新人\"},{\"name\":\"is_deleted\",\"type\":{\"Type\":14,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":8,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"是否删除;0：正常，1：删除\"},{\"name\":\"deleted_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":6,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"软删除时间\"},{\"name\":\"deleted_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"删除人\"},{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"账户ID\"},{\"name\":\"account\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(100)\"},\"comment\":\"账户\"},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(200)\"},\"comment\":\"密码\"},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(100)\"},\"comment\":\"邮箱\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":3,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"tinyint\"},\"comment\":\"状态:0失效,1正常\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"Comment\":{\"Text\":\"站点\"},\"EntSQL\":{\"table\":\"account\",\"with_comments\":true}}},{\"name\":\"Blog\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建时间\"},{\"name\":\"created_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建人\"},{\"name\":\"updated_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"update_default\":true,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新时间\"},{\"name\":\"updated_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新人\"},{\"name\":\"is_deleted\",\"type\":{\"Type\":14,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":8,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"是否删除;0：正常，1：删除\"},{\"name\":\"deleted_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":6,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"软删除时间\"},{\"name\":\"deleted_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"删除人\"},{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"博客ID\"},{\"name\":\"account_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"账户ID\"},{\"name\":\"title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(100)\"},\"comment\":\"标题\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(255)\"},\"comment\":\"描述\"},{\"name\":\"is_hidden\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":3,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"tinyint\"},\"comment\":\"是否隐藏:0否,1是\"},{\"name\":\"tags\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"标签\"},{\"name\":\"cover\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"varchar(200)\"},\"comment\":\"封面\"},{\"name\":\"content\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"text\"},\"comment\":\"内容\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"Comment\":{\"Text\":\"博客\"},\"EntSQL\":{\"table\":\"blogs\",\"with_comments\":true}}},{\"name\":\"BlogContent\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"博客ID\"},{\"name\":\"content\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"text\"},\"comment\":\"内容\"}],\"annotations\":{\"Comment\":{\"Text\":\"博客内容\"},\"EntSQL\":{\"table\":\"blog_content\",\"with_comments\":true}}},{\"name\":\"Comment\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建时间\"},{\"name\":\"created_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"创建人\"},{\"name\":\"updated_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1730003847,\"default_kind\":6,\"update_default\":true,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新时间\"},{\"name\":\"updated_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0},\"comment\":\"更新人\"},{\"name\":\"is_deleted\",\"type\":{\"Type\":14,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":8,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"是否删除;0：正常，1：删除\"},{\"name\":\"deleted_at\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":6,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"软删除时间\"},{\"name\":\"deleted_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"删除人\"},{\"name\":\"account_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"账户ID\"},{\"name\":\"blog_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"对应类型的ID\"},{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"评论ID\"},{\"name\":\"top_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"顶级ID\"},{\"name\":\"parent_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"父评论\"},{\"name\":\"level\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"第几楼\"},{\"name\":\"total\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"当前级有多少数据;状态未0的\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":3,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"tinyint\"},\"comment\":\"0显示,1隐藏\"},{\"name\":\"content\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"评论内容\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"Comment\":{\"Text\":\"博客评论\"},\"EntSQL\":{\"table\":\"comments\",\"with_comments\":true}}}],\"Features\":[\"intercept\",\"schema/snapshot\",\"sql/execquery\",\"sql/upsert\"]}"
