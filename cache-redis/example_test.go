package cache

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestFullCacheFlow(t *testing.T) {
	// 1. Connect (use Docker redis for local testing)
	// docker run -p 6379:6379 --name redis-test -d redis
	rdb, err := NewRedisClient("localhost:6379", "", 0, false)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// Clean up
	defer rdb.FlushDB(ctx)

	// -----------------------------------------------------------------
	// Simulated DB fetch (replace with real DB call in production)
	// -----------------------------------------------------------------
	fetchFromDB := func() (*UserPage, error) {
		// Simulate slow DB
		time.Sleep(200 * time.Millisecond)

		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Status: "active", CreatedAt: "2025-01-01T12:00:00Z"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Status: "active", CreatedAt: "2025-01-02T12:00:00Z"},
		}
		return &UserPage{
			Users:   users,
			Total:   42,
			Page:    1,
			Limit:   20,
			HasNext: true,
		}, nil
	}

	// -----------------------------------------------------------------
	// First request -> cache miss
	// -----------------------------------------------------------------
	page1, err := rdb.GetOrFetchUserPage(ctx, 1, 20, "john", "active", "name_asc", fetchFromDB)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("FIRST (miss) -> %d users, total=%d", len(page1.Users), page1.Total)

	// -----------------------------------------------------------------
	// Second request -> cache hit (instant)
	// -----------------------------------------------------------------
	start := time.Now()
	page2, err := rdb.GetOrFetchUserPage(ctx, 1, 20, "john", "active", "name_asc", fetchFromDB)
	if err != nil {
		t.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("SECOND (hit) -> %d users in %v", len(page2.Users), elapsed)

	if elapsed > 50*time.Millisecond {
		t.Error("Cache hit was too slow")
	}

	// -----------------------------------------------------------------
	// Force revalidation (simulate data change)
	// -----------------------------------------------------------------
	newFetch := func() (*UserPage, error) {
		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Status: "active", CreatedAt: "2025-01-01T12:00:00Z"},
			{ID: 3, Name: "Johnny New", Email: "new@example.com", Status: "active", CreatedAt: "2025-11-11T12:00:00Z"},
		}
		return &UserPage{
			Users:   users,
			Total:   43,
			Page:    1,
			Limit:   20,
			HasNext: true,
		}, nil
	}

	key := BuildUserListKey(1, 20, "john", "active", "name_asc")
	if err := rdb.RevalidateUserPage(ctx, key, newFetch); err != nil {
		t.Fatal(err)
	}

	// -----------------------------------------------------------------
	// Next request should return fresh data
	// -----------------------------------------------------------------
	page3, err := rdb.GetOrFetchUserPage(ctx, 1, 20, "john", "active", "name_asc", fetchFromDB)
	if err != nil {
		t.Fatal(err)
	}
	if page3.Total != 43 {
		t.Errorf("Expected total 43, got %d", page3.Total)
	}
	log.Println("Revalidation successful!")
}