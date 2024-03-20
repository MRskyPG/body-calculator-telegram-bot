CREATE TABLE food_products (
                               id SERIAL PRIMARY KEY,
                               product_name VARCHAR(100) NOT NULL,
                               calories_per_100g DECIMAL NOT NULL,
                               proteins DECIMAL NOT NULL,
                               fats DECIMAL NOT NULL,
                               carbohydrates DECIMAL NOT NULL,
                               product_type VARCHAR(50) NOT NULL
);

INSERT INTO food_products (product_name, calories_per_100g, proteins, fats, carbohydrates, product_type)
VALUES
    ('Свинина', 317, 27, 24, 0, 'Мясопродукты'),
    ('Говядина', 250, 26, 17, 0, 'Мясопродукты'),
    ('Курица', 239, 21, 15, 0, 'Мясопродукты'),
    ('Утка', 337, 19, 30, 0, 'Мясопродукты'),
    ('Индейка', 135, 29, 3, 0, 'Мясопродукты'),
    ('Фарш мясной', 332, 16, 28, 0, 'Мясопродукты'),
    ('Колбаса', 250, 13, 22, 1, 'Мясопродукты'),
    ('Тунец', 144, 30, 2, 0, 'Рыбопродукты'),
    ('Креветки', 85, 18, 1.7, 0, 'Рыбопродукты'),
    ('Семга', 208, 20, 14, 0, 'Рыбопродукты'),
    ('Хек', 76, 17, 1.3, 0, 'Рыбопродукты'),
    ('Лосось', 208, 20, 13, 0, 'Рыбопродукты'),
    ('Угорь', 184, 18, 12, 0, 'Рыбопродукты'),
    ('Мидии', 70, 12, 1, 3, 'Рыбопродукты'),
    ('Осетр', 146, 19, 7, 0, 'Рыбопродукты'),
    ('Минтай', 87, 17, 1.3, 0, 'Рыбопродукты'),
    ('Кальмар', 79, 15.6, 0.9, 2, 'Рыбопродукты'),
    ('Омлет', 154, 13, 11, 1, 'Яйца'),
    ('Куриное яйцо', 155, 13, 11, 1, 'Яйца'),
    ('Страусиное яйцо', 140, 13, 9, 2, 'Яйца'),
    ('Перепелиное яйцо', 168, 13, 11, 1, 'Яйца'),
    ('Яичный порошок', 548, 48, 38, 3, 'Яйца'),
    ('Пашот', 90, 10, 6, 1, 'Яйца'),
    ('Яичница', 149, 13, 11, 1, 'Яйца'),
    ('Молоко', 42, 3.4, 1, 5, 'Молочные продукты'),
    ('Творог', 174, 18, 18, 3.6, 'Молочные продукты'),
    ('Сметана', 208, 2.6, 20, 3.6, 'Молочные продукты'),
    ('Айран', 30, 1, 3, 1, 'Молочные продукты'),
    ('Кефир', 51, 3.3, 2, 4.7, 'Молочные продукты'),
    ('Сыр', 356, 24, 28, 2.2, 'Молочные продукты'),
    ('Хлеб', 265, 9, 1, 49, 'Хлебобулочные изделия'),
    ('Багет', 265, 9, 1, 49, 'Хлебобулочные изделия'),
    ('Булочка', 300, 6, 10, 50, 'Хлебобулочные изделия'),
    ('Пита', 275, 9, 1.5, 53, 'Хлебобулочные изделия'),
    ('Сдоба', 330, 7, 18, 36, 'Хлебобулочные изделия'),
    ('Чиабатта', 250, 8, 2, 50, 'Хлебобулочные изделия'),
    ('Крендель', 330, 9, 2, 70, 'Хлебобулочные изделия'),
    ('Ржаной хлеб', 165, 5, 1, 33, 'Хлебобулочные изделия'),
    ('Пампушка', 250, 7, 10, 35, 'Хлебобулочные изделия'),
    ('Перловка', 123, 12.6, 2.6, 62, 'Крупы'),
    ('Гречка', 123, 12.6, 2.6, 62, 'Крупы'),
    ('Рис', 130, 2.7, 0.3, 28, 'Крупы'),
    ('Пшено', 123, 9.3, 1.3, 73, 'Крупы'),
    ('Кукурузная крупа', 342, 9.3, 1.2, 74, 'Крупы'),
    ('Манная крупа', 334, 10.3, 0.6, 73, 'Крупы'),
    ('Ячневая крупа', 342, 8.3, 2.2, 67, 'Крупы'),
    ('Гороховая крупа', 335, 11.6, 1.6, 62, 'Крупы'),
    ('Овсянка', 342, 11.6, 6.6, 59, 'Крупы'),
    ('Перловая крупа', 342, 11.6, 0.6, 74, 'Крупы'),
    ('Фасоль', 333, 21, 1, 63, 'Бобовые'),
    ('Нут', 341, 20, 6, 60, 'Бобовые'),
    ('Чечевица', 116, 9, 0.4, 20, 'Бобовые'),
    ('Горох', 81, 5, 0.4, 14, 'Бобовые'),
    ('Помидор', 18, 0.9, 0.2, 3.9, 'Овощи'),
    ('Огурец', 15, 0.8, 0.1, 2.9, 'Овощи'),
    ('Картофель', 77, 2, 0.1, 17, 'Овощи'),
    ('Морковь', 41, 0.9, 0.2, 10, 'Овощи'),
    ('Баклажан', 24, 1.2, 0.1, 5.7, 'Овощи'),
    ('Капуста', 25, 1.3, 0.1, 5.8, 'Овощи'),
    ('Перец', 31, 1.3, 0.2, 5.3, 'Овощи'),
    ('Свекла', 43, 1.6, 0.1, 10, 'Овощи'),
    ('Кабачок', 17, 0.6, 0.3, 3.6, 'Овощи'),
    ('Лук', 40, 1.1, 0.1, 9.3, 'Овощи'),
    ('Яблоко', 52, 0.3, 0.2, 14, 'Фрукты'),
    ('Банан', 89, 1.1, 0.3, 23, 'Фрукты'),
    ('Апельсин', 43, 1, 0.2, 8.2, 'Фрукты'),
    ('Груша', 57, 0.4, 0.3, 15, 'Фрукты'),
    ('Персик', 39, 0.9, 0.1, 9.5, 'Фрукты'),
    ('Ананас', 50, 0.5, 0.1, 13, 'Фрукты'),
    ('Киви', 61, 1.1, 0.5, 14.7, 'Фрукты'),
    ('Манго', 60, 0.8, 0.4, 15, 'Фрукты'),
    ('Арбуз', 30, 0.6, 0.1, 6, 'Фрукты'),
    ('Персик', 39, 0.9, 0.1, 9.5, 'Фрукты'),
    ('Виноград', 69, 0.6, 0.2, 18, 'Фрукты'),
    ('Абрикос', 48, 1.4, 0.1, 11, 'Фрукты'),
    ('Клубника', 33, 0.8, 0.4, 7.7, 'Ягоды'),
    ('Малина', 52, 1.2, 0.3, 11.9, 'Ягоды'),
    ('Черника', 57, 1.1, 0.5, 12, 'Ягоды'),
    ('Голубика', 44, 0.7, 0.4, 8.2, 'Ягоды'),
    ('Земляника', 30, 0.8, 0.4, 5.4, 'Ягоды'),
    ('Клюква', 46, 0.4, 0.1, 9.6, 'Ягоды'),
    ('Ежевика', 32, 1.4, 0.5, 4.9, 'Ягоды'),
    ('Смородина', 44, 1, 0.4, 7.3, 'Ягоды'),
    ('Шиповник', 162, 1.6, 0.7, 38, 'Ягоды'),
    ('Крыжовник', 43, 1, 0.4, 7.7, 'Ягоды'),
    ('Жимолость', 58, 1.8, 0.4, 9.8, 'Ягоды'),
    ('Рябина', 83, 1.2, 0.3, 18, 'Ягоды'),
    ('Миндаль', 576, 21, 49, 22, 'Орехи'),
    ('Фундук', 650, 15, 61, 16, 'Орехи'),
    ('Кешью', 553, 18, 44, 30, 'Орехи'),
    ('Лесной орех', 674, 14, 65, 13, 'Орехи'),
    ('Грецкий орех', 654, 15, 65, 14, 'Орехи'),
    ('Кедровый орех', 673, 13, 61, 17, 'Орехи'),
    ('Фисташки', 562, 21, 45, 28, 'Орехи'),
    ('Кокосовая стружка', 354, 3.3, 33, 15, 'Орехи'),
    ('Бразильский орех', 656, 14, 66, 12, 'Орехи'),
    ('Макадамия', 718, 8, 76, 14, 'Орехи'),
    ('Пекан', 691, 9, 72, 14, 'Орехи'),
    ('Кунжут', 573, 17, 49, 23, 'Орехи'),
    ('Фисташки', 562, 21, 45, 28, 'Орехи'),
    ('Арахис', 567, 26, 49, 16, 'Орехи'),
    ('Жареный арахис', 567, 26, 49, 16, 'Орехи'),
    ('Шампиньоны', 22, 3.1, 0.3, 3.3, 'Грибы'),
    ('Лисички', 22, 2.5, 0.5, 3.5, 'Грибы'),
    ('Белые грибы', 27, 3.3, 0.5, 2.5, 'Грибы'),
    ('Маслята', 22, 1.3, 0.3, 2.5, 'Грибы'),
    ('Опята', 27, 3.5, 0.5, 2.5, 'Грибы'),
    ('Подберезовики', 20, 2.5, 0.5, 3.5, 'Грибы'),
    ('Моховик', 16, 2, 0.5, 1.5, 'Грибы'),
    ('Подосиновик', 22, 2.5, 0.5, 3.5, 'Грибы'),
    ('Пирожное', 453, 4.4, 26, 49, 'Кондитерские изделия'),
    ('Шоколад', 546, 5.4, 31, 60, 'Кондитерские изделия'),
    ('Мармелад', 260, 0.3, 0.2, 64, 'Кондитерские изделия'),
    ('Печенье', 440, 6, 19, 64, 'Кондитерские изделия'),
    ('Пончик', 452, 5.4, 25, 48, 'Кондитерские изделия'),
    ('Зефир', 308, 0.8, 0.2, 77, 'Кондитерские изделия'),
    ('Вафли', 440, 6, 19, 64, 'Кондитерские изделия'),
    ('Торт', 370, 3, 18, 52, 'Кондитерские изделия'),
    ('Кекс', 407, 4.7, 17, 58, 'Кондитерские изделия'),
    ('Мусс', 324, 2.8, 22, 29, 'Кондитерские изделия'),
    ('Оливковое масло', 884, 0, 100, 0, 'Пищевые жиры'),
    ('Сливочное масло', 717, 0.9, 81, 1.3, 'Пищевые жиры'),
    ('Кокосовое масло', 862, 0, 100, 0, 'Пищевые жиры'),
    ('Маргарин', 717, 0.3, 81, 0.1, 'Пищевые жиры'),
    ('Топленое масло', 876, 0.3, 99, 0, 'Пищевые жиры'),
    ('Льняное масло', 884, 0, 100, 0, 'Пищевые жиры'),
    ('Рапсовое масло', 884, 0, 100, 0, 'Пищевые жиры'),
    ('Масло подсолнечное', 884, 0, 100, 0, 'Пищевые жиры'),
    ('Чай', 1, 0, 0, 0, 'Напитки'),
    ('Кофе', 2, 0.1, 0, 0, 'Напитки'),
    ('Сок', 54, 0.3, 0, 13, 'Напитки'),
    ('Вода', 0, 0, 0, 0, 'Напитки'),
    ('Лимонад', 41, 0, 0, 10.3, 'Напитки'),
    ('Какао', 196, 1.9, 1.2, 20, 'Напитки'),
    ('Энергетический напиток', 45, 0, 0, 11, 'Напитки'),
    ('Минеральная вода', 0, 0, 0, 0, 'Напитки'),
    ('Смузи', 68, 1, 0.6, 15, 'Напитки'),
    ('Газировка', 40, 0, 0, 10, 'Напитки'),
    ('Компот', 60, 0.3, 0, 15, 'Напитки'),
    ('Квас', 33, 0.5, 0, 7, 'Напитки'),
    ('Лимонад', 54, 0.3, 0, 13, 'Напитки'),
    ('Матча', 340, 30, 5, 60, 'Напитки'),
    ('Морс', 45, 0, 0, 11, 'Напитки'),
    ('Фреш', 45, 0.4, 0.2, 10, 'Напитки');