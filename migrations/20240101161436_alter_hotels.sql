-- +goose Up
-- Step 1: Add the column without UNIQUE constraint
ALTER TABLE hotels ADD COLUMN api_key VARCHAR(64) DEFAULT (
    LOWER(HEX(RANDOMBLOB(32)))
);

-- Step 2: Update existing records with unique values
UPDATE hotels SET api_key = LOWER(HEX(RANDOMBLOB(32))) WHERE api_key IS NULL;

-- Step 3: Add UNIQUE constraint
CREATE UNIQUE INDEX idx_hotels_api_key ON hotels(api_key);

-- +goose Down
-- Step 1: Remove UNIQUE constraint
DROP INDEX idx_hotels_api_key;

-- Step 2: Remove the column
ALTER TABLE hotels DROP COLUMN api_key;
