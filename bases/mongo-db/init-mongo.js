db.createUser({
    user: "grupo33",
    pwd: "pass+1234",
    roles: [
      {
        role: "readWrite",
        db: "db_sopes",
      },
    ],
  });