# Body Calculator Telegram Bot

**Soon (will be working after finish project): <a href="https://t.me/body_calculator_bot">Body calculator bot</a>.**

This bot will help you find out your BMI, as well as your daily calorie intake.

Lib:
- <a href="https://github.com/go-telegram-bot-api/telegram-bot-api">telegram-bot-api</a>.

---

## Run:

Create **.env** file in the directory of this project, and paste your telegram token from <a href="http://t.me/BotFather">BotFather</a> and parameters for **Docker container** (See **.env.example**):


For the first run:

```console
make migrate
make build
make run
```

Then:

To stop: **Ctrl + C** in terminal

To stop container:
```console
make stop
```

To delete image:
```console
make migrate_down
```
---

## Help

For Makefile in Windows (<a href="https://www.gnu.org/software/make/#download">make</a>):
```
scoop install main/make
```

If scoop is not installed: <a href="https://scoop.sh/">scoop</a>

---
Mikhail Rogalsky 2024 (MIT LICENSE)
