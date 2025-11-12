-- +goose Up
CREATE TABLE expenses (
  id uuid PRIMARY KEY,
  expense_id UUID NOT NULL REFERENCES expense_type(id) ON DELETE CASCADE,
  total_expense INTEGER NOT NULL,
  transaction_date TIMESTAMP NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE expenses;
