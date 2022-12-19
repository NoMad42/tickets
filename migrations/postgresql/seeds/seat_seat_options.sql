INSERT INTO seats_seat_options (
    seats_id,
    seat_options_id
  )
VALUES (
    (select id from seats  limit 1),
    (select id from seat_options limit 1)
  ),(
    (select id from seats limit 1 offset 1),
    (select id from seat_options limit 1 offset 1)
  ),(
    (select id from seats offset 1 limit 1),
    (select id from seat_options offset 2 limit 1)
  ),(
    (select id from seats offset 2 limit 1),
    (select id from seat_options offset 1 limit 1)
  ),(
    (select id from seats  offset 2 limit 1),
    (select id from seat_options offset 2 limit 1)
  );
