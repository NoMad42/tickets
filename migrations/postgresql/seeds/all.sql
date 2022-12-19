INSERT INTO user_profiles (
    login,
    password,
    avatar
  )
VALUES (
    'admin',
    'secret',
    'https://example.com/avatar.jpg'
  );

INSERT INTO user_profiles (
    login,
    password,
    avatar,
    name,
    lastname,
    personal_id
) VALUES (
    'test',
    'test-secret',
    'test-avatar',
    'test-name',
    'test-lastname',
    'test-personal_id'
    'test'
  ), (
    'tets1',
    'tets1-secret',
    'tets1-avatar',
    'tets1-name',
    'tets1-lastname',
    'tets1-personal_id'
    'tets1'
  ), (
    'tets2',
    'tets2-secret',
    'tets2-avatar',
    'tets2-name',
    'tets2-lastname',
    'tets2-personal_id'
    'tets2'
  );

INSERT INTO airports (
    code,
    name,
    description,
    country,
    city
) VALUES (
    'test-code',
    'test-name',
    'test-description',
    'test-country',
    'test-city'
  ), (
    'test1-code',
    'test1-name',
    'test1-description',
    'test1-country',
    'test1-city'
  ), (
    'test2-code',
    'test2-name',
    'test2-description',
    'test2-country',
    'test2-city'
  );

insert into flights (
    code,
    from_airport_id,
    from_timestamp,
    to_airport_id,
    to_timestamp,
    status
)
values (
    'test1-code',
    (select id from airports order by random() limit 1),
    current_timestamp,
    (select id from airports order by random() limit 1),
    current_timestamp,
    'test2-status'
), (
    'test2-code',
    (select id from airports order by random() limit 1),
    current_timestamp,
    (select id from airports order by random() limit 1),
    current_timestamp,
    'test2-status'
), (
    'test3-code',
    (select id from airports order by random() limit 1),
    current_timestamp,
    (select id from airports order by random() limit 1),
    current_timestamp,
    'test3-status'
);

INSERT INTO seats (
    id,
    code,
    position,
    price,
    booking_status,
    seat_type
  )
VALUES (
    default,
    'test1-code',
    'middle',
    123,
    'free',
    'economy'
  ),
  (
    default,
    'test2-code',
    'aisle',
    234,
    'booking',
    'economy'
  ),
  (
    default,
    'test3-code',
    'window',
    345,
    'booked',
    'buisness'
  );

INSERT INTO seat_options (id, name, description, price)
VALUES (
    default,
    'Дополнительное место под багаж',
    'Дополнительное место под багаж. До 35 кг.',
    123
  ),(
    default,
    'Второй ланч',
    'Второй ланч. Очень Вкусный.',
    123
  ),(
    default,
    'Плед',
    'Плед. Чтобы было тёпленько',
    123
  ),(
    default,
    'Сувениры.',
    'Сувениры которые будут напоминать о прекрасном полёте',
    123
  ),(
    default,
    'Страховка',
    'Страховка. От всего для всего. Будьте спокойны.',
    123
  );

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
