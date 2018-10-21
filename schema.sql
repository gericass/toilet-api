CREATE TABLE toilets (
  `id`          BIGINT UNSIGNED       AUTO_INCREMENT PRIMARY KEY,
  `name`        VARCHAR(255) NOT NULL,
  `google_id`   VARCHAR(255) NOT NULL,
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