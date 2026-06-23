CREATE TABLE IF NOT EXISTS invoices (
	id UUID PRIMARY KEY,
	violation_id UUID NOT NULL UNIQUE,
	member_id UUID NOT NULL,
	amount NUMERIC(18,2) NOT NULL DEFAULT 0,
	status TEXT NOT NULL DEFAULT 'PENDING',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT fk_invoices_violation
		FOREIGN KEY (violation_id)
		REFERENCES violations (id)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	CONSTRAINT fk_invoices_member
		FOREIGN KEY (member_id)
		REFERENCES users (id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_invoices_deleted_at ON invoices (deleted_at);
CREATE INDEX IF NOT EXISTS idx_invoices_created_at ON invoices (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_invoices_member_id ON invoices (member_id);
CREATE INDEX IF NOT EXISTS idx_invoices_violation_id ON invoices (violation_id);

CREATE TABLE IF NOT EXISTS payment_transactions (
	id UUID PRIMARY KEY,
	invoice_id UUID NOT NULL UNIQUE,
	internal_transaction_id TEXT NOT NULL,
	amount NUMERIC(18,2) NOT NULL DEFAULT 0,
	status TEXT NOT NULL,
	scenario TEXT NOT NULL,
	paid_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT fk_payment_transactions_invoice
		FOREIGN KEY (invoice_id)
		REFERENCES invoices (id)
		ON UPDATE CASCADE
		ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_payment_transactions_deleted_at ON payment_transactions (deleted_at);
CREATE INDEX IF NOT EXISTS idx_payment_transactions_created_at ON payment_transactions (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_payment_transactions_invoice_id ON payment_transactions (invoice_id);
