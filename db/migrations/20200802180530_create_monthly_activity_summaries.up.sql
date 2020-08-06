CREATE TABLE user_stat_months
(
  period DATE PRIMARY KEY,
  total INTEGER NOT NULL,
  sport_type VARCHAR(45) NOT NULL,

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  FOREIGN KEY (sport_type) REFERENCES sport_types(name) ON DELETE CASCADE
);
