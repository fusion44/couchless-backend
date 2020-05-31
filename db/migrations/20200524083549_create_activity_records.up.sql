CREATE TABLE activity_records
(
  id BIGSERIAL PRIMARY KEY,
  "timestamp" TIMESTAMP NOT NULL,
  position_lat INTEGER,
  position_long INTEGER,
  distance FLOAT,
  time_from_course INTEGER,
  compressed_speed_distance bytea,
  heart_rate SMALLINT,
  altitude REAL,
  speed REAL,
  power INTEGER,
  grade INTEGER,
  cadence SMALLINT,
  fractional_cadence INTEGER,
  resistance SMALLINT,
  cycle_length SMALLINT,
  temperature SMALLINT,
  accumulated_power INTEGER,

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  activity_id BIGSERIAL REFERENCES activities(id) ON DELETE CASCADE
);