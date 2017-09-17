package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import umschlag "github.com/umschlag/umschlag-go/umschlag"

// ClientAPI is an autogenerated mock type for the ClientAPI type
type ClientAPI struct {
	mock.Mock
}

// AuthLogin provides a mock function with given fields: _a0, _a1
func (_m *ClientAPI) AuthLogin(_a0 string, _a1 string) (*umschlag.Token, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *umschlag.Token
	if rf, ok := ret.Get(0).(func(string, string) *umschlag.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthenticated provides a mock function with given fields:
func (_m *ClientAPI) IsAuthenticated() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OrgDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgGet provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgGet(_a0 string) (*umschlag.Org, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Org
	if rf, ok := ret.Get(0).(func(string) *umschlag.Org); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Org)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgList provides a mock function with given fields:
func (_m *ClientAPI) OrgList() ([]*umschlag.Org, error) {
	ret := _m.Called()

	var r0 []*umschlag.Org
	if rf, ok := ret.Get(0).(func() []*umschlag.Org); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.Org)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgPatch provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgPatch(_a0 *umschlag.Org) (*umschlag.Org, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Org
	if rf, ok := ret.Get(0).(func(*umschlag.Org) *umschlag.Org); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Org)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Org) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgPost provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgPost(_a0 *umschlag.Org) (*umschlag.Org, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Org
	if rf, ok := ret.Get(0).(func(*umschlag.Org) *umschlag.Org); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Org)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Org) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgTeamAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgTeamAppend(_a0 umschlag.OrgTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgTeamDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgTeamDelete(_a0 umschlag.OrgTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgTeamList provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgTeamList(_a0 umschlag.OrgTeamParams) ([]*umschlag.TeamOrg, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.TeamOrg
	if rf, ok := ret.Get(0).(func(umschlag.OrgTeamParams) []*umschlag.TeamOrg); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.TeamOrg)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.OrgTeamParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgTeamPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgTeamPerm(_a0 umschlag.OrgTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgUserAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgUserAppend(_a0 umschlag.OrgUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgUserDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgUserDelete(_a0 umschlag.OrgUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgUserList provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgUserList(_a0 umschlag.OrgUserParams) ([]*umschlag.UserOrg, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.UserOrg
	if rf, ok := ret.Get(0).(func(umschlag.OrgUserParams) []*umschlag.UserOrg); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.UserOrg)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.OrgUserParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgUserPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) OrgUserPerm(_a0 umschlag.OrgUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.OrgUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProfileGet provides a mock function with given fields:
func (_m *ClientAPI) ProfileGet() (*umschlag.Profile, error) {
	ret := _m.Called()

	var r0 *umschlag.Profile
	if rf, ok := ret.Get(0).(func() *umschlag.Profile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProfilePatch provides a mock function with given fields: _a0
func (_m *ClientAPI) ProfilePatch(_a0 *umschlag.Profile) (*umschlag.Profile, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Profile
	if rf, ok := ret.Get(0).(func(*umschlag.Profile) *umschlag.Profile); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Profile) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProfileToken provides a mock function with given fields:
func (_m *ClientAPI) ProfileToken() (*umschlag.Token, error) {
	ret := _m.Called()

	var r0 *umschlag.Token
	if rf, ok := ret.Get(0).(func() *umschlag.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) RegistryDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryGet provides a mock function with given fields: _a0
func (_m *ClientAPI) RegistryGet(_a0 string) (*umschlag.Registry, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Registry
	if rf, ok := ret.Get(0).(func(string) *umschlag.Registry); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryList provides a mock function with given fields:
func (_m *ClientAPI) RegistryList() ([]*umschlag.Registry, error) {
	ret := _m.Called()

	var r0 []*umschlag.Registry
	if rf, ok := ret.Get(0).(func() []*umschlag.Registry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryPatch provides a mock function with given fields: _a0
func (_m *ClientAPI) RegistryPatch(_a0 *umschlag.Registry) (*umschlag.Registry, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Registry
	if rf, ok := ret.Get(0).(func(*umschlag.Registry) *umschlag.Registry); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryPost provides a mock function with given fields: _a0
func (_m *ClientAPI) RegistryPost(_a0 *umschlag.Registry) (*umschlag.Registry, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Registry
	if rf, ok := ret.Get(0).(func(*umschlag.Registry) *umschlag.Registry); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistrySync provides a mock function with given fields: _a0
func (_m *ClientAPI) RegistrySync(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepoDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) RepoDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepoGet provides a mock function with given fields: _a0
func (_m *ClientAPI) RepoGet(_a0 string) (*umschlag.Repo, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Repo
	if rf, ok := ret.Get(0).(func(string) *umschlag.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepoList provides a mock function with given fields:
func (_m *ClientAPI) RepoList() ([]*umschlag.Repo, error) {
	ret := _m.Called()

	var r0 []*umschlag.Repo
	if rf, ok := ret.Get(0).(func() []*umschlag.Repo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetClient provides a mock function with given fields: client
func (_m *ClientAPI) SetClient(client *http.Client) {
	_m.Called(client)
}

// TagDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) TagDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagGet provides a mock function with given fields: _a0
func (_m *ClientAPI) TagGet(_a0 string) (*umschlag.Tag, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Tag
	if rf, ok := ret.Get(0).(func(string) *umschlag.Tag); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagList provides a mock function with given fields:
func (_m *ClientAPI) TagList() ([]*umschlag.Tag, error) {
	ret := _m.Called()

	var r0 []*umschlag.Tag
	if rf, ok := ret.Get(0).(func() []*umschlag.Tag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamGet provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamGet(_a0 string) (*umschlag.Team, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Team
	if rf, ok := ret.Get(0).(func(string) *umschlag.Team); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamList provides a mock function with given fields:
func (_m *ClientAPI) TeamList() ([]*umschlag.Team, error) {
	ret := _m.Called()

	var r0 []*umschlag.Team
	if rf, ok := ret.Get(0).(func() []*umschlag.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamOrgAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamOrgAppend(_a0 umschlag.TeamOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamOrgDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamOrgDelete(_a0 umschlag.TeamOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamOrgList provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamOrgList(_a0 umschlag.TeamOrgParams) ([]*umschlag.TeamOrg, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.TeamOrg
	if rf, ok := ret.Get(0).(func(umschlag.TeamOrgParams) []*umschlag.TeamOrg); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.TeamOrg)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.TeamOrgParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamOrgPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamOrgPerm(_a0 umschlag.TeamOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamPatch provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamPatch(_a0 *umschlag.Team) (*umschlag.Team, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Team
	if rf, ok := ret.Get(0).(func(*umschlag.Team) *umschlag.Team); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Team) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamPost provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamPost(_a0 *umschlag.Team) (*umschlag.Team, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.Team
	if rf, ok := ret.Get(0).(func(*umschlag.Team) *umschlag.Team); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.Team) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamUserAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamUserAppend(_a0 umschlag.TeamUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamUserDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamUserDelete(_a0 umschlag.TeamUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamUserList provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamUserList(_a0 umschlag.TeamUserParams) ([]*umschlag.TeamUser, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.TeamUser
	if rf, ok := ret.Get(0).(func(umschlag.TeamUserParams) []*umschlag.TeamUser); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.TeamUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.TeamUserParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamUserPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) TeamUserPerm(_a0 umschlag.TeamUserParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.TeamUserParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) UserDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGet provides a mock function with given fields: _a0
func (_m *ClientAPI) UserGet(_a0 string) (*umschlag.User, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.User
	if rf, ok := ret.Get(0).(func(string) *umschlag.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserList provides a mock function with given fields:
func (_m *ClientAPI) UserList() ([]*umschlag.User, error) {
	ret := _m.Called()

	var r0 []*umschlag.User
	if rf, ok := ret.Get(0).(func() []*umschlag.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserOrgAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) UserOrgAppend(_a0 umschlag.UserOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserOrgDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) UserOrgDelete(_a0 umschlag.UserOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserOrgList provides a mock function with given fields: _a0
func (_m *ClientAPI) UserOrgList(_a0 umschlag.UserOrgParams) ([]*umschlag.UserOrg, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.UserOrg
	if rf, ok := ret.Get(0).(func(umschlag.UserOrgParams) []*umschlag.UserOrg); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.UserOrg)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.UserOrgParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserOrgPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) UserOrgPerm(_a0 umschlag.UserOrgParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserOrgParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserPatch provides a mock function with given fields: _a0
func (_m *ClientAPI) UserPatch(_a0 *umschlag.User) (*umschlag.User, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.User
	if rf, ok := ret.Get(0).(func(*umschlag.User) *umschlag.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserPost provides a mock function with given fields: _a0
func (_m *ClientAPI) UserPost(_a0 *umschlag.User) (*umschlag.User, error) {
	ret := _m.Called(_a0)

	var r0 *umschlag.User
	if rf, ok := ret.Get(0).(func(*umschlag.User) *umschlag.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*umschlag.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*umschlag.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserTeamAppend provides a mock function with given fields: _a0
func (_m *ClientAPI) UserTeamAppend(_a0 umschlag.UserTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserTeamDelete provides a mock function with given fields: _a0
func (_m *ClientAPI) UserTeamDelete(_a0 umschlag.UserTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserTeamList provides a mock function with given fields: _a0
func (_m *ClientAPI) UserTeamList(_a0 umschlag.UserTeamParams) ([]*umschlag.TeamUser, error) {
	ret := _m.Called(_a0)

	var r0 []*umschlag.TeamUser
	if rf, ok := ret.Get(0).(func(umschlag.UserTeamParams) []*umschlag.TeamUser); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*umschlag.TeamUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(umschlag.UserTeamParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserTeamPerm provides a mock function with given fields: _a0
func (_m *ClientAPI) UserTeamPerm(_a0 umschlag.UserTeamParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(umschlag.UserTeamParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ umschlag.ClientAPI = (*ClientAPI)(nil)