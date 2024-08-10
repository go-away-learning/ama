-- name: GetRoom :one
SELECT *
  FROM rooms
 WHERE id = $1
 LIMIT 1;

-- name: ListRooms :many
  SELECT *
    FROM rooms;

-- name: CreateRoom :one
INSERT INTO rooms (
                theme
            )
     VALUES (
                $1
            )
  RETURNING id;

-- name: GetMessage :one
SELECT *
  FROM messages
 WHERE id = $1;

-- name: GetRoomMessages :many
SELECT *
  FROM messages
 WHERE room_id = $1;

-- name: CreateMessage :one
INSERT INTO messages (
                content, room_id
            )
     VALUES (
                $1, $2
            )
  RETURNING id;

-- name: IncreaseMessageReactions :one
   UPDATE messages
      SET reaction_count = reaction_count + 1
    WHERE id = $1
RETURNING reaction_count;

-- name: DecreaseMessageReactions :one
   UPDATE messages
      SET reaction_count = reaction_count - 1
    WHERE id = $1
RETURNING reaction_count;

-- name MarkMessageAsAnswered :one
   UPDATE messages
      SET answered = true
    WHERE id = $1;
