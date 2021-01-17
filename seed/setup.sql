use todo_db;
create table todos (
  id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  public_id VARCHAR(36),
  title varchar(255),
  content varchar(255),
  done boolean,
  created_at DATETIME
);
insert into todos (id, public_id, title, content, done, created_at)
values(
    1,
    '111111111111111111',
    'todo 1',
    'test todo',
    false,
    NOW()
  );