ALTER TABLE payment
ALTER COLUMN transaction_date_time SET DEFAULT NOW(),
ALTER COLUMN transaction_date_time SET NOT NULL;