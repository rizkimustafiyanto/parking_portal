-- Add violation type so fine calculation can use the business rule base amount.
ALTER TABLE violations
ADD COLUMN IF NOT EXISTS violation_type TEXT NOT NULL DEFAULT '';
