-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    category_id,
    title,
    type,
    description,
    value,
    date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;


-- name: GetAccount :one
SELECT * FROM accounts WHERE id =$1 LIMIT 1;

-- name: GetAccounts :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.category_id=$3
AND a.title Like $4 AND a.description Like $5 AND a.date=$6; 


-- name: GetAccountsByUserIdAndType :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2; 


-- name: GetAccountsByUserIdAndTypeAndCategoryId :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.category_id=$3; 


-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitle :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.category_id=$3
AND a.title Like $4;

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.category_id=$3
AND a.title Like $4 AND a.description Like $5; 

-- name: GetAccountsByUserIdAndTypeIdAndTitle :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.title Like $3;

-- name: GetAccountsByUserIdAndTypeIdAndDescription :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.description Like $3;
 
 -- name: GetAccountsByUserIdAndTypeIdAndDate :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id=a.category_id
WHERE a.user_id =$1 AND a.type =$2 AND a.date Like $3;

-- name: GetAccountsReports :one
SELECT SUM(value) FROM accounts WHERE user_id= $1 AND type=$2;

-- name: GetAccountsGraph :one
SELECT COUNT(*) AS sum_value FROM accounts WHERE user_id= $1 AND type=$2;


-- name: UpdateAccount :one
UPDATE accounts SET title =$2, description =$3, value=$4 WHERE id=$1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id =$1;