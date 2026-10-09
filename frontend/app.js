function switchFloor(floorNumber) {
    // 1. Находим все слои этажей и прячем их (стираем класс active)
    document.querySelectorAll('.floor-layer').forEach(layer => {
        layer.classList.remove('active');
    });
    
    // 2. Находим все кнопки и гасим их синюю подсветку
    document.querySelectorAll('.floor-btn').forEach(btn => {
        btn.classList.remove('active');
    });

    // 3. Включаем карту выбранного этажа
    const targetLayer = document.getElementById(`container-floor-${floorNumber}`);
    if (targetLayer) {
        targetLayer.classList.add('active');
    }

    // 4. Подсвечиваем синим нажатую кнопку в меню
    const allButtons = document.querySelectorAll('.floor-btn');
    if (allButtons[floorNumber - 1]) {
        allButtons[floorNumber - 1].classList.add('active');
    }
}
