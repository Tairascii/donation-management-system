-- name: Donate :exec
insert into donations (user_id, campaign_id, amount) values ($1, $2, $3);

-- name: DonationsByUserID :many
select id, campaign_id, amount, donated_at from donations where user_id = $1;

-- name: CampaignByID :one
with total as (
    select sum(amount) as total_amount, campaign_id
    from donations
    group by campaign_id
)
select c.id, c.organization_id, c.title, c.description, c.goal_amount, t.total_amount
from campaigns c left join total t on c.id = t.campaign_id
where c.id = $1;

-- name: CreateCampaign :one
insert into campaigns (organization_id, title, description, goal_amount, start_date, end_date)
       values ($1, $2, $3, $4, $5, $6) returning id;

-- name: DeleteCampaign :exec
delete from campaigns where id = $1;

