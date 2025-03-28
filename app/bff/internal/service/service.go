package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountService, NewBlogsService, NewBlogsCommentService,
	NewTravelService, NewToolsService, NewUserService, NewPalacesService)
