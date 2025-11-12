-- +goose Up
CREATE TABLE income (
  id uuid PRIMARY KEY,
  income_id UUID NOT NULL REFERENCES income_type(id) ON DELETE CASCADE,
  total_income INTEGER NOT NULL,
  transaction_date TIMESTAMP NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE income;
