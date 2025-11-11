package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------------------
// 1. Key builder – same pattern as real projects
// ---------------------------------------------------------------------
// Example: users:page:2:limit:20:search:john:status:active:sort:name_asc
func BuildUserListKey(page, limit int, search, status, sort string) string {
	parts := []string{
		"users",
		"page:" + strconv.Itoa(page),
		"limit:" + strconv.Itoa(limit),
	}
	if search != "" {
		parts = append(parts, "search:"+search)
	}
	if status != "" {
		parts = append(parts, "status:"+status)
	}
	if sort != "" {
		parts = append(parts, "sort:"+sort)
	}
	return strings.Join(parts, ":")
}

// ---------------------------------------------------------------------
// 2. Domain model
// ---------------------------------------------------------------------
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"` // ISO string
}

type UserPage struct {
	Users      []User `json:"users"`
	Total      int    `json:"total"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	HasNext    bool   `json:"has_next"`
	CacheUntil int64  `json:"-"` // unix timestamp for internal revalidation
}

// ---------------------------------------------------------------------
// 3. Cache operations
// ---------------------------------------------------------------------
const (
	defaultTTL      = 5 * time.Minute
	revalidateGrace = 30 * time.Second // allow stale data while background refresh
)

func (rc *RedisClient) SetUserPage(ctx context.Context, key string, page *UserPage, ttl time.Duration) error {
	data, err := json.Marshal(page)
	if err != nil {
		return err
	}
	// Use NX + EX to avoid overwriting newer data
	// Also store a "revalidate-until" score in a sorted set for background jobs
	err = rc.SetNX(ctx, key, data, ttl).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	// Keep a tiny sorted set for revalidation scheduling
	score := float64(time.Now().Add(ttl - revalidateGrace).Unix())
	_ = rc.ZAdd(ctx, "revalidate:users", redis.Z{Score: score, Member: key}).Err()
	return nil
}

func (rc *RedisClient) GetUserPage(ctx context.Context, key string) (*UserPage, bool, error) {
	data, err := rc.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, false, nil // cache miss
	}
	if err != nil {
		return nil, false, err
	}
	var page UserPage
	if err := json.Unmarshal(data, &err); err != nil {
		return nil, false, err
	}
	return &page, true, nil
}

// ---------------------------------------------------------------------
// 4. Revalidation helper (called from aync worker or HTTP handler)
// ---------------------------------------------------------------------
func (rc *RedisClient) RevalidateUserPage(ctx context.Context, key string, fetchFn func() (*UserPage, error)) error {
	// 1. Try to get fresh data
	fresh, err := fetchFn()
	if err != nil {
		return err
	}
	// 2. Overwrite cache (even if it was stale)
	return rc.SetUserPage(ctx, key, fresh, defaultTTL)
}

// ---------------------------------------------------------------------
// 5. Public API used by handlers
// ---------------------------------------------------------------------
func (rc *RedisClient) GetOrFetchUserPage(
	ctx context.Context,
	page, limit int,
	search, status, sort string,
	fetchFn func() (*UserPage, error),
) (*UserPage, error) {

	key := BuildUserListKey(page, limit, search, status, sort)

	// ---- 1. Cache hit? ----
	cached, hit, err := rc.GetUserPage(ctx, key)
	if err != nil {
		return nil, err
	}
	if hit {
		// Optional: stale-while-revalidate
		go rc.RevalidateUserPage(context.Background(), key, fetchFn)
		return cached, nil
	}

	// ---- 2. Cache miss -> fetch + store ----
	fresh, err := fetchFn()
	if err != nil {
		return nil, err
	}
	_ = rc.SetUserPage(ctx, key, fresh, defaultTTL)
	return fresh, nil
}