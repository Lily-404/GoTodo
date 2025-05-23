# Todo - 极简终端任务管理器

[English](REDME_zh.md) | [中文](README_zh.md)

Todo 是一个用 Go 语言编写的终端任务管理工具，专注于高效管理待办事项。

## 📦 安装

### 使用 Go Install（推荐）

```bash
go install github.com/Lily-404/todo@latest
```

### 从源码构建

1. 克隆仓库

```bash
git clone https://github.com/Lily-404/todo.git
cd todo
```

2. 构建项目

```bash
go build
```

## 🚀 使用方法

### 基本命令

- `todo task` 或 `todo t` - 交互式任务管理
- `todo add` 或 `todo a` - 添加任务
- `todo list` 或 `todo l` - 列出任务
- `todo remove` 或 `todo rm` - 删除任务
- `todo done` 或 `todo d` - 完成任务
- `todo clean` 或 `todo c` - 清理任务
- `todo lang`  - 切换语言设置

### 示例

1. 添加任务

```bash
todo add "完成项目文档"
```

2. 添加带优先级的任务

```bash
todo add "准备周会演示" -p high
```

3. 添加带截止日期的任务

```bash
todo add "提交季度报告" -d "2024-03-31"
```

4. 查看所有任务

```bash
todo list
```

5. 交互式管理任务

```bash
todo task
```

## 📁 数据存储

应用程序将数据存储在以下位置：

### macOS/Linux

- 配置文件：`~/.config/gotodo/config.json`
- 任务数据：`~/.local/share/gotodo/notes/notes.json`

### Windows

- 配置文件： `%APPDATA%\gotodo\config.json`
- 任务数据： `%APPDATA%\gotodo\notes\notes.json`

## 🎨 特色功能

- 优先级管理：支持 high、normal、low 三级优先级
- 交互式操作：通过交互式界面轻松管理任务
- 进度统计：直观显示任务完成进度
- 多彩输出：使用彩色文本增强可读性
- 国际化支持：支持中英文界面

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License
