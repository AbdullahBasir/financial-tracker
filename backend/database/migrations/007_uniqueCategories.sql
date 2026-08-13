-- +goose Up
ALTER TABLE categories
ADD CONSTRAINT unique_user_cayegory_name UNIQUE(user_id, name);

-- +goose Down
ALTER TABLE categories
DROP CONSTRAINT unique_user_category_name;