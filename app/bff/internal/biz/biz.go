package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBlogsUseCase, NewAccountUseCase, NewTravelUseCase,
	NewToolUseCase, NewUserUseCase, NewPalacesUseCase)
