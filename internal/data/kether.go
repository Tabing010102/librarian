package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedactionset"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"entgo.io/ent/dialect/sql"
)

type KetherRepo struct {
	data *Data
}

// NewKetherRepo .
func NewKetherRepo(data *Data) *KetherRepo {
	return &KetherRepo{
		data: data,
	}
}

func (k *KetherRepo) UpsertAccount(ctx context.Context, acc model.Account) error {
	return k.data.db.Account.Create().
		SetID(acc.ID).
		SetPlatform(acc.Platform).
		SetPlatformAccountID(acc.PlatformAccountID).
		SetName(acc.Name).
		SetProfileURL(acc.ProfileURL).
		SetAvatarURL(acc.AvatarURL).
		OnConflict(
			sql.ConflictColumns(account.FieldPlatform, account.FieldPlatformAccountID),
			resolveWithIgnores([]string{
				account.FieldID,
				account.FieldPlatform,
				account.FieldPlatformAccountID,
			}),
		).
		Exec(ctx)
}

func (k *KetherRepo) UpsertAppInfo( //nolint:gocognit //TODO
	ctx context.Context, ap *modelgebura.AppInfo, internal *modelgebura.AppInfo,
) error {
	return k.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.AppInfo.Create().
			SetID(ap.ID).
			SetSource(ap.Source).
			SetSourceAppID(ap.SourceAppID)
		if len(ap.SourceURL) > 0 {
			q.SetSourceURL(ap.SourceURL)
		}
		if len(ap.Name) > 0 {
			q.SetName(ap.Name)
		}
		if ap.Type != modelgebura.AppTypeUnspecified {
			q.SetType(converter.ToEntAppInfoType(ap.Type))
		}
		if len(ap.ShortDescription) > 0 {
			q.SetShortDescription(ap.ShortDescription)
		}
		if len(ap.IconImageURL) > 0 {
			q.SetIconImageURL(ap.IconImageURL)
		}
		if len(ap.BackgroundImageURL) > 0 {
			q.SetBackgroundImageURL(ap.BackgroundImageURL)
		}
		if len(ap.CoverImageURL) > 0 {
			q.SetCoverImageURL(ap.CoverImageURL)
		}
		if len(ap.Description) > 0 {
			q.SetDescription(ap.Description)
		}
		if len(ap.ReleaseDate) > 0 {
			q.SetReleaseDate(ap.ReleaseDate)
		}
		if len(ap.Developer) > 0 {
			q.SetDeveloper(ap.Developer)
		}
		if len(ap.Publisher) > 0 {
			q.SetPublisher(ap.Publisher)
		}
		q.OnConflict(
			sql.ConflictColumns(appinfo.FieldSource, appinfo.FieldSourceAppID),
			resolveWithIgnores([]string{
				appinfo.FieldID,
				appinfo.FieldSource,
				appinfo.FieldSourceAppID,
			}),
		)
		count, err := tx.AppInfo.Query().Where(
			appinfo.SourceEQ(ap.Source),
			appinfo.SourceAppIDEQ(ap.SourceAppID),
		).Count(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			err = tx.AppInfo.Create().
				SetID(internal.ID).
				SetSource(internal.Source).
				SetSourceAppID(internal.SourceAppID).
				SetName(internal.Name).
				SetType(converter.ToEntAppInfoType(internal.Type)).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return q.Exec(ctx)
	})
}

func (k *KetherRepo) UpsertAppInfos(ctx context.Context, al []*modelgebura.AppInfo) error {
	apps := make([]*ent.AppInfoCreate, len(al))
	for i, ap := range al {
		apps[i] = k.data.db.AppInfo.Create().
			SetID(ap.ID).
			SetSource(ap.Source).
			SetSourceAppID(ap.SourceAppID).
			SetSourceURL(ap.SourceURL).
			SetName(ap.Name).
			SetType(converter.ToEntAppInfoType(ap.Type)).
			SetShortDescription(ap.ShortDescription).
			SetIconImageURL(ap.IconImageURL).
			SetBackgroundImageURL(ap.BackgroundImageURL).
			SetCoverImageURL(ap.CoverImageURL).
			SetDescription(ap.Description).
			SetReleaseDate(ap.ReleaseDate).
			SetDeveloper(ap.Developer).
			SetPublisher(ap.Publisher)
	}
	return k.data.db.AppInfo.
		CreateBulk(apps...).
		OnConflict(
			sql.ConflictColumns(appinfo.FieldSource, appinfo.FieldSourceAppID),
			resolveWithIgnores([]string{
				appinfo.FieldID,
			}),
		).
		Exec(ctx)
}

// func (k *KetherRepo) AccountPurchaseAppInfos(
//	ctx context.Context, id model.InternalID, ids []model.InternalID,
// ) error {
//	return k.data.WithTx(ctx, func(tx *ent.Tx) error {
//		appIDs, err := tx.App.Query().Where(
//			app.IDIn(ids...),
//		).
//			IDs(ctx)
//		if err != nil {
//			return err
//		}
//		return k.data.db.Account.
//			UpdateOneID(id).
//			AddPurchasedAppIDs(appIDs...).
//			Exec(ctx)
//	})
//}

func (k *KetherRepo) UpsertFeed(ctx context.Context, f *modelfeed.Feed) error {
	return k.data.WithTx(ctx, func(tx *ent.Tx) error {
		conf, err := tx.FeedConfig.Query().
			Where(feedconfig.IDEQ(f.ID)).
			Only(ctx)
		if err != nil {
			return err
		}
		err = tx.Feed.Create().
			SetConfig(conf).
			SetID(f.ID).
			SetTitle(f.Title).
			SetDescription(f.Description).
			SetLink(f.Link).
			SetAuthors(f.Authors).
			SetLanguage(f.Language).
			SetImage(f.Image).
			OnConflict(
				sql.ConflictColumns(feed.FieldID),
				sql.ResolveWithNewValues(),
			).
			Exec(ctx)
		return err
	})
}

func (k *KetherRepo) CheckNewFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID model.InternalID,
) ([]string, error) {
	guids := make([]string, 0, len(items))
	for _, item := range items {
		guids = append(guids, item.GUID)
	}
	existItems, err := k.data.db.FeedItem.Query().Where(
		feeditem.FeedID(feedID),
		feeditem.GUIDIn(guids...),
	).Select(feeditem.FieldGUID).All(ctx)
	if err != nil {
		return nil, err
	}
	existItemMap := make(map[string]bool)
	res := make([]string, 0, len(items)-len(existItems))
	for _, item := range existItems {
		existItemMap[item.GUID] = true
	}
	for _, item := range items {
		if _, exist := existItemMap[item.GUID]; !exist {
			res = append(res, item.GUID)
		}
	}
	return res, nil
}

func (k *KetherRepo) UpsertFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID model.InternalID,
) error {
	il := make([]*ent.FeedItemCreate, len(items))
	for i, item := range items {
		il[i] = k.data.db.FeedItem.Create().
			SetFeedID(feedID).
			SetID(item.ID).
			SetTitle(item.Title).
			SetDescription(item.Description).
			SetContent(item.Content).
			SetLink(item.Link).
			SetUpdated(item.Updated).
			SetNillableUpdatedParsed(item.UpdatedParsed).
			SetPublished(item.Published).
			SetAuthors(item.Authors).
			SetGUID(item.GUID).
			SetImage(item.Image).
			SetEnclosures(item.Enclosures).
			SetPublishPlatform(item.PublishPlatform).
			SetDigestDescription(item.DigestDescription).
			SetDigestImages(item.DigestImages)
		if item.PublishedParsed != nil {
			il[i].SetPublishedParsed(*item.PublishedParsed)
		} else {
			il[i].SetPublishedParsed(time.Now())
		}
	}
	return k.data.db.FeedItem.CreateBulk(il...).
		OnConflict(
			sql.ConflictColumns(feeditem.FieldFeedID, feeditem.FieldGUID),
			//
			// Update feed item every time result in large disk writes
			//
			// resolveWithIgnores([]string{
			//	feeditem.FieldID,
			// }),
			sql.DoNothing(),
		).Exec(ctx)
}

func (k *KetherRepo) UpdateFeedPullStatus(ctx context.Context, conf *modelyesod.FeedConfig) error {
	c, err := k.data.db.FeedConfig.Query().Where(feedconfig.IDEQ(conf.ID)).Only(ctx)
	if err != nil {
		return err
	}
	return c.Update().
		SetLatestPullAt(conf.LatestPullTime).
		SetLatestPullStatus(converter.ToEntFeedConfigLatestPullStatus(conf.LatestPullStatus)).
		SetLatestPullMessage(conf.LatestPullMessage).
		SetNextPullBeginAt(conf.LatestPullTime.Add(c.PullInterval)).
		Exec(ctx)
}

func (k *KetherRepo) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, error) {
	item, err := k.data.db.FeedItem.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedItem(item), nil
}

func (k *KetherRepo) GetFeedActions(ctx context.Context, id model.InternalID) ([]*modelyesod.FeedActionSet, error) {
	actions, err := k.data.db.FeedActionSet.Query().
		Where(feedactionset.HasFeedConfigWith(feedconfig.IDEQ(id))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedActionSetList(actions), nil
}

func (k *KetherRepo) GetNotifyTargetItems(
	ctx context.Context,
	id model.InternalID,
	paging model.Paging,
) (*modelsupervisor.FeatureRequest, []*modelfeed.Item, error) {
	var fr *modelsupervisor.FeatureRequest
	var it []*modelfeed.Item
	err := k.data.WithTx(ctx, func(tx *ent.Tx) error {
		target, err := tx.NotifyTarget.Get(ctx, id)
		if err != nil {
			return err
		}
		fr = target.Destination
		ids, err := target.QueryNotifyFlow().IDs(ctx)
		if err != nil {
			return err
		}
		items, err := tx.FeedItem.Query().Where(
			feeditem.HasFeedItemCollectionWith(
				feeditemcollection.IDIn(ids...),
			),
		).Offset(paging.ToOffset()).Limit(paging.ToLimit()).All(ctx)
		if err != nil {
			return err
		}
		it = converter.ToBizFeedItemList(items)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return fr, it, nil
}

func (k *KetherRepo) AddFeedItemsToCollection(
	ctx context.Context,
	collectionID model.InternalID,
	itemIDs []model.InternalID,
) error {
	return k.data.db.FeedItemCollection.UpdateOneID(collectionID).
		AddFeedItemIDs(itemIDs...).
		Exec(ctx)
}
