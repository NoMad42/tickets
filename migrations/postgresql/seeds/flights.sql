insert into flights (
    code,
    from_airport_id,
    from_timestamp,
    to_airport_id,
    to_timestamp,
    status
)
values (
    'test4-code',
    (select id from airports order by random() limit 1),
    current_timestamp,
    (select id from airports order by random() limit 1),
    current_timestamp,
    'test4-status'
);