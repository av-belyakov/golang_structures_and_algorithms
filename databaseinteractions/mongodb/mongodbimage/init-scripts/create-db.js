databaseName = process.env.MONGO_INITDB_DATABASE;
userName = process.env.MONGO_INITDB_ROOT_USERNAME;
passwd = process.env.MONGO_INITDB_ROOT_PASSWORD;

// Создаётся в контексте базы MONGO_INITDB_DATABASE
db = db.getSiblingDB(databaseName);

// Создаём коллекцию для инициализации БД
db.createCollection("init");

// Создаём пользователя для этой БД
db.createUser({
  user: userName,
  pwd: passwd,
  roles: [
    {
      role: "readWrite",
      db: databaseName,
    },
  ],
});

print(
  'Database "' +
    databaseName +
    '" and user "' +
    userName +
    '" created successfully!',
);
