CREATE TABLE IF NOT EXISTS fine_rule_versions (
	id UUID PRIMARY KEY,
	version_number INTEGER NOT NULL,
	is_active BOOLEAN NOT NULL DEFAULT FALSE,
	published_by UUID NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT fk_fine_rule_versions_published_by
		FOREIGN KEY (published_by)
		REFERENCES users (id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_fine_rule_versions_deleted_at ON fine_rule_versions (deleted_at);
CREATE INDEX IF NOT EXISTS idx_fine_rule_versions_created_at ON fine_rule_versions (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_fine_rule_versions_published_by ON fine_rule_versions (published_by);
