// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"blog/internal/ent/account"
	"blog/internal/ent/blogs"
	"blog/internal/ent/blogscomment"
	"blog/internal/ent/blogscontent"
	"blog/internal/ent/files"
	"blog/internal/ent/filesextend"
	"blog/internal/ent/schema"
	"blog/internal/ent/travel"
	"blog/internal/ent/travelextend"
	"blog/internal/ent/user"
	"blog/internal/ent/userexperience"
	"blog/internal/ent/userproject"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	accountMixinHooks0 := accountMixin[0].Hooks()
	accountMixinHooks1 := accountMixin[1].Hooks()
	account.Hooks[0] = accountMixinHooks0[0]
	account.Hooks[1] = accountMixinHooks0[1]
	account.Hooks[2] = accountMixinHooks1[0]
	accountMixinInters1 := accountMixin[1].Interceptors()
	account.Interceptors[0] = accountMixinInters1[0]
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountMixinFields1 := accountMixin[1].Fields()
	_ = accountMixinFields1
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountMixinFields0[0].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(int64)
	// accountDescCreatedBy is the schema descriptor for created_by field.
	accountDescCreatedBy := accountMixinFields0[1].Descriptor()
	// account.DefaultCreatedBy holds the default value on creation for the created_by field.
	account.DefaultCreatedBy = accountDescCreatedBy.Default.(int64)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountMixinFields0[2].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(int64)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() int64)
	// accountDescUpdatedBy is the schema descriptor for updated_by field.
	accountDescUpdatedBy := accountMixinFields0[3].Descriptor()
	// account.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	account.DefaultUpdatedBy = accountDescUpdatedBy.Default.(int64)
	// accountDescDeletedAt is the schema descriptor for deleted_at field.
	accountDescDeletedAt := accountMixinFields1[0].Descriptor()
	// account.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	account.DefaultDeletedAt = accountDescDeletedAt.Default.(int64)
	// accountDescDeletedBy is the schema descriptor for deleted_by field.
	accountDescDeletedBy := accountMixinFields1[1].Descriptor()
	// account.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	account.DefaultDeletedBy = accountDescDeletedBy.Default.(int64)
	// accountDescNickname is the schema descriptor for nickname field.
	accountDescNickname := accountFields[1].Descriptor()
	// account.DefaultNickname holds the default value on creation for the nickname field.
	account.DefaultNickname = accountDescNickname.Default.(string)
	// accountDescAccount is the schema descriptor for account field.
	accountDescAccount := accountFields[2].Descriptor()
	// account.AccountValidator is a validator for the "account" field. It is called by the builders before save.
	account.AccountValidator = accountDescAccount.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[3].Descriptor()
	// account.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = accountDescPassword.Validators[0].(func(string) error)
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[4].Descriptor()
	// account.DefaultEmail holds the default value on creation for the email field.
	account.DefaultEmail = accountDescEmail.Default.(string)
	// accountDescDescription is the schema descriptor for description field.
	accountDescDescription := accountFields[5].Descriptor()
	// account.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	account.DescriptionValidator = accountDescDescription.Validators[0].(func(string) error)
	// accountDescAvatar is the schema descriptor for avatar field.
	accountDescAvatar := accountFields[6].Descriptor()
	// account.DefaultAvatar holds the default value on creation for the avatar field.
	account.DefaultAvatar = accountDescAvatar.Default.(string)
	// accountDescBlogNum is the schema descriptor for blog_num field.
	accountDescBlogNum := accountFields[7].Descriptor()
	// account.DefaultBlogNum holds the default value on creation for the blog_num field.
	account.DefaultBlogNum = accountDescBlogNum.Default.(int)
	// accountDescStatus is the schema descriptor for status field.
	accountDescStatus := accountFields[8].Descriptor()
	// account.DefaultStatus holds the default value on creation for the status field.
	account.DefaultStatus = accountDescStatus.Default.(int8)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.IDValidator is a validator for the "id" field. It is called by the builders before save.
	account.IDValidator = accountDescID.Validators[0].(func(int) error)
	blogsMixin := schema.Blogs{}.Mixin()
	blogsMixinHooks0 := blogsMixin[0].Hooks()
	blogsMixinHooks1 := blogsMixin[1].Hooks()
	blogs.Hooks[0] = blogsMixinHooks0[0]
	blogs.Hooks[1] = blogsMixinHooks0[1]
	blogs.Hooks[2] = blogsMixinHooks1[0]
	blogsMixinInters1 := blogsMixin[1].Interceptors()
	blogs.Interceptors[0] = blogsMixinInters1[0]
	blogsMixinFields0 := blogsMixin[0].Fields()
	_ = blogsMixinFields0
	blogsMixinFields1 := blogsMixin[1].Fields()
	_ = blogsMixinFields1
	blogsFields := schema.Blogs{}.Fields()
	_ = blogsFields
	// blogsDescCreatedAt is the schema descriptor for created_at field.
	blogsDescCreatedAt := blogsMixinFields0[0].Descriptor()
	// blogs.DefaultCreatedAt holds the default value on creation for the created_at field.
	blogs.DefaultCreatedAt = blogsDescCreatedAt.Default.(int64)
	// blogsDescCreatedBy is the schema descriptor for created_by field.
	blogsDescCreatedBy := blogsMixinFields0[1].Descriptor()
	// blogs.DefaultCreatedBy holds the default value on creation for the created_by field.
	blogs.DefaultCreatedBy = blogsDescCreatedBy.Default.(int64)
	// blogsDescUpdatedAt is the schema descriptor for updated_at field.
	blogsDescUpdatedAt := blogsMixinFields0[2].Descriptor()
	// blogs.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	blogs.DefaultUpdatedAt = blogsDescUpdatedAt.Default.(int64)
	// blogs.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	blogs.UpdateDefaultUpdatedAt = blogsDescUpdatedAt.UpdateDefault.(func() int64)
	// blogsDescUpdatedBy is the schema descriptor for updated_by field.
	blogsDescUpdatedBy := blogsMixinFields0[3].Descriptor()
	// blogs.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	blogs.DefaultUpdatedBy = blogsDescUpdatedBy.Default.(int64)
	// blogsDescDeletedAt is the schema descriptor for deleted_at field.
	blogsDescDeletedAt := blogsMixinFields1[0].Descriptor()
	// blogs.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	blogs.DefaultDeletedAt = blogsDescDeletedAt.Default.(int64)
	// blogsDescDeletedBy is the schema descriptor for deleted_by field.
	blogsDescDeletedBy := blogsMixinFields1[1].Descriptor()
	// blogs.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	blogs.DefaultDeletedBy = blogsDescDeletedBy.Default.(int64)
	// blogsDescIsHidden is the schema descriptor for is_hidden field.
	blogsDescIsHidden := blogsFields[4].Descriptor()
	// blogs.DefaultIsHidden holds the default value on creation for the is_hidden field.
	blogs.DefaultIsHidden = blogsDescIsHidden.Default.(int8)
	// blogsDescID is the schema descriptor for id field.
	blogsDescID := blogsFields[0].Descriptor()
	// blogs.IDValidator is a validator for the "id" field. It is called by the builders before save.
	blogs.IDValidator = blogsDescID.Validators[0].(func(int) error)
	blogscommentMixin := schema.BlogsComment{}.Mixin()
	blogscommentMixinHooks0 := blogscommentMixin[0].Hooks()
	blogscommentMixinHooks1 := blogscommentMixin[1].Hooks()
	blogscomment.Hooks[0] = blogscommentMixinHooks0[0]
	blogscomment.Hooks[1] = blogscommentMixinHooks0[1]
	blogscomment.Hooks[2] = blogscommentMixinHooks1[0]
	blogscommentMixinInters1 := blogscommentMixin[1].Interceptors()
	blogscomment.Interceptors[0] = blogscommentMixinInters1[0]
	blogscommentMixinFields0 := blogscommentMixin[0].Fields()
	_ = blogscommentMixinFields0
	blogscommentMixinFields1 := blogscommentMixin[1].Fields()
	_ = blogscommentMixinFields1
	blogscommentFields := schema.BlogsComment{}.Fields()
	_ = blogscommentFields
	// blogscommentDescCreatedAt is the schema descriptor for created_at field.
	blogscommentDescCreatedAt := blogscommentMixinFields0[0].Descriptor()
	// blogscomment.DefaultCreatedAt holds the default value on creation for the created_at field.
	blogscomment.DefaultCreatedAt = blogscommentDescCreatedAt.Default.(int64)
	// blogscommentDescCreatedBy is the schema descriptor for created_by field.
	blogscommentDescCreatedBy := blogscommentMixinFields0[1].Descriptor()
	// blogscomment.DefaultCreatedBy holds the default value on creation for the created_by field.
	blogscomment.DefaultCreatedBy = blogscommentDescCreatedBy.Default.(int64)
	// blogscommentDescUpdatedAt is the schema descriptor for updated_at field.
	blogscommentDescUpdatedAt := blogscommentMixinFields0[2].Descriptor()
	// blogscomment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	blogscomment.DefaultUpdatedAt = blogscommentDescUpdatedAt.Default.(int64)
	// blogscomment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	blogscomment.UpdateDefaultUpdatedAt = blogscommentDescUpdatedAt.UpdateDefault.(func() int64)
	// blogscommentDescUpdatedBy is the schema descriptor for updated_by field.
	blogscommentDescUpdatedBy := blogscommentMixinFields0[3].Descriptor()
	// blogscomment.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	blogscomment.DefaultUpdatedBy = blogscommentDescUpdatedBy.Default.(int64)
	// blogscommentDescDeletedAt is the schema descriptor for deleted_at field.
	blogscommentDescDeletedAt := blogscommentMixinFields1[0].Descriptor()
	// blogscomment.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	blogscomment.DefaultDeletedAt = blogscommentDescDeletedAt.Default.(int64)
	// blogscommentDescDeletedBy is the schema descriptor for deleted_by field.
	blogscommentDescDeletedBy := blogscommentMixinFields1[1].Descriptor()
	// blogscomment.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	blogscomment.DefaultDeletedBy = blogscommentDescDeletedBy.Default.(int64)
	// blogscommentDescAccountID is the schema descriptor for account_id field.
	blogscommentDescAccountID := blogscommentFields[0].Descriptor()
	// blogscomment.AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	blogscomment.AccountIDValidator = blogscommentDescAccountID.Validators[0].(func(int) error)
	// blogscommentDescBlogID is the schema descriptor for blog_id field.
	blogscommentDescBlogID := blogscommentFields[1].Descriptor()
	// blogscomment.BlogIDValidator is a validator for the "blog_id" field. It is called by the builders before save.
	blogscomment.BlogIDValidator = blogscommentDescBlogID.Validators[0].(func(int) error)
	// blogscommentDescTopID is the schema descriptor for top_id field.
	blogscommentDescTopID := blogscommentFields[3].Descriptor()
	// blogscomment.DefaultTopID holds the default value on creation for the top_id field.
	blogscomment.DefaultTopID = blogscommentDescTopID.Default.(int)
	// blogscommentDescParentID is the schema descriptor for parent_id field.
	blogscommentDescParentID := blogscommentFields[4].Descriptor()
	// blogscomment.DefaultParentID holds the default value on creation for the parent_id field.
	blogscomment.DefaultParentID = blogscommentDescParentID.Default.(int)
	// blogscommentDescLevel is the schema descriptor for level field.
	blogscommentDescLevel := blogscommentFields[5].Descriptor()
	// blogscomment.DefaultLevel holds the default value on creation for the level field.
	blogscomment.DefaultLevel = blogscommentDescLevel.Default.(int)
	// blogscommentDescTotal is the schema descriptor for total field.
	blogscommentDescTotal := blogscommentFields[6].Descriptor()
	// blogscomment.DefaultTotal holds the default value on creation for the total field.
	blogscomment.DefaultTotal = blogscommentDescTotal.Default.(int)
	// blogscommentDescStatus is the schema descriptor for status field.
	blogscommentDescStatus := blogscommentFields[7].Descriptor()
	// blogscomment.DefaultStatus holds the default value on creation for the status field.
	blogscomment.DefaultStatus = blogscommentDescStatus.Default.(int8)
	// blogscommentDescID is the schema descriptor for id field.
	blogscommentDescID := blogscommentFields[2].Descriptor()
	// blogscomment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	blogscomment.IDValidator = blogscommentDescID.Validators[0].(func(int) error)
	blogscontentFields := schema.BlogsContent{}.Fields()
	_ = blogscontentFields
	// blogscontentDescContent is the schema descriptor for content field.
	blogscontentDescContent := blogscontentFields[1].Descriptor()
	// blogscontent.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	blogscontent.ContentValidator = blogscontentDescContent.Validators[0].(func(string) error)
	filesMixin := schema.Files{}.Mixin()
	filesMixinHooks0 := filesMixin[0].Hooks()
	filesMixinHooks1 := filesMixin[1].Hooks()
	files.Hooks[0] = filesMixinHooks0[0]
	files.Hooks[1] = filesMixinHooks0[1]
	files.Hooks[2] = filesMixinHooks1[0]
	filesMixinInters1 := filesMixin[1].Interceptors()
	files.Interceptors[0] = filesMixinInters1[0]
	filesMixinFields0 := filesMixin[0].Fields()
	_ = filesMixinFields0
	filesMixinFields1 := filesMixin[1].Fields()
	_ = filesMixinFields1
	filesFields := schema.Files{}.Fields()
	_ = filesFields
	// filesDescCreatedAt is the schema descriptor for created_at field.
	filesDescCreatedAt := filesMixinFields0[0].Descriptor()
	// files.DefaultCreatedAt holds the default value on creation for the created_at field.
	files.DefaultCreatedAt = filesDescCreatedAt.Default.(int64)
	// filesDescCreatedBy is the schema descriptor for created_by field.
	filesDescCreatedBy := filesMixinFields0[1].Descriptor()
	// files.DefaultCreatedBy holds the default value on creation for the created_by field.
	files.DefaultCreatedBy = filesDescCreatedBy.Default.(int64)
	// filesDescUpdatedAt is the schema descriptor for updated_at field.
	filesDescUpdatedAt := filesMixinFields0[2].Descriptor()
	// files.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	files.DefaultUpdatedAt = filesDescUpdatedAt.Default.(int64)
	// files.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	files.UpdateDefaultUpdatedAt = filesDescUpdatedAt.UpdateDefault.(func() int64)
	// filesDescUpdatedBy is the schema descriptor for updated_by field.
	filesDescUpdatedBy := filesMixinFields0[3].Descriptor()
	// files.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	files.DefaultUpdatedBy = filesDescUpdatedBy.Default.(int64)
	// filesDescDeletedAt is the schema descriptor for deleted_at field.
	filesDescDeletedAt := filesMixinFields1[0].Descriptor()
	// files.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	files.DefaultDeletedAt = filesDescDeletedAt.Default.(int64)
	// filesDescDeletedBy is the schema descriptor for deleted_by field.
	filesDescDeletedBy := filesMixinFields1[1].Descriptor()
	// files.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	files.DefaultDeletedBy = filesDescDeletedBy.Default.(int64)
	// filesDescType is the schema descriptor for type field.
	filesDescType := filesFields[1].Descriptor()
	// files.DefaultType holds the default value on creation for the type field.
	files.DefaultType = filesDescType.Default.(string)
	// filesDescSize is the schema descriptor for size field.
	filesDescSize := filesFields[2].Descriptor()
	// files.DefaultSize holds the default value on creation for the size field.
	files.DefaultSize = filesDescSize.Default.(int64)
	// filesDescName is the schema descriptor for name field.
	filesDescName := filesFields[3].Descriptor()
	// files.DefaultName holds the default value on creation for the name field.
	files.DefaultName = filesDescName.Default.(string)
	// filesDescPath is the schema descriptor for path field.
	filesDescPath := filesFields[4].Descriptor()
	// files.DefaultPath holds the default value on creation for the path field.
	files.DefaultPath = filesDescPath.Default.(string)
	filesextendMixin := schema.FilesExtend{}.Mixin()
	filesextendMixinHooks0 := filesextendMixin[0].Hooks()
	filesextendMixinHooks1 := filesextendMixin[1].Hooks()
	filesextend.Hooks[0] = filesextendMixinHooks0[0]
	filesextend.Hooks[1] = filesextendMixinHooks0[1]
	filesextend.Hooks[2] = filesextendMixinHooks1[0]
	filesextendMixinInters1 := filesextendMixin[1].Interceptors()
	filesextend.Interceptors[0] = filesextendMixinInters1[0]
	filesextendMixinFields0 := filesextendMixin[0].Fields()
	_ = filesextendMixinFields0
	filesextendMixinFields1 := filesextendMixin[1].Fields()
	_ = filesextendMixinFields1
	filesextendFields := schema.FilesExtend{}.Fields()
	_ = filesextendFields
	// filesextendDescCreatedAt is the schema descriptor for created_at field.
	filesextendDescCreatedAt := filesextendMixinFields0[0].Descriptor()
	// filesextend.DefaultCreatedAt holds the default value on creation for the created_at field.
	filesextend.DefaultCreatedAt = filesextendDescCreatedAt.Default.(int64)
	// filesextendDescCreatedBy is the schema descriptor for created_by field.
	filesextendDescCreatedBy := filesextendMixinFields0[1].Descriptor()
	// filesextend.DefaultCreatedBy holds the default value on creation for the created_by field.
	filesextend.DefaultCreatedBy = filesextendDescCreatedBy.Default.(int64)
	// filesextendDescUpdatedAt is the schema descriptor for updated_at field.
	filesextendDescUpdatedAt := filesextendMixinFields0[2].Descriptor()
	// filesextend.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	filesextend.DefaultUpdatedAt = filesextendDescUpdatedAt.Default.(int64)
	// filesextend.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	filesextend.UpdateDefaultUpdatedAt = filesextendDescUpdatedAt.UpdateDefault.(func() int64)
	// filesextendDescUpdatedBy is the schema descriptor for updated_by field.
	filesextendDescUpdatedBy := filesextendMixinFields0[3].Descriptor()
	// filesextend.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	filesextend.DefaultUpdatedBy = filesextendDescUpdatedBy.Default.(int64)
	// filesextendDescDeletedAt is the schema descriptor for deleted_at field.
	filesextendDescDeletedAt := filesextendMixinFields1[0].Descriptor()
	// filesextend.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	filesextend.DefaultDeletedAt = filesextendDescDeletedAt.Default.(int64)
	// filesextendDescDeletedBy is the schema descriptor for deleted_by field.
	filesextendDescDeletedBy := filesextendMixinFields1[1].Descriptor()
	// filesextend.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	filesextend.DefaultDeletedBy = filesextendDescDeletedBy.Default.(int64)
	// filesextendDescFileID is the schema descriptor for file_id field.
	filesextendDescFileID := filesextendFields[1].Descriptor()
	// filesextend.DefaultFileID holds the default value on creation for the file_id field.
	filesextend.DefaultFileID = filesextendDescFileID.Default.(string)
	// filesextendDescUserID is the schema descriptor for user_id field.
	filesextendDescUserID := filesextendFields[2].Descriptor()
	// filesextend.DefaultUserID holds the default value on creation for the user_id field.
	filesextend.DefaultUserID = filesextendDescUserID.Default.(int)
	// filesextendDescID is the schema descriptor for id field.
	filesextendDescID := filesextendFields[0].Descriptor()
	// filesextend.IDValidator is a validator for the "id" field. It is called by the builders before save.
	filesextend.IDValidator = filesextendDescID.Validators[0].(func(int) error)
	travelMixin := schema.Travel{}.Mixin()
	travelMixinHooks0 := travelMixin[0].Hooks()
	travelMixinHooks1 := travelMixin[1].Hooks()
	travel.Hooks[0] = travelMixinHooks0[0]
	travel.Hooks[1] = travelMixinHooks0[1]
	travel.Hooks[2] = travelMixinHooks1[0]
	travelMixinInters1 := travelMixin[1].Interceptors()
	travel.Interceptors[0] = travelMixinInters1[0]
	travelMixinFields0 := travelMixin[0].Fields()
	_ = travelMixinFields0
	travelMixinFields1 := travelMixin[1].Fields()
	_ = travelMixinFields1
	travelFields := schema.Travel{}.Fields()
	_ = travelFields
	// travelDescCreatedAt is the schema descriptor for created_at field.
	travelDescCreatedAt := travelMixinFields0[0].Descriptor()
	// travel.DefaultCreatedAt holds the default value on creation for the created_at field.
	travel.DefaultCreatedAt = travelDescCreatedAt.Default.(int64)
	// travelDescCreatedBy is the schema descriptor for created_by field.
	travelDescCreatedBy := travelMixinFields0[1].Descriptor()
	// travel.DefaultCreatedBy holds the default value on creation for the created_by field.
	travel.DefaultCreatedBy = travelDescCreatedBy.Default.(int64)
	// travelDescUpdatedAt is the schema descriptor for updated_at field.
	travelDescUpdatedAt := travelMixinFields0[2].Descriptor()
	// travel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	travel.DefaultUpdatedAt = travelDescUpdatedAt.Default.(int64)
	// travel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	travel.UpdateDefaultUpdatedAt = travelDescUpdatedAt.UpdateDefault.(func() int64)
	// travelDescUpdatedBy is the schema descriptor for updated_by field.
	travelDescUpdatedBy := travelMixinFields0[3].Descriptor()
	// travel.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	travel.DefaultUpdatedBy = travelDescUpdatedBy.Default.(int64)
	// travelDescDeletedAt is the schema descriptor for deleted_at field.
	travelDescDeletedAt := travelMixinFields1[0].Descriptor()
	// travel.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	travel.DefaultDeletedAt = travelDescDeletedAt.Default.(int64)
	// travelDescDeletedBy is the schema descriptor for deleted_by field.
	travelDescDeletedBy := travelMixinFields1[1].Descriptor()
	// travel.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	travel.DefaultDeletedBy = travelDescDeletedBy.Default.(int64)
	// travelDescTitle is the schema descriptor for title field.
	travelDescTitle := travelFields[1].Descriptor()
	// travel.DefaultTitle holds the default value on creation for the title field.
	travel.DefaultTitle = travelDescTitle.Default.(string)
	// travelDescDescription is the schema descriptor for description field.
	travelDescDescription := travelFields[2].Descriptor()
	// travel.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	travel.DescriptionValidator = travelDescDescription.Validators[0].(func(string) error)
	// travelDescVideo is the schema descriptor for video field.
	travelDescVideo := travelFields[3].Descriptor()
	// travel.VideoValidator is a validator for the "video" field. It is called by the builders before save.
	travel.VideoValidator = travelDescVideo.Validators[0].(func(string) error)
	// travelDescIsHidden is the schema descriptor for is_hidden field.
	travelDescIsHidden := travelFields[4].Descriptor()
	// travel.DefaultIsHidden holds the default value on creation for the is_hidden field.
	travel.DefaultIsHidden = travelDescIsHidden.Default.(bool)
	// travelDescAccountID is the schema descriptor for account_id field.
	travelDescAccountID := travelFields[5].Descriptor()
	// travel.AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	travel.AccountIDValidator = travelDescAccountID.Validators[0].(func(int) error)
	travelextendMixin := schema.TravelExtend{}.Mixin()
	travelextendMixinHooks0 := travelextendMixin[0].Hooks()
	travelextendMixinHooks1 := travelextendMixin[1].Hooks()
	travelextend.Hooks[0] = travelextendMixinHooks0[0]
	travelextend.Hooks[1] = travelextendMixinHooks0[1]
	travelextend.Hooks[2] = travelextendMixinHooks1[0]
	travelextendMixinInters1 := travelextendMixin[1].Interceptors()
	travelextend.Interceptors[0] = travelextendMixinInters1[0]
	travelextendMixinFields0 := travelextendMixin[0].Fields()
	_ = travelextendMixinFields0
	travelextendMixinFields1 := travelextendMixin[1].Fields()
	_ = travelextendMixinFields1
	travelextendFields := schema.TravelExtend{}.Fields()
	_ = travelextendFields
	// travelextendDescCreatedAt is the schema descriptor for created_at field.
	travelextendDescCreatedAt := travelextendMixinFields0[0].Descriptor()
	// travelextend.DefaultCreatedAt holds the default value on creation for the created_at field.
	travelextend.DefaultCreatedAt = travelextendDescCreatedAt.Default.(int64)
	// travelextendDescCreatedBy is the schema descriptor for created_by field.
	travelextendDescCreatedBy := travelextendMixinFields0[1].Descriptor()
	// travelextend.DefaultCreatedBy holds the default value on creation for the created_by field.
	travelextend.DefaultCreatedBy = travelextendDescCreatedBy.Default.(int64)
	// travelextendDescUpdatedAt is the schema descriptor for updated_at field.
	travelextendDescUpdatedAt := travelextendMixinFields0[2].Descriptor()
	// travelextend.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	travelextend.DefaultUpdatedAt = travelextendDescUpdatedAt.Default.(int64)
	// travelextend.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	travelextend.UpdateDefaultUpdatedAt = travelextendDescUpdatedAt.UpdateDefault.(func() int64)
	// travelextendDescUpdatedBy is the schema descriptor for updated_by field.
	travelextendDescUpdatedBy := travelextendMixinFields0[3].Descriptor()
	// travelextend.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	travelextend.DefaultUpdatedBy = travelextendDescUpdatedBy.Default.(int64)
	// travelextendDescDeletedAt is the schema descriptor for deleted_at field.
	travelextendDescDeletedAt := travelextendMixinFields1[0].Descriptor()
	// travelextend.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	travelextend.DefaultDeletedAt = travelextendDescDeletedAt.Default.(int64)
	// travelextendDescDeletedBy is the schema descriptor for deleted_by field.
	travelextendDescDeletedBy := travelextendMixinFields1[1].Descriptor()
	// travelextend.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	travelextend.DefaultDeletedBy = travelextendDescDeletedBy.Default.(int64)
	// travelextendDescAccountID is the schema descriptor for account_id field.
	travelextendDescAccountID := travelextendFields[0].Descriptor()
	// travelextend.AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	travelextend.AccountIDValidator = travelextendDescAccountID.Validators[0].(func(int) error)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	userMixinHooks1 := userMixin[1].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	user.Hooks[1] = userMixinHooks0[1]
	user.Hooks[2] = userMixinHooks1[0]
	userMixinInters1 := userMixin[1].Interceptors()
	user.Interceptors[0] = userMixinInters1[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(int64)
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(int64)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(int64)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() int64)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userMixinFields0[3].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(int64)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields1[0].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(int64)
	// userDescDeletedBy is the schema descriptor for deleted_by field.
	userDescDeletedBy := userMixinFields1[1].Descriptor()
	// user.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	user.DefaultDeletedBy = userDescDeletedBy.Default.(int64)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[2].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescProfessional is the schema descriptor for professional field.
	userDescProfessional := userFields[4].Descriptor()
	// user.DefaultProfessional holds the default value on creation for the professional field.
	user.DefaultProfessional = userDescProfessional.Default.(string)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[5].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(string)
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[7].Descriptor()
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = userDescDescription.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int) error)
	userexperienceMixin := schema.UserExperience{}.Mixin()
	userexperienceMixinHooks0 := userexperienceMixin[0].Hooks()
	userexperienceMixinHooks1 := userexperienceMixin[1].Hooks()
	userexperience.Hooks[0] = userexperienceMixinHooks0[0]
	userexperience.Hooks[1] = userexperienceMixinHooks0[1]
	userexperience.Hooks[2] = userexperienceMixinHooks1[0]
	userexperienceMixinInters1 := userexperienceMixin[1].Interceptors()
	userexperience.Interceptors[0] = userexperienceMixinInters1[0]
	userexperienceMixinFields0 := userexperienceMixin[0].Fields()
	_ = userexperienceMixinFields0
	userexperienceMixinFields1 := userexperienceMixin[1].Fields()
	_ = userexperienceMixinFields1
	userexperienceFields := schema.UserExperience{}.Fields()
	_ = userexperienceFields
	// userexperienceDescCreatedAt is the schema descriptor for created_at field.
	userexperienceDescCreatedAt := userexperienceMixinFields0[0].Descriptor()
	// userexperience.DefaultCreatedAt holds the default value on creation for the created_at field.
	userexperience.DefaultCreatedAt = userexperienceDescCreatedAt.Default.(int64)
	// userexperienceDescCreatedBy is the schema descriptor for created_by field.
	userexperienceDescCreatedBy := userexperienceMixinFields0[1].Descriptor()
	// userexperience.DefaultCreatedBy holds the default value on creation for the created_by field.
	userexperience.DefaultCreatedBy = userexperienceDescCreatedBy.Default.(int64)
	// userexperienceDescUpdatedAt is the schema descriptor for updated_at field.
	userexperienceDescUpdatedAt := userexperienceMixinFields0[2].Descriptor()
	// userexperience.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userexperience.DefaultUpdatedAt = userexperienceDescUpdatedAt.Default.(int64)
	// userexperience.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userexperience.UpdateDefaultUpdatedAt = userexperienceDescUpdatedAt.UpdateDefault.(func() int64)
	// userexperienceDescUpdatedBy is the schema descriptor for updated_by field.
	userexperienceDescUpdatedBy := userexperienceMixinFields0[3].Descriptor()
	// userexperience.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	userexperience.DefaultUpdatedBy = userexperienceDescUpdatedBy.Default.(int64)
	// userexperienceDescDeletedAt is the schema descriptor for deleted_at field.
	userexperienceDescDeletedAt := userexperienceMixinFields1[0].Descriptor()
	// userexperience.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	userexperience.DefaultDeletedAt = userexperienceDescDeletedAt.Default.(int64)
	// userexperienceDescDeletedBy is the schema descriptor for deleted_by field.
	userexperienceDescDeletedBy := userexperienceMixinFields1[1].Descriptor()
	// userexperience.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	userexperience.DefaultDeletedBy = userexperienceDescDeletedBy.Default.(int64)
	// userexperienceDescID is the schema descriptor for id field.
	userexperienceDescID := userexperienceFields[0].Descriptor()
	// userexperience.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userexperience.IDValidator = userexperienceDescID.Validators[0].(func(int) error)
	userprojectMixin := schema.UserProject{}.Mixin()
	userprojectMixinHooks0 := userprojectMixin[0].Hooks()
	userprojectMixinHooks1 := userprojectMixin[1].Hooks()
	userproject.Hooks[0] = userprojectMixinHooks0[0]
	userproject.Hooks[1] = userprojectMixinHooks0[1]
	userproject.Hooks[2] = userprojectMixinHooks1[0]
	userprojectMixinInters1 := userprojectMixin[1].Interceptors()
	userproject.Interceptors[0] = userprojectMixinInters1[0]
	userprojectMixinFields0 := userprojectMixin[0].Fields()
	_ = userprojectMixinFields0
	userprojectMixinFields1 := userprojectMixin[1].Fields()
	_ = userprojectMixinFields1
	userprojectFields := schema.UserProject{}.Fields()
	_ = userprojectFields
	// userprojectDescCreatedAt is the schema descriptor for created_at field.
	userprojectDescCreatedAt := userprojectMixinFields0[0].Descriptor()
	// userproject.DefaultCreatedAt holds the default value on creation for the created_at field.
	userproject.DefaultCreatedAt = userprojectDescCreatedAt.Default.(int64)
	// userprojectDescCreatedBy is the schema descriptor for created_by field.
	userprojectDescCreatedBy := userprojectMixinFields0[1].Descriptor()
	// userproject.DefaultCreatedBy holds the default value on creation for the created_by field.
	userproject.DefaultCreatedBy = userprojectDescCreatedBy.Default.(int64)
	// userprojectDescUpdatedAt is the schema descriptor for updated_at field.
	userprojectDescUpdatedAt := userprojectMixinFields0[2].Descriptor()
	// userproject.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userproject.DefaultUpdatedAt = userprojectDescUpdatedAt.Default.(int64)
	// userproject.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userproject.UpdateDefaultUpdatedAt = userprojectDescUpdatedAt.UpdateDefault.(func() int64)
	// userprojectDescUpdatedBy is the schema descriptor for updated_by field.
	userprojectDescUpdatedBy := userprojectMixinFields0[3].Descriptor()
	// userproject.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	userproject.DefaultUpdatedBy = userprojectDescUpdatedBy.Default.(int64)
	// userprojectDescDeletedAt is the schema descriptor for deleted_at field.
	userprojectDescDeletedAt := userprojectMixinFields1[0].Descriptor()
	// userproject.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	userproject.DefaultDeletedAt = userprojectDescDeletedAt.Default.(int64)
	// userprojectDescDeletedBy is the schema descriptor for deleted_by field.
	userprojectDescDeletedBy := userprojectMixinFields1[1].Descriptor()
	// userproject.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	userproject.DefaultDeletedBy = userprojectDescDeletedBy.Default.(int64)
	// userprojectDescID is the schema descriptor for id field.
	userprojectDescID := userprojectFields[0].Descriptor()
	// userproject.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userproject.IDValidator = userprojectDescID.Validators[0].(func(int) error)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)
