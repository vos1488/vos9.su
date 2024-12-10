package main

import (
	"fmt"
	"net/http"
)

const htmlContent = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Добро пожаловать на vos9.su</title>
    <style>
        body {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            background-color: #f0f0f0;
            margin: 0;
        }
        .heart {
            font-size: 100px;
            color: red;
            animation: pulse 1s ease infinite;
            cursor: pointer;
        }
        @keyframes pulse {
            0% { transform: scale(1); }
            50% { transform: scale(1.3); }
            100% { transform: scale(1); }
        }
        .message {
            margin-top: 20px;
            font-size: 18px;
            display: none;
        }
        footer {
            width: 100%;
            background-color: #ddd;
            text-align: center;
            padding: 10px 0;
            position: fixed;
            bottom: 0;
            left: 0;
            display: flex;
            align-items: center;
        }
        .gitea-button {
            margin-left: 20px;
            font-size: 18px;
            padding: 10px 20px;
            cursor: pointer;
        }
        footer p {
            flex-grow: 1;
            margin: 0;
        }
    </style>
</head>
<body>
    <div class="heart" onclick="showMessage()">❤️</div>
    <div id="message" class="message">Спасибо за посещение сайта, скоро всё будет</div>

    <!-- Убираем кнопку Gitea с прежнего места -->
    <!--
    <a href="https://gitea.vos9.su" class="gitea-button">
        <img src="https://gitea.vos9.su/assets/img/logo.svg" alt="Gitea Logo">	
    </a>
    -->

    <!-- Добавляем футер -->
    <footer>
        <!-- Перемещаем кнопку Gitea в футер -->
        <a href="https://gitea.vos9.su" class="gitea-button">
            <img src="https://gitea.vos9.su/assets/img/logo.svg" alt="Gitea Logo">	
        </a>
        <p>© 2024 vos9.su. Все права защищены.</p>
    </footer>

    <script>
        function showMessage() {
            document.getElementById('message').style.display = 'block';
        }
        function openGitea() {
            window.location.href = 'https://gitea.vos9.su';
        }
    </script>
</body>
</html>
`

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlContent)
}

func main() {
	http.HandleFunc("/", handleRoot)
	fmt.Println("Сервер запущен на http://localhost:8999")
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s\n", err)
	}
}
