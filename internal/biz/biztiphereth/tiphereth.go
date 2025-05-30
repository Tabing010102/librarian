package biztiphereth

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/service/supervisor"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTiphereth,
	NewUserCountCache,
	NewPorterInstanceCache,
	NewPorterContextCache,
)

type Tiphereth struct {
	app                 *libapp.Settings
	auth                *libauth.Auth
	repo                *data.TipherethRepo
	supv                *supervisor.Supervisor
	id                  *libidgenerator.IDGenerator
	search              libsearch.Search
	pullAccount         *libmq.Topic[model.PullAccountInfo]
	userCountCache      *libcache.Key[model.UserCount]
	porterInstanceCache *libcache.Map[string, modelsupervisor.PorterInstance]
}

func NewTiphereth(
	app *libapp.Settings,
	repo *data.TipherethRepo,
	auth *libauth.Auth,
	supv *supervisor.Supervisor,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pullAccount *libmq.Topic[model.PullAccountInfo],
	cron *libcron.Cron,
	userCountCache *libcache.Key[model.UserCount],
	porterInstanceCache *libcache.Map[string, modelsupervisor.PorterInstance],
) (*Tiphereth, error) {
	t := &Tiphereth{
		app:                 app,
		auth:                auth,
		repo:                repo,
		supv:                supv,
		id:                  id,
		search:              search,
		pullAccount:         pullAccount,
		userCountCache:      userCountCache,
		porterInstanceCache: porterInstanceCache,
	}
	err := cron.Duration(
		"TipherethUpdatePorter",
		supv.GetHeartbeatInterval(),
		t.updatePorters, context.Background(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register cron: %w", err)
	}
	return t, nil
}

const (
	demoAdminUserName = "admin"
	demoAdminPassword = "admin"
)

func (t *Tiphereth) CreateConfiguredAdmin() {
	ctx := context.Background()
	if !t.app.EnvExist(libapp.EnvDemoMode) && !t.app.EnvExist(libapp.EnvCreateAdminUserName) {
		return
	}
	user := &model.User{
		ID:       0,
		Username: demoAdminUserName,
		Password: demoAdminPassword,
		Type:     model.UserTypeAdmin,
		Status:   model.UserStatusActive,
	}
	if username, err := t.app.Env(libapp.EnvCreateAdminUserName); err == nil && username != "" {
		user.Username = username
	}
	if password, err := t.app.Env(libapp.EnvCreateAdminPassword); err == nil && password != "" {
		user.Password = password
	}
	password, err := t.auth.GeneratePassword(user.Password)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return
	}
	user.Password = password
	id, err := t.id.New()
	if err != nil {
		return
	}
	user.ID = id
	// if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
	//	{
	//		Vid:  int64(user.ID),
	//		Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
	//		Prop: nil,
	//	},
	// }}); err != nil {
	//	return
	//}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return
	}
}

func NewUserCountCache(
	t *data.TipherethRepo,
	store libcache.Store,
) *libcache.Key[model.UserCount] {
	return libcache.NewKey[model.UserCount](
		store,
		"UserCount",
		func(ctx context.Context) (*model.UserCount, error) {
			res, err := t.GetUserCount(ctx)
			if err != nil {
				return nil, err
			}
			return &model.UserCount{Count: res}, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

func NewPorterInstanceCache(
	t *data.TipherethRepo,
	store libcache.Store,
) *libcache.Map[string, modelsupervisor.PorterInstance] {
	return libcache.NewMap[string, modelsupervisor.PorterInstance](
		store,
		"PorterInstanceCache",
		func(s string) string {
			return s
		},
		func(ctx context.Context, s string) (*modelsupervisor.PorterInstance, error) {
			return t.FetchPorterByAddress(ctx, s)
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

func NewPorterContextCache(
	t *data.TipherethRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelsupervisor.PorterContext] {
	return libcache.NewMap[model.InternalID, modelsupervisor.PorterContext](
		store,
		"PorterContextCache",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, k model.InternalID) (*modelsupervisor.PorterContext, error) {
			return t.FetchPorterContext(ctx, k)
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}
