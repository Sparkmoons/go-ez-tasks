## Сети, процессы, потоки итд

Тут я попытаюсь собрать информацию, которую нужно и полезно знать разработчику.

___1 января 1983 - День рождения Интернета___

### 1. Процессы и потоки

Процесс - базовая исполняемая единица в ОС. Состоит из одного или несколько потоков, текущего состояния и множества используемых ресурсов. 

___Процесс___ - программа в состоянии выполнения. Каждому процессу системой присваивается уникальный идентификационный номер (PID). Процесс состоит из: 1) код программы, 2) данные для функционирования кода, 3) контекст процесса (вся информация для управления процессом). 

При запуске одной программы многократно, например калькулятор, в ОС создается столько же процессов, каждый из которых независим друг от друга и изолирован. Обладает собственным адресным пространством и средой выполнения. 

Благодаря поддержке Copy on write данные не дублируются в физической памяти без необходимости. Копия создается лишь в тот момент, когда происходит ее изменения.  (Отсюда и название - механизм копирования при записи). 

При загрузке ОС (Windows) происходит создание базовых процессов. Первым загружается ядро.
Ядро вызывает создание несколько служебных процессов: 
1) Бездействие системы (system idle PID=0)
2) Первый настоящий процесс PID=4 - представляет ядро операционной системы. 
3) Инициализационная задача ядра для запуска системного менеджера сессий. SMSS (Session Manager Subsystem Service)
4) SMSS загружает остальные процессы для полной инициализации системы.
5) Далее, в режиме пользователя, SMSS загружает остальные процессы для полной инициализации системы.

__Unix подобные системы__

В Unix ядро не является процессом. В Linux, например, сначала запускаются разные системные процессы (0-sched, 1-init)

Основным механизмом создания новых процессов являются Fork() и Exec().

Fork() - копия вызывающего процесса (дочерний процесс). Создает два параллельных потока выполнения, вызывается один раз, а выполняется два раза. 

Exec() - замещает образ дочернего процесса  новой командой для выполнения. 

После fork родительский и дочерние процессы обладают раздельными адресными пространствами. 

В Unix всегда присутствует четкая иерархия процессов.


__Windows__ 

За создание новых процессов отвечает единый системный вызов - NtCreateProcess(). В нем много параметров, благодаря чему он обладает высокой мощностью и гибкостью в управлении процессами. 

___Потоки___

Это легковесные подзадачи, которые могут выполняться на ядрах процессора. У каждого как минимум один основной поток и, при необходимости, N дополнительных. Тредов может быть сколь угодно много, но в один момент времени может выполняться столько, сколько есть ядер. 

Основная идея состоит в отделении понятия процесса, то есть адресного пространства и ресурсов ОС от базовой единицы выполнения, известной как поток. 
Современные ОС придерживаются этой концепции и поддерживают два основных объекта: 1) процесс - который определяет общее адресное пространство и общие атрибуты процесса, 2) поток - который определяет последовательный поток выполнения в рамках этого процесса. 

Поток привязываются к одному процессу, но в одном адресном пространстве может быть множество потоков, что обеспечивает легкий доступ к общим данным. Создание потоков занимает очень мало времени, поскольку не требуется заново выделять адресное пространство или создавать объемные блоки управления процессами PCB, достаточно лишь информации для управления потоками. 

Таким образом объектом управления в ОС становятся потоки, а процессы служат контейнерами для их выполнения. 

В ядре ОС присутствуют системные вызовы, предназначенные для создания новых потоков. Создают блок управления потоками TCB (Thread Control Block) и присваивают им уникальный идентификатор потока TID (Thread IDentifier). 
User Mode Threads - потоки, управляемые по режиму пользователя. 

Поток (thread) — это, сущность операционной системы, процесс выполнения на процессоре набора инструкций, точнее говоря программного кода. Общее назначение потоков — параллельное выполнение на процессоре двух или более различных задач. Как можно догадаться, потоки были первым шагом на пути к многозадачным ОС. Планировщик ОС, руководствуясь приоритетом потока, распределяет кванты времени между разными потоками и ставит потоки на выполнение.

### 2. TCP и UDP
TCP (Transfer Control Protocol) - транспортный протокол передачи данных в сетях TCP/IP, предварительно устанавливающий соединение с сетью. Требует подтверждения получения пакета данных, для этого необходимо установленное заранее соединение. Считается надежным, исключает потери данных, дублирование, задержки и перемешивание пакетов. Контролирует загруженность соединения. Подходит для почтовых программ.

UDP (User Datagram Protocol) - транспортный протокол, передающий сообщения-датаграммы без необходимости установки соединения в IP-сети. Ненадежный протокол, но гораздо быстрее чем TCP. Ничего не контролирует, кроме целостности полученных данных. Подходит для сетевых/браузерных игр, приложений просмотра видео и видеосвязи. (Так как он быстрый и рассчитан на широкую пропускную способность)

### 3. DNS
DNS (Domain NAme System) - позволяет по доменному имени определить ip-адрес компьютера. 
 ( Поиск начинается с корневого домена, отправляется запрос на доменное имя. Корневой адрес отправляет ответ - не знает его ip-адрес, но знает ip-адрес днс сервера, которому делегировано управление зоной. ИТД)
Может работать в двух режимах:
1) Итеративный (Если сервер отвечает за ту зону, к которой пришел запрос - возвращает ответ. Если нет - возвращает адрес DNS сервера к которому нужно обратиться с запросом)
2) Рекурсивный (Сервер сам отправляет запросы всем DNS серверам пока не получит ответ)
Когда найден ip-адрес некоторого доменного имени - его помещают в кэш.
Состоит из:
1) Заголовок (Идентификатор запроса, Флаги, Количество запросов, Количество ответов, Количество авторитетных ответов, Количество дополнительных ответов)
2) Данные (Запросы DNS((Имя, Тип записи, Класс записи)), Ответы DNS((Имя, Тип записи, Класс записи, Время жизни(TTL), Длина данных, Данные)), Авторитетные серверы, Дополнительная информация)

### 4. HTTP/HTTPS
HTTP (HyperText Transfer Protocol) - протокол передачи гипертекста. Прикладной протокол передачи данных в сети.
(Стартовая строка - Заголов - Тело сообщения) Все эти данные передаются в открытом виде.

HTTPS (HyperText Transfer Protocol Secure) - расширение протокола HTTP, безопасный протокол передачи гипертекста поддерживающий шифрование посредством криптографических протоколов SSL и TLS.
  SSL (Secure Sockets Layer) - при подключении к сайту просим идентифицировать его и получаем в ответ копию его SSL сертификата, доказывающий, что общаться с ним безопасно. Если все хорошо - начнется обмен шифрованными данными.
  TLS (Transport Layer Security) - также проверяет подлинность сервера и шифрует данные. Сейчас используют TLS1.2 и TLS1.3. Конфиденциальность, Целостность, Аутентификация (Ссылка начинается с https - соединение по TLS).

Различия HTTP и HTTPS:
1) HTTPS не является отдельным протоколом передачи данных, а представляет собой расширение протокола HTTP с надстройкой шифрования
2) Передаваемые по протоколу HTTP данные не защищены, HTTPS обеспечивает конфиденциальность информации путем ее шифрования
3) HTTP использует порт 80, HTTPS - порт 443.

### 5. IPv4 и IPv6

IPv4 (Internet Protocol version 4) и IPv6 (Internet Protocol version 6) - это две разные версии Интернет-протокола, который определяет способ передачи и получения данных через Интернет.

IPv4 - это первоначальная версия Интернет-протокола, а IPv6 - более новая версия, разработанная для устранения ограничений IPv4. Одним из основных ограничений IPv4 является ограниченное адресное пространство, в котором можно использовать только 32-разрядные десятичные числа.

В отличие от этого, адреса IPv6 являются 128-разрядными, что обеспечивает большие блоки и практически неограниченное количество уникальных адресов. Увеличение числа доступных адресов в IPv6 позволяет подключать к Интернету большее количество устройств, поскольку мы движемся к более связанному будущему с растущим числом устройств.

Большинство современных устройств поддерживают как IPv4, так и IPv6, однако некоторые старые аппаратные средства могут поддерживать только один или другой вариант. Таким образом, переход на IPv6 становится все более важным для обеспечения непрерывной связи.

Пример адреса IPv4: 192:168:0:0

Пример адреса IPv6: 2001:0db8:85a3:0000:0000:8a2e:0370:7334. Группу 0 можно заменить на ::, тогда получаем: 2001:0db8:85a3::8a2e:0370:7334. 


### 6. NAT
NAT (Network Address Translation) - переводит приватные адреса в общедоступные. Это позволяет устройству с частным адресом IPv4 обращаться к ресурсам за пределами его частной сети. NAT в сочетании с частными адресами IPv4 оказался полезным методом сохранения общедоступных IPv4-адресов. Один общедоступный IPv4-адрес может быть использован сотнями, даже тысячами устройств, каждый из которых имеет частный IPv4-адрес. NAT имеет дополнительное преимущество, заключающееся в добавлении степени конфиденциальности и безопасности в сеть, поскольку он скрывает внутренние IPv4-адреса из внешних сетей.

Сети обычно проектируются с использованием частных IP адресов. Это адреса 10.0.0.0/8, 172.16.0.0/12 и 192.168.0.0/16. Эти частные адреса используются внутри организации или площадки, чтобы позволить устройствам общаться локально, и они не маршрутизируются в интернете. Чтобы позволить устройству с приватным IPv4-адресом обращаться к устройствам и ресурсам за пределами локальной сети, приватный адрес сначала должен быть переведен на общедоступный публичный адрес.

### 7. Виды деплоя

<img src="https://github.com/Sparkmoons/go_ez_tasks/blob/main/img/Deploys.png">

### 8. 12FA

1) Кодовая база (Одна кодовая база, отслеживаемая в системе контроля версий, множество развертываний)

2) Зависимости (Явно объявляйте и изолируйте поля)
   
3) Конфигурация (Сохраняйте конфигурацию в среде выполнения)
   
4) Сторонние сервисы (Backing Services) (Считайте сторонние службы подключаемыми ресурсами)

5) Сборка, релиз, выполнение (Строго разделяйте стадии сборки и выполнения)

6) Процессы (Запускайте приложения как оидн или несколько процессов несохраняющих внутреннее состояние (Stateless))

7) Связка портов (Port Binding) (Экспортируйте сервисы через привязку портов)

8) Параллелизм (Масштабируйте приложение с помощью процееов)

9) Утилизируемость (Disposability) (Максимизируйте надежность с помощью быстрого запуска и корректного завершения работы)

10) Паритет разработки/работы приложения (Держите окружение разработки промежуточного развертывания (Staging) и рабочего развертывания (Production) максимально похожими)

11) Журанлирование (Logs) (Рассматривайте журнал как поток событий)

12) Задачи администрирования (Выполняйте задачи администрирования/управления с помощью разовых процессов)

### 9. Свойства архитектуры

1) Доступность

2) Надежность

3) Тестируемость

4) Масштабируемость

5) Безопасность

6) Гибкость

7) Отказоустойчивость

8) Адаптируемость

9) Восстанавливаемость

10) Производительность

11) Раазвертываемость

12) Обучаемость

13) (Проверяемость, Производительность, Безопасность, Требования, Данные, Допустимость, Масштабируемость)

### 10. Полезные аббревиатуры

DHCP - Dynamic Host Configuration Protocol (Протокол динамической настройки узла)

CSMA/CA - Carrier Sense Multiple Access with Collision Avoidance (Множественный доступ с прослушиванием несущей частоты и избежанием коллизий)

CSMA/CD - Carrier Sense Multiple Access with Collision Detection (Множественный доступ с прослушиванием несущей частоты и обнаружением коллизий)

API - Application Programming Interface (Интерфейс программирования приложения)

Rest API - REpresentational State Transfer (Передача репрезентативного состояния)

URL - Uniform Resource Locator (Единый указатель ресурса)

CRUD - Create Read Update Delete (Создать, Почитать, Обновить, Удалить)

KISS - Keep It Simple, Stupid (Будь проще)

DRY - Don't Repeat Yourself (Не повторяйся)

YAGNI - You Aren't Gonna Need It (Тебе это не понадобится)

BDUF - Big Data Up Front (Глобальное проектирование прежде всего)

APO - Avoid Premature Optimization (Избегайте преждевременной оптимизации)

SOLID:
 - Single-responsubulity principle (Принцип единственной ответственности)
 - Open/closed principle (Принцип открытости/закрытости)
 - Liskov substitution principle (Принцип подстановки Лисков)
 - Interface segregation principle (Принцип разделения интерфейсов)
 - Dependency inversion principle (Принцип инверсии зависимостей)

### 11. БД

___Индекс___ - это объект, который создается с целью повышения производительности поиска данных. Объект БД содержащий упорядоченные значения указанных столбцов таблицы и ссылки на физическое размещение записи с данными значениями. Индекс позвоялет ускорить поиск данных в таблице и упорядочивание данных. Обычно пользователи или админы БД сами создают необходимые для этого индексы с помощью оператора CREATE INDEX.

___Транзакция___ - это группа последовательных операций с БД, которая представляет собой логическую единицу работы с данными. Транзакция может быть выполнена либо целиком и успешно, соблюдая целостность данных и независимо от параллельно идущих других траназкций, либо не выполнена вообще, и тогда она не должна произвести никакого эффекта.

Транзакции обрабатываются транзакционными системами, в прцоессе работы которых создаётся история транзакций.

Различают следующие виды транзакций: 
- Последовательные (обычные)
- Параллельные
- Распределенные. Подразумевают использование более чем одной транзакционной системы и требуют намного более сложной логики.


