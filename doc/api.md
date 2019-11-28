# API Reference

## 模型

### 设置

```json
{
   "student_id":  学生id
   "mode":        设置主题模式（dark/light）
   "saveTodoIn":  ToDo项存储在哪（server/client）
}
```

### 期望的POST输入

```json
{
    "mode": 'dark' | 'light';
    "saveSettingsIn": 'server' | 'client';
    "saveTodoIn": 'server' | 'client';
}
```

## web api

- `GET /ping`

  检查服务是否可用，应该直接返回`pong`。

- `GET /config`

  返回头部Authorization字段内JWT用户对应的Config。

- `POST /config`

  更新或新增Authorization字段内JWT用户对应的Config。
  