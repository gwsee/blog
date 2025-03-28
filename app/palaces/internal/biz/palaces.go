package biz

import (
	"blog/api/global"
	pb "blog/api/palaces/v1"
	"blog/internal/constx"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

// Palaces is a Palaces model.
type Palaces struct {
	Hello string
}
type Memo struct {
	Id        int
	AccountId int
	Name      string
	Content   string
	Status    int8
	CreatedAt int64
	UpdatedAt int64
}
type MemoQuery struct {
	*global.PageInfo
	Name      string
	Status    *int64
	AccountId int
}

type Todo struct {
	Id        int
	Theme     string
	Type      int64
	From      int64
	To        int64
	Num       int64
	Sort      int64
	Content   string
	Status    int64
	CreatedAt int64
	UpdatedAt int64
	DoneId    int64
	AccountId int
}

type TodoQuery struct {
	*global.PageInfo
	AccountId int
	Theme     string
	Type      int64
	From      int64
	Status    *int64
}

// PalacesRepo is a Greater repo.
type PalacesRepo interface {
	CreatePalaces(context.Context, *Palaces) error
	UpdatePalaces(context.Context, *Palaces) error
	DeletePalaces(context.Context, *global.ID) error
	GetPalaces(context.Context, *global.ID) (*Palaces, error)
	ListPalaces(context.Context, *Palaces) (int64, []*Palaces, error)

	//备忘录

	SaveMemo(context.Context, *Memo) error
	DoneMemo(context.Context, *global.State) error
	DeleteMemo(context.Context, *global.ID) error
	ListMemo(context.Context, *MemoQuery) (int64, []*Memo, error)

	//TODOList 相关
	ValidTodo(context.Context, *Todo) error
	SaveTodo(context.Context, *Todo) error
	DoneTodo(context.Context, *global.ID) error
	DeleteTodo(context.Context, *global.ID) error
	DeleteTodoDone(context.Context, *global.ID) error
	DeleteTodoRecord(context.Context, *global.ID) error
	ListTodo(context.Context, *TodoQuery) (int64, []*Todo, error)
	ListTodoDone(context.Context, *TodoQuery) (int64, []*Todo, error)
}

// PalacesUseCase is a Palaces useCase.
type PalacesUseCase struct {
	repo PalacesRepo
	log  *log.Helper
}

// NewPalacesUseCase new a Palaces useCase.
func NewPalacesUseCase(repo PalacesRepo, logger log.Logger) *PalacesUseCase {
	return &PalacesUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *PalacesUseCase) CreatePalaces(ctx context.Context, req *pb.CreatePalacesRequest) (*pb.CreatePalacesReply, error) {
	return &pb.CreatePalacesReply{}, s.repo.CreatePalaces(ctx, &Palaces{})
}
func (s *PalacesUseCase) UpdatePalaces(ctx context.Context, req *pb.UpdatePalacesRequest) (*pb.UpdatePalacesReply, error) {
	return &pb.UpdatePalacesReply{}, s.repo.UpdatePalaces(ctx, &Palaces{})
}
func (s *PalacesUseCase) DeletePalaces(ctx context.Context, req *pb.DeletePalacesRequest) (*pb.DeletePalacesReply, error) {
	return &pb.DeletePalacesReply{}, s.repo.DeletePalaces(ctx, &global.ID{})
}
func (s *PalacesUseCase) GetPalaces(ctx context.Context, req *pb.GetPalacesRequest) (*pb.GetPalacesReply, error) {
	item, err := s.repo.GetPalaces(ctx, &global.ID{})
	if err != nil {
		return nil, err
	}
	fmt.Println(item.Hello)
	return &pb.GetPalacesReply{}, err
}
func (s *PalacesUseCase) ListPalaces(ctx context.Context, req *pb.ListPalacesRequest) (*pb.ListPalacesReply, error) {
	total, list, err := s.repo.ListPalaces(ctx, &Palaces{})
	fmt.Println(total, len(list))
	return &pb.ListPalacesReply{}, err
}

func (s *PalacesUseCase) SaveMemo(ctx context.Context, req *pb.SaveMemoRequest) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.SaveMemo(ctx, &Memo{
		Id:        int(req.Id),
		AccountId: int(u.Id),
		Name:      req.Name,
		Content:   req.Content,
		Status:    int8(req.Status),
	})
}
func (s *PalacesUseCase) DoneMemo(ctx context.Context, req *global.State) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DoneMemo(ctx, req)
}
func (s *PalacesUseCase) DeleteMemo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DeleteMemo(ctx, req)
}
func (s *PalacesUseCase) ListMemo(ctx context.Context, req *pb.ListMemoRequest) (*pb.ListMemoReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	total, list, err := s.repo.ListMemo(ctx, &MemoQuery{
		PageInfo:  req.Page,
		Name:      req.Name,
		Status:    req.Status,
		AccountId: int(u.Id),
	})
	if err != nil {
		return nil, err
	}
	resp := &pb.ListMemoReply{
		Total: total,
	}
	for _, m := range list {
		resp.List = append(resp.List, &pb.SaveMemoRequest{
			Id:        int64(m.Id),
			Name:      m.Name,
			Content:   m.Content,
			Status:    int64(m.Status),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	return resp, nil
}
func (s *PalacesUseCase) SaveTodo(ctx context.Context, req *pb.SaveTodoItem) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	params := &Todo{
		Id:        int(req.Id),
		Theme:     req.Theme,
		Type:      req.Type,
		From:      req.From,
		To:        req.To,
		Num:       req.Num,
		Sort:      req.Sort,
		Content:   req.Content,
		Status:    req.Status,
		AccountId: int(u.Id),
	}
	if err = s.repo.ValidTodo(ctx, params); err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.SaveTodo(ctx, params)
}
func (s *PalacesUseCase) DoneTodo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DoneTodo(ctx, req)
}
func (s *PalacesUseCase) DeleteTodo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DeleteTodo(ctx, req)
}
func (s *PalacesUseCase) DeleteTodoDone(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DeleteTodoDone(ctx, req)
}
func (s *PalacesUseCase) DeleteTodoRecord(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DeleteTodoRecord(ctx, req)
}
func (s *PalacesUseCase) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	total, list, err := s.repo.ListTodo(ctx, &TodoQuery{
		PageInfo:  req.Page,
		AccountId: int(u.Id),
		Theme:     req.Theme,
		Type:      req.Type,
		From:      req.From,
		Status:    req.Status,
	})
	if err != nil {
		return nil, err
	}
	resp := &pb.ListTodoReply{
		Total: total,
	}
	for _, m := range list {
		resp.List = append(resp.List, &pb.SaveTodoItem{
			Id:        int64(m.Id),
			Theme:     m.Theme,
			Type:      m.Type,
			From:      m.From,
			To:        m.To,
			Num:       m.Num,
			Sort:      m.Sort,
			Content:   m.Content,
			Status:    m.Status,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			//DoneId:    &m.DoneId,
		})
	}

	return resp, nil
}
func (s *PalacesUseCase) ListTodoDone(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	total, list, err := s.repo.ListTodoDone(ctx, &TodoQuery{
		PageInfo:  req.Page,
		AccountId: int(u.Id),
		Theme:     req.Theme,
		Type:      req.Type,
		From:      req.From,
		Status:    req.Status,
	})
	resp := &pb.ListTodoReply{
		Total: total,
	}
	for _, m := range list {
		resp.List = append(resp.List, &pb.SaveTodoItem{
			Id:        int64(m.Id),
			Theme:     m.Theme,
			Type:      m.Type,
			From:      m.From,
			To:        m.To,
			Num:       m.Num,
			Sort:      m.Sort,
			Content:   m.Content,
			Status:    m.Status,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			DoneId:    &m.DoneId,
		})
	}

	return resp, nil
}
