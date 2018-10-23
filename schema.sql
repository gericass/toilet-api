-- +migrate Up
CREATE TABLE toilets (
  `id`          BIGINT UNSIGNED       AUTO_INCREMENT PRIMARY KEY,
  `name`        VARCHAR(255) NOT NULL,
  `uid`         VARCHAR(255) NOT NULL,
  `lat`         FLOAT        NOT NULL,
  `lng`         FLOAT        NOT NULL,
  `geolocation` VARCHAR(255) NOT NULL,
  `image_path`  TEXT,
  `description` TEXT,
  `valuation`   FLOAT,
  `updated_at`  TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
  `id`         BIGINT UNSIGNED       AUTO_INCREMENT PRIMARY KEY,
  `name`       VARCHAR(255) NOT NULL,
  `google_id`  VARCHAR(255) NOT NULL,
  `icon_path`  TEXT,
  `created_at` TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE users_toilets (
  `id`         BIGINT UNSIGNED    AUTO_INCREMENT PRIMARY KEY,
  `user_id`    BIGINT    NOT NULL,
  `toilet_id`  BIGINT    NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
  `updated_at` TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE reviews (
  `id`         BIGINT UNSIGNED    AUTO_INCREMENT PRIMARY KEY,
  `toilet_id`  BIGINT    NOT NULL,
  `user_id`    BIGINT    NOT NULL,
  `valuation`  FLOAT     NOT NULL,
  `message`    TEXT,
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE toilets;
DROP TABLE users;
DROP TABLE users_toilets;
DROP TABLE reviews;