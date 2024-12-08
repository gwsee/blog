package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate  --feature intercept,schema/snapshot,sql/execquery,sql/upsert ./schema
// 在最外面 blog 根目录执行 go run -mod=mod entgo.io/ent/cmd/ent generate  --feature intercept,schema/snapshot,sql/execquery,sql/upsert ./internal/ent/schema  //上面不知道为什么出现问题了
//import (
//	_ "blog/internal/ent/runtime"
//)

//./schema/paging ./schema/tracker

// 需要将 runtime目录内的内容手动copy到外面    或者 直接import 到外面避免循环引用
