package atlas

import (
	"context"
	"os"
	"testing"
	"time"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/helper/logging"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	envVarRunAccTests    = "VAULT_ACC"
	envVarPrivateKey     = "ATLAS_PRIVATE_KEY"
	envVarPublicKey      = "ATLAS_PUBLIC_KEY"
	envVarProjectID      = "ATLAS_PROJECT_ID"
	envVarOrganizationID = "ATLAS_ORGANIZATION_ID"
)

var runAcceptanceTests = os.Getenv(envVarRunAccTests) == "1"

func TestAcceptanceDatabaseUser(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add role", acceptanceTestEnv.AddRole)
	t.Run("read database user creds", acceptanceTestEnv.ReadDatabaseUserCreds)
	t.Run("renew database user creds", acceptanceTestEnv.RenewDatabaseUserCreds)
	t.Run("revoke database user creds", acceptanceTestEnv.RevokeDatabaseUsersCreds)
}

func TestAcceptanceProgrammaticAPIKey(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRole)
	t.Run("read progammatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew progammatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke progammatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)

}

func TestAcceptanceProgrammaticAPIKey_WithProjectID(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithProjectID)
	t.Run("read progammatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew progammatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke progammatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)

}

func TestAcceptanceDatabaseUser_WithCustomTTL(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add role with ttl", acceptanceTestEnv.AddRoleWithTTL)
	t.Run("read database user creds", acceptanceTestEnv.ReadDatabaseUserCreds)
	t.Run("renew database user creds", acceptanceTestEnv.RenewDatabaseUserCreds)
	t.Run("revoke database user creds", acceptanceTestEnv.RevokeDatabaseUsersCreds)
}

func TestAcceptanceProgrammaticAPIKey_WithIPWhitelist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithIP)
	t.Run("read progammatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew progammatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke progammatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)

}

func TestAcceptanceProgrammaticAPIKey_WithCIDRWhitelist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithCIDR)
	t.Run("read progammatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew progammatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke progammatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)

}

func newAcceptanceTestEnv() (*testEnv, error) {
	ctx := context.Background()
	conf := &logical.BackendConfig{
		System: &logical.StaticSystemView{
			DefaultLeaseTTLVal: time.Hour,
			MaxLeaseTTLVal:     time.Hour,
		},
		Logger: logging.NewVaultLogger(log.Debug),
	}
	b, err := Factory(ctx, conf)
	if err != nil {
		return nil, err
	}
	return &testEnv{
		PublicKey:      os.Getenv(envVarPublicKey),
		PrivateKey:     os.Getenv(envVarPrivateKey),
		ProjectID:      os.Getenv(envVarProjectID),
		OrganizationID: os.Getenv(envVarOrganizationID),
		Backend:        b,
		Context:        ctx,
		Storage:        &logical.InmemStorage{},
	}, nil
}
