package data

import (
	"blog/internal/ent"
	"blog/internal/ent/userexperience"
	"blog/internal/ent/userproject"
	"context"
	sql2 "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"errors"

	"blog/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *userRepo) SaveUser(ctx context.Context, data *biz.User) (err error) {
	err = o.data.db.User.Create().SetID(data.ID).
		SetName(data.Name).SetAvatar(data.Avatar).
		SetEmail(data.Email).SetProfessional(data.Professional).
		SetAddress(data.Address).SetSkills(data.Skills).SetDescription(data.Description).
		OnConflict().UpdateNewValues().Exec(ctx)
	return
}
func (o *userRepo) GetUser(ctx context.Context, data *biz.User) (user *biz.User, err error) {
	info, err := o.data.db.User.Get(ctx, data.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			err = nil
			user = &biz.User{ID: data.ID}
			return
		}
		return
	}
	user = &biz.User{}
	user.Name = info.Name
	user.Avatar = info.Avatar
	user.Email = info.Email
	user.Professional = info.Professional
	user.Address = info.Address
	user.Skills = info.Skills
	user.Description = info.Description
	user.Experience = info.Experience
	user.Project = info.Project
	user.CreatedAt = info.CreatedAt
	user.UpdatedAt = info.UpdatedAt
	return
}
func (o *userRepo) SaveProject(ctx context.Context, data *biz.Project) (err error) {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	if data.ID > 0 {
		_, err = tx.UserProject.UpdateOneID(data.ID).Where(userproject.UserIDEQ(data.UserId)).
			SetUserID(data.UserId).
			SetExperienceID(data.ExperienceId).
			SetTitle(data.Title).
			SetDescription(data.Description).
			SetSkills(data.Skills).
			SetEnd(data.End).
			SetStart(data.Start).
			SetLink(data.Link).
			SetPhotos(data.Photos).
			Save(ctx)
		return
	}
	_, err = tx.UserProject.Create().
		SetUserID(data.UserId).
		SetExperienceID(data.ExperienceId).
		SetTitle(data.Title).
		SetDescription(data.Description).
		SetSkills(data.Skills).
		SetEnd(data.End).
		SetStart(data.Start).
		SetLink(data.Link).
		SetPhotos(data.Photos).
		Save(ctx)
	if err != nil {
		return
	}
	_, err = tx.User.UpdateOneID(data.UserId).AddProject(1).Save(ctx)
	return
}
func (o *userRepo) GetProject(ctx context.Context, data *biz.Project) (project *biz.Project, err error) {
	info, err := o.data.db.UserProject.Get(ctx, data.ID)
	if err != nil {
		return
	}
	project = &biz.Project{
		ID:           info.ID,
		UserId:       info.UserID,
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
		ExperienceId: info.ExperienceID,
		Title:        info.Title,
		Description:  info.Description,
		Skills:       info.Skills,
		Start:        info.Start,
		End:          info.End,
		Link:         info.Link,
		Photos:       info.Photos,
	}
	return
}
func (o *userRepo) DeleteProject(ctx context.Context, data *biz.Project) (err error) {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	row, err := tx.UserProject.Delete().Where(userproject.IDEQ(int(data.ID))).Where(userproject.UserIDEQ(int(data.UserId))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("project not found")
	}
	_, err = tx.User.UpdateOneID(int(data.UserId)).AddProject(-1).Save(ctx)
	return
}
func (o *userRepo) ListProject(ctx context.Context, data *biz.ProjectQuery) (total int, list []*biz.Project, err error) {
	tx := o.data.db.UserProject.Query().Where(userproject.UserIDEQ(int(data.UserId)))
	if len(data.Title) > 0 {
		tx = tx.Where(userproject.TitleContains(data.Title))
	}
	if len(data.Time) == 2 {
		tx = tx.Where(userproject.StartLTE(data.Time[0]), userproject.EndGTE(data.Time[1]))
	}
	if len(data.Skills) > 0 {
		tx = tx.Where(func(selector *sql2.Selector) {
			selector.Where(sqljson.ValueContains(userproject.FieldSkills, data.Skills))
		})
	}
	total, err = tx.Count(ctx)
	if err != nil {
		return
	}
	all, err := tx.Limit(int(data.GetPageSize())).Offset(int(data.GetOffset())).
		Order(ent.Desc(userproject.FieldID)).All(ctx)
	if err != nil {
		return
	}
	list = make([]*biz.Project, 0, len(all))
	for _, info := range all {
		project := &biz.Project{
			ID:           info.ID,
			UserId:       info.UserID,
			CreatedAt:    info.CreatedAt,
			UpdatedAt:    info.UpdatedAt,
			ExperienceId: info.ExperienceID,
			Title:        info.Title,
			Description:  info.Description,
			Skills:       info.Skills,
			Start:        info.Start,
			End:          info.End,
			Link:         info.Link,
			Photos:       info.Photos,
		}
		list = append(list, project)
	}
	return
}
func (o *userRepo) SaveExperience(ctx context.Context, data *biz.Experience) (err error) {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	if data.ID > 0 {
		_, err = tx.UserExperience.UpdateOneID(data.ID).Where(userexperience.UserIDEQ(data.UserId)).
			SetUserID(int(data.UserId)).
			SetCompany(data.Company).
			SetRole(data.Role).
			SetLocation(data.Location).
			SetStart(data.Start).
			SetEnd(data.End).
			SetDescription(data.Description).
			SetResponsibilities(data.Responsibilities).
			SetAchievements(data.Achievements).
			SetSkills(data.Skills).
			Save(ctx)
		return
	}
	_, err = tx.UserExperience.Create().
		SetUserID(int(data.UserId)).
		SetCompany(data.Company).
		SetRole(data.Role).
		SetLocation(data.Location).
		SetStart(data.Start).
		SetEnd(data.End).
		SetDescription(data.Description).
		SetResponsibilities(data.Responsibilities).
		SetAchievements(data.Achievements).
		SetSkills(data.Skills).
		Save(ctx)
	_, err = tx.User.UpdateOneID(int(data.UserId)).AddExperience(1).Save(ctx)
	return
}
func (o *userRepo) GetExperience(ctx context.Context, data *biz.Experience) (project *biz.Experience, err error) {
	info, err := o.data.db.UserExperience.Get(ctx, data.ID)
	if err != nil {
		return
	}
	project = &biz.Experience{
		ID:               info.ID,
		UserId:           info.UserID,
		CreatedAt:        info.CreatedAt,
		UpdatedAt:        info.UpdatedAt,
		Company:          info.Company,
		Role:             info.Role,
		Location:         info.Location,
		Start:            info.Start,
		End:              info.End,
		Description:      info.Description,
		Responsibilities: info.Responsibilities,
		Achievements:     info.Achievements,
		Skills:           info.Skills,
	}
	return
}
func (o *userRepo) DeleteExperience(ctx context.Context, data *biz.Experience) (err error) {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	row, err := tx.UserExperience.Delete().Where(userexperience.IDEQ(int(data.ID))).Where(userexperience.UserIDEQ(int(data.UserId))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("experience not found")
	}
	_, err = tx.User.UpdateOneID(int(data.UserId)).AddExperience(-1).Save(ctx)
	return
}
func (o *userRepo) ListExperience(ctx context.Context, data *biz.ExperienceQuery) (total int, list []*biz.Experience, err error) {
	tx := o.data.db.UserExperience.Query().Where(userexperience.UserIDEQ(int(data.UserId)))
	if len(data.Company) > 0 {
		tx = tx.Where(userexperience.CompanyContains(data.Company))
	}
	total, err = tx.Count(ctx)
	if err != nil {
		return
	}
	all, err := tx.Limit(int(data.GetPageSize())).Offset(int(data.GetOffset())).
		Order(ent.Desc(userexperience.FieldID)).All(ctx)
	if err != nil {
		return
	}
	list = make([]*biz.Experience, 0, len(all))
	for _, info := range all {
		project := &biz.Experience{
			ID:               info.ID,
			UserId:           info.UserID,
			CreatedAt:        info.CreatedAt,
			UpdatedAt:        info.UpdatedAt,
			Company:          info.Company,
			Role:             info.Role,
			Location:         info.Location,
			Start:            info.Start,
			End:              info.End,
			Description:      info.Description,
			Responsibilities: info.Responsibilities,
			Achievements:     info.Achievements,
			Skills:           info.Skills,
		}
		list = append(list, project)
	}
	return
}
