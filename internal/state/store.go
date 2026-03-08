package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const currentConfigVersion = 1

type Config struct {
	Version   int       `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AppProfile struct {
	Env          string    `json:"env"`
	ClientID     string    `json:"client_id"`
	Secret       string    `json:"secret"`
	ClientName   string    `json:"client_name"`
	Language     string    `json:"language"`
	CountryCodes []string  `json:"country_codes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AccountSummary struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	Mask      string `json:"mask,omitempty"`
	Type      string `json:"type,omitempty"`
	Subtype   string `json:"subtype,omitempty"`
}

type ItemRecord struct {
	ItemID          string           `json:"item_id"`
	AccessToken     string           `json:"access_token"`
	InstitutionID   string           `json:"institution_id,omitempty"`
	InstitutionName string           `json:"institution_name,omitempty"`
	LinkToken       string           `json:"link_token,omitempty"`
	Products        []string         `json:"products,omitempty"`
	Accounts        []AccountSummary `json:"accounts,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type plaidAccount interface {
	GetAccountID() string
	GetName() string
	GetMask() string
	GetType() string
	GetSubtype() string
}

type Store struct {
	dir string
}

func New(dir string) *Store {
	return &Store{dir: dir}
}

func DefaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home directory: %w", err)
	}
	return filepath.Join(home, ".plaid-cli"), nil
}

func GetenvAny(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func DefaultConfig() Config {
	now := time.Now().UTC()
	return Config{
		Version:   currentConfigVersion,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p AppProfile) Validate() error {
	env := strings.TrimSpace(strings.ToLower(p.Env))
	if !slices.Contains([]string{"sandbox", "development", "production"}, env) {
		return fmt.Errorf("unsupported env %q", p.Env)
	}
	if strings.TrimSpace(p.ClientID) == "" {
		return errors.New("client_id is required")
	}
	if strings.TrimSpace(p.Secret) == "" {
		return errors.New("secret is required")
	}
	if strings.TrimSpace(p.ClientName) == "" {
		return errors.New("client_name is required")
	}
	if strings.TrimSpace(p.Language) == "" {
		return errors.New("language is required")
	}
	if len(p.CountryCodes) == 0 {
		return errors.New("at least one country code is required")
	}
	return nil
}

func (s *Store) Ensure() error {
	for _, dir := range []string{
		s.dir,
		filepath.Join(s.dir, "items"),
		filepath.Join(s.dir, "cache"),
		filepath.Join(s.dir, "logs"),
	} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return fmt.Errorf("create %s: %w", dir, err)
		}
	}
	return nil
}

func (s *Store) SaveConfig(cfg Config) error {
	if err := s.Ensure(); err != nil {
		return err
	}
	existing, err := s.LoadConfig()
	if err == nil && !existing.CreatedAt.IsZero() {
		cfg.CreatedAt = existing.CreatedAt
	}
	if cfg.CreatedAt.IsZero() {
		cfg.CreatedAt = time.Now().UTC()
	}
	cfg.UpdatedAt = time.Now().UTC()
	return writeJSONFile(s.ConfigPath(), cfg)
}

func (s *Store) LoadConfig() (Config, error) {
	var cfg Config
	err := readJSONFile(s.ConfigPath(), &cfg)
	return cfg, err
}

func (s *Store) SaveAppProfile(profile AppProfile) error {
	if err := s.Ensure(); err != nil {
		return err
	}
	existing, err := s.LoadAppProfile()
	if err == nil && !existing.CreatedAt.IsZero() {
		profile.CreatedAt = existing.CreatedAt
	}
	if profile.CreatedAt.IsZero() {
		profile.CreatedAt = time.Now().UTC()
	}
	profile.UpdatedAt = time.Now().UTC()
	return writeJSONFile(s.AppProfilePath(), profile)
}

func (s *Store) LoadAppProfile() (AppProfile, error) {
	var profile AppProfile
	err := readJSONFile(s.AppProfilePath(), &profile)
	return profile, err
}

func (s *Store) SaveItem(record ItemRecord) error {
	if err := s.Ensure(); err != nil {
		return err
	}
	existing, err := s.LoadItem(record.ItemID)
	if err == nil && !existing.CreatedAt.IsZero() {
		record.CreatedAt = existing.CreatedAt
	}
	if record.CreatedAt.IsZero() {
		record.CreatedAt = time.Now().UTC()
	}
	record.UpdatedAt = time.Now().UTC()
	return writeJSONFile(s.ItemPath(record.ItemID), record)
}

func (s *Store) LoadItem(itemID string) (ItemRecord, error) {
	var record ItemRecord
	err := readJSONFile(s.ItemPath(itemID), &record)
	return record, err
}

func (s *Store) DeleteItem(itemID string) error {
	if err := os.Remove(s.ItemPath(itemID)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("delete item %s: %w", itemID, err)
	}
	return nil
}

func (s *Store) ListItems() ([]ItemRecord, error) {
	entries, err := os.ReadDir(filepath.Join(s.dir, "items"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("read items directory: %w", err)
	}

	records := make([]ItemRecord, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		record, err := s.LoadItem(strings.TrimSuffix(entry.Name(), ".json"))
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	slices.SortFunc(records, func(a, b ItemRecord) int {
		return strings.Compare(a.ItemID, b.ItemID)
	})
	return records, nil
}

func (s *Store) FindItemByAccessToken(accessToken string) (*ItemRecord, error) {
	items, err := s.ListItems()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.AccessToken == accessToken {
			itemCopy := item
			return &itemCopy, nil
		}
	}
	return nil, os.ErrNotExist
}

func (s *Store) ConfigPath() string {
	return filepath.Join(s.dir, "config.json")
}

func (s *Store) AppProfilePath() string {
	return filepath.Join(s.dir, "app-profile.json")
}

func (s *Store) ItemPath(itemID string) string {
	return filepath.Join(s.dir, "items", itemID+".json")
}

func AccountSummariesFromPlaid[T plaidAccount](accounts []T) []AccountSummary {
	out := make([]AccountSummary, 0, len(accounts))
	for _, account := range accounts {
		out = append(out, AccountSummary{
			AccountID: account.GetAccountID(),
			Name:      account.GetName(),
			Mask:      account.GetMask(),
			Type:      account.GetType(),
			Subtype:   account.GetSubtype(),
		})
	}
	return out
}

func writeJSONFile(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create parent directory: %w", err)
	}

	file, err := os.CreateTemp(filepath.Dir(path), ".tmp-*.json")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tempName := file.Name()
	cleanup := func() {
		_ = os.Remove(tempName)
	}
	defer cleanup()

	if err := file.Chmod(0o600); err != nil {
		_ = file.Close()
		return fmt.Errorf("chmod temp file: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		_ = file.Close()
		return fmt.Errorf("encode json: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close temp file: %w", err)
	}
	if err := os.Rename(tempName, path); err != nil {
		return fmt.Errorf("rename temp file: %w", err)
	}
	return nil
}

func readJSONFile(path string, value any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, value); err != nil {
		return fmt.Errorf("decode json %s: %w", path, err)
	}
	return nil
}
