package room

import (
	"errors"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	pb "github.com/synctv-org/synctv/proto"
	"github.com/zijiren233/gencontainer/dllist"
	"github.com/zijiren233/stream"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	room     *Room
	name     string
	password []byte
	version  uint64
	admin    bool
	lastAct  int64
}

var (
	ErrUsernameEmpty       = errors.New("user name is empty")
	ErrUsernameTooLong     = errors.New("user name is too long")
	ErrUserPasswordEmpty   = errors.New("user password is empty")
	ErrUserPasswordTooLong = errors.New("user password is too long")
)

type UserConf func(*User)

func WithUserVersion(version uint64) UserConf {
	return func(u *User) {
		u.version = version
	}
}

func WithUserAdmin(admin bool) UserConf {
	return func(u *User) {
		u.admin = admin
	}
}

func NewUser(id string, password string, room *Room, conf ...UserConf) (*User, error) {
	if id == "" {
		return nil, ErrUsernameEmpty
	} else if len(id) > 32 {
		return nil, ErrUsernameTooLong
	}
	if password == "" {
		return nil, ErrUserPasswordEmpty
	} else if len(password) > 32 {
		return nil, ErrUserPasswordTooLong
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(stream.StringToBytes(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{
		room:     room,
		name:     id,
		password: hashedPassword,
		lastAct:  time.Now().UnixMicro(),
	}
	for _, c := range conf {
		c(u)
	}
	if u.version == 0 {
		u.version = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
	}
	return u, nil
}

func (u *User) LastAct() int64 {
	return atomic.LoadInt64(&u.lastAct)
}

func (u *User) LastActTime() time.Time {
	return time.UnixMicro(u.LastAct())
}

func (u *User) UpdateLastAct() int64 {
	return atomic.SwapInt64(&u.lastAct, time.Now().UnixMicro())
}

func (u *User) Version() uint64 {
	return atomic.LoadUint64(&u.version)
}

func (u *User) CheckVersion(version uint64) bool {
	return u.Version() == version
}

func (u *User) SetVersion(version uint64) {
	atomic.StoreUint64(&u.version, version)
}

func (u *User) updateVersion() uint64 {
	return atomic.AddUint64(&u.version, 1)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.password, stream.StringToBytes(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(stream.StringToBytes(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = hashedPassword
	u.updateVersion()
	return nil
}

func (u *User) CloseHub() {
	c, loaded := u.room.hub.clients.LoadAndDelete(u.name)
	if loaded {
		c.Close()
	}
}

func (u *User) IsRoot() bool {
	return u.room.rootUser == u
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Room() *Room {
	return u.room
}

func (u *User) IsAdmin() bool {
	return u.admin
}

func (u *User) SetAdmin(admin bool) {
	u.admin = admin
}

func (u *User) NewMovie(url string, name string, type_ string, live bool, proxy bool, rtmpSource bool, headers map[string]string, conf ...MovieConf) (*Movie, error) {
	return u.NewMovieWithBaseMovie(BaseMovieInfo{
		Url:        url,
		Name:       name,
		Live:       live,
		Proxy:      proxy,
		RtmpSource: rtmpSource,
		Type:       type_,
		Headers:    headers,
	}, conf...)
}

func (u *User) NewMovieWithBaseMovie(baseMovie BaseMovieInfo, conf ...MovieConf) (*Movie, error) {
	return NewMovieWithBaseMovie(atomic.AddUint64(&u.room.mid, 1), baseMovie, append(conf, WithCreator(u))...)
}

func (u *User) Movie(id uint64) (*pb.MovieInfo, error) {
	u.room.movies.lock.RLock()
	defer u.room.movies.lock.RUnlock()

	m, err := u.room.movies.GetMovie(id)
	if err != nil {
		return nil, err
	}
	movie := &pb.MovieInfo{
		Id:         m.Id(),
		Url:        m.Url,
		Name:       m.Name,
		Live:       m.Live,
		Proxy:      m.Proxy,
		RtmpSource: m.RtmpSource,
		Type:       m.Type,
		Headers:    m.Headers,
		PullKey:    m.PullKey,
		CreatedAt:  m.CreatedAt,
		Creator:    m.Creator().Name(),
	}
	if movie.Proxy && u.name != movie.Creator {
		m.Headers = nil
	}
	return movie, nil
}

func (u *User) Movies() []*MovieInfo {
	u.room.movies.lock.RLock()
	defer u.room.movies.lock.RUnlock()

	movies := make([]*MovieInfo, 0, u.room.movies.l.Len())
	u.room.movies.range_(func(e *dllist.Element[*Movie]) bool {
		m := &MovieInfo{
			Id: e.Value.Id(),
			BaseMovieInfo: BaseMovieInfo{
				Url:        e.Value.Url,
				Name:       e.Value.Name,
				Live:       e.Value.Live,
				Proxy:      e.Value.Proxy,
				RtmpSource: e.Value.RtmpSource,
				Type:       e.Value.Type,
				Headers:    e.Value.Headers,
			},
			PullKey:   e.Value.PullKey,
			CreatedAt: e.Value.CreatedAt,
			Creator:   e.Value.Creator().Name(),
		}
		if e.Value.Proxy && u.name != m.Creator {
			m.Headers = nil
		}
		movies = append(movies, m)
		return true
	})
	return movies
}

func (u *User) EditMovie(id uint64, movie BaseMovieInfo) error {
	m, err := u.room.movies.GetMovie(id)
	if err != nil {
		return err
	}
	if !u.IsAdmin() && !u.IsRoot() && m.Creator().name != u.name {
		return errors.New("you are not the creator of this movie, you can't edit it")
	}
	pre := m.BaseMovieInfo
	m.BaseMovieInfo = movie
	switch {
	case pre.RtmpSource && !movie.RtmpSource:
		u.room.rtmpa.DelChannel(m.PullKey)
		m.PullKey = ""
		// TODO: live proxy
	case pre.Proxy && !movie.Proxy:
		m.PullKey = ""
	}
	return nil
}

func (u *User) RegClient(conn *websocket.Conn) (*Client, error) {
	return u.room.RegClient(u, conn)
}

func (u *User) Broadcast(msg Message, conf ...BroadcastConf) error {
	return u.room.Broadcast(msg, append(conf, WithSender(u.name))...)
}
