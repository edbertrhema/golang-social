CREATE TABLE IF NOT EXISTS followers (
    user_id bigint NOT NULL,
    follower_id bigint NOT NULL,
    created_at timestamp (0) with time zone NOT NULL DEFAULT now(),

    primary key (user_id, follower_id),
    foreign key (user_id) references users (id) ON DELETE CASCADE ,
    foreign key (follower_id) references users (id) ON DELETE CASCADE


)