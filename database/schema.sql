CREATE TABLE rooms (
    "id"    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "theme" TEXT NOT NULL
);

CREATE TABLE messages (
    "id"             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "content"        TEXT NOT NULL,
    "reaction_count" INT NOT NULL DEFAULT 0,
    "answered"       BOOLEAN NOT NULL DEFAULT false,
    "room_id"        UUID NOT NULL REFERENCES rooms(id)
)
