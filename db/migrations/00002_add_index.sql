-- +goose Up
-- +goose StatementBegin
create index user_index on donations(user_id);
create index campaign_index on donations(campaign_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index user_index;
drop index campaign_index;
-- +goose StatementEnd
