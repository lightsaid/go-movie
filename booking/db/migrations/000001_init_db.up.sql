CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "avatar_url" varchar,
  "password" varchar,
  "phone" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "oauths" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "oauth_type" varchar NOT NULL,
  "oauth_id" varchar NOT NULL,
  "unionid" varchar,
  "credential" varchar
);

CREATE TABLE "movie" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "release_at" timestamptz NOT NULL,
  "cover_url" varchar NOT NULL,
  "duration" int,
  "language" varchar,
  "genre" varchar,
  "rating" float DEFAULT '0.0',
  "director" varchar,
  "desc" varchar,
  "status" int NOT NULL DEFAULT '0',
  "star" varchar,
  "wish_count" int
);

CREATE TABLE "city" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "lat" decimal NOT NULL,
  "long" decimal NOT NULL
);

CREATE TABLE "cinema" (
  "id" bigserial PRIMARY KEY,
  "city_id" bigint NOT NULL,
  "name" varchar(50) NOT NULL,
  "lat" decimal NOT NULL,
  "long" decimal NOT NULL,
  "total_cinema_halls" int NOT NULL
);

CREATE TABLE "cinema_hall" (
  "id" bigserial PRIMARY KEY,
  "cinema_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "total_seats" int NOT NULL
);

CREATE TABLE "cinema_seat" (
  "id" bigserial PRIMARY KEY,
  "type" int NOT NULL DEFAULT '0',
  "cinema_hall_id" bigint NOT NULL,
  "seat_number" varchar NOT NULL
);

CREATE TABLE "show" (
  "id" bigserial PRIMARY KEY,
  "date" timestamptz NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "cinema_hall_id" bigint NOT NULL,
  "movie_id" bigint NOT NULL
);

CREATE TABLE "show_seat" (
  "id" bigserial PRIMARY KEY,
  "cinema_seat_id" bigint NOT NULL,
  "show_id" bigint NOT NULL,
  "booking_id" bigint NOT NULL,
  "status" int NOT NULL,
  "price" decimal(8,2) NOT NULL
);

CREATE TABLE "booking" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "show_id" bigint NOT NULL,
  "seat_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "status" int NOT NULL
);

CREATE TABLE "payment" (
  "id" bigserial PRIMARY KEY,
  "booking_id" bigint NOT NULL,
  "amount" decimal(8,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "user" ("phone");

CREATE UNIQUE INDEX ON "city" ("name");

COMMENT ON COLUMN "oauths"."oauth_type" IS '?????????????????????wechat\qq\weibo???';

COMMENT ON COLUMN "oauths"."oauth_id" IS '???????????????????????????';

COMMENT ON COLUMN "oauths"."unionid" IS '????????????unionid';

COMMENT ON COLUMN "oauths"."credential" IS '????????????/access_token';

COMMENT ON COLUMN "movie"."title" IS '?????????';

COMMENT ON COLUMN "movie"."release_at" IS '????????????';

COMMENT ON COLUMN "movie"."cover_url" IS '??????';

COMMENT ON COLUMN "movie"."duration" IS '??????(????????????)';

COMMENT ON COLUMN "movie"."language" IS '??????';

COMMENT ON COLUMN "movie"."genre" IS '??????/??????';

COMMENT ON COLUMN "movie"."rating" IS '??????';

COMMENT ON COLUMN "movie"."director" IS '??????';

COMMENT ON COLUMN "movie"."desc" IS '??????';

COMMENT ON COLUMN "movie"."status" IS '0:???????????????1:????????????';

COMMENT ON COLUMN "movie"."star" IS '????????????';

COMMENT ON COLUMN "movie"."wish_count" IS '????????????';

COMMENT ON COLUMN "cinema_seat"."type" IS '??????/??????/??????';

COMMENT ON COLUMN "cinema_seat"."seat_number" IS '?????????????????????1???2??????';

ALTER TABLE "oauths" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "cinema" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "cinema_hall" ADD FOREIGN KEY ("cinema_id") REFERENCES "cinema" ("id");

ALTER TABLE "cinema_seat" ADD FOREIGN KEY ("cinema_hall_id") REFERENCES "cinema_hall" ("id");

ALTER TABLE "show" ADD FOREIGN KEY ("cinema_hall_id") REFERENCES "cinema_hall" ("id");

ALTER TABLE "show" ADD FOREIGN KEY ("movie_id") REFERENCES "movie" ("id");

ALTER TABLE "show_seat" ADD FOREIGN KEY ("cinema_seat_id") REFERENCES "cinema_seat" ("id");

ALTER TABLE "show_seat" ADD FOREIGN KEY ("show_id") REFERENCES "show" ("id");

ALTER TABLE "show_seat" ADD FOREIGN KEY ("booking_id") REFERENCES "booking" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("show_id") REFERENCES "show" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("booking_id") REFERENCES "booking" ("id");
