INSERT INTO
    `users` (
        `id`,
        `instance_id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `username`,
        `password`,
        `status`,
        `nickname`,
        `email`,
        `phone`,
        `bio`,
        `company`,
        `location`,
        `profile_url`
    )
VALUES (
        0,
        1,
        NOW(3),
        NOW(3),
        NULL,
        'admin',
        '$2a$10$SU8KidtYjdrevrMzhW6c0efPcegk0Tv9OOenG0SjDJrnbSbJ0VjV2',
        'admin',
        'admin',
        'admin@example.com',
        '11122223333',
        'admin user',
        'admin company',
        'admin location',
        'https://example.com/admin'
    );