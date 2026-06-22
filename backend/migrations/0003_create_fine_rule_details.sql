CREATE TABLE IF NOT EXISTS fine_rule_details (
	id UUID PRIMARY KEY,
	rule_version_id UUID NOT NULL,
	rule_type TEXT NOT NULL,
	key TEXT NOT NULL,
	value TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT fk_fine_rule_details_rule_version
		FOREIGN KEY (rule_version_id)
		REFERENCES fine_rule_versions (id)
		ON UPDATE CASCADE
		ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_fine_rule_details_deleted_at ON fine_rule_details (deleted_at);
CREATE INDEX IF NOT EXISTS idx_fine_rule_details_created_at ON fine_rule_details (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_fine_rule_details_rule_version_id ON fine_rule_details (rule_version_id);
