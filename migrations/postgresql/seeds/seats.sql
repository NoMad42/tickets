INSERT INTO seats (
    id,
    flight_id,
    code,
    position,
    price,
    booking_status,
    seat_type
  )
VALUES (
    default,
    (select id from flights order by random() limit 1),
    'test1-code',
    'middle',
    123,
    'free',
    'economy'
  ),
  (
    default,
    (select id from flights order by random() limit 1),
    'test2-code',
    'aisle',
    234,
    'booking',
    'economy'
  ),
  (
    default,
    (select id from flights order by random() limit 1),
    'test3-code',
    'window',
    345,
    'booked',
    'buisness'
  );
