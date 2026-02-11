package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"

	"gin-gorm-postgres-crud-itest/internal/app"
	"gin-gorm-postgres-crud-itest/internal/db"

	"github.com/stretchr/testify/require"
	//tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
)

// func startPostgres(t *testing.T) (context.Context, *tcpostgres.PostgresContainer, string) {
// 	t.Helper()

// 	ctx := context.Background()
// 	pg, err := tcpostgres.Run(ctx,
// 		"postgres:16-alpine",
// 		tcpostgres.WithDatabase("usersdb"),
// 		tcpostgres.WithUsername("postgres"),
// 		tcpostgres.WithPassword("postgres"),
// 		//tcpostgres.WithStartupTimeout(90*time.Second)
// 		///tcpostgres.
// 	)
// 	require.NoError(t, err)

// 	dsn, err := pg.ConnectionString(ctx, "sslmode=disable")
// 	require.NoError(t, err)
// 	return ctx, pg, dsn
// }

var containerMap map[string]string

func init() {
	containerMap = make(map[string]string) // just to store the ids in the map for understanding .. here in this context there is no need of map
	// there is only one container that is postgres
}
func startPostgressContainer() (string, func() error, error) {
	// docker run -d --name pg -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=usersdb postgres:16
	cmd := exec.Command("docker", "run", "-d", "--name", "pg", "-p", "5432:5432", "-e", "POSTGRES_USER=postgres", "-e", "POSTGRES_PASSWORD=postgres", "-e", "POSTGRES_DB=usersdb", "postgres:16")
	if bytes, err := cmd.CombinedOutput(); err != nil {
		return "", nil, err
	} else {
		if containerMap != nil {
			containerMap["postgres"] = string(bytes)
		}
		return string(bytes), func() error {
			cmd := exec.Command("docker", "rm", "-f", containerMap["postgres"])
			_, err := cmd.CombinedOutput()
			return err
		}, nil
	}
}

func waitForPostgres(t *testing.T, dsn string) {
	t.Helper()

	deadline := time.Now().Add(30 * time.Second)
	var lastErr error

	for time.Now().Before(deadline) {
		gdb, err := db.Connect(dsn)
		if err == nil {
			// Close underlying sql.DB to avoid leaks if your db.Connect exposes it.
			// If db.Connect returns *gorm.DB, you can do:
			sqlDB, _ := gdb.DB()
			_ = sqlDB.Close()
			return
		}
		lastErr = err
		time.Sleep(300 * time.Millisecond)
	}
	require.NoError(t, lastErr)
}

func TestUsersCRUD_Integration(t *testing.T) {
	// ctx, pg, dsn := startPostgres(t)
	// t.Cleanup(func() { _ = pg.Terminate(ctx) })

	//db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})

	_, rmContainer, err := startPostgressContainer()
	require.NoError(t, err)
	// defer func() {
	// 	rmContainer()
	// }()
	defer rmContainer()

	dsn := `host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`

	waitForPostgres(t, dsn)

	gdb, err := db.Connect(dsn)

	require.NoError(t, err)
	require.NoError(t, db.Migrate(gdb))

	router := app.NewRouter(gdb)
	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	// 1) Create
	createReq := map[string]any{"name": "JP", "email": "jp@example.com"}
	body, _ := json.Marshal(createReq)

	resp, err := http.Post(srv.URL+"/v1/users", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var created map[string]any
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))
	id := int(created["id"].(float64))
	require.Greater(t, id, 0)

	// 2) Get
	resp2, err := http.Get(srv.URL + "/v1/users/" + itoa(id))
	require.NoError(t, err)
	defer resp2.Body.Close()
	require.Equal(t, http.StatusOK, resp2.StatusCode)

	var got map[string]any
	require.NoError(t, json.NewDecoder(resp2.Body).Decode(&got))
	require.Equal(t, "JP", got["name"])
	require.Equal(t, "jp@example.com", got["email"])

	// 3) Update
	newName := "JP Updated"
	updReq := map[string]any{"name": newName}
	updBody, _ := json.Marshal(updReq)

	req, _ := http.NewRequest(http.MethodPut, srv.URL+"/v1/users/"+itoa(id), bytes.NewReader(updBody))
	req.Header.Set("content-type", "application/json")
	resp3, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp3.Body.Close()
	require.Equal(t, http.StatusOK, resp3.StatusCode)

	var updated map[string]any
	require.NoError(t, json.NewDecoder(resp3.Body).Decode(&updated))
	require.Equal(t, newName, updated["name"])

	// 4) List
	resp4, err := http.Get(srv.URL + "/v1/users")
	require.NoError(t, err)
	defer resp4.Body.Close()
	require.Equal(t, http.StatusOK, resp4.StatusCode)

	var list []map[string]any
	require.NoError(t, json.NewDecoder(resp4.Body).Decode(&list))
	require.GreaterOrEqual(t, len(list), 1)

	// 5) Delete
	reqDel, _ := http.NewRequest(http.MethodDelete, srv.URL+"/v1/users/"+itoa(id), nil)
	resp5, err := http.DefaultClient.Do(reqDel)
	require.NoError(t, err)
	defer resp5.Body.Close()
	require.Equal(t, http.StatusNoContent, resp5.StatusCode)

	// 6) Get again -> 404
	resp6, err := http.Get(srv.URL + "/v1/users/" + itoa(id))
	require.NoError(t, err)
	defer resp6.Body.Close()
	require.Equal(t, http.StatusNotFound, resp6.StatusCode)
}

// tiny int->string without fmt to keep test minimal
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var b [32]byte
	i := len(b)
	n := v
	for n > 0 {
		i--
		b[i] = byte('0' + (n % 10))
		n /= 10
	}
	return string(b[i:])
}
