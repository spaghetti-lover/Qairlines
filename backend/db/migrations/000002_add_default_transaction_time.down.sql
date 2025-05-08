ALTER TABLE payment
ALTER COLUMN transaction_date_time DROP DEFAULT,
ALTER COLUMN transaction_date_time DROP NOT NULL;
