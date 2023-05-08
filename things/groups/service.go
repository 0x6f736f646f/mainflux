package groups

import (
	"context"
	"time"

	"github.com/mainflux/mainflux"
	"github.com/mainflux/mainflux/internal/apiutil"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/mainflux/users/policies"
	upolicies "github.com/mainflux/mainflux/users/policies"
)

// Possible token types are access and refresh tokens.
const (
	thingsObjectKey   = "things"
	updateRelationKey = "g_update"
	listRelationKey   = "g_list"
	deleteRelationKey = "g_delete"
	entityType        = "group"
)

var (
	// ErrInvalidStatus indicates invalid status.
	ErrInvalidStatus = errors.New("invalid groups status")

	// ErrEnableGroup indicates error in enabling group.
	ErrEnableGroup = errors.New("failed to enable group")

	// ErrDisableGroup indicates error in disabling group.
	ErrDisableGroup = errors.New("failed to disable group")

	// ErrStatusAlreadyAssigned indicated that the group has already been assigned the status.
	ErrStatusAlreadyAssigned = errors.New("status already assigned")
)

type service struct {
	auth       upolicies.AuthServiceClient
	groups     Repository
	idProvider mainflux.IDProvider
}

// NewService returns a new Clients service implementation.
func NewService(auth upolicies.AuthServiceClient, g Repository, idp mainflux.IDProvider) Service {
	return service{
		auth:       auth,
		groups:     g,
		idProvider: idp,
	}
}

func (svc service) CreateGroups(ctx context.Context, token string, gs ...Group) ([]Group, error) {
	res, err := svc.auth.Identify(ctx, &upolicies.Token{Value: token})
	if err != nil {
		return []Group{}, errors.Wrap(errors.ErrAuthentication, err)
	}

	var grps []Group
	for _, g := range gs {
		if g.ID == "" {
			groupID, err := svc.idProvider.ID()
			if err != nil {
				return []Group{}, err
			}
			g.ID = groupID
		}
		if g.Owner == "" {
			g.Owner = res.GetId()
		}

		if g.Status != EnabledStatus && g.Status != DisabledStatus {
			return []Group{}, apiutil.ErrInvalidStatus
		}

		g.CreatedAt = time.Now()
		g.UpdatedAt = g.CreatedAt
		g.UpdatedBy = g.Owner
		grp, err := svc.groups.Save(ctx, g)
		if err != nil {
			return []Group{}, err
		}
		grps = append(grps, grp)
	}
	return grps, nil
}

func (svc service) ViewGroup(ctx context.Context, token string, id string) (Group, error) {
	if err := svc.authorize(ctx, token, id, listRelationKey); err != nil {
		return Group{}, errors.Wrap(errors.ErrNotFound, err)
	}
	return svc.groups.RetrieveByID(ctx, id)
}

func (svc service) ListGroups(ctx context.Context, token string, gm GroupsPage) (GroupsPage, error) {
	res, err := svc.auth.Identify(ctx, &upolicies.Token{Value: token})
	if err != nil {
		return GroupsPage{}, errors.Wrap(errors.ErrAuthentication, err)
	}

	// If the user is admin, fetch all channels from the database.
	if err := svc.authorize(ctx, token, thingsObjectKey, listRelationKey); err == nil {
		page, err := svc.groups.RetrieveAll(ctx, gm)
		if err != nil {
			return GroupsPage{}, err
		}
		return page, err
	}

	gm.Subject = res.GetId()
	gm.OwnerID = res.GetId()
	gm.Action = "g_list"
	return svc.groups.RetrieveAll(ctx, gm)
}

func (svc service) ListMemberships(ctx context.Context, token, clientID string, gm GroupsPage) (MembershipsPage, error) {
	res, err := svc.auth.Identify(ctx, &upolicies.Token{Value: token})
	if err != nil {
		return MembershipsPage{}, errors.Wrap(errors.ErrAuthentication, err)
	}

	// If the user is admin, fetch all channels from the database.
	if err := svc.authorize(ctx, token, thingsObjectKey, listRelationKey); err == nil {
		return svc.groups.Memberships(ctx, clientID, gm)
	}

	gm.OwnerID = res.GetId()
	return svc.groups.Memberships(ctx, clientID, gm)
}

func (svc service) UpdateGroup(ctx context.Context, token string, g Group) (Group, error) {
	res, err := svc.auth.Identify(ctx, &upolicies.Token{Value: token})
	if err != nil {
		return Group{}, errors.Wrap(errors.ErrAuthentication, err)
	}

	if err := svc.authorize(ctx, token, g.ID, updateRelationKey); err != nil {
		return Group{}, errors.Wrap(errors.ErrNotFound, err)
	}

	g.Owner = res.GetId()
	g.UpdatedAt = time.Now()
	g.UpdatedBy = res.GetId()

	return svc.groups.Update(ctx, g)
}

func (svc service) EnableGroup(ctx context.Context, token, id string) (Group, error) {
	group := Group{
		ID:        id,
		Status:    EnabledStatus,
		UpdatedAt: time.Now(),
	}
	group, err := svc.changeGroupStatus(ctx, token, group)
	if err != nil {
		return Group{}, errors.Wrap(ErrEnableGroup, err)
	}
	return group, nil
}

func (svc service) DisableGroup(ctx context.Context, token, id string) (Group, error) {
	group := Group{
		ID:        id,
		Status:    DisabledStatus,
		UpdatedAt: time.Now(),
	}
	group, err := svc.changeGroupStatus(ctx, token, group)
	if err != nil {
		return Group{}, errors.Wrap(ErrDisableGroup, err)
	}
	return group, nil
}

func (svc service) changeGroupStatus(ctx context.Context, token string, group Group) (Group, error) {
	res, err := svc.auth.Identify(ctx, &upolicies.Token{Value: token})
	if err != nil {
		return Group{}, errors.Wrap(errors.ErrAuthentication, err)
	}
	if err := svc.authorize(ctx, token, group.ID, deleteRelationKey); err != nil {
		return Group{}, errors.Wrap(errors.ErrNotFound, err)
	}
	dbGroup, err := svc.groups.RetrieveByID(ctx, group.ID)
	if err != nil {
		return Group{}, err
	}

	if dbGroup.Status == group.Status {
		return Group{}, ErrStatusAlreadyAssigned
	}
	group.UpdatedBy = res.GetId()
	return svc.groups.ChangeStatus(ctx, group)
}

func (svc service) identifyUser(ctx context.Context, token string) (string, error) {
	req := &policies.Token{Value: token}
	res, err := svc.auth.Identify(ctx, req)
	if err != nil {
		return "", errors.Wrap(errors.ErrAuthorization, err)
	}
	return res.GetId(), nil
}

func (svc service) authorize(ctx context.Context, subject, object string, relation string) error {
	// Check if the client is the owner of the group.
	userID, err := svc.identifyUser(ctx, subject)
	if err != nil {
		return err
	}
	dbGroup, err := svc.groups.RetrieveByID(ctx, object)
	if err != nil {
		return err
	}
	if dbGroup.Owner == userID {
		return nil
	}
	req := &upolicies.AuthorizeReq{
		Sub:        subject,
		Obj:        object,
		Act:        relation,
		EntityType: entityType,
	}
	res, err := svc.auth.Authorize(ctx, req)
	if err != nil {
		return errors.Wrap(errors.ErrAuthorization, err)
	}
	if !res.GetAuthorized() {
		return errors.ErrAuthorization
	}
	return nil
}