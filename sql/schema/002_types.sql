-- +goose Up
CREATE TABLE expense_type (
  id UUID PRIMARY KEY,
  description TEXT NOT NULL
);


CREATE TABLE income_type (
  id UUID PRIMARY KEY,
  description TEXT NOT NULL
);

-- +goose Down
DROP TABLE expense_type;
DROP TABLE income_type;
