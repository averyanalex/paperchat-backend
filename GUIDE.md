Как принимать участие в разработке
Клонируем
1. git clone https://github.com/paper-chat/nnm
Открываем каталог
2. cd nnm
3. Откройте VSCode
Создаём новую ветвь
4. git checkout -b feature/add-some-bruh
5. Делайте изменения
Чтобы гит увидел, что нужно коммитить
6. Stage changes в разделе source control
Но верифи нужен в app, ибо сейчас токо так
7. git commit -m "Add some bruh" --no-verify (Можно и в VSCode)
Загружаем ветвь
8. git push origin feature/add-some-bruh
9. Создаете pull request на гитхабе
10. git checkout main (переходим на главую ветвь)
11. git pull --ff (получаем изменения, --ff что бы не было лишних merge commit)
12. Перейдите к шагу 4
