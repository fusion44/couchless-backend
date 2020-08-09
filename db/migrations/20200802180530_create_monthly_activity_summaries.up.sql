CREATE TABLE user_stat_months
(
  period DATE,
  total INTEGER NOT NULL,
  sport_type VARCHAR(45) NOT NULL,

  UNIQUE(period, sport_type),

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  FOREIGN KEY (sport_type) REFERENCES sport_types(name) ON DELETE CASCADE
);
