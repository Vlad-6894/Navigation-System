document.addEventListener("DOMContentLoaded", () => {
    // 1. Перехватываем форму на текущей странице (хоть вход, хоть регистрация)
    const authForm = document.querySelector("form");
    if (!authForm) return;

    authForm.addEventListener("submit", async (event) => {
        // Отменяем стандартное поведение браузера (перезагрузку страницы)
        event.preventDefault();

        // 2. Автоматически собираем все данные из полей формы
        const formData = new FormData(authForm);
        const jsonData = {};

        formData.forEach((value, key) => {
            // КРИТИЧЕСКИ ВАЖНО ДЛЯ GO: приводим числовые поля к типу Number, 
            // чтобы твой json.NewDecoder не падал на структурах int64!
            if (key === "school_id" || key === "age") {
                jsonData[key] = Number(value);
            } else {
                jsonData[key] = value;
            }
        });

        // Берем URL бэкенда прямо из атрибута action нашей HTML-формы
        const requestUrl = authForm.getAttribute("action");

        try {
            // 3. Отправляем асинхронный HTTPS запрос с JSON-телом
            const response = await fetch(requestUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(jsonData)
            });

            // Читаем ответ от твоего кастомного пакета pkg_http_response
            const result = await response.json();

            // 4. Логика обработки ответов бэкенда
            if (response.ok) {
                // Выводим красивое уведомление об успехе
                alert(result.message || "Операция успешно выполнена!");

                // РЕЖИМ №1: Если это был эндпоинт регистрации
                if (requestUrl.includes("register")) {
                    // Перенаправляем пользователя на окно входа
                    window.location.href = "login.html";
                } 
                // РЕЖИМ №2: Если это был эндпоинт авторизации (логина)
                else if (requestUrl.includes("login")) {
                    // Надежно сохраняем выданные бэкендом токены в локальный кэш браузера
                    if (result.access_token) localStorage.setItem("access_token", result.access_token);
                    if (result.refresh_token) localStorage.setItem("refresh_token", result.refresh_token);

                    // Переводим пользователя на страницу с интерактивной картой школы
                    window.location.href = "map.html"; 
                }
            } else {
                // Если бэкенд прислал ошибку (например, "fail decode dto!")
                alert(`Ошибка: ${result.error || "Произошел сбой при обработке запроса"}`);
            }

        } catch (error) {
            console.error("Сетевая ошибка:", error);
            alert("Не удалось связаться с сервером. Проверь, запущен ли твой Go-бэкенд на порту 8443.");
        }
    });
});
