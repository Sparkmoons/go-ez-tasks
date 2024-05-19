<div align="center">
  <img src="https://github.com/Sparkmoons/go_ez_tasks/blob/main/img/gogogo.png">
  <h1>Некоторые данные для подготовки к собеседованию на Golang-разработчика</h1>
</div>


Всем привет. 
Осенью 2023 мне рассказали про Golang, до этого начинал учить плюсы и писал драйвера на простом С, также были небольшие истории с Python. Я начал его учить и тут решил собрать задачи, про которые слышал на собеседованиях, да и в целом они будут полезны. 
Файл с теорией, пока что, не структурирован, не исключено, что я это сделаю позже. Сейчас это просто заметка из моего девайса, в которую я закидывал информацию, которую слышал где-либо и считал ее полезной.
Постепенно все будет расширяться и добавляться :)



1. #### [Теория (язык)](readme/theory1)
   - [Тред, Процессор, Горутина](readme/theory1#1-тред-процессор-горутина)
   - [Плюсы и минусы](readme/theory1#2-плюсы-и-минусы)
   - [Типы данных](readme/theory1#3-типы-данных)
   - [Ключевые слова](readme/theory1#4-ключевые-слова)
   - [ООП](readme/theory1#5-ооп)
   - [Массив и слайс](readme/theory1#6-массив-и-слайс)
   - [Интерфейсы](readme/theory1#7-интерфейсы)
   - [Указатели](readme/theory1#8-указатели)
   - [Map](readme/theory1#9-map)
   - [Функции](readme/theory1#10-функции)
   - [Замыкание](readme/theory1#11-замыкание)
   - [Defer](readme/theory1#12-defer)
   - [Методы](readme/theory1#13-методы)
   - [Mutex](readme/theory1#14-mutex)
   - [strconv](readme/theory1#15-strconv)
   - [iota](readme/theory1#16-iota)
   - [Каналы](readme/theory1#17-каналы)
   - [JSON](readme/theory1#18-json)
   - [Сериализация](readme/theory1#19-сериализация)
   - [Context](readme/theory1#20-context)
   - [Recover Panic](readme/theory1#21-recover-panic)
   - [Немного про память](readme/theory1#22-немного-про-память)
   - [Сборщик мусора](readme/theory1#23-сборщик-мусора)
   - [Каналы и пайпы](readme/theory1#24-каналы-и-пайпы)
   - [Docker](readme/theory1#25-docker)
2. #### [Теория (надо знать)](readme/theory2)
   - [Процессы и потоки](readme/theory2#1-процессы-и-потоки)
   - [TCP и UDP](readme/theory2#2-tcp-и-udp)
   - [DNS](readme/theory2#3-dns)
   - [HTTP/HTTPS](readme/theory2#4-http/https)
   - [IPv4 и IPv6](readme/theory2#5-ipv4-и-ipv6)
   - [NAT](readme/theory2#6-nat)
   - [Виды деплоя](readme/theory2#7-виды-деплоя)
   - [12FA](readme/theory2#8-12fa)
   - [Свойства архитектуры](readme/theory2#9-свойства-архитектуры)
   - [Полезные аббревиатуры](readme/theory2#10-полезные-аббревиатуры)


4. #### [Задачи](#tasks)
    1. [Задача на префикс.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/PREFIX.md)
    2. [Проверка слова на палиндром.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/PALINDROM.md)
    3. [Проверка слайса на уникальность компонентов.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/UNIQUE.md)
    4. [Вывод пересечения двух слайсов.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/CROSS.md)
    5. [Слияние двух слайсов в уникальный.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/UNIQUE_SL.md)
    6. [Поиск числа в слайсе.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/FIND_NUMB.md)
    7. [Сортировка.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/SORT.md)
    8. [Вывод числа Фибоначчи.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/FIBNUM.md)
    9. [Вывод номера числа Фибоначчи.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/FIBNUMNUM.md)
    10. [Вывод последовательности чисел Фибоначчи.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/FIBSEQ.md)
    11. [Вывод всех возможных комбинаций из букв.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/ALLCOMB.md)
    12. [Слияние n каналов в один.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/MERGE_CH.md)
    13. [Слияние двух числовых каналов.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/TWOINTONE.md)
    14. [Реализация LIFO.](https://github.com/Sparkmoons/go_ez_tasks/blob/main/readme/LIFO.md)
    15. [Рандомайзер.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/RANDOMIZER.md)
    16. [Конвейер чисел.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/CONVEYER.md)
    17. [Разворот слов в строке.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/REVERSEWORD.md)
    18. [Нахождение цифрового корня числа.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/DROOT.md)
    19. [Поиск самой длинной уникальной подстроки.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/slide.md)
    20. [Перенос чисел в конец слайса.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/moveNumb.md)
    21. [Сложение двоичных чисел.](https://github.com/Sparkmoons/go-ez-tasks/blob/main/readme/bin_sum.md)
