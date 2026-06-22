CREATE TABLE IF NOT EXISTS violations (
	id UUID PRIMARY KEY,
	plate_number TEXT NOT NULL,
	location TEXT NOT NULL,
	occurred_at TIMESTAMPTZ NOT NULL,
	photo_url TEXT NOT NULL,
	officer_id UUID NOT NULL,
	fine_rule_version_id UUID NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT fk_violations_officer
		FOREIGN KEY (officer_id)
		REFERENCES users (id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT,
	CONSTRAINT fk_violations_fine_rule_version
		FOREIGN KEY (fine_rule_version_id)
		REFERENCES fine_rule_versions (id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_violations_deleted_at ON violations (deleted_at);
CREATE INDEX IF NOT EXISTS idx_violations_created_at ON violations (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_violations_officer_id ON violations (officer_id);
CREATE INDEX IF NOT EXISTS idx_violations_fine_rule_version_id ON violations (fine_rule_version_id);
